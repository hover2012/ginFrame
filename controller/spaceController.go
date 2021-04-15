package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
var secrets = gin.H{
	"wang": gin.H{"email":"w@meidaifu.com","phone":6610},
	"liu":gin.H{"email":"l@meidaifu.com","phone":6611},
}
func ShowSpaceList(c *gin.Context )  {
	var msg struct{
		Name string `json:"user"`
		Message string
		Number int
	}
	msg.Name = "meidaifu"
	msg.Message = "success"
	msg.Number =100
	c.JSON(http.StatusOK,msg)
}