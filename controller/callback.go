package controller

import (
	"linebotim/models"
	"linebotim/services"
	"linebotim/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

// linebot接收
func Callback(c *gin.Context) {
	config, err := utils.LoadConfig("./.")
	if err != nil {
		log.Fatal("load config err = ", err)
	}

	bot, err := linebot.New(
		config.ChannelSecret,
		config.ChannelToken,
	)
	if err != nil {
		log.Fatal(err)
	}

	events, err := bot.ParseRequest(c.Request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	//Get UserInfo And Message
	userMesInfo, err := services.GetUserAndMes(events)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	//insert User Mes
	models.InsertUserInfo(userMesInfo)

	c.JSON(http.StatusOK, nil)
}
