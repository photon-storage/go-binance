package portfolio

import (
	"context"

	"github.com/photon-storage/go-binance/v2/delivery"
)

// CreateOrderService create order
type CreateOrderService struct {
	ds *delivery.CreateOrderService
}

// Do send request
func (s *CreateOrderService) Do(
	ctx context.Context,
	opts ...delivery.RequestOption,
) (*delivery.CreateOrderResponse, error) {
	opts = append(opts, delivery.WithEndpoint("/papi/v1/cm/order"))
	return s.ds.Do(ctx, opts...)
}

// GetOrderService get an order
type GetOrderService struct {
	ds *delivery.GetOrderService
}

func (s *GetOrderService) Do(
	ctx context.Context,
	opts ...delivery.RequestOption,
) (*delivery.Order, error) {
	opts = append(opts, delivery.WithEndpoint("/papi/v1/cm/order"))
	return s.ds.Do(ctx, opts...)
}
