package httpclient

import (
	"testing"
)

func TestA(t *testing.T) {
	client := Init()
	client.TestClient()
	//client := NewClient("http://127.0.0.1:8080", "ws://127.0.0.1:8080")

	//user := types.Register{
	//	Email:    "374856123123@gmail.com",
	//	PassWord: "12312312312312",
	//	UserName: "zzzzzzz",
	//}
	//client.Register(user)

	//login := types.UserLoginReq{
	//	Email:    "374856123123@gmail.com",
	//	PassWord: "12312312312312",
	//}
	//client.Login(login)

}
