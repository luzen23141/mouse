package chain

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rotisserie/eris"
	"github.com/shopspring/decimal"
	"mouse/pkg/blockchain/model"
	"mouse/pkg/lib/cyptolib"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
)

type EvmChain struct{}

func NewEvmChain() *EvmChain {
	return &EvmChain{}
}

func (*EvmChain) GenAddr() (string, string, error) {
	// 1. 生成私鑰
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return "", "", err
	}

	// 2. 從私鑰獲取公鑰
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", "", err
	}

	// 3. 從公鑰獲取地址
	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	return strings.ToLower(address.Hex()), hex.EncodeToString(privateKey.D.Bytes()), nil
}

func (s *EvmChain) GenHdAddr() (string, string, error) {
	mnemonic, _, err := cryptolib.NewMnemonic()
	if err != nil {
		return "", "", err
	}

	publicKeyECDSA, _, err := cryptolib.MnemonicToEcdsaPubKey(mnemonic, "m/44'/60'/0'/0/0")
	if err != nil {
		return "", "", err
	}

	return crypto.PubkeyToAddress(*publicKeyECDSA).Hex(), mnemonic, nil
}

func (s *EvmChain) GetAddrBalance(addr string, cur model.CurrencyContract) (decimal.Decimal, error) {
	if cur.IsGov {
		conn, err := s.getClient()
		if err != nil {
			return decimal.Zero, eris.Wrap(err, "failed to connect ethclient")
		}

		// Get the balance of an account
		account := common.HexToAddress(addr)
		balance, err := conn.BalanceAt(context.Background(), account, nil)
		if err != nil {
			return decimal.Zero, eris.Wrap(err, "failed to get balance")
		}

		return decimal.NewFromBigInt(balance, cur.Decimal), nil
	}

	conn, err := s.getClient()
	if err != nil {
		return decimal.Zero, eris.Wrap(err, "failed to connect ethclient")
	}

	// 建立 USDT 合約實例
	contractModel, err := NewEip20(common.HexToAddress(cur.Addr), conn)
	if err != nil {
		return decimal.Zero, eris.Wrap(err, "failed to create contract model")
	}

	// 呼叫 balanceOf 函數獲取餘額
	balance, err := contractModel.BalanceOf(&bind.CallOpts{}, common.HexToAddress(addr))
	if err != nil {
		return decimal.Zero, eris.Wrap(err, "failed to call balanceOf")
	}

	return decimal.NewFromBigInt(balance, cur.Decimal), nil
}

func (*EvmChain) getClient() (*ethclient.Client, error) {
	return ethclient.Dial("https://ethereum-rpc.publicnode.com")
}
