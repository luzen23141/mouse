package routes

import (
	"mouse/pkg"
	"mouse/pkg/controller"
	"mouse/pkg/helper"

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
