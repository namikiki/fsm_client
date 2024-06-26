package httpclient

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"fsm_client/pkg/ent"
	"fsm_client/pkg/types"

	"github.com/google/go-querystring/query"
)

func (c *Client) FileCreate(file *ent.File, fileIO io.ReadCloser) error {

	defer fileIO.Close()
	values, _ := query.Values(file)

	request, _ := http.NewRequest("POST", c.BaseUrl+"/file?"+values.Encode(), fileIO)
	resp, err := c.HttpClient.Do(request)
	if err != nil {
		return err
	}

	var res types.ApiResult
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return err
	}

	if res.Code >= 500 {
		return errors.New(res.Message)
	}

	return json.Unmarshal(res.Data, file)
}

func (c *Client) GetFile(fileID string) (io.ReadCloser, error) {
	request, _ := http.NewRequest("GET", c.BaseUrl+"/file/"+fileID, nil)
	resp, err := c.HttpClient.Do(request)
	return resp.Body, err
}

func (c *Client) GetAllFileBySyncID(syncID string) ([]ent.File, error) {
	var files []ent.File
	res, err := c.deserialization("GET", "/files/"+syncID, nil)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &files); err != nil {
		return nil, err
	}
	return files, nil
}

func (c *Client) FileDelete(file ent.File) error {
	_, err := c.deserialization("DELETE", "/file", file)
	return err
}

func (c *Client) FileUpdate(file *ent.File, fileIO io.ReadCloser) error {
	defer fileIO.Close()
	values, _ := query.Values(file)

	request, _ := http.NewRequest("PUT", c.BaseUrl+"/file?"+values.Encode(), fileIO)
	resp, err := c.HttpClient.Do(request)
	if err != nil {
		return err
	}

	var res types.ApiResult
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return err
	}

	if res.Code >= 500 {
		return errors.New(res.Message)
	}

	return json.Unmarshal(res.Data, file)
}

func (c *Client) FileRename(file ent.File) error {
	_, err := c.deserialization("PUT", "/file/name", file)
	return err
}
