package model

type Currency struct {
	Name   string
	Symbol string
	Remark string
	Sort   int8
	Chain  map[string]CurrencyContract
}

type CurrencyContract struct {
	Addr    string
	Name    string
	Decimal int32
	IsGov   bool
}
