package controller

import (
	"mouse/pkg/blockchain"
	"mouse/pkg/helper"
	"mouse/pkg/model/modelapi"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/maps"
)

var CfgSupport = &_cfgSupport{}

type _cfgSupport struct{}

func (*_cfgSupport) GetSupportChain(g *gin.Context) {
	res := modelapi.CfgSupportRes{
		Main: maps.Keys(blockchain.ChainMap),
		Test: []string{},
	}

	helper.Success(g, res)
}
