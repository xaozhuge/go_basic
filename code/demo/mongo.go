package main

import (
	"fmt"
	"context"
	"time"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
)

type Person struct {
	Name  string
	Age   int
	City  string
	State string
}


func main() {
	// 设置 MongoDB 连接信息
	clientOptions := options.Client().ApplyURI("mongodb://admin:admin@172.16.1.218:27017")

	// 如果需要其他配置，可以继续设置
	clientOptions = clientOptions.
		SetConnectTimeout(10 * time.Second).
		SetSocketTimeout(10 * time.Second)

	// 连接到 MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	// 选择数据库和集合
	database := client.Database("your_database1")
	collection := database.Collection("your_collection")

	// 构建查询条件
	filter := bson.D{{"name", "John"}}

	// 执行查询
	var result Person
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(result)
	fmt.Printf("Name: %s, Age: %d, City: %s, State: %s\n", result.Name, result.Age, result.City, result.State)


}
