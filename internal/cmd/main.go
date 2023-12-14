package main

import (
	"context"
	"log/slog"

	"github.com/Rhymond/go-money"
	"github.com/google/uuid"

	coffeeco "coffeeco/internal"
	"coffeeco/internal/payment"
	"coffeeco/internal/purchase"
	"coffeeco/internal/store"
)

func main() {
	ctx := context.Background()

	// 模拟的支付API key
	mamaTestAPIKey := "sk_test_4eC39HqLyjWDarjtT1zdp7dc"
	// test token
	cardToken := "tok_visa"
	// 初始化支付服务
	csvc, err := payment.NewPayService(mamaTestAPIKey)
	if err != nil {
		panic(err)
	}

	mongoConString := "mongodb://admin:adqwe123@localhost:27017"
	prepo, err := purchase.NewMongoRepo(ctx, mongoConString)
	if err != nil {
		panic(err)
	}
	if err := prepo.Ping(ctx); err != nil {
		panic(err)
	}

	sRepo, err := store.NewMongoRepo(ctx, mongoConString)
	if err != nil {
		panic(err)
	}
	if err := sRepo.Ping(ctx); err != nil {
		panic(err)
	}

	sSvc := store.NewService(sRepo)

	svc := purchase.NewService(csvc, prepo, sSvc)

	someStoreID := uuid.New()

	pur := &purchase.Purchase{
		CardToken: &cardToken,
		Store: store.Store{
			ID: someStoreID,
		},
		ProductsToPurchase: []coffeeco.Product{{
			ItemName:  "item1",
			BasePrice: *money.New(250, "USD"),
		}},
		PaymentMeans: payment.MEANS_CARD,
	}
	if err := svc.CompletePurchase(ctx, someStoreID, pur, nil); err != nil {
		panic(err)
	}

	slog.Info("purchase was successful")
}
