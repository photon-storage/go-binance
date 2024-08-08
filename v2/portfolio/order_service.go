package portfolio

import (
	"context"

	"github.com/photon-storage/go-binance/v2/delivery"
)

// CreateOrderService create order
type CreateOrderServiceCM struct {
	Ds *delivery.CreateOrderService
}

// Do send request
func (s *CreateOrderServiceCM) Do(
	ctx context.Context,
	opts ...delivery.RequestOption,
) (*delivery.CreateOrderResponse, error) {
	opts = append(opts, delivery.WithEndpoint("/papi/v1/cm/order"))
	return s.Ds.Do(ctx, opts...)
}

// GetOrderService get an order
type GetOrderServiceCM struct {
	Ds *delivery.GetOrderService
}

func (s *GetOrderServiceCM) Do(
	ctx context.Context,
	opts ...delivery.RequestOption,
) (*delivery.Order, error) {
	opts = append(opts, delivery.WithEndpoint("/papi/v1/cm/order"))
	return s.Ds.Do(ctx, opts...)
}
