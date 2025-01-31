package portfolio

import (
	"context"
	"net/http"
)

type FutureRepayService struct {
	c *Client
}

func (s *FutureRepayService) Do(ctx context.Context, opts ...RequestOption) error {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/papi/v1/repay-futures-negative-balance",
		secType:  secTypeSigned,
	}
	_, err := s.c.callAPI(ctx, r, opts...)
	return err
}

type AutoCollectionService struct {
	c *Client
}

func (s *AutoCollectionService) Do(ctx context.Context, opts ...RequestOption) error {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/papi/v1/auto-collection",
		secType:  secTypeSigned,
	}
	_, err := s.c.callAPI(ctx, r, opts...)
	return err
}

type BnbTransferService struct {
	c      *Client
	amount float64
	toUM   bool
}

func (s *BnbTransferService) Amount(v float64) *BnbTransferService {
	s.amount = v
	return s
}

func (s *BnbTransferService) ToUM(v bool) *BnbTransferService {
	s.toUM = v
	return s
}

func (s *BnbTransferService) Do(ctx context.Context, opts ...RequestOption) error {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/papi/v1/bnb-transfer",
		secType:  secTypeSigned,
	}
	r.setFormParams(params{
		"amount":       s.amount,
		"transferSide": map[bool]string{true: "TO_UM", false: "FROM_UM"}[s.toUM],
	})
	_, err := s.c.callAPI(ctx, r, opts...)
	return err
}
