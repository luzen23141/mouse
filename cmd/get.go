package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// getAddressCmd represents the serve command
var getCmd = &cobra.Command{
	Use:     `get`,
	Aliases: []string{"g"},
	Short:   `獲取特定內容`,
	Long:    `獲取特定內容`,
	Hidden:  true,
	Run: func(cmd *cobra.Command, args []string) {
		for _, c := range cmd.Root().Commands() {
			if c.GroupID == "get" {
				fmt.Println(" -", c.Short)
				fmt.Println("      ", c.Use)
			}
		}
	},
}

func getCmdInit(cmd *cobra.Command) {
	cmd.AddCommand(getCmd)
}
