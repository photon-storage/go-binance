package portfolio

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/photon-storage/go-binance/v2/delivery"
)

type GetCommissionRateService struct {
	ds *delivery.GetCommissionRateService
}

func (s *GetCommissionRateService) Symbol(symbol string) *GetCommissionRateService {
	s.ds.Symbol(symbol)
	return s
}

func (s *GetCommissionRateService) Do(
	ctx context.Context,
	opts ...delivery.RequestOption,
) (*delivery.CommissionRate, error) {
	opts = append(opts, delivery.WithEndpoint("/papi/v1/cm/commissionRate"))
	return s.ds.Do(ctx, opts...)
}

type GetAccountService struct {
	c *Client
}

func (s *GetAccountService) Do(ctx context.Context, opts ...RequestOption) (*Account, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/account",
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := new(Account)
	if err = json.Unmarshal(data, res); err != nil {
		return nil, err
	}
	return res, nil
}

type Account struct {
	UniMMR                   string `json:"uniMMR"`
	AccountEquity            string `json:"accountEquity"`
	ActualEquity             string `json:"actualEquity"`
	AccountInitialMargin     string `json:"accountInitialMargin"`
	AccountMaintMargin       string `json:"accountMaintMargin"`
	AccountStatus            string `json:"accountStatus"`
	VirtualMaxWithdrawAmount string `json:"virtualMaxWithdrawAmount"`
	TotalAvailableBalance    string `json:"totalAvailableBalance"`
	TotalMarginOpenLoss      string `json:"totalMarginOpenLoss"`
	UpdateTime               int64  `json:"updateTime"`
}

// GetBalanceService get account balance
type GetBalanceService struct {
	c *Client
}

// Do send request
func (s *GetBalanceService) Do(ctx context.Context, opts ...RequestOption) ([]*Balance, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/balance",
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := make([]*Balance, 0)
	return res, json.Unmarshal(data, &res)
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

type GetLeverageBracketServiceCM struct {
	c      *Client
	symbol *string
}

func (s *GetLeverageBracketServiceCM) Do(ctx context.Context, opts ...RequestOption) ([]*LeverageBrackets, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/cm/leverageBracket",
		secType:  secTypeSigned,
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := make([]*LeverageBrackets, 0)
	return res, json.Unmarshal(data, &res)
}

func (s *GetLeverageBracketServiceCM) Symbol(symbol string) *GetLeverageBracketServiceCM {
	s.symbol = &symbol
	return s
}

type GetLeverageBracketServiceUM struct {
	c      *Client
	symbol *string
}

func (s *GetLeverageBracketServiceUM) Do(ctx context.Context, opts ...RequestOption) ([]*LeverageBrackets, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/um/leverageBracket",
		secType:  secTypeSigned,
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := make([]*LeverageBrackets, 0)
	return res, json.Unmarshal(data, &res)
}

func (s *GetLeverageBracketServiceUM) Symbol(symbol string) *GetLeverageBracketServiceUM {
	s.symbol = &symbol
	return s
}

type LeverageBrackets struct {
	Symbol   string `json:"symbol"`
	Brackets []struct {
		Bracket                int     `json:"bracket"`
		InitialLeverage        int     `json:"initialLeverage"`
		NotionalFloor          float64 `json:"notionalFloor"`
		NotionalCap            float64 `json:"notionalCap"`
		QtyFloor               float64 `json:"qtyFloor"`
		QtyCap                 float64 `json:"qtyCap"`
		MaintenanceMarginRatio float64 `json:"maintMarginRatio"`
	} `json:"brackets"`
}

type GetMarginMaxBorrowableService struct {
	c     *Client
	asset string
}

func (s *GetMarginMaxBorrowableService) Asset(asset string) *GetMarginMaxBorrowableService {
	s.asset = asset
	return s
}

type GetMarginMaxBorrowableResponse struct {
	Amount      string `json:"amount"`
	BorrowLimit string `json:"borrowLimit"`
}

// Weight: 100
func (s *GetMarginMaxBorrowableService) Do(
	ctx context.Context,
	opts ...RequestOption,
) (*GetMarginMaxBorrowableResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/margin/maxBorrowable",
		secType:  secTypeSigned,
	}
	r.setParam("asset", s.asset)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	var res GetMarginMaxBorrowableResponse
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

type MarginBorrowService struct {
	c      *Client
	asset  string
	amount float64
}

func (s *MarginBorrowService) Asset(asset string) *MarginBorrowService {
	s.asset = asset
	return s
}

func (s *MarginBorrowService) Amount(amount float64) *MarginBorrowService {
	s.amount = amount
	return s
}

type MarginBorrowResponse struct {
	ID int64 `json:"tranId"`
}

// Weight: 100
func (s *MarginBorrowService) Do(
	ctx context.Context,
	opts ...RequestOption,
) (*MarginBorrowResponse, error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/papi/v1/marginLoan",
		secType:  secTypeSigned,
	}
	r.setParam("asset", s.asset)
	r.setParam("amount", s.amount)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	var res MarginBorrowResponse
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

type MarginRepayService struct {
	c      *Client
	asset  string
	amount float64
}

func (s *MarginRepayService) Asset(asset string) *MarginRepayService {
	s.asset = asset
	return s
}

func (s *MarginRepayService) Amount(amount float64) *MarginRepayService {
	s.amount = amount
	return s
}

type MarginRepayResponse struct {
	ID int64 `json:"tranId"`
}

// Weight: 100
func (s *MarginRepayService) Do(
	ctx context.Context,
	opts ...RequestOption,
) (*MarginRepayResponse, error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/papi/v1/repayLoan",
		secType:  secTypeSigned,
	}
	r.setParam("asset", s.asset)
	r.setParam("amount", s.amount)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	var res MarginRepayResponse
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
