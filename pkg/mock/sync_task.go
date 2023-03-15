package mock

import (
	"time"

	"fsm_client/pkg/ent"
)

func NewSyncTask() ent.SyncTask {
	return ent.SyncTask{
		Type:       "two",
		Name:       "mediawa",
		RootDir:    "/dawdawdaw",
		Deleted:    false,
		CreateTime: time.Now(),
	}
}
func df() {

}
