package chain

import (
	"errors"
	"github.com/btcsuite/btcd/btcec/v2"
	cryptolib "mouse/pkg/lib/cyptolib"

	"github.com/btcsuite/btcd/btcec/v2/schnorr"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
)

type BtcChain struct {
	netParams *chaincfg.Params
	addrType  string
}

const (
	_btcAddrLegacy = "Legacy"

	_btcAddrNestedSegWit49 = "Nested SegWit (m/49)"
	_btcAddrNestedSegWit44 = "Nested SegWit (m/44)"

	_btcAddrNativeSegWit84 = "Native SegWit (m/84)"
	_btcAddrNativeSegWit44 = "Native SegWit (m/44)"

	_btcAddrTaproot86 = "Taproot (m/86)"
	_btcAddrTaproot44 = "Taproot (m/44)"
)

var _btcAddrTypePath = map[string]string{
	_btcAddrLegacy: "m/44'/0'/0'/0/0",

	_btcAddrNestedSegWit49: "m/49'/0'/0'/0/0",
	_btcAddrNestedSegWit44: "m/44'/0'/0'/0/0",

	_btcAddrNativeSegWit84: "m/84'/0'/0'/0/0",
	_btcAddrNativeSegWit44: "m/44'/0'/0'/0/0",

	_btcAddrTaproot86: "m/86'/0'/0'/0/0",
	_btcAddrTaproot44: "m/44'/0'/0'/0/0",
}

func NewBtcChain() *BtcChain {
	return &BtcChain{
		netParams: &chaincfg.MainNetParams,
		addrType:  _btcAddrLegacy,
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

	switch s.addrType {
	case _btcAddrLegacy:
		pubKeyAddr, err := btcutil.NewAddressPubKey(pubKey.SerializeCompressed(), s.netParams)
		if err != nil {
			return "", "", err
		}
		return pubKeyAddr.AddressPubKeyHash().EncodeAddress(), wif.String(), nil

	case _btcAddrNestedSegWit49:
		fallthrough
	case _btcAddrNestedSegWit44:
		// 產生 Witness program
		witnessProgram := btcutil.Hash160(pubKey.SerializeCompressed())

		// 將 Witness program 轉換為 btcutil.Address
		witnessAddr, err := btcutil.NewAddressWitnessPubKeyHash(witnessProgram, &chaincfg.MainNetParams)
		if err != nil {
			return "", "", err
		}

		// 使用 witnessAddr 產生 P2SH 地址
		script, err := txscript.PayToAddrScript(witnessAddr)
		if err != nil {
			return "", "", err
		}
		addr, err := btcutil.NewAddressScriptHash(script, &chaincfg.MainNetParams)
		if err != nil {
			return "", "", err
		}
		return addr.EncodeAddress(), wif.String(), nil

	case _btcAddrNativeSegWit84:
		fallthrough
	case _btcAddrNativeSegWit44:
		addr, err := btcutil.NewAddressWitnessPubKeyHash(
			btcutil.Hash160(pubKey.SerializeCompressed()),
			s.netParams,
		)
		if err != nil {
			return "", "", err
		}
		return addr.EncodeAddress(), wif.String(), nil

	case _btcAddrTaproot86:
		fallthrough
	case _btcAddrTaproot44:
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

	default:
		return "", "", errors.New("invalid address type")
	}
}

func (s *BtcChain) GenHdAddr() (string, string, error) {
	mnemonic, _, err := cryptolib.NewMnemonic()
	if err != nil {
		return "", "", err
	}
	//mnemonic := "outside harbor seed crumble ginger broccoli excite cloth post wait label snow family humble gas toilet fit blur lecture connect end turn walnut craft"

	path := _btcAddrTypePath[s.addrType]
	pubKey, privKey, err := cryptolib.MnemonicToBtcEcKey(mnemonic, path)
	if err != nil {
		return "", "", err
	}

	switch s.addrType {
	case _btcAddrLegacy:
		pubKeyAddr, err := btcutil.NewAddressPubKey(pubKey.SerializeCompressed(), s.netParams)
		if err != nil {
			return "", "", err
		}
		return pubKeyAddr.AddressPubKeyHash().EncodeAddress(), mnemonic, nil

	case _btcAddrNestedSegWit49:
		fallthrough
	case _btcAddrNestedSegWit44:
		// 產生 Witness program
		witnessProgram := btcutil.Hash160(pubKey.SerializeCompressed())

		// 將 Witness program 轉換為 btcutil.Address
		witnessAddr, err := btcutil.NewAddressWitnessPubKeyHash(witnessProgram, &chaincfg.MainNetParams)
		if err != nil {
			return "", "", err
		}

		// 使用 witnessAddr 產生 P2SH 地址
		script, err := txscript.PayToAddrScript(witnessAddr)
		if err != nil {
			return "", "", err
		}
		addr, err := btcutil.NewAddressScriptHash(script, &chaincfg.MainNetParams)
		if err != nil {
			return "", "", err
		}
		return addr.EncodeAddress(), mnemonic, nil

	case _btcAddrNativeSegWit84:
		fallthrough
	case _btcAddrNativeSegWit44:
		addr, err := btcutil.NewAddressWitnessPubKeyHash(
			btcutil.Hash160(pubKey.SerializeCompressed()),
			s.netParams,
		)
		if err != nil {
			return "", "", err
		}
		return addr.EncodeAddress(), mnemonic, nil

	case _btcAddrTaproot86:
		fallthrough
	case _btcAddrTaproot44:
		wif, err := btcutil.NewWIF(privKey, s.netParams, true)
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
		return taprootAddr.EncodeAddress(), mnemonic, nil

	default:
		return "", "", errors.New("invalid address type")
	}
}
