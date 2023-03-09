package main

import (
	"log"
	"os"

	"fsm_client/pkg/client/http"
	"fsm_client/pkg/database"
	"fsm_client/pkg/sync"
)

func main() {

	client := http.NewCustomHttpClient()
	connect := database.NewGormSQLiteConnect()
	syncer := sync.NewSyncer(client, connect, os.Args[1])

	go syncer.WebSocketConn()

	err := syncer.CreateSync("filetest", "/Users/zylzyl/Desktop/GolangProjects/fsm/test/filetest")
	if err != nil {
		log.Println(err)
	}

	select {}
}
