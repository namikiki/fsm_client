package mock

import (
	"time"

	"fsm_client/pkg/ent"
)

func NewDir() ent.Dir {
	return ent.Dir{
		ID:         "763f106a-392c-4a69-ba82-b51812f64be6",
		SyncID:     "ab7dfaa0-721c-47de-95d1-3fbb2af7a7ad",
		Dir:        "/synctest",
		Level:      233,
		Deleted:    false,
		CreateTime: time.Now().Unix(),
		ModTime:    time.Now().Unix(),
	}
}
