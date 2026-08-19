package pokeapi

import (
	"fmt"
	"encoding/json"
	"io"
)

func (c *Client) ListLocationAreas(url string) (APIResourceList, error) {
	var apiResource APIResourceList
	var data []byte
	data, ok := c.cacheStruct.Get(url)
	if !ok {
		resp, err := c.client.Get(url)
		if err != nil {
			return apiResource, fmt.Errorf("error when getting next url: %s", err)
		}
		info, err := io.ReadAll(resp.Body)
		if err != nil {
			return apiResource, fmt.Errorf("error when io readeing from resp.Body: %s", err)
		}
		data = info
		c.cacheStruct.Add(url, data)
		resp.Body.Close()
	}
	err := json.Unmarshal(data, &apiResource)
	if err != nil {
		return apiResource, fmt.Errorf("error when unmarshaling data into location: %s", err)
	}
	return apiResource, nil
}