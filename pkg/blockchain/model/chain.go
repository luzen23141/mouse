package model

import "github.com/btcsuite/btcd/chaincfg"

type BtcCfg struct {
	URL       string
	IsTest    bool
	NetParams *chaincfg.Params
}
