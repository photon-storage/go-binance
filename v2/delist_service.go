package binance

import (
	"context"
	"net/http"
)

type GetDelistService struct {
	c *Client
}

// Weight: 50
func (s *GetDelistService) Do(ctx context.Context, opts ...RequestOption) ([]*DelistBatch, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/spot/delist-schedule",
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	var res []*DelistBatch
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type MarginGetDelistService struct {
	c *Client
}

type DelistBatch struct {
	DelistTime int64    `json:"delistTime"`
	Symbols    []string `json:"symbols"`
}

// Weight: 100
func (s *MarginGetDelistService) Do(ctx context.Context, opts ...RequestOption) ([]*MarginDelistBatch, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/margin/delist-schedule",
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	var res []*MarginDelistBatch
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type MarginDelistBatch struct {
	DelistTime            int64    `json:"delistTime"`
	CrossMarginAssets     []string `json:"crossMarginAssets"`
	IsolatedMarginSymbols []string `json:"isolatedMarginSymbols"`
}
