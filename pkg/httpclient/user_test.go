package httpclient

import (
	"log"
	"testing"
)

func TestWebsocket(t *testing.T) {
	client := Init()
	connect, err := client.WebSocketConnect()
	if err != nil {
		log.Println(err)
	}
	log.Println(connect.RemoteAddr())
}
