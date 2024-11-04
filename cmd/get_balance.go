package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"golang.org/x/exp/maps"
	"mouse/pkg/blockchain"
	"mouse/pkg/blockchain/model"
	"sort"
)

// getBalanceCmd represents the serve command
var getBalanceCmd = &cobra.Command{
	Use:   `get:balance`,
	Short: `獲取餘額`,
	Long:  `獲取餘額`,
	RunE:  getBalanceExec,
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
	var curChainCfg model.CurrencyContract
	if curCfg, ok := blockchain.CurMap[args[0]]; !ok {
		curList := maps.Values(blockchain.CurMap)
		sort.Slice(curList, func(i, j int) bool {
			return curList[i].Sort < curList[j].Sort
		})
		curList = curList[0:min(10, len(curList))]
		return fmt.Errorf("invalid argument %q for %q, valid args: %v ...etc", args[0], cmd.CommandPath(), curList)
	} else {
		if curChainCfg, ok = curCfg.Chain[args[1]]; !ok {
			return fmt.Errorf("invalid argument %q for %q, valid args: %v ...etc", args[1], cmd.CommandPath(), maps.Keys(curCfg.Chain))
		}
	}

	chainSer, ok := blockchain.ChainMap[args[1]]
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