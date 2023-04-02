package httpclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fsm_client/pkg/sec"
	"log"
	"net/http"

	"fsm_client/pkg/types"

	"github.com/gorilla/websocket"
)

func (c *Client) Register(user types.UserRegister) error {
	marshal, _ := json.Marshal(user)
	resp, err := http.Post(c.BaseUrl+"/register", "application/json", bytes.NewBuffer(marshal))
	if err != nil {
		return err
	}

	var res types.ApiResult
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return err
	}

	if res.Code >= 500 {
		return errors.New(res.Message)
	}
	return err
}

func (c *Client) Login(user types.UserLoginReq) error {
	marshal, _ := json.Marshal(user)

	resp, err := http.Post(c.BaseUrl+"/login", "application/json", bytes.NewBuffer(marshal))
	if err != nil {
		return err
	}

	return c.loginRes(resp)
}

func (c *Client) LoginByJWT() error {
	jwt, err := sec.ReadJWT()
	if err != nil {
		log.Println(err)
		return err
	}

	request, _ := http.NewRequest("POST", c.BaseUrl+"/jwt", nil)
	request.Header.Add("jwt", string(jwt))
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}

	return c.loginRes(resp)
}

func (c *Client) loginRes(resp *http.Response) error {

	var res types.ApiResult
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return err
	}
	defer resp.Body.Close()

	var lr types.LoginRes
	if err := json.Unmarshal(res.Data, &lr); err != nil || lr.UserID == "" || lr.Token == "" {
		return errors.New("登陆失败" + err.Error())
	}

	if err := sec.SaveJWT(lr.Token); err != nil {
		return err
	}

	c.JWT = lr.Token
	c.UserID = lr.UserID
	c.HttpClient = newHttpClient(lr.Token, c.Conf.Device.ClientID)
	c.Ch <- 1
	c.Ch <- 1
	return nil
}

func (c *Client) WebSocketConnect() (*websocket.Conn, error) {
	// 使用 HTTP 客户端与 WebSocket 服务器建立连接
	headers := make(http.Header)
	headers.Set("authorization", c.JWT)
	headers.Set("clientID", c.ClientID)

	dial, resp, err := websocket.DefaultDialer.Dial(c.WebsocketUrl+"/websocket/connect", headers)
	if resp.StatusCode != 101 {
		log.Println("websocket connect fail")
	}
	log.Println("websocket connect success")
	return dial, err
}
