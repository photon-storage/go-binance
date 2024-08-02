package binance

import (
	"context"
	"net/http"
)

type GetEarnAccountService struct {
	c *Client
}

// Do send request
func (s *GetEarnAccountService) Do(ctx context.Context, opts ...RequestOption) (*EarnAccount, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/simple-earn/account",
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	var res EarnAccount
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

type EarnAccount struct {
	TotalInBTC     string `json:"totalAmountInBTC"`
	TotalInUSDT    string `json:"totalAmountInUSDT"`
	FlexibleInBTC  string `json:"totalFlexibleAmountInBTC"`
	FlexibleInUSDT string `json:"totalFlexibleAmountInUSDT"`
	LockedInBTC    string `json:"totalLockedInBTC"`
	LockedInUSDT   string `json:"totalLockedInUSDT"`
}
