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
