package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"time"

	"fsm_client/pkg/types"
)

func NewHttpClient(jwt, client string) *http.Client {
	return &http.Client{
		Transport: MyRoundTripper{r: http.DefaultTransport, JWT: jwt, Client: client},
		Timeout:   time.Second * 20,
	}
}

type MyRoundTripper struct {
	r      http.RoundTripper
	JWT    string
	Client string
}

func (mrt MyRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	r.Header.Add("authorization", mrt.JWT)
	r.Header.Add("client", mrt.Client)
	return mrt.r.RoundTrip(r)
}

type Client struct {
	Conf   *types.Config
	H      *http.Client
	DL     map[string]struct{}
	BSU    string // base URL
	WSU    string // ws URL
	UserID string
}

func NewClient(bus, wsu string, conf *types.Config) *Client { // todo
	return &Client{
		H:    nil,
		DL:   map[string]struct{}{},
		BSU:  bus,
		WSU:  wsu,
		Conf: conf,
	}
}

func (c *Client) deserialization(method string, url string, d interface{}) ([]byte, error) {
	marshal, _ := json.Marshal(d)

	request, _ := http.NewRequest(method, c.BSU+url, bytes.NewBuffer(marshal))
	resp, err := c.H.Do(request)
	if err != nil {
		return nil, err
	}

	var res types.ApiResult
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		log.Println("aaaaaaaaaa", err)
		return nil, err
	}

	if res.Code >= 500 {
		return nil, errors.New(res.Message)
	}

	return res.Data, nil

}

func (c *Client) TestClient() {
	resp, err := c.H.Get(c.BSU + "/test")
	if err != nil {
		log.Println(err)
	}

	all, _ := io.ReadAll(resp.Body)
	log.Println(string(all))
}
