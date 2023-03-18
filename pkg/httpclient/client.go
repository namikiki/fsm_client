package httpclient

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

func newHttpClient(jwt, client string) *http.Client {
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
	r.Header.Add("clientID", mrt.Client)
	return mrt.r.RoundTrip(r)
}

type Client struct {
	Conf         *types.Config
	HttpClient   *http.Client
	DL           map[string]struct{}
	BaseUrl      string // base URL
	WebsocketUrl string // ws URL
	UserID       string
	ClientID     string
	JWT          string
}

func NewClient(conf *types.Config) *Client { // todo
	log.Println("clientID = ", conf.Device.ClientID)

	return &Client{
		HttpClient:   nil,
		DL:           map[string]struct{}{},
		BaseUrl:      conf.Server.BaseUrl,
		WebsocketUrl: conf.Server.WebSocketUrl,
		Conf:         conf,
		ClientID:     conf.Device.ClientID,
	}
}

func (c *Client) deserialization(method string, url string, d interface{}) ([]byte, error) {
	marshal, _ := json.Marshal(d)

	request, _ := http.NewRequest(method, c.BaseUrl+url, bytes.NewBuffer(marshal))
	resp, err := c.HttpClient.Do(request)
	if err != nil {
		return nil, err
	}

	var res types.ApiResult
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}

	if res.Code >= 500 {
		return nil, errors.New(res.Message)
	}

	return res.Data, nil

}

func (c *Client) TestClient() {
	resp, err := c.HttpClient.Get(c.BaseUrl + "/test")
	log.Println(resp.StatusCode)
	if err != nil {
		log.Println(err)
	}

	all, _ := io.ReadAll(resp.Body)
	log.Println(string(all))
}
