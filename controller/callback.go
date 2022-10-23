package controller

import (
	"linebotim/models"
	"linebotim/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

// linebot接收
func Callback(c *gin.Context) {

	bot, err := linebot.New(
		"e448b3588974cb7a8f0ec8086b3df5b2",
		"dqBvQ+aHd2Mnw4ga+JswaFuMbFBqrX+neVVv0ogcdvPeDWjccHshmD5hll7jhpJcuEPBRhrWFuBj1YRMXLWJZCf7Y8lGF8c+BMgPvnwb8NyymHKA7/X92YfqZKts4UpXKRV41QzlensLoq11j9YenAdB04t89/1O/w1cDnyilFU=",
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
