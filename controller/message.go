package controller

import (
	"fmt"
	"linebotim/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserMes(c *gin.Context) {

	userid := c.Query("userid")
	fmt.Println(userid)
	Messages, err := models.GetMessageByUserID(userid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"UserMes": Messages,
	})

}
