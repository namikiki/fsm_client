package mock

import (
	"time"

	"fsm_client/pkg/ent"
)

func NewSyncTask() ent.SyncTask {
	return ent.SyncTask{
		//ID:         "",
		//UserID:     "",
		Type:       "two",
		Name:       "media",
		RootDir:    "/dawdawdaw",
		Deleted:    false,
		CreateTime: time.Now(),
	}
}
