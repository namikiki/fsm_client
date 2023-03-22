package database

import (
	"context"
	"log"

	"fsm_client/pkg/ent"

	"github.com/jmoiron/sqlx"
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

func memDBMigrate(db *sqlx.DB) {
	db.MustExec(createDBDir)
	db.MustExec(createDriveDir)
	db.MustExec(createCloudDir)

	db.MustExec(createDBFile)
	db.MustExec(createDriveFile)
	db.MustExec(createCloudFile)
}

func ResetTable(db *sqlx.DB) {
	db.MustExec(dropSql)
	db.MustExec(createDBDir)
}

func NewSqliteMemoryDB() *sqlx.DB {
	//"file::memory:?cache=shared"
	db, err := sqlx.Connect("sqlite3", "test.db")
	if err != nil {
		log.Fatalln(err)
	}

	memDBMigrate(db)
	ResetTable(db)
	return db
}

//func CreateTable(tableName string) {
//	createTabSql := fmt.Sprintf(createDBDir, tableName)
//
//}

//var files []ent.File
//if err = gm.Select(&files, "SELECT * FROM files"); err != nil {
//log.Println(err)
//}
//log.Println(files)
//
//res, err := test.NamedExec(`INSERT INTO local_files (id, sync_id, name,parent_dir_id,level,hash,size,deleted,create_time,mod_time)
//        VALUES (:id, :sync_id, :name, :parent_dir_id, :level, :hash, :size,:deleted,:create_time,:mod_time)`, files)
