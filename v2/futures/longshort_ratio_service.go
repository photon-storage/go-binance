package futures

import (
	"context"
	"encoding/json"
	"net/http"
)

// LongShortRatioService list open history data of a symbol.
type LongShortRatioService struct {
	c         *Client
	symbol    string
	period    string
	limit     *int
	startTime *int64
	endTime   *int64
}

// Symbol set symbol
func (s *LongShortRatioService) Symbol(symbol string) *LongShortRatioService {
	s.symbol = symbol
	return s
}

// Period set period interval
func (s *LongShortRatioService) Period(period string) *LongShortRatioService {
	s.period = period
	return s
}

// Limit set limit
func (s *LongShortRatioService) Limit(limit int) *LongShortRatioService {
	s.limit = &limit
	return s
}

// StartTime set startTime
func (s *LongShortRatioService) StartTime(startTime int64) *LongShortRatioService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *LongShortRatioService) EndTime(endTime int64) *LongShortRatioService {
	s.endTime = &endTime
	return s
}

// Do send request
func (s *LongShortRatioService) Do(ctx context.Context, opts ...RequestOption) (res []*LongShortRatio, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/futures/data/globalLongShortAccountRatio",
	}

	r.setParam("symbol", s.symbol)
	r.setParam("period", s.period)

	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*LongShortRatio{}, err
	}

	res = make([]*LongShortRatio, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*LongShortRatio{}, err
	}

	return res, nil
}

type LongShortRatio struct {
	Symbol         string `json:"symbol"`
	LongShortRatio string `json:"longShortRatio"`
	LongAccount    string `json:"longAccount"`
	ShortAccount   string `json:"shortAccount"`
	Timestamp      int64  `json:"timestamp"`
}

// TopTraderLongShortPositionRatioService list open history data of a symbol.
type TopTraderLongShortPositionRatioService struct {
	c         *Client
	symbol    string
	period    string
	limit     *int
	startTime *int64
	endTime   *int64
}

// Symbol set symbol
func (s *TopTraderLongShortPositionRatioService) Symbol(symbol string) *TopTraderLongShortPositionRatioService {
	s.symbol = symbol
	return s
}

// Period set period interval
func (s *TopTraderLongShortPositionRatioService) Period(period string) *TopTraderLongShortPositionRatioService {
	s.period = period
	return s
}

// Limit set limit
func (s *TopTraderLongShortPositionRatioService) Limit(limit int) *TopTraderLongShortPositionRatioService {
	s.limit = &limit
	return s
}

// StartTime set startTime
func (s *TopTraderLongShortPositionRatioService) StartTime(startTime int64) *TopTraderLongShortPositionRatioService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *TopTraderLongShortPositionRatioService) EndTime(endTime int64) *TopTraderLongShortPositionRatioService {
	s.endTime = &endTime
	return s
}

// Do send request
func (s *TopTraderLongShortPositionRatioService) Do(ctx context.Context, opts ...RequestOption) (res []*LongShortRatio, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/futures/data/topLongShortPositionRatio",
	}

	r.setParam("symbol", s.symbol)
	r.setParam("period", s.period)

	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*LongShortRatio{}, err
	}

	res = make([]*LongShortRatio, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*LongShortRatio{}, err
	}

	return res, nil
}

// TopTraderLongShortAccountRatioService list open history data of a symbol.
type TopTraderLongShortAccountRatioService struct {
	c         *Client
	symbol    string
	period    string
	limit     *int
	startTime *int64
	endTime   *int64
}

// Symbol set symbol
func (s *TopTraderLongShortAccountRatioService) Symbol(symbol string) *TopTraderLongShortAccountRatioService {
	s.symbol = symbol
	return s
}

// Period set period interval
func (s *TopTraderLongShortAccountRatioService) Period(period string) *TopTraderLongShortAccountRatioService {
	s.period = period
	return s
}

// Limit set limit
func (s *TopTraderLongShortAccountRatioService) Limit(limit int) *TopTraderLongShortAccountRatioService {
	s.limit = &limit
	return s
}

// StartTime set startTime
func (s *TopTraderLongShortAccountRatioService) StartTime(startTime int64) *TopTraderLongShortAccountRatioService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *TopTraderLongShortAccountRatioService) EndTime(endTime int64) *TopTraderLongShortAccountRatioService {
	s.endTime = &endTime
	return s
}

// Do send request
func (s *TopTraderLongShortAccountRatioService) Do(ctx context.Context, opts ...RequestOption) (res []*LongShortRatio, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/futures/data/topLongShortAccountRatio",
	}

	r.setParam("symbol", s.symbol)
	r.setParam("period", s.period)

	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*LongShortRatio{}, err
	}

	res = make([]*LongShortRatio, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*LongShortRatio{}, err
	}

	return res, nil
}
