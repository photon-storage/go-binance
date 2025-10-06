package portfolio

import (
	"context"

	"github.com/photon-storage/go-binance/v2"
)

type DustTransferService struct {
	Ms *binance.DustTransferService
}

func (s *DustTransferService) Asset(asset []string) *DustTransferService {
	s.Ms.Asset(asset)
	return s
}

func (s *DustTransferService) MarginAccount() *DustTransferService {
	s.Ms.MarginAccount()
	return s
}

func (s *DustTransferService) Do(
	ctx context.Context,
	opts ...binance.RequestOption,
) (*binance.DustTransferResponse, error) {
	return s.Ms.Do(ctx)
}

type DustListService struct {
	Ms *binance.ListDustService
}

func (s *DustListService) MarginAccount() *DustListService {
	s.Ms.MarginAccount()
	return s
}

func (s *DustListService) Do(
	ctx context.Context,
	opts ...binance.RequestOption,
) (*binance.ListDustResponse, error) {
	return s.Ms.Do(ctx)
}

type ConvertDustLiabilityService struct {
	Ms *binance.ConvertDustLiabilityService
}

func (s *ConvertDustLiabilityService) Asset(asset []string) *ConvertDustLiabilityService {
	s.Ms.Asset(asset)
	return s
}

func (s *ConvertDustLiabilityService) Do(
	ctx context.Context,
	opts ...binance.RequestOption,
) error {
	return s.Ms.Do(ctx)
}

type ListDustLiabilityService struct {
	Ms *binance.ListDustLiabilityService
}

func (s *ListDustLiabilityService) Do(
	ctx context.Context,
	opts ...binance.RequestOption,
) ([]*binance.ListDustLiability, error) {
	return s.Ms.Do(ctx)
}
