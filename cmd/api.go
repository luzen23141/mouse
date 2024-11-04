package cmd

import (
	"fmt"
	"mouse/pkg"
	"mouse/pkg/helper"
	"mouse/pkg/lib/log"
	"mouse/pkg/lib/validate"
	"mouse/pkg/middleware"
	"mouse/pkg/routes"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/spf13/cobra"
)

// apiCmd represents the serve command
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: `api伺服器`,
	Long:  `api伺服器`,
	//Run: func(cmd *cobra.Command, args []string) {
	//	fmt.Println("serve called")
	//},
	RunE: apiExec,
	// SilenceUsage: true,
}

func apiCmdInit(cmd *cobra.Command) {
	cmd.AddCommand(apiCmd)
}

func apiExec(cmd *cobra.Command, args []string) error {
	fmt.Println("api server starting...")
	log.InfoF(log.FileMainInfo, "api server starting..., GitVersion:%s,BuildDate:%s,GoVersion:%s", pkg.Version, pkg.BuildDate, pkg.GoVersion)

	ginMode := helper.EnvCfg.Gin.Debug
	if ginMode {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// 创建一个 Gin 实例
	g := gin.New()
	g.Use(gin.Logger(), gin.Recovery(), middleware.URLExists())
	binding.Validator = new(validate.DefaultValidator) // 設定validator 版本
	routes.Routing(g)                                  // 建立route

	// 運行 server
	apiPort := helper.EnvCfg.Gin.Port
	server := &http.Server{
		Addr:    ":" + strconv.Itoa(apiPort),
		Handler: g,
	}

	// 启动服务器
	log.InfoF(log.FileMainInfo, "api server start, port:%d", apiPort)
	if err := server.ListenAndServe(); err != nil {
		cmd.SilenceUsage = true // 是否要打印指令的說明，如果是參數帶錯才要，如果是運行錯誤的不要
		fmt.Printf("Server in fatal. %v\n", err)
		return err
	}

	return nil
}
