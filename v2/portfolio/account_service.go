package portfolio

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/photon-storage/go-binance/v2/delivery"
)

type GetCommissionRateService struct {
	*delivery.GetCommissionRateService
}

func (s *GetCommissionRateService) Symbol(symbol string) *GetCommissionRateService {
	s.GetCommissionRateService.Symbol(symbol)
	return s
}

func (s *GetCommissionRateService) Do(
	ctx context.Context,
	opts ...delivery.RequestOption,
) (*delivery.CommissionRate, error) {
	opts = append(opts, delivery.WithEndpoint("/papi/v1/cm/commissionRate"))
	return s.GetCommissionRateService.Do(ctx, opts...)
}

// GetBalanceService get account balance
type GetBalanceService struct {
	c *Client
}

// Do send request
func (s *GetBalanceService) Do(ctx context.Context, opts ...RequestOption) (res []*Balance, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/balance",
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*Balance{}, err
	}
	res = make([]*Balance, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*Balance{}, err
	}
	return res, nil
}

// Balance define user balance of your account
type Balance struct {
	Asset               string `json:"asset"`
	TotalWalletBalance  string `json:"totalWalletBalance"`
	CrossMarginAsset    string `json:"crossMarginAsset"`
	CrossMarginBorrowed string `json:"crossMarginBorrowed"`
	CrossMarginFree     string `json:"crossMarginFree"`
	CrossMarginInterest string `json:"crossMarginInterest"`
	CrossMarginLocked   string `json:"crossMarginLocked"`
	UmWalletBalance     string `json:"umWalletBalance"`
	UmUnrealizedPNL     string `json:"umUnrealizedPNL"`
	CmWalletBalance     string `json:"cmWalletBalance"`
	CmUnrealizedPNL     string `json:"cmUnrealizedPNL"`
	NegativeBalance     string `json:"negativeBalance"`
	UpdateTime          int64  `json:"updateTime"`
}
