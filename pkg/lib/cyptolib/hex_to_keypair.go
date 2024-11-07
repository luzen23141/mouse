package cryptolib

import (
	"crypto/ecdsa"
	"crypto/ed25519"
	"errors"
	"strconv"
	"strings"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/mr-tron/base58"
	"github.com/rotisserie/eris"
)

func StrKeyToEcdsaKeyPair(privKeyStr string) (*ecdsa.PublicKey, *ecdsa.PrivateKey, error) {
	if strings.HasPrefix(privKeyStr, "0x") {
		return HexToEcdsaKeyPair(privKeyStr)
	} else {
		return Base58ToEcdsaKeyPair(privKeyStr)
	}
}

func StrKeyToEd25519KeyPair(privKeyStr string) (ed25519.PublicKey, ed25519.PrivateKey, error) {
	if strings.HasPrefix(privKeyStr, "0x") {
		return HexToEd25519KeyPair(privKeyStr)
	} else {
		return Base58ToEd25519KeyPair(privKeyStr)
	}
}

func HexToEcdsaKeyPair(hexKey string) (*ecdsa.PublicKey, *ecdsa.PrivateKey, error) {
	privKey, err := HexToEcdsaPriv(hexKey)
	if err != nil {
		return nil, nil, err
	}
	publicKey := privKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, nil, eris.New("failed to load public key")
	}

	return publicKeyECDSA, privKey, nil
}

func HexToEcdsaPriv(hexKey string) (*ecdsa.PrivateKey, error) {
	return crypto.HexToECDSA(strings.TrimPrefix(hexKey, "0x"))
}

func HexToEd25519KeyPair(hexKey string) (ed25519.PublicKey, ed25519.PrivateKey, error) {
	privKey, err := HexToEd25519Priv(hexKey)
	if err != nil {
		return nil, nil, err
	}
	pubKey, ok := privKey.Public().(ed25519.PublicKey)
	if !ok {
		return nil, nil, errors.New("failed to get public key")
	}
	return pubKey, privKey, nil
}

func HexToEd25519Priv(hexKey string) (ed25519.PrivateKey, error) {
	seed, err := hexutil.Decode(hexKey)
	if err != nil {
		return nil, err
	}
	privateKey := ed25519.NewKeyFromSeed(seed)
	if privateKey == nil {
		return nil, errors.New("hexKey seed conversion failed")
	}

	return privateKey, nil
}

func Base58ToEcdsaKeyPair(base58Key string) (*ecdsa.PublicKey, *ecdsa.PrivateKey, error) {
	wif, err := btcutil.DecodeWIF(base58Key)
	if err != nil {
		return nil, nil, err
	}
	return wif.PrivKey.PubKey().ToECDSA(), wif.PrivKey.ToECDSA(), nil
}

func Base58ToEcdsaPriv(base58Key string) (*ecdsa.PrivateKey, error) {
	wif, err := btcutil.DecodeWIF(base58Key)
	if err != nil {
		return nil, err
	}
	return wif.PrivKey.ToECDSA(), nil
}

func Base58ToEd25519KeyPair(base58Key string) (ed25519.PublicKey, ed25519.PrivateKey, error) {
	privateKey, err := Base58ToEd25519Priv(base58Key)
	if err != nil {
		return nil, nil, err
	}
	pubKey, ok := privateKey.Public().(ed25519.PublicKey)
	if !ok {
		return nil, nil, errors.New("failed to get public key")
	}
	return pubKey, privateKey, nil
}

func Base58ToEd25519Priv(base58Key string) (ed25519.PrivateKey, error) {
	seed, err := base58.Decode(base58Key)
	if err != nil {
		return nil, err
	}
	if len(seed) == 64 {
		seed = seed[:32]
	}
	if len(seed) != ed25519.SeedSize {
		return nil, errors.New("ed25519: bad seed length: " + strconv.Itoa(len(seed)))
	}
	return ed25519.NewKeyFromSeed(seed[:32]), nil
}
