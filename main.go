package main

import (
	"fsm_client/pkg/config"
	"fsm_client/pkg/http"
	"fsm_client/pkg/types"
)

func main() {

	cfg, _ := config.ReadConfigFile()
	client := http.NewClient("http://127.0.0.1:8080", "ws://127.0.0.1:8080ß", cfg)

	//userReg := types.UserRegister{
	//	Email:    "1231231231@gmail.com",
	//	PassWord: "1231der232342423",
	//	UserName: "zylzyl",
	//}
	//
	//client.Register(userReg)

	userLog := types.UserLoginReq{
		Email:    "1231231231@gmail.com",
		PassWord: "1231der232342423",
	}

	err := client.Login(userLog)
	if err != nil {
		return
	}

	client.TestClient()
	//client := http.NewCustomHttpClient("123", "232")
	//connect := database.NewGormSQLiteConnect()
	//syncer := sync.NewSyncer(client, connect, os.Args[1])
	//
	//go syncer.WebSocketConn()
	//
	//err := syncer.GetSyncTaskToDownload("6fd35b63-95fb-46b8-a7fa-394ccec20b01", "/Users/zylzyl/Desktop/GolangProjects/fsm/test/filetest")
	//if err != nil {
	//	log.Println(err)
	//}

	//err := syncer.CreateSync("filetest", "/Users/zylzyl/Desktop/GolangProjects/fsm/test/filetest")
	//if err != nil {
	//	log.Println(err)
	//}
	//
	//select {}
}
