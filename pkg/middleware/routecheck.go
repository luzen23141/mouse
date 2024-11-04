package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// URLExists 檢查是否輸入不存在url/uri
func URLExists() gin.HandlerFunc {
	return func(g *gin.Context) {
		// 404 即是頁面不存在
		if g.Writer.Status() == 404 {
			g.AbortWithStatus(http.StatusNotFound)
			return
		}
	}
}
