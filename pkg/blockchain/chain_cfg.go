package blockchain

import (
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/luzen23141/mouse/pkg/blockchain/model"
)

var BtcCfg = model.BtcCfg{
	URL:       "https://api.blockcypher.com",
	IsTest:    false,
	NetParams: &chaincfg.MainNetParams,
}
