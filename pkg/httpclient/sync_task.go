package httpclient

import (
	"encoding/json"

	"fsm_client/pkg/ent"
)

func (c *Client) SyncTaskCreate(task *ent.SyncTask) error {
	resp, err := c.deserialization("POST", "/synctask/create", task)
	if err != nil {
		return err
	}
	return json.Unmarshal(resp, task)
}

func (c *Client) SyncTaskDelete(syncID string) error {
	//var sts []ent.SyncTask
	_, err := c.deserialization("GET", "/synctask/delete/"+syncID, nil)
	return err
}

func (c *Client) SyncTaskGetAll() ([]ent.SyncTask, error) {
	var sts []ent.SyncTask
	resp, err := c.deserialization("GET", "/synctask/getAll", nil)
	if err != nil {
		return nil, err
	}

	return sts, json.Unmarshal(resp, &sts)
}
