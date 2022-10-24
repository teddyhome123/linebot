package models

import (
	"context"
	"linebotim/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// 存放User資料的struct
type UserMesInfo struct {
	UserID  string `json:"userid" bson:"userid"`
	MesTime string `json:"mestime" bson:"mesTime"`
	Message string `json:"message" bson:"message"`
}

// insert一筆資料
func InsertUserInfo(ub *UserMesInfo) error {
	coll := utils.MongoDB.Collection("userinfo_mes")
	doc := bson.D{{"userid", ub.UserID}, {"mestime", ub.MesTime}, {"Message", ub.Message}}
	_, err := coll.InsertOne(context.TODO(), doc)
	if err != nil {
		return err
	}
	//fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
	return nil
}

// select All 所有訊息
func GetAllMessage() ([]primitive.D, error) {
	coll := utils.MongoDB.Collection("userinfo_mes")

	filter := bson.D{}

	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	var results []bson.D
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	return results, nil
}

// select User 所有訊息
func GetMessageByUserID(userid string) ([]primitive.D, error) {
	coll := utils.MongoDB.Collection("userinfo_mes")

	filter := bson.D{{"userid", userid}}

	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	var results []bson.D
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil,err
	}

	return results, nil
}
