package cryptolib

import (
	"crypto/ecdsa"
	"crypto/ed25519"
	"errors"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/rotisserie/eris"
)

func hexToEcdsaKeyPair(hexKey string) (*ecdsa.PublicKey, *ecdsa.PrivateKey, error) {
	privKey, err := hexToEcdsaPriv(hexKey)
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

func hexToEcdsaPriv(hexKey string) (*ecdsa.PrivateKey, error) {
	return crypto.HexToECDSA(strings.TrimPrefix(hexKey, "0x"))
}

func hexToEcdsaPub(hexKey string) (*ecdsa.PublicKey, error) {
	privKey, err := hexToEcdsaPriv(hexKey)
	if err != nil {
		return nil, err
	}
	publicKey := privKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, eris.New("failed to load public key")
	}

	return publicKeyECDSA, nil
}

func hexToEd25519KeyPair(hexKey string) (ed25519.PublicKey, ed25519.PrivateKey, error) {
	privKey, err := hexToEd25519Priv(hexKey)
	if err != nil {
		return nil, nil, err
	}
	pubKey, ok := privKey.Public().(ed25519.PublicKey)
	if !ok {
		return nil, nil, errors.New("failed to get public key")
	}
	return pubKey, privKey, nil
}

func hexToEd25519Pub(hexKey string) (ed25519.PublicKey, error) {
	privKey, err := hexToEd25519Priv(hexKey)
	if err != nil {
		return nil, err
	}
	pubKey, ok := privKey.Public().(ed25519.PublicKey)
	if !ok {
		return nil, errors.New("failed to get public key")
	}
	return pubKey, nil
}

func hexToEd25519Priv(hexKey string) (ed25519.PrivateKey, error) {
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
