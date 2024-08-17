package portfolio

import (
	"context"

	"github.com/photon-storage/go-binance/v2/delivery"
	"github.com/photon-storage/go-binance/v2/futures"
)

type ChangeLeverageServiceCM struct {
	ds *delivery.ChangeLeverageService
}

func (s *ChangeLeverageServiceCM) Symbol(symbol string) *ChangeLeverageServiceCM {
	s.ds.Symbol(symbol)
	return s
}

func (s *ChangeLeverageServiceCM) Leverage(leverage int) *ChangeLeverageServiceCM {
	s.ds.Leverage(leverage)
	return s
}

func (s *ChangeLeverageServiceCM) Do(
	ctx context.Context,
	opts ...delivery.RequestOption,
) (*delivery.SymbolLeverage, error) {
	opts = append(opts, delivery.WithEndpoint("/papi/v1/cm/leverage"))
	return s.ds.Do(ctx, opts...)
}

type ChangeLeverageServiceUM struct {
	fs *futures.ChangeLeverageService
}

func (s *ChangeLeverageServiceUM) Symbol(symbol string) *ChangeLeverageServiceUM {
	s.fs.Symbol(symbol)
	return s
}

func (s *ChangeLeverageServiceUM) Leverage(leverage int) *ChangeLeverageServiceUM {
	s.fs.Leverage(leverage)
	return s
}

func (s *ChangeLeverageServiceUM) Do(
	ctx context.Context,
	opts ...futures.RequestOption,
) (*futures.SymbolLeverage, error) {
	opts = append(opts, futures.WithEndpoint("/papi/v1/um/leverage"))
	return s.fs.Do(ctx, opts...)
}

type ChangePositionModeServiceCM struct {
	ds *delivery.ChangePositionModeService
}

func (s *ChangePositionModeServiceCM) DualSide(dualSide bool) *ChangePositionModeServiceCM {
	s.ds.DualSide(dualSide)
	return s
}

func (s *ChangePositionModeServiceCM) Do(
	ctx context.Context,
	opts ...delivery.RequestOption,
) error {
	opts = append(opts, delivery.WithEndpoint("/papi/v1/cm/positionSide/dual"))
	return s.ds.Do(ctx, opts...)
}

type ChangePositionModeServiceUM struct {
	fs *futures.ChangePositionModeService
}

func (s *ChangePositionModeServiceUM) DualSide(dualSide bool) *ChangePositionModeServiceUM {
	s.fs.DualSide(dualSide)
	return s
}

func (s *ChangePositionModeServiceUM) Do(
	ctx context.Context,
	opts ...futures.RequestOption,
) error {
	opts = append(opts, futures.WithEndpoint("/papi/v1/um/positionSide/dual"))
	return s.fs.Do(ctx, opts...)
}

type GetPositionModeServiceCM struct {
	ds *delivery.GetPositionModeService
}

func (s *GetPositionModeServiceCM) Do(
	ctx context.Context,
	opts ...delivery.RequestOption,
) (*delivery.PositionMode, error) {
	opts = append(opts, delivery.WithEndpoint("/papi/v1/cm/positionSide/dual"))
	return s.ds.Do(ctx, opts...)
}

type GetPositionModeServiceUM struct {
	fs *futures.GetPositionModeService
}

func (s *GetPositionModeServiceUM) Do(
	ctx context.Context,
	opts ...futures.RequestOption,
) (*futures.PositionMode, error) {
	opts = append(opts, futures.WithEndpoint("/papi/v1/um/positionSide/dual"))
	return s.fs.Do(ctx, opts...)
}
