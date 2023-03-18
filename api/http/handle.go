package httpapi

import (
	"fsm_client/pkg/sync"

	"gorm.io/gorm"
)

type Handle struct {
	DB   *gorm.DB
	Sync *sync.Syncer
}

func New(sync *sync.Syncer, db *gorm.DB) Handle {
	return Handle{Sync: sync, DB: db}
}
