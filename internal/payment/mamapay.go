package payment

import (
	"context"
	"errors"
	"fmt"

	"github.com/Rhymond/go-money"
)

type MamaPayApi struct {
	key string
}

type SourceParams struct {
	Token string `form:"source"`
}

type ChargeParams struct {
	Amount   int64         `form:"amount"`
	Currency string        `form:"currency"`
	Source   *SourceParams `form:"*"`
}

type PayService struct {
	mamapayClient *MamaPayApi // 假设的mama支付API
}

func NewPayService(apiKey string) (*PayService, error) {
	if apiKey == "" {
		return nil, errors.New("API key cannot be nil ")
	}
	sc := &MamaPayApi{key: apiKey}
	return &PayService{mamapayClient: sc}, nil
}

func (s PayService) Charges(cp *ChargeParams) (bool, error) {
	return true, nil
}

func (s PayService) ChargeCard(ctx context.Context, amount money.Money, cardToken string) error {
	params := &ChargeParams{
		Amount:   amount.Amount(),
		Currency: "usd",
		Source:   &SourceParams{Token: cardToken},
	}
	_, err := s.Charges(params)
	if err != nil {
		return fmt.Errorf("failed to create a charge:%w", err)
	}
	return nil
}
