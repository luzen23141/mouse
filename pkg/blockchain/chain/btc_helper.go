package chain

import (
	"errors"
	"github.com/btcsuite/btcd/btcec/v2/schnorr"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
)

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

func _btcGenAddr(addrType string, pubKey *secp256k1.PublicKey, net *chaincfg.Params) (string, error) {
	switch addrType {
	case _btcAddrLegacy:
		pubKeyAddr, err := btcutil.NewAddressPubKey(pubKey.SerializeCompressed(), net)
		if err != nil {
			return "", err
		}
		return pubKeyAddr.AddressPubKeyHash().EncodeAddress(), nil

	case _btcAddrNestedSegWit49:
		fallthrough
	case _btcAddrNestedSegWit44:
		// 產生 Witness program
		witnessProgram := btcutil.Hash160(pubKey.SerializeCompressed())

		// 將 Witness program 轉換為 btcutil.Address
		witnessAddr, err := btcutil.NewAddressWitnessPubKeyHash(witnessProgram, &chaincfg.MainNetParams)
		if err != nil {
			return "", err
		}

		// 使用 witnessAddr 產生 P2SH 地址
		script, err := txscript.PayToAddrScript(witnessAddr)
		if err != nil {
			return "", err
		}
		addr, err := btcutil.NewAddressScriptHash(script, &chaincfg.MainNetParams)
		if err != nil {
			return "", err
		}
		return addr.EncodeAddress(), nil

	case _btcAddrNativeSegWit84:
		fallthrough
	case _btcAddrNativeSegWit44:
		addr, err := btcutil.NewAddressWitnessPubKeyHash(
			btcutil.Hash160(pubKey.SerializeCompressed()), net,
		)
		if err != nil {
			return "", err
		}
		return addr.EncodeAddress(), nil

	case _btcAddrTaproot86:
		fallthrough
	case _btcAddrTaproot44:
		taprootAddr, err := btcutil.NewAddressTaproot(
			schnorr.SerializePubKey(
				txscript.ComputeTaprootKeyNoScript(pubKey),
			), net)
		if err != nil {
			return "", err
		}
		return taprootAddr.EncodeAddress(), nil

	default:
		return "", errors.New("invalid address type")
	}
}
