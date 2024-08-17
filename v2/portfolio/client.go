package portfolio

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/photon-storage/go-binance/v2/common"
	"github.com/photon-storage/go-binance/v2/delivery"
	"github.com/photon-storage/go-binance/v2/futures"
)

const (
	apiUrl = "https://papi.binance.com"

	timestampKey  = "timestamp"
	signatureKey  = "signature"
	recvWindowKey = "recvWindow"
)

func currentTimestamp() int64 {
	return int64(time.Nanosecond) * time.Now().UnixNano() / int64(time.Millisecond)
}

// NewClient initialize an API client instance with API key and secret key.
// You should always call this function before using this SDK.
// Services will be created by the form client.NewXXXService().
func NewClient(apiKey, secretKey string) *Client {
	return &Client{
		APIKey:     apiKey,
		SecretKey:  secretKey,
		BaseURL:    apiUrl,
		HTTPClient: http.DefaultClient,
		Logger:     log.New(os.Stderr, "Binance-golang ", log.LstdFlags),
	}
}

type doFunc func(req *http.Request) (*http.Response, error)

// Client define API client
type Client struct {
	APIKey     string
	SecretKey  string
	BaseURL    string
	HTTPClient *http.Client
	Debug      bool
	Logger     *log.Logger
	TimeOffset int64
	do         doFunc
}

func (c *Client) debug(format string, v ...interface{}) {
	if c.Debug {
		c.Logger.Printf(format, v...)
	}
}

func (c *Client) parseRequest(r *request, opts ...RequestOption) (err error) {
	// set request options from user
	for _, opt := range opts {
		opt(r)
	}
	err = r.validate()
	if err != nil {
		return err
	}

	fullURL := fmt.Sprintf("%s%s", c.BaseURL, r.endpoint)
	if r.recvWindow > 0 {
		r.setParam(recvWindowKey, r.recvWindow)
	}
	if r.secType == secTypeSigned {
		r.setParam(timestampKey, currentTimestamp()-c.TimeOffset)
	}
	queryString := r.query.Encode()
	body := &bytes.Buffer{}
	bodyString := r.form.Encode()
	header := http.Header{}
	if r.header != nil {
		header = r.header.Clone()
	}
	if bodyString != "" {
		header.Set("Content-Type", "application/x-www-form-urlencoded")
		body = bytes.NewBufferString(bodyString)
	}
	if r.secType == secTypeAPIKey || r.secType == secTypeSigned {
		header.Set("X-MBX-APIKEY", c.APIKey)
	}

	if r.secType == secTypeSigned {
		raw := fmt.Sprintf("%s%s", queryString, bodyString)
		mac := hmac.New(sha256.New, []byte(c.SecretKey))
		_, err = mac.Write([]byte(raw))
		if err != nil {
			return err
		}
		v := url.Values{}
		v.Set(signatureKey, fmt.Sprintf("%x", (mac.Sum(nil))))
		if queryString == "" {
			queryString = v.Encode()
		} else {
			queryString = fmt.Sprintf("%s&%s", queryString, v.Encode())
		}
	}
	if queryString != "" {
		fullURL = fmt.Sprintf("%s?%s", fullURL, queryString)
	}
	c.debug("full url: %s, body: %s", fullURL, bodyString)

	r.fullURL = fullURL
	r.header = header
	r.body = body
	return nil
}

func (c *Client) callAPI(ctx context.Context, r *request, opts ...RequestOption) (data []byte, err error) {
	err = c.parseRequest(r, opts...)
	if err != nil {
		return []byte{}, err
	}
	req, err := http.NewRequest(r.method, r.fullURL, r.body)
	if err != nil {
		return []byte{}, err
	}
	req = req.WithContext(ctx)
	req.Header = r.header
	c.debug("request: %#v", req)
	f := c.do
	if f == nil {
		f = c.HTTPClient.Do
	}
	res, err := f(req)
	if err != nil {
		return []byte{}, err
	}
	data, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}
	defer func() {
		cerr := res.Body.Close()
		// Only overwrite the retured error if the original error was nil and an
		// error occurred while closing the body.
		if err == nil && cerr != nil {
			err = cerr
		}
	}()
	c.debug("response: %#v", res)
	c.debug("response body: %s", string(data))
	c.debug("response status code: %d", res.StatusCode)

	if res.StatusCode >= http.StatusBadRequest {
		apiErr := new(common.APIError)
		e := json.Unmarshal(data, apiErr)
		if e != nil {
			c.debug("failed to unmarshal json: %s", e)
		}
		return nil, apiErr
	}
	return data, nil
}

