package main

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Name  string `bson:"name"`
	Age   int    `bson:"age"`
	Email string `bson:"email"`
}

func main() {
	ctx := context.Background()
	// MongoDB 连接字符串
	connectionString := "mongodb://admin:adqwe123@localhost:27017"

	// 连接 MongoDB
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		slog.Error(err.Error())
	}

	// 创建 context 和取消函数
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 确保在函数执行完毕后断开连接
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			slog.Error(err.Error())
		}
	}()

	// 选择数据库
	databaseName := "testdb"
	db := client.Database(databaseName)

	// 选择集合（表）
	collectionName := "users"
	collection := db.Collection(collectionName)

	// 插入测试数据
	user := User{
		Name:  "John Doe",
		Age:   30,
		Email: "john@example.com",
	}

	insertResult, err := collection.InsertOne(ctx, user)
	if err != nil {
		slog.Error(err.Error())
	}

	fmt.Printf("Inserted document with ID: %v\n", insertResult.InsertedID)
}
