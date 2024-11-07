package main

import (
	"testing"

	"github.com/luzen23141/mouse/pkg/blockchain"
	"github.com/luzen23141/mouse/pkg/blockchain/_const"
)

func TestGetAddressByPrivKey(t *testing.T) {
	// 定義測試案例
	testCases := []struct {
		chain   string
		name    string
		key     string
		address string
	}{
		{
			_const.BtcChain,
			"btc_base58",
			"L4ekbXpema8Cv1sPFibE2fa2aLwUi1hhp1iLzVMt7EvNskxMu1Jz",
			"19rB3Eym8Jpbk938uBLYifqcnjm8pMxrFF",
		},
		{
			_const.EthChain,
			"eth_hex",
			"0x933e850d639a1783f604c06e1727b69248029ce3411d90ae833f0e0341b13478",
			"0x3029894fcc847bbdc9ef25776a976949c6b982f7",
		},
		{
			_const.InjChain,
			"inj_hex",
			"0xcd2d220726d1d0a9b1df128edab1d5d00cb5237be506edc5cab1ab9536723a17",
			"inj1kn2gzlkakqukxzcee9ywe54rwywr92tf0wp4qx",
		},
		{
			_const.SolChain,
			"sol_base58",
			"uJUJMxc62f4BMqJcYrWX99sGdxUu5MGzxZqJyEgUdsNtSYPqpkK2vx3BLgs3w3FZ9n4xJZaqRQMzVgeiVEinMwK",
			"qx8tyknvYV6CCfctrcw9nZQA1maGPunZwkBi31RNm73",
		},
		{
			_const.SuiChain,
			"sui_hex",
			"0xd6c65df9fea3d38f3dda8d3952b80381ab97590d721dc29d50c620e1fb23ca73",
			"0x388560521cce9502ddc831f342df292bbcc8043c80277abc082934e5b80907e9",
		},
		{
			_const.TonChain,
			"ton_hex",
			"0x35e234f872db1396bdbb178ecb012947fae97f059119f2062cbe18a988a388bc",
			"UQDJAESRvJUIaa68BGKMNRFneOWPZwY293Vjy-SBphvL3bpa",
		},
		{
			_const.TronChain,
			"tron_hex",
			"0xf2b989b42bb350d97528e0ea24203e4e13c0cfa6d157f2f1b7d4395ed975cbfb",
			"TGGUbBbeZhGihsQ7yHpzZiQve15iDT2aGE",
		},
	}

	// 迴圈執行測試案例
	for _, tc := range testCases {
		chain := tc.chain
		c, ok := blockchain.ChainMap[chain]
		if !ok {
			t.Errorf("%s 不支援的鏈", tc.name)
			continue
		}

		var (
			addr string
			err  error
		)
		addr, _, err = c.GetAddrByPrivKey(tc.key)
		if err != nil {
			t.Errorf("%s convert err:%s", tc.name, err)
			continue
		}
		if addr != tc.address {
			t.Errorf("%s 輸入 %s, 預期 %s", tc.name, addr, tc.address)
			continue
		}
		t.Logf("%s success", tc.name)
	}
}

func TestGetAddress(t *testing.T) {
	// 定義測試案例
	testCases := []struct {
		input  string
		expect string
	}{
		{_const.BtcChain, "15Ha6WbhruHw7ZvBtny68QCUkWTUQnnFrT"},
		{_const.EthChain, "0xAF04F011956e79aeE7488A2C9f7fA3a8E9B30645"},
		{_const.InjChain, "inj14uz0qyv4deu6ae6g3gkf7lar4r5mxpj90ljqls"},
		{_const.TonChain, "UQBMvIookoGcfOF_6zzvQIYwWOKllV2y-SOm_mNZ5QAalChX"},
		{_const.SolChain, "G3CD82mbCdckXHEXLZepEYhpMioHiPXJhNQktSk7qKAz"},
		{_const.SuiChain, "0x61a6dc73a9d3d30a20f8049f3042f6e40dd23afa7d83b2cd8c3bc4a0f3c1a3c6"},
		{_const.TronChain, "TEonA4MhaAVEgXH3Qmhg3P42JDrDT2anTr"},
	}
	testKey := "outside harbor seed crumble ginger broccoli excite cloth post wait label snow " +
		"family humble gas toilet fit blur lecture connect end turn walnut craft"

	// 迴圈執行測試案例
	for _, tc := range testCases {
		chain := tc.input
		c, ok := blockchain.ChainMap[chain]
		if !ok {
			t.Errorf("%s 不支援的鏈", chain)
			continue
		}

		var (
			addr string
			err  error
		)
		addr, _, err = c.GetAddrByMnemonic(testKey)
		if err != nil {
			t.Errorf("%s convert err:%s", chain, err)
			continue
		}
		if addr != tc.expect {
			t.Errorf("%s 輸入 %s, 預期 %s", chain, addr, tc.expect)
			continue
		}
		t.Logf("%s success", chain)
	}
}
