package cmd

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/luzen23141/mouse/pkg/blockchain"
	"github.com/spf13/cobra"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

// apiCmd represents the serve command
var genAddrCmd = &cobra.Command{
	Use:   `gen`,
	Short: `產地址`,
	Long:  `產地址`,
	RunE:  genAddrExec,
	Args:  genAddrValidArgs,
}

var _genAddrUsePriv = false

func genAddrValidArgs(cmd *cobra.Command, args []string) error {
	if len(args) < 1 || len(args) > 2 {
		return fmt.Errorf("accepts between 1 and 2 arg(s), received %d", len(args))
	}

	chainList := maps.Keys(blockchain.ChainMap)
	if !slices.Contains(chainList, args[0]) {
		return fmt.Errorf("invalid argument %q for %q, valid args: %v", args[0], cmd.CommandPath(), chainList)
	}

	return nil
}

func genAddrCmdInit(cmd *cobra.Command) {
	genAddrCmd.Flags().Uint8P("concurrent", "c", 1, "併發執行數量")
	genAddrCmd.Flags().BoolVar(&_genAddrUsePriv, "priv", false, "返回私鑰，預設為返回為助記詞")
	cmd.AddCommand(genAddrCmd)
}

func genAddrExec(cmd *cobra.Command, args []string) error {
	chain := args[0] // 已驗證過數量
	chainSer, ok := blockchain.ChainMap[chain]
	if !ok {
		return errors.New("%s 鏈沒有支援產生地址功能")
	}

	genAddrFunc := chainSer.GenHdAddr
	if _genAddrUsePriv {
		genAddrFunc = chainSer.GenAddr
	}

	postfix := ""
	if len(args) > 1 {
		postfix = args[1]
	}
	if postfix == "" {
		addr, mnemonic, err := genAddrFunc()
		if err != nil {
			return err
		}
		fmt.Printf("\nprivKey: \n%s\n", mnemonic)
		fmt.Printf("account %s: %s\n\n", chain, addr)
		return nil
	}
	postfix = strings.ToLower(postfix)

	times := int64(0)
	concurrentCount, _ := cmd.Flags().GetUint8("concurrent")
	if concurrentCount == 0 {
		concurrentCount = 1
	}

	fmt.Println("concurrentCount:", concurrentCount)
	startTime := time.Now().Unix()

	for i := uint8(0); i < concurrentCount; i++ {
		go func() {
			for {
				times++
				if times%100000 == 0 {
					nowTime := time.Now().Unix()
					execTime := nowTime - startTime

					fmt.Printf("times: %d，平均每秒執行次數: %d\n", times, times/execTime)
				}

				addr, mnemonic, err := genAddrFunc()
				if err != nil {
					fmt.Println(err)
					continue
				}

				// 檢查addr postfix
				if !strings.HasSuffix(strings.ToLower(addr), postfix) {
					continue
				}

				fmt.Printf("\nmnemonic: \n%s\n", mnemonic)
				fmt.Printf("account %s: %s\n\n", chain, addr)
			}
		}()
	}

	select {}
}
