package chain

import (
	"crypto/ed25519"
	"github.com/blocto/solana-go-sdk/pkg/hdwallet"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/luzen23141/mouse/pkg/blockchain/model"
	cryptolib "github.com/luzen23141/mouse/pkg/lib/cyptolib"
	"github.com/rotisserie/eris"
	"github.com/shopspring/decimal"
	"github.com/tyler-smith/go-bip39"
	"github.com/xssnick/tonutils-go/ton/wallet"
)

type TonChain struct {
	tonVersion _tonAddrConfig
}

func NewTonChain() *TonChain {
	return &TonChain{
		tonVersion: _tonAddrConfigMap["v5"],
	}
}

// 跟btc一樣，有很多版本的地址跟變數，每個錢包使用的版本或變數不同
var _tonAddrConfigMap = map[string]_tonAddrConfig{
	// 2024-11-04 trustWallet 為v4
	//            bitgetWallet 為v5,v4 可切換
	//
	"v5": {
		versionCfg: wallet.ConfigV5R1Final{
			NetworkGlobalID: wallet.MainnetGlobalID,
		},
		path:      `m/44'/607'/0'`,
		subWallet: uint32(0),
	},
	"v4": {
		versionCfg: wallet.V4R2,
		path:       `m/44'/607'/0'`,
		subWallet:  wallet.DefaultSubwallet,
	},
	"v3": {
		versionCfg: wallet.V3R2,
		path:       `m/44'/607'/0'/0'`,
		subWallet:  wallet.DefaultSubwallet,
	},
}

type _tonAddrConfig struct {
	versionCfg wallet.VersionConfig
	subWallet  uint32
	path       string
}

// GenAddr 产生地址
func (s *TonChain) GenAddr() (string, string, error) {
	// 生成新的密钥对
	pubKey, privateKey, err := ed25519.GenerateKey(nil)
	if err != nil {
		return "", "", err
	}

	addrConfig := _tonAddrConfigMap["default"]
	address, err := wallet.AddressFromPubKey(pubKey, addrConfig.versionCfg, addrConfig.subWallet)
	if err != nil {
		return "", "", err
	}

	// Bounce(false) => UQ 開頭，Bounce(true) => EQ 開頭
	return address.Bounce(false).String(), hexutil.Encode(privateKey.Seed()), nil
}

// GenHdAddr 产生Hd wallet 地址，返回的key為mnemonic
func (s *TonChain) GenHdAddr() (string, string, error) {
	// 生成新的密钥对
	mnemonic, _, err := cryptolib.NewMnemonic()
	if err != nil {
		return "", "", err
	}

	addrConfig := _tonAddrConfigMap["v5"]
	derivedKey, err := hdwallet.Derived(addrConfig.path, bip39.NewSeed(mnemonic, ""))
	if err != nil {
		return "", "", err
	}

	priKey := ed25519.NewKeyFromSeed(derivedKey.PrivateKey)
	pubKey := priKey.Public().(ed25519.PublicKey)

	address, err := wallet.AddressFromPubKey(pubKey, addrConfig.versionCfg, addrConfig.subWallet)
	if err != nil {
		return "", "", err
	}

	// Bounce(false) => UQ 開頭，Bounce(true) => EQ 開頭
	return address.Bounce(false).String(), mnemonic, nil
}

func (s *TonChain) GetAddrBalance(addr string, cur model.CurrencyContract) (decimal.Decimal, error) {
	return decimal.Zero, eris.New("not support")
}
