package chain

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"

	"github.com/InjectiveLabs/sdk-go/client/common"
	exchangeclient "github.com/InjectiveLabs/sdk-go/client/exchange"
	"github.com/btcsuite/btcd/btcutil/bech32"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/luzen23141/mouse/pkg/blockchain/model"
	cryptolib "github.com/luzen23141/mouse/pkg/lib/cyptolib"
	"github.com/rotisserie/eris"
	"github.com/shopspring/decimal"
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

func (s *InjChain) GetAddrByPrivKey(privKeyStr string) (string, string, error) {
	publicKey, privateKey, err := cryptolib.StrKeyToEcdsaKeyPair(privKeyStr)
	if err != nil {
		return "", "", err
	}

	ethAddress := crypto.PubkeyToAddress(*publicKey)
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

	return s.GetAddrByMnemonic(mnemonic)
}

func (s *InjChain) GetAddrByMnemonic(mnemonic string) (string, string, error) {
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
	exchangeClient, err := s.getClient()
	if err != nil {
		return decimal.Zero, err
	}

	res, err := exchangeClient.GetAccountPortfolioBalances(context.Background(), addr)
	if err != nil {
		return decimal.Zero, err
	}
	if res.GetPortfolio().AccountAddress != addr { // 檢查資料有沒有成功取回來
		return decimal.Zero, eris.New("not support")
	}

	allBalances := res.GetPortfolio().GetBankBalances()
	if len(allBalances) == 0 {
		return decimal.Zero, nil
	}

	for _, v := range allBalances {
		if v.GetDenom() == cur.Addr {
			balance, err := decimal.NewFromString(v.GetAmount())
			if err != nil {
				return decimal.Zero, err
			}
			return balance.Mul(decimal.New(1, cur.Decimal)), nil
		}
	}

	return decimal.Zero, err
}

func (s *InjChain) getClient() (exchangeclient.ExchangeClient, error) {
	network := common.LoadNetwork("mainnet", "lb")
	exchangeClient, err := exchangeclient.NewExchangeClient(network)
	if err != nil {
		return nil, err
	}
	return exchangeClient, nil
}
