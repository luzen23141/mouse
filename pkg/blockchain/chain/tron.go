package chain

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"github.com/luzen23141/mouse/pkg/blockchain/model"
	"github.com/luzen23141/mouse/pkg/lib/cyptolib"
	"github.com/rotisserie/eris"
	"github.com/shopspring/decimal"
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/mr-tron/base58"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type TronChain struct{}

func NewTronChain() *TronChain {
	return &TronChain{}
}

// GenAddr 产生TRON地址
func (s *TronChain) GenAddr() (string, string, error) {
	privateKey, _ := crypto.GenerateKey()
	privateKeyBytes := crypto.FromECDSA(privateKey)
	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	return s.pubKeyToAddr(*publicKeyECDSA), hexutil.Encode(privateKeyBytes)[2:], nil
}

func (s *TronChain) GenHdAddr() (string, string, error) {
	mnemonic, _, err := cryptolib.NewMnemonic()
	if err != nil {
		return "", "", err
	}

	publicKeyECDSA, _, err := cryptolib.MnemonicToEcdsaPubKey(mnemonic, "m/44'/195'/0'/0/0")
	if err != nil {
		return "", "", err
	}

	return s.pubKeyToAddr(*publicKeyECDSA), mnemonic, nil
}

func (s *TronChain) GetAddrBalance(addr string, cur model.CurrencyContract) (decimal.Decimal, error) {

	conn, err := s.getClient()
	if err != nil {
		return decimal.Zero, err
	}
	if cur.IsGov {
		account, err := conn.GetAccount(addr)
		if err != nil {
			if err.Error() == "account not found" {
				return decimal.Zero, nil
			}
			return decimal.Zero, eris.Wrap(err, "Failed to get balance")
		}

		return decimal.NewFromBigInt(big.NewInt(account.Balance), cur.Decimal), nil
	}

	balance, err := conn.TRC20ContractBalance(addr, cur.Addr)
	if err != nil {
		fmt.Printf("Failed to get balance: %v", err)
		return decimal.Zero, eris.Wrapf(err, "Failed to get balance, addr:%s, contractAddr:%s", addr, cur.Addr)
	}

	return decimal.NewFromBigInt(balance, cur.Decimal), nil
}

func (s *TronChain) getClient() (*client.GrpcClient, error) {
	grpcClient := client.NewGrpcClient("grpc.trongrid.io:50051")
	err := grpcClient.Start(grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return grpcClient, nil
}

func (s *TronChain) pubKeyToAddr(p ecdsa.PublicKey) string {
	addr := crypto.PubkeyToAddress(p).Hex()
	addrStr := "41" + addr[2:]
	addrByte, _ := hex.DecodeString(addrStr)
	firstHash := sha256.Sum256(addrByte)
	secondHash := sha256.Sum256(firstHash[:])
	secret := secondHash[:4]
	addrByte = append(addrByte, secret...)
	return base58.Encode(addrByte)
}
