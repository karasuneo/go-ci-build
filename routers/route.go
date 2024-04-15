package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DefineRoutes(r gin.IRouter) {
	r.GET("/", func(ctx *gin.Context) {
		// HTTPステータス200,レスポンスのdataに"Hello World"を返却
		ctx.JSONP(http.StatusOK, gin.H{
			"message": "ok",
			"data":    "Hello World",
		})
	})
}
