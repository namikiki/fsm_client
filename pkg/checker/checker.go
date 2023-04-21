package checker

import (
	"fmt"

	"fsm_client/pkg/ent"
	"fsm_client/pkg/handle"
	"fsm_client/pkg/httpclient"
	"fsm_client/pkg/ignore"

	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
)

type Checker struct {
	MemDB  *sqlx.DB
	DB     *gorm.DB
	Client *httpclient.Client
	//Ignore     ignore.Ignore
	Handle *handle.Handle
}

func NewChecker(memdb *sqlx.DB, db *gorm.DB, client *httpclient.Client, handle *handle.Handle, ignore *ignore.Ignore) *Checker {
	return &Checker{
		MemDB:  memdb,
		DB:     db,
		Client: client,
		Handle: handle,
		//Ignore:     ignore,
	}
}

const (
	DBFile = "db_file"
	DBDir  = "db_dir"

	DriveFile = "drive_file"
	DriveDir  = "drive_dir"

	cloudFile = "cloud_file"
	cloudDir  = "cloud_dir"
)

var (
	insertFile = "INSERT INTO %s (id, sync_id, name,parent_dir_id,level,hash,size,deleted,create_time,mod_time)" +
		" VALUES (:id, :sync_id, :name, :parent_dir_id, :level, :hash, :size,:deleted,:create_time,:mod_time)"

	insertDir = "INSERT INTO %s (id, sync_id, dir, level, deleted, create_time ,mod_time)" +
		" VALUES (:id, :sync_id, :dir, :level, :deleted, :create_time, :mod_time)"

	GetDirChange = "SELECT %s.* FROM %s LEFT JOIN %s ON %s.dir = %s.dir and %s.sync_id = %s.sync_id WHERE %s.dir IS NULL;"

	GetFileChange = func(lt, rt string) string {
		return fmt.Sprintf("SELECT %s.* FROM %s LEFT JOIN %s ON %s.name = %s.name and %s.sync_id = %s.sync_id and %s.level = %s.level  WHERE %s.name IS NULL;",
			lt, lt, rt, lt, rt, lt, rt, lt, rt, rt)
	}

	getFileUpdate = func(lt, rt string) string {
		return fmt.Sprintf("SELECT %s.* FROM %s LEFT JOIN %s ON %s.name = %s.name and %s.sync_id = %s.sync_id and %s.level = %s.level  WHERE  %s.mod_time != %s.mod_time;",
			lt, lt, rt, lt, rt, lt, rt, lt, rt, lt, rt)
	}
)

func (c *Checker) insertDirs(tableName string, dirs []ent.Dir) error {
	insertSql := fmt.Sprintf(insertDir, tableName)
	_, err := c.MemDB.NamedExec(insertSql, dirs)
	return err
}

func (c *Checker) insertDir(tableName string, dir ent.Dir) error {
	insertSql := fmt.Sprintf(insertDir, tableName)
	_, err := c.MemDB.NamedExec(insertSql, &dir)
	return err
}

func (c *Checker) insertFiles(tableName string, files []ent.File) error {
	insertSql := fmt.Sprintf(insertFile, tableName)
	_, err := c.MemDB.NamedExec(insertSql, files)
	return err
}

func (c *Checker) insertFile(tableName string, file ent.File) error {
	insertSql := fmt.Sprintf(insertFile, tableName)
	_, err := c.MemDB.NamedExec(insertSql, &file)
	return err
}
