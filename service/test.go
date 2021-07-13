package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary 获取消息列表数据
// @Description get all messages
// @Accept  json
// @Produce json
// @Success 200 {string} string "success"
// @Router /message [get]
func GetMessage(c *gin.Context) {
	fmt.Fprint(c.Writer, "sdf")
}

// @Summary Hello接口
// @Description Hello接口
// @Tags 用户信息
// @Success 200 {string} json "{"message":"success"}"
// @Router /hello [get]
func Hello(c *gin.Context) {
	// 当响应码为200时，返回JSON格式数据
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
