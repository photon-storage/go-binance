package futures

import (
	"context"
	"encoding/json"
	"net/http"
)

// GetPositionRiskService get account balance
type GetPositionRiskService struct {
	c      *Client
	symbol string
}

// Symbol set symbol
func (s *GetPositionRiskService) Symbol(symbol string) *GetPositionRiskService {
	s.symbol = symbol
	return s
}

// Do send request
func (s *GetPositionRiskService) Do(ctx context.Context, opts ...RequestOption) (res []*PositionRisk, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v2/positionRisk",
		secType:  secTypeSigned,
	}
	if s.symbol != "" {
		r.setParam("symbol", s.symbol)
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*PositionRisk{}, err
	}
	res = make([]*PositionRisk, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*PositionRisk{}, err
	}
	return res, nil
}

// PositionRisk define position risk info
type PositionRisk struct {
	EntryPrice       string `json:"entryPrice"`
	BreakEvenPrice   string `json:"breakEvenPrice"`
	MarginType       string `json:"marginType"`
	IsAutoAddMargin  string `json:"isAutoAddMargin"`
	IsolatedMargin   string `json:"isolatedMargin"`
	Leverage         string `json:"leverage"`
	LiquidationPrice string `json:"liquidationPrice"`
	MarkPrice        string `json:"markPrice"`
	MaxNotionalValue string `json:"maxNotionalValue"`
	PositionAmt      string `json:"positionAmt"`
	Symbol           string `json:"symbol"`
	UnRealizedProfit string `json:"unRealizedProfit"`
	PositionSide     string `json:"positionSide"`
	Notional         string `json:"notional"`
	IsolatedWallet   string `json:"isolatedWallet"`
}

// GetPositionRiskV3Service get account balance
type GetPositionRiskV3Service struct {
	c      *Client
	symbol string
}

// Symbol set symbol
func (s *GetPositionRiskV3Service) Symbol(symbol string) *GetPositionRiskV3Service {
	s.symbol = symbol
	return s
}

// Do send request
func (s *GetPositionRiskV3Service) Do(ctx context.Context, opts ...RequestOption) (res []*PositionRiskV3, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v3/positionRisk",
		secType:  secTypeSigned,
	}
	if s.symbol != "" {
		r.setParam("symbol", s.symbol)
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*PositionRiskV3{}, err
	}
	res = make([]*PositionRiskV3, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*PositionRiskV3{}, err
	}
	return res, nil
}

// PositionRiskV3 define position risk info
type PositionRiskV3 struct {
	Symbol                 string `json:"symbol"`
	PositionSide           string `json:"positionSide"`
	PositionAmt            string `json:"positionAmt"`
	EntryPrice             string `json:"entryPrice"`
	BreakEvenPrice         string `json:"breakEvenPrice"`
	MarkPrice              string `json:"markPrice"`
	UnRealizedProfit       string `json:"unRealizedProfit"`
	LiquidationPrice       string `json:"liquidationPrice"`
	IsolatedMargin         string `json:"isolatedMargin"`
	Notional               string `json:"notional"`
	MarginAsset            string `json:"marginAsset"`
	IsolatedWallet         string `json:"isolatedWallet"`
	InitialMargin          string `json:"initialMargin"`
	MaintMargin            string `json:"maintMargin"`
	PositionInitialMargin  string `json:"positionInitialMargin"`
	OpenOrderInitialMargin string `json:"openOrderInitialMargin"`
	Adl                    int    `json:"adl"`
	UpdateTime             int64  `json:"updateTime"`
}

type GetPositionAdlService struct {
	c      *Client
	symbol string
}

// Symbol set symbol
func (s *GetPositionAdlService) Symbol(symbol string) *GetPositionAdlService {
	s.symbol = symbol
	return s
}

// Do send request
func (s *GetPositionAdlService) Do(ctx context.Context, opts ...RequestOption) (res []*PositionAdl, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/adlQuantile",
		secType:  secTypeSigned,
	}
	if s.symbol != "" {
		r.setParam("symbol", s.symbol)
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*PositionAdl{}, err
	}
	res = make([]*PositionAdl, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*PositionAdl{}, err
	}
	return res, nil
}

type AdlQuantile struct {
	Long  int  `json:"LONG"`
	Short int  `json:"SHORT"`
	Hedge int  `json:"HEDGE"`
	Both  *int `json:"BOTH"`
}

type PositionAdl struct {
	Symbol      string       `json:"symbol"`
	AdlQuantile *AdlQuantile `json:"adlQuantile"`
}
