package blockchain

import (
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/luzen23141/mouse/pkg/blockchain/model"
)

var (
	BtcCfg = model.BtcCfg{
		URL:       "https://api.blockcypher.com/v1/btc/main",
		IsTest:    false,
		NetParams: &chaincfg.MainNetParams,
	}
	BtcTestCfg = model.BtcCfg{
		URL:       "https://api.blockcypher.com/v1/btc/test3",
		IsTest:    true,
		NetParams: &chaincfg.TestNet3Params,
	}
)
