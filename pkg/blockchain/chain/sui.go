package chain

import (
	"context"
	"crypto/ed25519"
	"encoding/hex"
	"github.com/block-vision/sui-go-sdk/constant"
	"github.com/block-vision/sui-go-sdk/models"
	"github.com/block-vision/sui-go-sdk/sui"
	"math/big"
	"strconv"

	"github.com/luzen23141/mouse/pkg/blockchain/model"
	cryptolib "github.com/luzen23141/mouse/pkg/lib/cyptolib"
	"github.com/shopspring/decimal"

	"github.com/block-vision/sui-go-sdk/common/keypair"
	"github.com/block-vision/sui-go-sdk/signer"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/rotisserie/eris"
	"golang.org/x/crypto/blake2b"
)

type SuiChain struct{}

func NewSuiChain() *SuiChain {
	return &SuiChain{}
}

// GenAddr 产生地址
func (s *SuiChain) GenAddr() (string, string, error) {
	// 生成新的密钥对
	pubKey, privateKey, err := ed25519.GenerateKey(nil)
	if err != nil {
		return "", "", err
	}

	tmp := []byte{byte(keypair.Ed25519Flag)}
	tmp = append(tmp, pubKey...)
	addrBytes := blake2b.Sum256(tmp)
	addr := "0x" + hex.EncodeToString(addrBytes[:])[:signer.AddressLength]

	return addr, hexutil.Encode(privateKey.Seed()), nil
}

// GenHdAddr 产生Hd wallet 地址，返回的key為mnemonic
func (s *SuiChain) GenHdAddr() (string, string, error) {
	// 生成新的密钥对
	mnemonic, _, err := cryptolib.NewMnemonic()
	if err != nil {
		return "", "", err
	}

	suiSigner, err := signer.NewSignertWithMnemonic(mnemonic)
	if err != nil {
		return "", "", eris.Wrap(err, "failed to generate signer")
	}

	return suiSigner.Address, mnemonic, nil
}

func (s *SuiChain) GetAddrBalance(addr string, cur model.CurrencyContract) (decimal.Decimal, error) {
	conn, err := s.getClient()
	if err != nil {
		return decimal.Zero, err
	}

	rsp, err := conn.SuiXGetBalance(context.Background(), models.SuiXGetBalanceRequest{
		Owner:    addr,
		CoinType: cur.Addr,
	})
	if err != nil {
		return decimal.Zero, err
	}

	balanceInt, err := strconv.Atoi(rsp.TotalBalance)
	if err != nil {
		return decimal.Zero, err
	}
	return decimal.NewFromBigInt(big.NewInt(int64(balanceInt)), cur.Decimal), nil
}

func (*SuiChain) getClient() (sui.ISuiAPI, error) {
	return sui.NewSuiClient(constant.SuiMainnetEndpoint), nil
}
