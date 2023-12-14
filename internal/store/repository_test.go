package store

import (
	"context"
	"testing"
)

var ctx = context.Background()

// MongoDB 连接字符串
var connectionString = "mongodb://admin:adqwe1234@localhost:27017"

func TestNewMongoDB(t *testing.T) {
	sRepo, err := NewMongoRepo(ctx, connectionString)
	if err != nil {
		t.Fatal(err.Error())
	}
	if err := sRepo.Ping(ctx); err != nil {
		t.Fatal(err)
	}
}
