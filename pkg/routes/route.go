package routes

import (
	"github.com/luzen23141/mouse/pkg"
	"github.com/luzen23141/mouse/pkg/controller"
	"github.com/luzen23141/mouse/pkg/helper"

	"github.com/gin-gonic/gin"
)

// Routing : 路由
func Routing(router *gin.Engine) {
	apiRouter := router.Group("api")

	apiRouter.GET("version", func(g *gin.Context) {
		helper.Success(g, gin.H{
			"version":   pkg.Version,
			"buildDate": pkg.BuildDate,
			"goVersion": pkg.GoVersion,
		})
	})

	configRouter := apiRouter.Group("config")
	{
		configRouter.GET("support/chains", controller.CfgSupport.GetSupportChain)
	}
}
