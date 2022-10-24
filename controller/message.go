package controller

import (
	"linebotim/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllUserMes(c *gin.Context) {

	Messages, err := models.GetAllMessage()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"UserMes": Messages,
	})

}

func GetUserMes(c *gin.Context) {

	//userid := c.Query("userid")
	userid := c.Param("userid")
	
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
