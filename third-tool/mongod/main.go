package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

var client *mongo.Client

func init() {
	var err error
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://47.98.199.80:27017"))
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	//ping()
	queryRecord()
	//saveRecord()
}

func ping() {
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	err := client.Ping(ctx, readpref.Primary())
	fmt.Println(err)
}

// 添加一条记录
func saveRecord() {
	collection := client.Database("test").Collection("user")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := collection.InsertOne(ctx, bson.M{"name": "dazuo", "age": 22, "gender": 1, "createTime": time.Now()})
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(result.InsertedID) // 5ed786735246e0ba823fac38
}

// 查询文档
func queryRecord() {
	collection := client.Database("test").Collection("user")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	objectId, _ := primitive.ObjectIDFromHex("5ed786735246e0ba823fac38")
	cur, err := collection.Find(ctx, bson.D{{"_id", objectId}, {"name", "dazuo"}})
	if err != nil {
		log.Panic(err)
	}
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			log.Panic(err)
		}
		fmt.Println(result)
	}
	if err := cur.Err(); err != nil {
		log.Panic(err)
	}
}
