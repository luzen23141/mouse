package chain

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"
	"mouse/pkg/lib/cyptolib"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/mr-tron/base58"
)

type TronChain struct{}

func NewTronChain() *TronChain {
	return &TronChain{}
}

// GenAddr 产生TRON地址
func (s *TronChain) GenAddr() (addr, key string, err error) {
	privateKey, _ := crypto.GenerateKey()
	privateKeyBytes := crypto.FromECDSA(privateKey)
	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	addrStr := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	addrStr = "41" + addrStr[2:]
	addrByte, _ := hex.DecodeString(addrStr)
	firstHash := sha256.Sum256(addrByte)
	secondHash := sha256.Sum256(firstHash[:])
	secret := secondHash[:4]
	addrByte = append(addrByte, secret...)
	return base58.Encode(addrByte), hexutil.Encode(privateKeyBytes)[2:], nil
}

func (s *TronChain) GenHdAddr() (string, string, error) {
	mnemonic, _, err := cryptolib.NewMnemonic()
	if err != nil {
		return "", "", err
	}

	publicKeyECDSA, _, err := cryptolib.MnemonicToEcdsaPubKey(mnemonic, "m/44'/195'/0'/0/0")
	if err != nil {
		return "", "", err
	}

	return addrConvEthToTron(crypto.PubkeyToAddress(*publicKeyECDSA).Hex()), mnemonic, nil
}

func addrConvEthToTron(addr string) string {
	addrStr := "41" + addr[2:]
	addrByte, _ := hex.DecodeString(addrStr)
	firstHash := sha256.Sum256(addrByte)
	secondHash := sha256.Sum256(firstHash[:])
	secret := secondHash[:4]
	addrByte = append(addrByte, secret...)
	return base58.Encode(addrByte)
}
