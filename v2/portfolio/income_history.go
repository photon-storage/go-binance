package portfolio

import (
	"context"

	"github.com/photon-storage/go-binance/v2/futures"
)

type GetIncomeHistoryServiceUM struct {
	Fs *futures.GetIncomeHistoryService
}

func (s *GetIncomeHistoryServiceUM) Do(
	ctx context.Context,
	opts ...futures.RequestOption,
) ([]*futures.IncomeHistory, error) {
	opts = append(opts, futures.WithEndpoint("/papi/v1/um/income"))
	return s.Fs.Do(ctx, opts...)
}
