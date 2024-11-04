package chain

import (
	cryptolib "mouse/pkg/lib/cyptolib"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcec/v2/schnorr"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
)

type BtcChain struct {
	netParams *chaincfg.Params
	addrType  string
}

const _btcAddrLegacy = "legacy"

var _btcAddrTypePath = map[string]string{
	_btcAddrLegacy: "m/44'/0'/0'/0/0",
}

func NewBtcChain() *BtcChain {
	return &BtcChain{
		netParams: &chaincfg.MainNetParams,
		addrType:  _btcAddrLegacy,
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

	serializedKey := pubKey.SerializeCompressed()
	pubKeyAddr, err := btcutil.NewAddressPubKey(serializedKey, s.netParams)
	if err != nil {
		return "", "", err
	}
	addr := pubKeyAddr.AddressPubKeyHash()

	return addr.EncodeAddress(), mnemonic, nil
}
