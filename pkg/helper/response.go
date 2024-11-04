package helper

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// APIResponse 一般的回覆狀態
type APIResponse struct {
	Sign    int64       `json:"sign"`
	Code    uint32      `json:"code"`
	Message string      `json:"msg"`
	Result  interface{} `json:"data"`
}

// SuccessResponse 成功的回覆
type SuccessResponse struct {
	Sign   int64       `json:"sign"`
	Result interface{} `json:"data"`
}

func Success(g *gin.Context, result interface{}) {
	res := SuccessResponse{
		Sign:   time.Now().UnixMilli(),
		Result: result,
	}
	if res.Result == nil {
		res.Result = ""
	}

	g.JSON(http.StatusOK, &res)
}

func Fail(g *gin.Context, errorCode uint32, message string, result ...interface{}) {
	res := APIResponse{
		Sign:    time.Now().UnixMilli(),
		Code:    errorCode,
		Message: message,
		Result:  result,
	}

	g.JSON(http.StatusBadRequest, res)
}

func StatusCode(g *gin.Context, statusCode int) {
	g.Status(statusCode)
}
