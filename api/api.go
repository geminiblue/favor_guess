package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//userLogin 用户登录
func userLogin(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
	})
}

//getUserInfo 获取用户信息
func getUserInfo(c *gin.Context) {

}
