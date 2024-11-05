package cryptolib

import (
	"crypto/ecdsa"
	"errors"

	"github.com/btcsuite/btcd/btcec/v2"

	"github.com/btcsuite/btcd/btcutil/hdkeychain"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tyler-smith/go-bip39"
)

func NewMnemonic() (string, *ecdsa.PrivateKey, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return "", nil, err
	}
	mnemonic, err := bip39.NewMnemonic(crypto.FromECDSA(privateKey))
	if err != nil {
		return "", nil, err
	}

	return mnemonic, privateKey, nil
}

func MnemonicToBtcEcKey(mnemonic string, pathStr string) (*btcec.PublicKey, *btcec.PrivateKey, error) {
	if mnemonic == "" {
		return nil, nil, errors.New("mnemonic is required")
	}
	seed, err := bip39.NewSeedWithErrorChecking(mnemonic, "")
	if err != nil {
		return nil, nil, err
	}

	key, err := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
	if err != nil {
		return nil, nil, err
	}

	path, err := accounts.ParseDerivationPath(pathStr)
	if err != nil {
		return nil, nil, err
	}

	for _, n := range path {
		key, err = key.Derive(n)
		if err != nil {
			return nil, nil, err
		}
	}

	privateKey, err := key.ECPrivKey()
	if err != nil {
		return nil, nil, err
	}

	return privateKey.PubKey(), privateKey, nil
}

func MnemonicToEcdsaPubKey(mnemonic string, pathStr string) (*ecdsa.PublicKey, *ecdsa.PrivateKey, error) {
	if mnemonic == "" {
		return nil, nil, errors.New("mnemonic is required")
	}
	seed, err := bip39.NewSeedWithErrorChecking(mnemonic, "")
	if err != nil {
		return nil, nil, err
	}

	key, err := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
	if err != nil {
		return nil, nil, err
	}

	path, err := accounts.ParseDerivationPath(pathStr)
	if err != nil {
		return nil, nil, err
	}

	for _, n := range path {
		key, err = key.Derive(n)
		if err != nil {
			return nil, nil, err
		}
	}

	privateKey, err := key.ECPrivKey()
	if err != nil {
		return nil, nil, err
	}

	privateKeyECDSA := privateKey.ToECDSA()
	publicKey := privateKeyECDSA.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, nil, errors.New("failed to get public key")
	}

	return publicKeyECDSA, privateKeyECDSA, nil
}
