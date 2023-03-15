package httpclient

import (
	"log"
	"testing"

	"fsm_client/pkg/mock"
)

func TestWebsocket(t *testing.T) {
	client := Init()
	connect, err := client.WebSocketConnect()
	if err != nil {
		log.Println(err)
	}
	log.Println(connect.RemoteAddr())
}

func TestRegister(t *testing.T) {
	client := Init()
	user := mock.NewRegis()
	err := client.Register(user)
	if err != nil {
		log.Println(err)
	}
}
