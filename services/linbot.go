package services

import (
	"errors"
	"linebotim/models"
	"time"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func GetUserAndMes(events []*linebot.Event) (*models.UserMesInfo, error) {
	//取得訊息裡面User和Mes
	for _, event := range events {
		userID := event.Source.UserID
		//timestamp := event.Timestamp
		//timestamp to string
		time := time.Now().Format("2006-01-02 15:04:05")
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				UserMesInfo := &models.UserMesInfo{
					UserID:  userID,
					MesTime: time,
					Message: message.Text,
				}
				return UserMesInfo, nil

			default:
				err := errors.New("type error")
				return nil, err

			}
		}
	}
	return nil, nil
}
