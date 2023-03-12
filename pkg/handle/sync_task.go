package handle

import (
	"encoding/json"

	"fsm_client/pkg/ent"
)

func (h *Handle) SyncTaskCreate(data []byte) {
	var synctask ent.SyncTask
	json.Unmarshal(data, &synctask)
	h.DB.Create(&synctask)
}

func (h *Handle) SyncTaskDelete(data []byte) {
	var synctask ent.SyncTask
	json.Unmarshal(data, &synctask)
	h.DB.Delete(&synctask)
}
