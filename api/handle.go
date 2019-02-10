package api

import (
	"github.com/geminiblue/favor_guess/config"
	"github.com/gin-gonic/gin"
)

//Run 运行api服务
func Run() {
	app := gin.Default()
	v1 := app.Group("/v1")
	{
		v1.POST("/auth", userLogin)
	}
	app.Run(config.ApiAddr)
}
