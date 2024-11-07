package cmd

import (
	"errors"
	"fmt"

	"github.com/luzen23141/mouse/pkg/blockchain"
	"github.com/luzen23141/mouse/pkg/blockchain/model"
	"github.com/spf13/cobra"
	"golang.org/x/exp/maps"
)

// getBalanceCmd represents the serve command
var getBalanceCmd = &cobra.Command{
	Use:     `get:balance {currency} {chain} {addr}`,
	Aliases: []string{"gb", "get:b", "g:balance"},
	Short:   `獲取餘額`,
	Long:    `獲取餘額`,
	GroupID: `get`,
	RunE:    getBalanceExec,
	Args: func(cmd *cobra.Command, args []string) error {
		needArgs := 3
		if len(args) != needArgs {
			return fmt.Errorf("accepts %d arg(s), received %d", needArgs, len(args))
		}
		return nil
	},
}

func getBalanceCmdInit(cmd *cobra.Command) {
	cmd.AddCommand(getBalanceCmd)
}

func getBalanceExec(cmd *cobra.Command, args []string) error {
	cmd.SilenceUsage = true // 是否要打印指令的說明，如果是參數帶錯才要，如果是運行錯誤的不要

	var curChainCfg model.CurrencyContract
	if curCfg, ok := blockchain.CurMap[args[0]]; !ok {
		curList := maps.Keys(blockchain.CurMap)
		curList = curList[0:min(10, len(curList))]
		return fmt.Errorf("invalid argument %q for %q, valid args: %v ...etc", args[0], cmd.CommandPath(), curList)
	} else {
		if curChainCfg, ok = curCfg.Chain[args[1]]; !ok {
			msgF := "invalid argument %q for %q, valid args: %v ...etc"
			return fmt.Errorf(msgF, args[1], cmd.CommandPath(), maps.Keys(curCfg.Chain))
		}
	}

	chainSer, ok := blockchain.ChainMap[curChainCfg.Chain]
	if !ok {
		return errors.New("%s 鏈未支援")
	}

	balance, err := chainSer.GetAddrBalance(args[2], curChainCfg)
	if err != nil {
		return err
	}

	fmt.Println("balance:", balance)
	return nil
}
