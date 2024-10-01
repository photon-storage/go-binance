package portfolio

import (
	"context"

	"github.com/photon-storage/go-binance/v2"
	"github.com/photon-storage/go-binance/v2/delivery"
	"github.com/photon-storage/go-binance/v2/futures"
)

type CreateOrderServiceMargin struct {
	Ms *binance.CreateMarginOrderService
}

func (s *CreateOrderServiceMargin) Do(
	ctx context.Context,
	opts ...binance.RequestOption,
) (*binance.CreateOrderResponse, error) {
	opts = append(opts, binance.WithEndpoint("/papi/v1/margin/order"))
	return s.Ms.Do(ctx, opts...)
}

type CreateOrderServiceCM struct {
	Ds *delivery.CreateOrderService
}

func (s *CreateOrderServiceCM) Do(
	ctx context.Context,
	opts ...delivery.RequestOption,
) (*delivery.CreateOrderResponse, error) {
	opts = append(opts, delivery.WithEndpoint("/papi/v1/cm/order"))
	return s.Ds.Do(ctx, opts...)
}

type CreateOrderServiceUM struct {
	Fs *futures.CreateOrderService
}

func (s *CreateOrderServiceUM) Do(
	ctx context.Context,
	opts ...futures.RequestOption,
) (*futures.CreateOrderResponse, error) {
	opts = append(opts, futures.WithEndpoint("/papi/v1/um/order"))
	return s.Fs.Do(ctx, opts...)
}

type GetOrderServiceMargin struct {
	Ms *binance.GetMarginOrderService
}

func (s *GetOrderServiceMargin) Do(
	ctx context.Context,
	opts ...binance.RequestOption,
) (*binance.Order, error) {
	opts = append(opts, binance.WithEndpoint("/papi/v1/margin/order"))
	return s.Ms.Do(ctx, opts...)
}

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

type GetOrderServiceUM struct {
	Fs *futures.GetOrderService
}

func (s *GetOrderServiceUM) Do(
	ctx context.Context,
	opts ...futures.RequestOption,
) (*futures.Order, error) {
	opts = append(opts, futures.WithEndpoint("/papi/v1/um/order"))
	return s.Fs.Do(ctx, opts...)
}
