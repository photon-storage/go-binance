package portfolio

import (
	"context"

	"github.com/photon-storage/go-binance/v2/delivery"
)

type ChangeLeverageService struct {
	ds *delivery.ChangeLeverageService
}

func (s *ChangeLeverageService) Do(
	ctx context.Context,
	opts ...delivery.RequestOption,
) (*delivery.SymbolLeverage, error) {
	opts = append(opts, delivery.WithEndpoint("/papi/v1/cm/leverage"))
	return s.ds.Do(ctx, opts...)
}

type ChangePositionModeService struct {
	ds *delivery.ChangePositionModeService
}

func (s *ChangePositionModeService) Do(
	ctx context.Context,
	opts ...delivery.RequestOption,
) error {
	opts = append(opts, delivery.WithEndpoint("/papi/v1/cm/positionSide/dual"))
	return s.ds.Do(ctx, opts...)
}
