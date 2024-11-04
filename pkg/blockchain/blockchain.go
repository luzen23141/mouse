package blockchain

import "mouse/pkg/blockchain/chain"

var ChainMap = map[string]ChainInterface{
	BtcChain:     chain.NewBtcChain(),
	TrxChain:     chain.NewTronChain(),
	EthChain:     chain.NewEvmChain(),
	BscChain:     chain.NewEvmChain(),
	PolygonChain: chain.NewEvmChain(),
	SuiChain:     chain.NewSuiChain(),
	InjChain:     chain.NewInjChain(),
	SolChain:     chain.NewSolChain(),
}

type ChainInterface interface {
	GenAddr() (string, string, error)
	GenHdAddr() (string, string, error)
}
