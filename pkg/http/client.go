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

	"github.com/gorilla/websocket"
)

func NewCustomHttpClient(jwt, client string) *http.Client {
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

func (c *Client) deserialization(r *http.Request) (interface{}, error) {

	resp, err := c.H.Do(r)
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

func (c *Client) Register(user types.UserRegister) {
	marshal, _ := json.Marshal(user)
	resp, err := http.Post(c.BSU+"/register", "application/json", bytes.NewBuffer(marshal))
	if err != nil {
		log.Println(err)
	}
	all, _ := io.ReadAll(resp.Body)
	log.Println(string(all))
}

func (c *Client) Login(user types.UserLoginReq) error {
	marshal, _ := json.Marshal(user)

	resp, err := http.Post(c.BSU+"/login", "application/json", bytes.NewBuffer(marshal))
	if err != nil {
		return err
	}

	var res types.ApiResult
	json.NewDecoder(resp.Body).Decode(&res)

	c.UserID = res.Data.UserID
	c.H = NewCustomHttpClient(res.Data.Token, c.Conf.Device.ClientID)

	log.Println(res.Data.UserID)
	return err

}

func (c *Client) LoginOut() {

}

func (c *Client) TestClient() {
	resp, err := c.H.Get(c.BSU + "/test")
	if err != nil {
		log.Println(err)
	}

	all, _ := io.ReadAll(resp.Body)
	log.Println(string(all))
}

func (c *Client) WebSocketConnect(clientID string) (*websocket.Conn, error) {
	// 使用 HTTP 客户端与 WebSocket 服务器建立连接
	dial, _, err := websocket.DefaultDialer.Dial(c.WSU, nil)
	return dial, err
}

func (c *Client) CreateDir() {

}

func (c *Client) DeleteDir() {

}
