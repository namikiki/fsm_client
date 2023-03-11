package http

import (
	"encoding/json"

	"fsm_client/pkg/ent"
)

func (c *Client) DirCreate(dir *ent.Dir) error {
	res, err := c.deserialization("POST", "/dir/create", dir)
	if err != nil {
		return err
	}

	return json.Unmarshal(res, &dir)
}

func (c *Client) DirDelete(dir ent.Dir) error {
	_, err := c.deserialization("GET", "/dir/delete", dir)
	return err
}
