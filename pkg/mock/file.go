package mock

import (
	"time"

	"fsm_client/pkg/ent"
)

func NewFile() ent.File {
	return ent.File{
		ID:          "6caffa80-a1a3-4d94-8397-1557a69f17e6",
		SyncID:      "ab7dfaa0-721c-47de-95d1-3fbb2af7a7ad",
		Name:        "surflabom",
		ParentDirID: "cb52b2d6-14c5-4869-8395-abd3e618e801",
		Level:       2,
		Hash:        "",
		Size:        123,
		Deleted:     false,
		CreateTime:  time.Now(),
		ModTime:     time.Now(),
	}
}
