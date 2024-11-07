package chain

import (
	"context"
	"crypto/ed25519"
	"errors"

	"github.com/blocto/solana-go-sdk/pkg/hdwallet"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/luzen23141/mouse/pkg/blockchain/model"
	cryptolib "github.com/luzen23141/mouse/pkg/lib/cyptolib"
	"github.com/shopspring/decimal"
	"github.com/tyler-smith/go-bip39"
	"github.com/xssnick/tonutils-go/address"
	"github.com/xssnick/tonutils-go/liteclient"
	"github.com/xssnick/tonutils-go/tlb"
	"github.com/xssnick/tonutils-go/ton"
	"github.com/xssnick/tonutils-go/ton/jetton"
	"github.com/xssnick/tonutils-go/ton/wallet"
)

type TonChain struct {
	tonVersion _tonAddrConfig
}

func NewTonChain() *TonChain {
	return &TonChain{
		tonVersion: _tonAddrConfigMap["v5"],
	}
}

// 跟btc一樣，有很多版本的地址跟變數，每個錢包使用的版本或變數不同
var _tonAddrConfigMap = map[string]_tonAddrConfig{
	// 2024-11-04 trustWallet 為v4
	//            bitgetWallet 為v5,v4 可切換
	//
	"v5": {
		versionCfg: wallet.ConfigV5R1Final{
			NetworkGlobalID: wallet.MainnetGlobalID,
		},
		path:      `m/44'/607'/0'`,
		subWallet: uint32(0),
	},
	"v4": {
		versionCfg: wallet.V4R2,
		path:       `m/44'/607'/0'`,
		subWallet:  wallet.DefaultSubwallet,
	},
	"v3": {
		versionCfg: wallet.V3R2,
		path:       `m/44'/607'/0'/0'`,
		subWallet:  wallet.DefaultSubwallet,
	},
}

type _tonAddrConfig struct {
	versionCfg wallet.VersionConfig
	subWallet  uint32
	path       string
}

// GenAddr 产生地址
func (s *TonChain) GenAddr() (string, string, error) {
	// 生成新的密钥对
	pubKey, privateKey, err := ed25519.GenerateKey(nil)
	if err != nil {
		return "", "", err
	}

	address, err := wallet.AddressFromPubKey(pubKey, s.tonVersion.versionCfg, s.tonVersion.subWallet)
	if err != nil {
		return "", "", err
	}

	// Bounce(false) => UQ 開頭，Bounce(true) => EQ 開頭
	return address.Bounce(false).String(), hexutil.Encode(privateKey.Seed()), nil
}

func (s *TonChain) GetAddrByPrivKey(privKeyStr string) (string, string, error) {
	pubKey, privateKey, err := cryptolib.StrKeyToEd25519KeyPair(privKeyStr)
	if err != nil {
		return "", "", err
	}
	address, err := wallet.AddressFromPubKey(pubKey, s.tonVersion.versionCfg, s.tonVersion.subWallet)
	if err != nil {
		return "", "", err
	}

	// Bounce(false) => UQ 開頭，Bounce(true) => EQ 開頭
	return address.Bounce(false).String(), hexutil.Encode(privateKey.Seed()), nil
}

// GenHdAddr 产生Hd wallet 地址，返回的key為mnemonic
func (s *TonChain) GenHdAddr() (string, string, error) {
	// 生成新的密钥对
	mnemonic, _, err := cryptolib.NewMnemonic()
	if err != nil {
		return "", "", err
	}

	return s.GetAddrByMnemonic(mnemonic)
}

func (s *TonChain) GetAddrByMnemonic(mnemonic string) (string, string, error) {
	derivedKey, err := hdwallet.Derived(s.tonVersion.path, bip39.NewSeed(mnemonic, ""))
	if err != nil {
		return "", "", err
	}

	priKey := ed25519.NewKeyFromSeed(derivedKey.PrivateKey)
	pubKey := priKey.Public().(ed25519.PublicKey)

	address, err := wallet.AddressFromPubKey(pubKey, s.tonVersion.versionCfg, s.tonVersion.subWallet)
	if err != nil {
		return "", "", err
	}

	// Bounce(false) => UQ 開頭，Bounce(true) => EQ 開頭
	return address.Bounce(false).String(), mnemonic, nil
}

func (s *TonChain) GetAddrBalance(addrStr string, cur model.CurrencyContract) (decimal.Decimal, error) {
	conn, ctx, err := s.getClient()
	if err != nil {
		return decimal.Zero, err
	}

	if cur.IsGov {
		// we need fresh block info to run get methods
		b, err := conn.CurrentMasterchainInfo(ctx)
		if err != nil {
			return decimal.Zero, err
		}

		// we use WaitForBlock to make sure block is ready,
		// it is optional but escapes us from liteserver block not ready errors
		addr := address.MustParseAddr(addrStr)
		res, err := conn.WaitForBlock(b.SeqNo).GetAccount(ctx, b, addr)
		if err != nil {
			return decimal.Zero, err
		}

		if !res.IsActive {
			return decimal.Zero, errors.New("account not active")
		}
		if res.State.Status != tlb.AccountStatusActive {
			return decimal.Zero, errors.New("account not active, active:" + string(res.State.Status))
		}

		return decimal.NewFromBigInt(res.State.Balance.Nano(), cur.Decimal), nil
	}

	// jetton contract address
	master := jetton.NewJettonMasterClient(conn, address.MustParseAddr(cur.Addr))

	// get jetton wallet for account
	ownerAddr := address.MustParseAddr(addrStr)
	jettonWallet, err := master.GetJettonWallet(context.Background(), ownerAddr)
	if err != nil {
		return decimal.Zero, err
	}

	jettonBalance, err := jettonWallet.GetBalance(context.Background())
	if err != nil {
		return decimal.Zero, err
	}

	return decimal.NewFromBigInt(jettonBalance, cur.Decimal), nil
}

func (*TonChain) getClient() (ton.APIClientWrapped, context.Context, error) {
	client := liteclient.NewConnectionPool()

	// connect to mainnet lite servers
	err := client.AddConnectionsFromConfigUrl(context.Background(), "https://ton.org/global.config.json")
	if err != nil {
		return nil, nil, err
	}

	// initialize ton api lite connection wrapper
	ctx := client.StickyContext(context.Background())
	return ton.NewAPIClient(client, ton.ProofCheckPolicyFast).WithRetry(), ctx, nil
}
