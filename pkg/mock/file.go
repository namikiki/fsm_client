package mock

import (
	"time"

	"fsm_client/pkg/ent"
)

func NewFile() ent.File {
	return ent.File{
		//ID:          "",
		SyncID:      "fielsync1",
		Name:        "file1",
		ParentDirID: "dawdawdawd",
		Level:       1,
		Hash:        "",
		Size:        123123,
		Deleted:     false,
		CreateTime:  time.Now(),
		ModTime:     time.Now(),
	}
}
