package futures

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type positionRiskServiceTestSuite struct {
	baseTestSuite
}

func TestPositionRiskTestService(t *testing.T) {
	suite.Run(t, new(positionRiskServiceTestSuite))
}

func (s *positionRiskServiceTestSuite) TestGetPositionRisk() {
	data := []byte(`[
		{
			"entryPrice": "10359.38000",
			"breakEvenPrice": "10387.38000",
			"marginType": "isolated",
			"isAutoAddMargin": "false",
			"isolatedMargin": "3.15899368",
			"leverage": "125",
			"liquidationPrice": "9332.61",
			"markPrice": "10348.27548846",
			"maxNotionalValue": "50000",
			"positionAmt": "0.003",
			"symbol": "BTCUSDT",
			"unRealizedProfit": "-0.03331353",
			"positionSide": "BOTH"
		}
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSDT"
	recvWindow := int64(1000)
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol":     symbol,
			"recvWindow": recvWindow,
		})
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewGetPositionRiskService().Symbol(symbol).
		Do(newContext(), WithRecvWindow(recvWindow))
	r := s.r()
	r.NoError(err)
	r.Len(res, 1)
	e := &PositionRisk{
		EntryPrice:       "10359.38000",
		BreakEvenPrice:   "10387.38000",
		MarginType:       "isolated",
		IsAutoAddMargin:  "false",
		IsolatedMargin:   "3.15899368",
		Leverage:         "125",
		LiquidationPrice: "9332.61",
		MarkPrice:        "10348.27548846",
		MaxNotionalValue: "50000",
		PositionAmt:      "0.003",
		Symbol:           "BTCUSDT",
		UnRealizedProfit: "-0.03331353",
		PositionSide:     "BOTH",
	}
	s.assertPositionRiskEqual(e, res[0])
}

func (s *positionRiskServiceTestSuite) assertPositionRiskEqual(e, a *PositionRisk) {
	r := s.r()
	r.Equal(e.EntryPrice, a.EntryPrice, "EntryPrice")
	r.Equal(e.BreakEvenPrice, a.BreakEvenPrice, "BreakEvenPrice")
	r.Equal(e.MarginType, a.MarginType, "MarginType")
	r.Equal(e.IsAutoAddMargin, a.IsAutoAddMargin, "IsAutoAddMargin")
	r.Equal(e.IsolatedMargin, a.IsolatedMargin, "IsolatedMargin")
	r.Equal(e.Leverage, a.Leverage, "Leverage")
	r.Equal(e.LiquidationPrice, a.LiquidationPrice, "LiquidationPrice")
	r.Equal(e.MarkPrice, a.MarkPrice, "MarkPrice")
	r.Equal(e.MaxNotionalValue, a.MaxNotionalValue, "MaxNotionalValue")
	r.Equal(e.PositionAmt, a.PositionAmt, "PositionAmt")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.UnRealizedProfit, a.UnRealizedProfit, "UnRealizedProfit")
	r.Equal(e.PositionSide, a.PositionSide, "PositionSide")
}

type positionRiskServiceV3TestSuite struct {
	baseTestSuite
}

func TestPositionRiskV3TestService(t *testing.T) {
	suite.Run(t, new(positionRiskServiceV3TestSuite))
}

func (s *positionRiskServiceV3TestSuite) TestGetPositionRisk() {
	data := []byte(`[
		{
		"symbol": "ADAUSDT",
		"positionSide": "BOTH",
		"positionAmt": "30",
		"entryPrice": "0.385",
		"breakEvenPrice": "0.385077",
		"markPrice": "0.41047590",
		"unRealizedProfit": "0.76427700",
		"liquidationPrice": "0",
		"isolatedMargin": "0",
		"notional": "12.31427700",
		"marginAsset": "USDT",
		"isolatedWallet": "0",
		"initialMargin": "0.61571385",
		"maintMargin": "0.08004280",
		"positionInitialMargin": "0.61571385",
		"openOrderInitialMargin": "0",
		"adl": 2,
		"bidNotional": "0",
		"askNotional": "0",
		"updateTime": 1720736417660
		}
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "ADAUSDT"
	recvWindow := int64(1000)
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol":     symbol,
			"recvWindow": recvWindow,
		})
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewGetPositionRiskV3Service().Symbol(symbol).
		Do(newContext(), WithRecvWindow(recvWindow))
	r := s.r()
	r.NoError(err)
	r.Len(res, 1)
	e := &PositionRiskV3{
		Symbol:                 "ADAUSDT",
		PositionSide:           "BOTH",
		PositionAmt:            "30",
		EntryPrice:             "0.385",
		BreakEvenPrice:         "0.385077",
		MarkPrice:              "0.41047590",
		UnRealizedProfit:       "0.76427700",
		LiquidationPrice:       "0",
		IsolatedMargin:         "0",
		Notional:               "12.31427700",
		MarginAsset:            "USDT",
		IsolatedWallet:         "0",
		InitialMargin:          "0.61571385",
		MaintMargin:            "0.08004280",
		PositionInitialMargin:  "0.61571385",
		OpenOrderInitialMargin: "0",
		Adl:                    2,
		UpdateTime:             1720736417660,
	}
	s.assertPositionRiskEqual(e, res[0])
}

func (s *positionRiskServiceV3TestSuite) assertPositionRiskEqual(e, a *PositionRiskV3) {
	r := s.r()
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.PositionSide, a.PositionSide, "PositionSide")
	r.Equal(e.PositionAmt, a.PositionAmt, "PositionAmt")
	r.Equal(e.EntryPrice, a.EntryPrice, "EntryPrice")
	r.Equal(e.BreakEvenPrice, a.BreakEvenPrice, "BreakEvenPrice")
	r.Equal(e.MarkPrice, a.MarkPrice, "MarkPrice")
	r.Equal(e.UnRealizedProfit, a.UnRealizedProfit, "UnRealizedProfit")
	r.Equal(e.LiquidationPrice, a.LiquidationPrice, "LiquidationPrice")
	r.Equal(e.IsolatedMargin, a.IsolatedMargin, "IsolatedMargin")
	r.Equal(e.Notional, a.Notional, "Notional")
	r.Equal(e.MarginAsset, a.MarginAsset, "MarginAsset")
	r.Equal(e.IsolatedWallet, a.IsolatedWallet, "IsolatedWallet")
	r.Equal(e.InitialMargin, a.InitialMargin, "InitialMargin")
	r.Equal(e.MaintMargin, a.MaintMargin, "MaintMargin")
	r.Equal(e.PositionInitialMargin, a.PositionInitialMargin, "PositionInitialMargin")
	r.Equal(e.OpenOrderInitialMargin, a.OpenOrderInitialMargin, "OpenOrderInitialMargin")
	r.Equal(e.Adl, a.Adl, "Adl")
	r.Equal(e.UpdateTime, a.UpdateTime, "UpdateTime")
}
