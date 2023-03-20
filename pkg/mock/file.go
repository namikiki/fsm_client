package mock

import (
	"time"

	"fsm_client/pkg/ent"
)

func NewFile() ent.File {
	return ent.File{
		SyncID:      "123",
		ParentDirID: "123",
		Name:        "123",
		Level:       123,
		Deleted:     false,
		CreateTime:  time.Now(),
		ModTime:     time.Now(),
	}
}
