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

// Get order
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

// Cancel order
type CancelOrderServiceMargin struct {
	Ms *binance.CancelMarginOrderService
}

func (s *CancelOrderServiceMargin) Do(
	ctx context.Context,
	opts ...binance.RequestOption,
) (*binance.CancelMarginOrderResponse, error) {
	opts = append(opts, binance.WithEndpoint("/papi/v1/margin/order"))
	return s.Ms.Do(ctx, opts...)
}

type CancelOrderServiceCM struct {
	Ds *delivery.CancelOrderService
}

func (s *CancelOrderServiceCM) Do(
	ctx context.Context,
	opts ...delivery.RequestOption,
) (*delivery.CancelOrderResponse, error) {
	opts = append(opts, delivery.WithEndpoint("/papi/v1/cm/order"))
	return s.Ds.Do(ctx, opts...)
}

type CancelOrderServiceUM struct {
	Fs *futures.CancelOrderService
}

func (s *CancelOrderServiceUM) Do(
	ctx context.Context,
	opts ...futures.RequestOption,
) (*futures.CancelOrderResponse, error) {
	opts = append(opts, futures.WithEndpoint("/papi/v1/um/order"))
	return s.Fs.Do(ctx, opts...)
}

// Cancel open orders
type CancelOpenOrdersServiceMargin struct {
	Ms *binance.CancelMarginOpenOrdersService
}

func (s *CancelOpenOrdersServiceMargin) Do(
	ctx context.Context,
	opts ...binance.RequestOption,
) (*binance.CancelOpenOrdersResponse, error) {
	opts = append(opts, binance.WithEndpoint("/papi/v1/margin/allOpenOrders"))
	return s.Ms.Do(ctx, opts...)
}

type CancelOpenOrdersServiceCM struct {
	Ds *delivery.CancelAllOpenOrdersService
}

func (s *CancelOpenOrdersServiceCM) Do(
	ctx context.Context,
	opts ...delivery.RequestOption,
) error {
	opts = append(opts, delivery.WithEndpoint("/papi/v1/cm/allOpenOrders"))
	return s.Ds.Do(ctx, opts...)
}

type CancelOpenOrdersServiceUM struct {
	Fs *futures.CancelAllOpenOrdersService
}

func (s *CancelOpenOrdersServiceUM) Do(
	ctx context.Context,
	opts ...futures.RequestOption,
) error {
	opts = append(opts, futures.WithEndpoint("/papi/v1/um/allOpenOrders"))
	return s.Fs.Do(ctx, opts...)
}

// List orders
type ListOrdersServiceMargin struct {
	Ms *binance.ListMarginOrdersService
}

func (s *ListOrdersServiceMargin) Do(
	ctx context.Context,
	opts ...binance.RequestOption,
) ([]*binance.Order, error) {
	opts = append(opts, binance.WithEndpoint("/papi/v1/margin/allOrders"))
	return s.Ms.Do(ctx, opts...)
}

type ListOrdersServiceCM struct {
	Ds *delivery.ListOrdersService
}

func (s *ListOrdersServiceCM) Do(
	ctx context.Context,
	opts ...delivery.RequestOption,
) ([]*delivery.Order, error) {
	opts = append(opts, delivery.WithEndpoint("/papi/v1/cm/allOrders"))
	return s.Ds.Do(ctx, opts...)
}

type ListOrdersServiceUM struct {
	Fs *futures.ListOrdersService
}

func (s *ListOrdersServiceUM) Do(
	ctx context.Context,
	opts ...futures.RequestOption,
) ([]*futures.Order, error) {
	opts = append(opts, futures.WithEndpoint("/papi/v1/um/allOrders"))
	return s.Fs.Do(ctx, opts...)
}

// List open orders
type ListOpenOrdersServiceMargin struct {
	Ms *binance.ListMarginOpenOrdersService
}

func (s *ListOpenOrdersServiceMargin) Do(
	ctx context.Context,
	opts ...binance.RequestOption,
) ([]*binance.Order, error) {
	opts = append(opts, binance.WithEndpoint("/papi/v1/margin/openOrders"))
	return s.Ms.Do(ctx, opts...)
}

type ListOpenOrdersServiceCM struct {
	Ds *delivery.ListOpenOrdersService
}

func (s *ListOpenOrdersServiceCM) Do(
	ctx context.Context,
	opts ...delivery.RequestOption,
) ([]*delivery.Order, error) {
	opts = append(opts, delivery.WithEndpoint("/papi/v1/cm/openOrders"))
	return s.Ds.Do(ctx, opts...)
}

type ListOpenOrdersServiceUM struct {
	Fs *futures.ListOpenOrdersService
}

func (s *ListOpenOrdersServiceUM) Do(
	ctx context.Context,
	opts ...futures.RequestOption,
) ([]*futures.Order, error) {
	opts = append(opts, futures.WithEndpoint("/papi/v1/um/openOrders"))
	return s.Fs.Do(ctx, opts...)
}

type ListLiquidationOrdersServiceCM struct {
	Ds *delivery.ListLiquidationOrdersService
}

func (s *ListLiquidationOrdersServiceCM) Do(
	ctx context.Context,
	opts ...delivery.RequestOption,
) ([]*delivery.LiquidationOrder, error) {
	opts = append(opts, delivery.WithEndpoint("/papi/v1/cm/forceOrders"))
	return s.Ds.Do(ctx, opts...)
}

type ListLiquidationOrdersServiceUM struct {
	Fs *futures.ListLiquidationOrdersService
}

func (s *ListLiquidationOrdersServiceUM) Do(
	ctx context.Context,
	opts ...futures.RequestOption,
) ([]*futures.LiquidationOrder, error) {
	opts = append(opts, futures.WithEndpoint("/papi/v1/um/forceOrders"))
	return s.Fs.Do(ctx, opts...)
}
