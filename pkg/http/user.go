package http

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"fsm_client/pkg/types"

	"github.com/gorilla/websocket"
)

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
	if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return err
	}

	var lr types.LoginRes

	json.Unmarshal(res.Data, &lr)

	c.UserID = lr.UserID
	c.H = NewHttpClient(lr.Token, c.Conf.Device.ClientID)

	return err

}

func (c *Client) WebSocketConnect(clientID string) (*websocket.Conn, error) {
	// 使用 HTTP 客户端与 WebSocket 服务器建立连接
	dial, _, err := websocket.DefaultDialer.Dial(c.WSU, nil)
	return dial, err
}

func (c *Client) LoginOut() {

}
