package database

//createDBDir = "CREATE TABLE %s (id TEXT,sync_id TEXT,dir TEXT,level INTEGER,deleted numeric,create_time datetime,mod_time datetime)"

var (
	dropSql = "drop table db_dir"

	createDBDir = `create table db_dir
    (
    id          TEXT,
    sync_id     TEXT,
    dir         TEXT,
    level       INTEGER,
    deleted     numeric,
    create_time datetime,
    mod_time    datetime
);`
	createDriveDir = `create table drive_dir
    (
    id          TEXT,
    sync_id     TEXT,
    dir         TEXT,
    level       INTEGER,
    deleted     numeric,
    create_time datetime,
    mod_time    datetime
);`

	createCloudDir = `create table cloud_dir
    (
    id          TEXT,
    sync_id     TEXT,
    dir         TEXT,
    level       INTEGER,
    deleted     numeric,
    create_time datetime,
    mod_time    datetime
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
    create_time   datetime,
    mod_time      datetime
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
    create_time   datetime,
    mod_time      datetime
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
    create_time   datetime,
    mod_time      datetime
);`
)
