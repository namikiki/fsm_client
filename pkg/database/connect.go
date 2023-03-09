package database

import (
	"context"
	"log"

	"fsm_client/pkg/ent"

	_ "github.com/mattn/go-sqlite3"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func EntConnect() *ent.Client {
	conn, err := ent.Open("sqlite3", "file:ent.db?mode=rwc&cache=shared&_fk=1&_cache_size=20000")
	if err != nil {
		log.Println(err)
	}

	// 自动迁移 ent/schema 对象结构
	if err := conn.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return conn
}

func NewGormSQLiteConnect() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(&ent.Dir{}, &ent.File{}, &ent.SyncTask{}); err != nil {
		panic(err)
	}

	return db
}
