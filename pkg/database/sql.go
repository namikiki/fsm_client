package database

//createDBDir = "CREATE TABLE %s (id TEXT,sync_id TEXT,dir TEXT,level INTEGER,deleted numeric,create_time datetime,mod_time datetime)"

var (
	dropDirSql  = "drop table db_dir"
	dropFileSql = "drop table db_file"

	createDBDir = `create table db_dir
    (
    id          TEXT,
    sync_id     TEXT,
    dir         TEXT,
    level       INTEGER,
    deleted     numeric,
    create_time INTEGER,
    mod_time    INTEGER
);`
	createDriveDir = `create table drive_dir
    (
    id          TEXT,
    sync_id     TEXT,
    dir         TEXT,
    level       INTEGER,
    deleted     numeric,
    create_time INTEGER,
    mod_time    INTEGER
);`

	createCloudDir = `create table cloud_dir
    (
    id          TEXT,
    sync_id     TEXT,
    dir         TEXT,
    level       INTEGER,
    deleted     numeric,
    create_time INTEGER,
    mod_time    INTEGER
);`

	createDBFile = `create table db_file
(
    id            TEXT,
    sync_id       TEXT,
    name          TEXT,
    parent_dir_id TEXT,
    level         INTEGER,
    hash          TEXT,
    size          INTEGER,
    deleted       numeric,
    create_time   INTEGER,
    mod_time      INTEGER
);`

	createDriveFile = `create table drive_file
(
    id            TEXT,
    sync_id       TEXT,
    name          TEXT,
    parent_dir_id TEXT,
    level         INTEGER,
    hash          TEXT,
    size          INTEGER,
    deleted       numeric,
    create_time   INTEGER,
    mod_time      INTEGER
);`

	createCloudFile = `create table cloud_file
(
    id            TEXT,
    sync_id       TEXT,
    name          TEXT,
    parent_dir_id TEXT,
    level         INTEGER,
    hash          TEXT,
    size          INTEGER,
    deleted       numeric,
    create_time   INTEGER,
    mod_time      INTEGER
);`
)
