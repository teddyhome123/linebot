package test

import (
	"context"
	"fmt"
	"linebotim/models"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestDb(t *testing.T) {
	fmt.Println("testdb")
	//t.Run("Find", testFind)
	//t.Run("獲取:", testGetMessageByUserID)
	t.Run("獲取All:", testGetAllMessage)
}

func testFind(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		t.Fatal(err)
	}

	db := client.Database("im")
	ub := &models.UserMesInfo{}
	db.Collection("userinfo_mes").FindOne(context.Background(), bson.D{}).Decode(ub)
	fmt.Println(ub)
}

func testGetMessageByUserID(t *testing.T) {
	res, _ := models.GetMessageByUserID("Ue77abf9b8c942c140f04554d95367db0")
	for _, v := range res {
		fmt.Println(v)
	}
}

func testGetAllMessage(t *testing.T) {
	res, _ := models.GetAllMessage()
	for _, v := range res {
		fmt.Println(v)
	}
}