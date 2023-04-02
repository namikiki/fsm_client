package types

type NewSyncTask struct {
	Name   string `json:"name"`
	Path   string `json:"path"`
	Type   string `json:"type"`
	Ignore bool   `json:"ignore"`
}

type RecSyncTask struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Path   string `json:"path"`
	Ignore bool   `json:"ignore"`
}

type DeleteSyncTask struct {
	ID       string `json:"id"`
	DelLocal bool   `json:"del_local"`
	DelCloud bool   `json:"del_cloud"`
}
