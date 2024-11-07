package chain

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/bytedance/sonic"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/go-resty/resty/v2"
	"github.com/luzen23141/mouse/pkg/blockchain/model"
	cryptolib "github.com/luzen23141/mouse/pkg/lib/cyptolib"
	"github.com/rotisserie/eris"
	"github.com/shopspring/decimal"
)

type BtcChain struct {
	addrType string
	cfg      model.BtcCfg
}

func NewBtcChain(cfg model.BtcCfg) *BtcChain {
	cfg.URL = strings.TrimSuffix(cfg.URL, "/")
	return &BtcChain{
		addrType: _btcAddrLegacy,
		cfg:      cfg,
	}
}

func (s *BtcChain) GenAddr() (string, string, error) {
	privKeyEc, err := btcec.NewPrivateKey()
	if err != nil {
		return "", "", err
	}
	wif, err := btcutil.NewWIF(privKeyEc, s.cfg.NetParams, true)
	if err != nil {
		return "", "", err
	}
	privKey := wif.PrivKey
	pubKey := privKey.PubKey()

	addr, err := _btcGenAddr(s.addrType, pubKey, s.cfg.NetParams)
	if err != nil {
		return "", "", err
	}

	return addr, wif.String(), nil
}

func (s *BtcChain) GetAddrByPrivKey(privKeyStr string) (string, string, error) {
	var (
		wif *btcutil.WIF
		err error
	)
	if strings.HasPrefix(privKeyStr, "0x") {
		seed, err := hexutil.Decode(privKeyStr)
		if err != nil {
			return "", "", err
		}

		privKeyEc, _ := btcec.PrivKeyFromBytes(seed)
		wif, err = btcutil.NewWIF(privKeyEc, s.cfg.NetParams, true)
		if err != nil {
			return "", "", err
		}
	} else {
		wif, err = btcutil.DecodeWIF(privKeyStr)
		if err != nil {
			return "", "", err
		}
	}

	privKey := wif.PrivKey
	pubKey := privKey.PubKey()

	addr, err := _btcGenAddr(s.addrType, pubKey, s.cfg.NetParams)
	if err != nil {
		return "", "", err
	}

	return addr, wif.String(), nil
}

func (s *BtcChain) GenHdAddr() (string, string, error) {
	mnemonic, _, err := cryptolib.NewMnemonic()
	if err != nil {
		return "", "", err
	}

	return s.GetAddrByMnemonic(mnemonic)
}

func (s *BtcChain) GetAddrByMnemonic(mnemonic string) (string, string, error) {
	path := _btcAddrTypePath[s.addrType]
	pubKey, _, err := cryptolib.MnemonicToBtcEcKey(mnemonic, path)
	if err != nil {
		return "", "", err
	}

	addr, err := _btcGenAddr(s.addrType, pubKey, s.cfg.NetParams)
	if err != nil {
		return "", "", err
	}

	return addr, mnemonic, nil
}

func (s *BtcChain) GetAddrBalance(addr string, cur model.CurrencyContract) (decimal.Decimal, error) {
	url := fmt.Sprintf("%s/addrs/%s/balance", s.cfg.URL, addr)
	resp, err := resty.New().R().Get(url)
	if err != nil {
		return decimal.Zero, err
	}

	if resp.StatusCode() != 200 {
		return decimal.Zero, eris.Errorf("status: %s", resp.Status())
	}

	jsonNode, err := sonic.Get(resp.Body(), "balance")
	if err != nil {
		return decimal.Zero, eris.Wrap(err, "failed to get blocks")
	}

	balanceSatoshi, err := jsonNode.Int64()
	if err != nil {
		return decimal.Zero, eris.Wrap(err, "failed to get blocksHigh")
	}

	return decimal.NewFromBigInt(big.NewInt(balanceSatoshi), cur.Decimal), nil
}
