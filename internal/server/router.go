package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter()*gin.Engine {
	r:=gin.Default()

	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"ok":true,
			"status":"healthy ✅",
		})
	})
	return r
}

// 4:30:00
