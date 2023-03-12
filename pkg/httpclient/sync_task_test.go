package httpclient

import (
	"log"
	"testing"

	"fsm_client/pkg/mock"
)

func TestSyncTaskCreate(t *testing.T) {
	client := Init()
	task := mock.NewSyncTask()

	err := client.SyncTaskCreate(&task)
	if err != nil {
		log.Println(err)
	}
	log.Println(task)

}

func TestSyncTaskDelete(t *testing.T) {
	init := Init()
	syncID := "6081ae73-9c9e-4028-90d4-b30624b80072"

	err := init.SyncTaskDelete(syncID)
	log.Println(err)
}

func BenchmarkName(b *testing.B) {
	client := Init()
	for i := 0; i < b.N; i++ {
		_, err := client.SyncTaskGetAll()
		if err != nil {
			panic(err)
		}
	}
}

func TestSyncTaskGetAll(t *testing.T) {
	client := Init()

	syncTasks, err := client.SyncTaskGetAll()
	if err != nil {
		log.Println()
	}
	log.Println(syncTasks)

}
