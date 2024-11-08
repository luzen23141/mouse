package chain

import (
	"context"
	"crypto/ed25519"

	"github.com/blocto/solana-go-sdk/client"
	"github.com/blocto/solana-go-sdk/pkg/hdwallet"
	"github.com/blocto/solana-go-sdk/rpc"
	"github.com/blocto/solana-go-sdk/types"
	"github.com/luzen23141/mouse/pkg/blockchain/model"
	cryptolib "github.com/luzen23141/mouse/pkg/lib/cyptolib"
	"github.com/mr-tron/base58"
	"github.com/shopspring/decimal"
	"github.com/tyler-smith/go-bip39"
)

type SolChain struct{}

func NewSolChain() *SolChain {
	return &SolChain{}
}

// GenAddr 产生地址
func (s *SolChain) GenAddr() (string, string, error) {
	// 生成新的密钥对
	pubKey, privateKey, err := ed25519.GenerateKey(nil)
	if err != nil {
		return "", "", err
	}

	return base58.Encode(pubKey), base58.Encode(privateKey), nil
}

func (s *SolChain) GetAddrByPrivKey(privKeyStr string) (string, string, error) {
	publicKey, privateKey, err := cryptolib.StrKeyToEd25519KeyPair(privKeyStr)
	if err != nil {
		return "", "", err
	}
	return base58.Encode(publicKey), base58.Encode(privateKey), nil
}

// GenHdAddr 产生Hd wallet 地址，返回的key為mnemonic
func (s *SolChain) GenHdAddr() (string, string, error) {
	mnemonic, _, err := cryptolib.NewMnemonic()
	if err != nil {
		return "", "", err
	}

	return s.GetAddrByMnemonic(mnemonic)
}

func (s *SolChain) GetAddrByMnemonic(mnemonic string) (string, string, error) {
	derivedKey, err := hdwallet.Derived(`m/44'/501'/0'/0'`, bip39.NewSeed(mnemonic, ""))
	if err != nil {
		return "", "", err
	}

	account, err := types.AccountFromSeed(derivedKey.PrivateKey)
	if err != nil {
		return "", "", err
	}

	return account.PublicKey.ToBase58(), mnemonic, nil
}

func (s *SolChain) GetAddrBalance(addr string, cur model.CurrencyContract) (decimal.Decimal, error) {
	conn, err := s.getClient()
	if err != nil {
		return decimal.Zero, err
	}

	if cur.IsGov {
		// get balance
		balance, err := conn.GetBalance(context.TODO(), addr)
		if err != nil {
			return decimal.Zero, err
		}

		return decimal.NewFromUint64(balance).Mul(decimal.New(1, cur.Decimal)), nil
	}

	tokenAccs, err := conn.GetTokenAccountsByOwnerByMint(context.TODO(), addr, cur.Addr)
	if err != nil {
		return decimal.Zero, err
	}
	balance := uint64(0)
	for _, v := range tokenAccs {
		balance += v.Amount
	}
	return decimal.NewFromUint64(balance).Mul(decimal.New(1, cur.Decimal)), nil
}

func (s *SolChain) getClient() (*client.Client, error) {
	return client.NewClient(rpc.MainnetRPCEndpoint), nil
}
