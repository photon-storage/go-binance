package portfolio

import (
	"context"

	"github.com/photon-storage/go-binance/v2/delivery"
	"github.com/photon-storage/go-binance/v2/futures"
)

type GetPositionRiskServiceCM struct {
	ds *delivery.GetPositionRiskService
}

func (s *GetPositionRiskServiceCM) Do(
	ctx context.Context,
	opts ...delivery.RequestOption,
) ([]*delivery.PositionRisk, error) {
	opts = append(opts, delivery.WithEndpoint("/papi/v1/cm/positionRisk"))
	return s.ds.Do(ctx, opts...)
}

type GetPositionRiskServiceUM struct {
	fs *futures.GetPositionRiskService
}

func (s *GetPositionRiskServiceUM) Do(
	ctx context.Context,
	opts ...futures.RequestOption,
) ([]*futures.PositionRisk, error) {
	opts = append(opts, futures.WithEndpoint("/papi/v1/um/positionRisk"))
	return s.fs.Do(ctx, opts...)
}

type GetPositionAdlServiceCM struct {
	ds *delivery.GetPositionAdlService
}

func (s *GetPositionAdlServiceCM) Do(
	ctx context.Context,
	opts ...delivery.RequestOption,
) ([]*delivery.PositionAdl, error) {
	opts = append(opts, delivery.WithEndpoint("/papi/v1/cm/adlQuantile"))
	return s.ds.Do(ctx, opts...)
}

type GetPositionAdlServiceUM struct {
	fs *futures.GetPositionAdlService
}

func (s *GetPositionAdlServiceUM) Do(
	ctx context.Context,
	opts ...futures.RequestOption,
) ([]*futures.PositionAdl, error) {
	opts = append(opts, futures.WithEndpoint("/papi/v1/um/adlQuantile"))
	return s.fs.Do(ctx, opts...)
}
