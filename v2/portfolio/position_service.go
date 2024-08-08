package portfolio

import (
	"context"

	"github.com/photon-storage/go-binance/v2/delivery"
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
