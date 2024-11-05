package chain

import (
	"fmt"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/bytedance/sonic"
	"github.com/go-resty/resty/v2"
	"github.com/luzen23141/mouse/pkg/blockchain/model"
	cryptolib "github.com/luzen23141/mouse/pkg/lib/cyptolib"
	"github.com/rotisserie/eris"
	"github.com/shopspring/decimal"
	"math/big"
	"strings"
)

type BtcChain struct {
	netParams *chaincfg.Params
	addrType  string
	cfg       model.BtcCfg
}

func NewBtcChain(cfg model.BtcCfg) *BtcChain {
	cfg.Url = strings.TrimSuffix(cfg.Url, "/")

	if cfg.IsTest {
		return &BtcChain{
			netParams: &chaincfg.TestNet3Params,
			addrType:  _btcAddrLegacy,
			cfg:       cfg,
		}
	}

	return &BtcChain{
		netParams: &chaincfg.MainNetParams,
		addrType:  _btcAddrLegacy,
		cfg:       cfg,
	}
}

func (s *BtcChain) GenAddr() (string, string, error) {
	privKeyEc, err := btcec.NewPrivateKey()
	if err != nil {
		return "", "", err
	}
	wif, err := btcutil.NewWIF(privKeyEc, s.netParams, true)
	if err != nil {
		return "", "", err
	}
	//wifStr := "L4ekbXpema8Cv1sPFibE2fa2aLwUi1hhp1iLzVMt7EvNskxMu1Jz"
	//wif, err := btcutil.DecodeWIF(wifStr)
	//if err != nil {
	//	return "", "", err
	//}
	privKey := wif.PrivKey
	pubKey := privKey.PubKey()

	addr, err := _btcGenAddr(s.addrType, pubKey, s.netParams)
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
	//mnemonic := "outside harbor seed crumble ginger broccoli excite cloth post wait label snow family humble gas toilet fit blur lecture connect end turn walnut craft"

	path := _btcAddrTypePath[s.addrType]
	pubKey, _, err := cryptolib.MnemonicToBtcEcKey(mnemonic, path)
	if err != nil {
		return "", "", err
	}

	addr, err := _btcGenAddr(s.addrType, pubKey, s.netParams)
	if err != nil {
		return "", "", err
	}

	return addr, mnemonic, nil
}

func (s *BtcChain) GetAddrBalance(addr string, cur model.CurrencyContract) (decimal.Decimal, error) {
	url := fmt.Sprintf("%s/v1/btc/main/addrs/%s/balance", s.cfg.Url, addr)
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
