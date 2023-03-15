package handle

import (
	"fsm_client/pkg/ent"
)

func (h *Handle) SyncTaskCreate(synctask ent.SyncTask) {
	h.DB.Create(&synctask)
}

func (h *Handle) SyncTaskDelete(synctask ent.SyncTask) {
	h.DB.Delete(&synctask)
}
