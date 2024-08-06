package portfolio

import (
	"context"

	"github.com/photon-storage/go-binance/v2/delivery"
)

// CreateOrderService create order
type CreateOrderService struct {
	*delivery.CreateOrderService
}

// Do send request
func (s *CreateOrderService) Do(
	ctx context.Context,
	opts ...delivery.RequestOption,
) (res *delivery.CreateOrderResponse, err error) {
	opts = append(opts, delivery.WithEndpoint("/papi/v1/cm/order"))
	return s.CreateOrderService.Do(ctx, opts...)
}
