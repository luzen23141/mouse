package blockchain

import (
	"github.com/luzen23141/mouse/pkg/blockchain/_const"
	"github.com/luzen23141/mouse/pkg/blockchain/model"
)

var CurMap = map[string]model.Currency{
	_const.BTC: {
		Name:   "Bitcoin",
		Symbol: "BTC",
		Remark: "比特幣",
		Sort:   1,
		Chain: map[string]model.CurrencyContract{
			_const.BtcChain: {
				Chain:   _const.BtcChain,
				Addr:    "",
				Name:    "Bitcoin",
				Decimal: -8,
				IsGov:   true,
			},
		},
	},
	_const.ETH: {
		Name:   "Ethereum",
		Symbol: "ETH",
		Remark: "以太坊",
		Sort:   2,
		Chain: map[string]model.CurrencyContract{
			_const.EthChain: {
				Chain:   _const.EthChain,
				Addr:    "",
				Name:    "Ethereum",
				Decimal: -18,
				IsGov:   true,
			},
		},
	},
	_const.USDT: {
		Name:   "Tether",
		Symbol: "USDT",
		Remark: "USDT",
		Sort:   3,
		Chain: map[string]model.CurrencyContract{
			"erc20": {
				Chain:   _const.EthChain,
				Addr:    "0xdAC17F958D2ee523a2206206994597C13D831ec7",
				Name:    "Tether",
				Decimal: -6,
				IsGov:   false,
			},
			"trc20": {
				Chain:   _const.TronChain,
				Addr:    "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t",
				Name:    "Tether",
				Decimal: -6,
				IsGov:   false,
			},
			"sol": {
				Chain:   _const.SolChain,
				Addr:    "Es9vMFrzaCERmJfrF4H2FYD4KCoNkY11McCe8BenwNYB",
				Name:    "Tether",
				Decimal: -6,
				IsGov:   false,
			},
			"sui": {
				Chain:   _const.SuiChain,
				Addr:    "0xc060006111016b8a020ad5b33834984a437aaa7d3c74c18e09a95d48aceab08c::coin::COIN",
				Name:    "Tether",
				Decimal: -6,
				IsGov:   false,
			},
		},
	},
	_const.TRX: {
		Name:   "Tron",
		Symbol: "TRX",
		Remark: "Tron",
		Sort:   4,
		Chain: map[string]model.CurrencyContract{
			_const.TronChain: {
				Chain:   _const.TronChain,
				Addr:    "",
				Name:    "Tron",
				Decimal: -6,
				IsGov:   true,
			},
		},
	},
	_const.SOL: {
		Name:   "Solana",
		Symbol: "SOL",
		Remark: "Solana",
		Sort:   5,
		Chain: map[string]model.CurrencyContract{
			_const.SolChain: {
				Chain:   _const.SolChain,
				Addr:    "",
				Name:    "Solana",
				Decimal: -9,
				IsGov:   true,
			},
		},
	},
	_const.SUI: {
		Name:   "Sui",
		Symbol: "SUI",
		Remark: "Sui",
		Sort:   6,
		Chain: map[string]model.CurrencyContract{
			_const.SuiChain: {
				Chain:   _const.SuiChain,
				Addr:    "0x2::sui::SUI",
				Name:    "Sui",
				Decimal: -9,
				IsGov:   true,
			},
		},
	},
}
