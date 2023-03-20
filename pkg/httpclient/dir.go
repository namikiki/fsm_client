package httpclient

import (
	"encoding/json"
	"log"

	"fsm_client/pkg/ent"
)

func (c *Client) DirCreate(dir *ent.Dir) error {
	log.Println(dir.Dir)
	res, err := c.deserialization("POST", "/dir", dir)
	if err != nil {
		return err
	}

	return json.Unmarshal(res, dir)
}

func (c *Client) DirDelete(dir ent.Dir) error {
	_, err := c.deserialization("DELETE", "/dir", dir)
	return err
}

func (c *Client) GetAllDirBySyncID(syncID string) ([]ent.Dir, error) {
	var dirs []ent.Dir
	res, err := c.deserialization("GET", "/dirs/sid/"+syncID, nil)
	if err != nil {
		return nil, err
	}

	return dirs, json.Unmarshal(res, &dirs)
}

func (c *Client) DirRename(dir ent.Dir) error {
	_, err := c.deserialization("PUT", "/dir/name", dir)
	return err
}
