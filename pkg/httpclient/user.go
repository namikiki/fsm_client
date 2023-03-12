package httpclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
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

	var res types.ApiResult
	if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return err
	}

	var lr types.LoginRes

	json.Unmarshal(res.Data, &lr)

	c.JWT = lr.Token
	c.UserID = lr.UserID
	c.HttpClient = newHttpClient(lr.Token, c.Conf.Device.ClientID)

	return err

}

func (c *Client) WebSocketConnect() (*websocket.Conn, error) {
	// 使用 HTTP 客户端与 WebSocket 服务器建立连接
	headers := make(http.Header)
	headers.Set("authorization", c.JWT)
	headers.Set("clientID", c.ClientID)

	dial, resp, err := websocket.DefaultDialer.Dial(c.WebsocketUrl+"/websocket/connect", headers)
	//log.Println(resp.Body)
	all, _ := io.ReadAll(resp.Body)
	log.Println(string(all))

	return dial, err
}

func (c *Client) LoginOut() {

}
