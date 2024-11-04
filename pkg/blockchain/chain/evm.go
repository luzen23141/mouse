package chain

import (
	"crypto/ecdsa"
	"encoding/hex"
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
	return decimal.Zero, eris.New("not support")
}
