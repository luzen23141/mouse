package chain

import (
	"crypto/ecdsa"
	"encoding/hex"
	"github.com/rotisserie/eris"
	"github.com/shopspring/decimal"
	"mouse/pkg/blockchain/model"
	"mouse/pkg/lib/cyptolib"

	"github.com/btcsuite/btcd/btcutil/bech32"
	"github.com/ethereum/go-ethereum/crypto"
)

type InjChain struct{}

func NewInjChain() *InjChain {
	return &InjChain{}
}

// GenAddr 产生地址
func (s *InjChain) GenAddr() (string, string, error) {
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
	ethAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// 4. 從eth地址格式轉換成inj格式
	address, err := bech32.EncodeFromBase256("inj", ethAddress.Bytes())
	if err != nil {
		panic(err)
	}

	return address, hex.EncodeToString(privateKey.D.Bytes()), nil
}

// GenHdAddr 产生Hd wallet 地址，返回的key為mnemonic
func (s *InjChain) GenHdAddr() (string, string, error) {
	mnemonic, _, err := cryptolib.NewMnemonic()
	if err != nil {
		return "", "", err
	}

	publicKeyECDSA, _, err := cryptolib.MnemonicToEcdsaPubKey(mnemonic, "m/44'/60'/0'/0/0")
	if err != nil {
		return "", "", err
	}

	// 4. 從eth地址格式轉換成inj格式
	address, err := bech32.EncodeFromBase256("inj", crypto.PubkeyToAddress(*publicKeyECDSA).Bytes())
	if err != nil {
		return "", "", err
	}

	return address, mnemonic, nil
}

func (s *InjChain) GetAddrBalance(addr string, cur model.CurrencyContract) (decimal.Decimal, error) {
	return decimal.Zero, eris.New("not support")
}
