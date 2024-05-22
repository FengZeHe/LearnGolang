package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type Users struct {
	Name  string
	Score int
}

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(err.Error())
	}
	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connect mongodb success")
	defer client.Disconnect(context.Background())

	collection := client.Database("game").Collection("users")
	u1 := Users{"大伟", 66}

	res, err := collection.InsertOne(context.TODO(), u1)
	if err != nil {
		log.Println("error", err)
	}
	fmt.Println(res.InsertedID)

	// 设置过滤条件
	filter := bson.M{"name": "大伟"}

	//设置更新操作
	update := bson.M{"$set": bson.M{"name": "大伟!!"}}

	// 执行更新操作
	insertRes, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		fmt.Println("ERROR", err)
		return
	}
	fmt.Println(insertRes.UpsertedCount)

	// 执行查询操作
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("ERROR", err)
	}
	defer cur.Close(context.Background())
	// 遍历查询结果
	for cur.Next(context.Background()) {
		var user Users
		if err := cur.Decode(&user); err != nil {
			fmt.Println("Error decoding document:", err)
			continue
		}
		fmt.Println("user=", user)
	}

	delFilter := bson.M{"name": "大伟"}
	// 执行删除操作
	delRes, err := collection.DeleteOne(context.TODO(), delFilter)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(delRes.DeletedCount)
}
