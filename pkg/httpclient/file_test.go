package httpclient

import (
	"io"
	"log"
	"os"
	"testing"

	"fsm_client/pkg/mock"
)

func TestFileCreate(t *testing.T) {
	client := Init()
	file := mock.NewFile()
	fileIO, err := os.Open("/Users/zylzyl/go/src/fsm_client/pkg/mock/testfile")
	if err != nil {
		log.Println(err)
	}

	err = client.FileCreate(&file, fileIO)
	log.Println(err)
}

func TestGetFile(t *testing.T) {
	client := Init()

	fileIO, err := client.GetFile("8639a95a-4a9e-4594-ba7c-2294e0020473")
	if err != nil {
		log.Println(err)
	}

	file, err := os.Create("download")
	if err != nil {
		log.Println(err)
	}

	io.Copy(file, fileIO)

	file.Close()
	fileIO.Close()

}

func TestGetAllFileBySyncID(t *testing.T) {
	client := Init()
	file := mock.NewFile()

	syncIDs, err := client.GetAllFileBySyncID(file.SyncID)
	if err != nil {
		log.Println(err)
	}
	log.Println(syncIDs)

}
