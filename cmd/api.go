package cmd

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/luzen23141/mouse/pkg"
	"github.com/luzen23141/mouse/pkg/helper"
	"github.com/luzen23141/mouse/pkg/lib/log"
	"github.com/luzen23141/mouse/pkg/lib/validate"
	"github.com/luzen23141/mouse/pkg/middleware"
	"github.com/luzen23141/mouse/pkg/routes"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/spf13/cobra"
)

// apiCmd represents the serve command
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: `api伺服器`,
	Long:  `api伺服器`,
	RunE:  apiExec,
	// SilenceUsage: true,
}

func apiCmdInit(cmd *cobra.Command) {
	cmd.AddCommand(apiCmd)
}

func apiExec(cmd *cobra.Command, args []string) error {
	fmt.Println("api server starting...")
	log.InfoF(log.FileMainInfo,
		"api server starting..., GitVersion:%s,BuildDate:%s,GoVersion:%s",
		pkg.Version, pkg.BuildDate, pkg.GoVersion)

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
		Addr:         ":" + strconv.Itoa(apiPort),
		Handler:      g,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
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
