package cmd

import (
	"fmt"
	"strings"

	"github.com/luzen23141/mouse/pkg/blockchain"
	"github.com/spf13/cobra"
)

// getAddressCmd represents the serve command
var getAddressCmd = &cobra.Command{
	Use:     `get:address {mnemonic|privKey}`,
	Aliases: []string{"ga", "get:a", "g:address"},
	Short:   `輸入私鑰或助記詞，返回對應的地址`,
	Long:    `輸入私鑰或助記詞，返回對應的地址`,
	GroupID: `get`,
	RunE:    getAddressExec,
}

func getAddressCmdInit(cmd *cobra.Command) {
	cmd.AddCommand(getAddressCmd)
}

func getAddressExec(cmd *cobra.Command, args []string) error {
	cmd.SilenceUsage = true // 是否要打印指令的說明，如果是參數帶錯才要，如果是運行錯誤的不要

	isPriv := false
	key := ""
	switch len(args) {
	case 1:
		isPriv = true
		key = args[0]
		fmt.Println("這是私鑰")
	case 12:
		fallthrough
	case 24:
		key = strings.Join(args, " ")
		fmt.Println("這是助記詞")
	default:
		fmt.Println("私鑰或助記詞格式錯誤")
	}

	for chain, c := range blockchain.ChainMap {
		if isPriv {
			addr, _, err := c.GetAddrByPrivKey(key)
			if err != nil {
				fmt.Printf("%s convert err:%s\n", chain, err)
				continue
			}
			fmt.Printf("%s 地址: %s\n", chain, addr)
		} else {
			addr, _, err := c.GetAddrByMnemonic(key)
			if err != nil {
				fmt.Printf("%s convert err:%s\n", chain, err)
				continue
			}
			fmt.Printf("%s: %s\n", chain, addr)
		}
	}

	return nil
}
