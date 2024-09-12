package portfolio

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/photon-storage/go-binance/v2/delivery"
	"github.com/photon-storage/go-binance/v2/futures"
)

// Endpoints
const (
	baseWsMainUrl = "wss://fstream.binance.com/pm/ws"
)

var (
	// WebsocketTimeout is an interval for sending ping/pong messages if WebsocketKeepalive is enabled
	WebsocketTimeout = time.Second * 60
	// WebsocketKeepalive enables sending ping/pong messages to check the connection stability
	WebsocketKeepalive = false
)

// WsUserDataEvent define user data event
type WsUserDataEvent struct {
	Event            UserDataEventType `json:"e"`
	FutureSubtype    FutureSubtype     `json:"fs"`
	Time             int64             `json:"E"`
	TransactionTime  int64             `json:"T"`
	OrderTradeUpdate WsOrderTradeUpdate
}

type WsOrderTradeUpdate struct {
	CM *delivery.WsOrderTradeUpdate
	UM *futures.WsOrderTradeUpdate
}

// WsUserDataHandler handle WsUserDataEvent
type WsUserDataHandler func(event *WsUserDataEvent)

// WsUserDataServe serve user data handler with listen key
func WsUserDataServe(listenKey string, handler WsUserDataHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/%s", baseWsMainUrl, listenKey)
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		event := new(WsUserDataEvent)
		if err := json.Unmarshal(message, event); err != nil {
			errHandler(err)
			return
		}

		if event.FutureSubtype == FutureSubtypeCM {
			cmEvent := new(delivery.WsUserDataEvent)
			if err := json.Unmarshal(message, cmEvent); err != nil {
				errHandler(err)
				return
			}

			event.OrderTradeUpdate.CM = &cmEvent.OrderTradeUpdate
		}

		if event.FutureSubtype == FutureSubtypeUM {
			umEvent := new(futures.WsUserDataEvent)
			if err := json.Unmarshal(message, umEvent); err != nil {
				errHandler(err)
				return
			}

			event.OrderTradeUpdate.UM = &umEvent.OrderTradeUpdate
		}

		handler(event)
	}
	return wsServe(cfg, wsHandler, errHandler)
}
