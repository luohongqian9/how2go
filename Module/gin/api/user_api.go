package api

import (
	"ginlearn/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserHandler(c *gin.Context) {
	id := c.Param("id")

	user, err := service.GetUserInfo(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}
