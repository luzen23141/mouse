package chain

import (
	"crypto/ed25519"
	"github.com/rotisserie/eris"
	"github.com/shopspring/decimal"
	"mouse/pkg/blockchain/model"
	"mouse/pkg/lib/cyptolib"

	"github.com/blocto/solana-go-sdk/pkg/hdwallet"
	"github.com/blocto/solana-go-sdk/types"
	"github.com/mr-tron/base58"
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

// GenHdAddr 产生Hd wallet 地址，返回的key為mnemonic
func (s *SolChain) GenHdAddr() (string, string, error) {
	mnemonic, _, err := cryptolib.NewMnemonic()
	if err != nil {
		return "", "", err
	}

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
	return decimal.Zero, eris.New("not support")
}
