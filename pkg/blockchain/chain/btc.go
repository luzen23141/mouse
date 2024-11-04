package chain

import (
	"errors"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcec/v2/schnorr"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
)

type BtcChain struct {
	netParams *chaincfg.Params
}

func NewBtcChain() *BtcChain {
	return &BtcChain{
		netParams: &chaincfg.MainNetParams,
	}
}

func (s *BtcChain) GenAddr() (string, string, error) {
	privateKey, err := btcec.NewPrivateKey()
	if err != nil {
		return "", "", err
	}

	wif, err := btcutil.NewWIF(privateKey, s.netParams, true)
	if err != nil {
		return "", "", err
	}

	taprootAddr, err := btcutil.NewAddressTaproot(
		schnorr.SerializePubKey(
			txscript.ComputeTaprootKeyNoScript(
				wif.PrivKey.PubKey(),
			),
		),
		s.netParams)
	if err != nil {
		return "", "", err
	}

	return taprootAddr.EncodeAddress(), wif.String(), nil
}

func (s *BtcChain) GenHdAddr() (addr, mnemonic string, err error) {
	return "", "", errors.New("not support")
}
