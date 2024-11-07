package cmd

import (
	"os"

	"github.com/luzen23141/mouse/pkg"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   `mouse`,
	Short: `金灰鼠`,
	Long:  `金灰鼠的後端`,
	// Version: pkg.BuildDate + " " + pkg.GoVersion + " " + pkg.Version,
	Version: "編譯時間：" + pkg.BuildDate + " ,編譯Go版本：" + pkg.GoVersion + " ,git版本：" + pkg.Version,
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	apiCmdInit(rootCmd)
	genAddrCmdInit(rootCmd)
	getBalanceCmdInit(rootCmd)

	rootCmd.SetVersionTemplate(`{{printf "%s，%s\n" .Long .Version}}`)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