func (c *Client) newDeliveryClient() *delivery.Client {
	dc := delivery.NewClient(c.APIKey, c.SecretKey)
	dc.BaseURL = c.BaseURL
	dc.Debug = c.Debug
	return dc
}

func (c *Client) newFutureClient() *futures.Client {
	fc := futures.NewClient(c.APIKey, c.SecretKey)
	fc.BaseURL = c.BaseURL
	fc.Debug = c.Debug
	return fc
}

func (c *Client) NewCreateOrderServiceCM() *CreateOrderServiceCM {
	dc := c.newDeliveryClient()
	return &CreateOrderServiceCM{Ds: dc.NewCreateOrderService()}
}

func (c *Client) NewCreateOrderServiceUM() *CreateOrderServiceUM {
	fc := c.newFutureClient()
	return &CreateOrderServiceUM{Fs: fc.NewCreateOrderService()}
}

func (c *Client) NewChangeLeverageServiceCM() *ChangeLeverageServiceCM {
	dc := c.newDeliveryClient()
	return &ChangeLeverageServiceCM{ds: dc.NewChangeLeverageService()}
}

func (c *Client) NewChangeLeverageServiceUM() *ChangeLeverageServiceUM {
	fc := c.newFutureClient()
	return &ChangeLeverageServiceUM{fs: fc.NewChangeLeverageService()}
}

func (c *Client) NewChangePositionModeServiceCM() *ChangePositionModeServiceCM {
	dc := c.newDeliveryClient()
	return &ChangePositionModeServiceCM{ds: dc.NewChangePositionModeService()}
}

func (c *Client) NewChangePositionModeServiceUM() *ChangePositionModeServiceUM {
	fc := c.newFutureClient()
	return &ChangePositionModeServiceUM{fs: fc.NewChangePositionModeService()}
}

func (c *Client) NewGetCommissionRateService() *GetCommissionRateService {
	dc := c.newDeliveryClient()
	return &GetCommissionRateService{ds: dc.NewGetCommissionRateService()}
}

func (c *Client) NewGetOrderServiceCM() *GetOrderServiceCM {
	dc := c.newDeliveryClient()
	return &GetOrderServiceCM{Ds: dc.NewGetOrderService()}
}

func (c *Client) NewGetOrderServiceUM() *GetOrderServiceUM {
	fc := c.newFutureClient()
	return &GetOrderServiceUM{Fs: fc.NewGetOrderService()}
}

func (c *Client) NewGetPositionModeServiceCM() *GetPositionModeServiceCM {
	dc := c.newDeliveryClient()
	return &GetPositionModeServiceCM{ds: dc.NewGetPositionModeService()}
}

func (c *Client) NewGetPositionModeServiceUM() *GetPositionModeServiceUM {
	fc := c.newFutureClient()
	return &GetPositionModeServiceUM{fs: fc.NewGetPositionModeService()}
}

func (c *Client) NewGetAccountService() *GetAccountService {
	return &GetAccountService{c: c}
}

func (c *Client) NewGetBalanceService() *GetBalanceService {
	return &GetBalanceService{c: c}
}

func (c *Client) NewGetPositionRiskService() *GetPositionRiskService {
	dc := c.newDeliveryClient()
	return &GetPositionRiskService{ds: dc.NewGetPositionRiskService()}
}

func (c *Client) NewFutureRepayService() *FutureRepayService {
	return &FutureRepayService{c: c}
}
