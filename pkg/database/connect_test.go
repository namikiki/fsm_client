package database

import (
	"fsm_client/pkg/ent"
	"log"
	"testing"
)

func TestT1(t *testing.T) {
	NewSqliteMemoryDB()
}

func TestT2(t *testing.T) {
	connect := NewGormSQLiteConnect()

	var results []map[string]interface{}
	connect.Model(&ent.SyncTask{}).Find(&results)

	for i, result := range results {
		log.Println(i, result)
	}

}
