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

func TestFile(t *testing.T) {
	client := Init()
	//file := ent.File{
	//	ID:          "7679a111-cc4b-4718-ab73-13d1a92cd058",
	//	SyncID:      "ab7dfaa0-721c-47de-95d1-3fbb2af7a7ad",
	//	Name:        "file2",
	//	ParentDirID: "cb52b2d6-14c5-4869-8395-abd3e618e801",
	//	Level:       2,
	//	Hash:        "a9c7dbdc61936620ff3204326f8065a6-1",
	//	Size:        24,
	//	Deleted:     false,
	//	CreateTime:  time.Now(),
	//	ModTime:     time.Now(),
	//}

	//t.Run("Test file delete", func(t *testing.T) {
	//	err := client.FileDelete(file)
	//	log.Println(err)
	//})

	t.Run("Test file update", func(t *testing.T) {
		fileIO, err := os.Open("/Users/zylzyl/go/src/fsm_client/pkg/mock/testfile")
		if err != nil {
			log.Println(err)
		}
		file := mock.NewFile()

		client.FileUpdate(&file, fileIO)
	})

}

func TestFIleRename(t *testing.T) {
	client := Init()

	file := mock.NewFile()
	fileIO, err := os.Open("/Users/zylzyl/go/src/fsm_client/pkg/mock/包子")
	if err != nil {
		log.Println(err)
		return
	}

	err = client.FileCreate(&file, fileIO)
	if err != nil {
		log.Println(err)
		return
	}

	file.Name = "小学博士"
	err = client.FileRename(file)
	if err != nil {
		log.Println(err)
		return
	}
}
