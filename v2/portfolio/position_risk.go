package portfolio

import (
	"context"

	"github.com/photon-storage/go-binance/v2/delivery"
)

type GetPositionRiskService struct {
	ds *delivery.GetPositionRiskService
}

func (s *GetPositionRiskService) Do(
	ctx context.Context,
	opts ...delivery.RequestOption,
) ([]*delivery.PositionRisk, error) {
	opts = append(opts, delivery.WithEndpoint("/papi/v1/cm/positionRisk"))
	return s.ds.Do(ctx, opts...)
}
