package blockchain

import "mouse/pkg/blockchain/model"

var CurMap = map[string]model.Currency{
	BTC: {
		Name:   "Bitcoin",
		Symbol: "BTC",
		Remark: "比特幣",
		Sort:   1,
		Chain: map[string]model.CurrencyContract{
			BtcChain: {
				Addr:    "",
				Name:    "Bitcoin",
				Decimal: -8,
				IsGov:   true,
			},
		},
	},
	ETH: {
		Name:   "Ethereum",
		Symbol: "ETH",
		Remark: "以太坊",
		Sort:   2,
		Chain: map[string]model.CurrencyContract{
			EthChain: {
				Addr:    "",
				Name:    "Ethereum",
				Decimal: -18,
				IsGov:   true,
			},
		},
	},
	USDT: {
		Name:   "Tether",
		Symbol: "USDT",
		Remark: "USDT",
		Sort:   3,
		Chain: map[string]model.CurrencyContract{
			EthChain: {
				Addr:    "0xdAC17F958D2ee523a2206206994597C13D831ec7",
				Name:    "Tether",
				Decimal: -6,
				IsGov:   false,
			},
			TronChain: {
				Addr:    "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t",
				Name:    "Tether",
				Decimal: -6,
				IsGov:   false,
			},
		},
	},
}
