package mock

import (
	"time"

	"fsm_client/pkg/ent"
)

func NewDir() ent.Dir {
	return ent.Dir{
		SyncID:     "2312",
		Dir:        "test1231231",
		Level:      12312,
		Deleted:    false,
		CreateTime: time.Now(),
		ModTime:    time.Now(),
	}
}
