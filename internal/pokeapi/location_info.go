package pokeapi

import (
	"fmt"
	"encoding/json"
	"io"
)

func (c *Client) LocationArea(url string) (LocationArea, error) {
	var locationArea LocationArea
	var data []byte
	data, ok := c.cacheStruct.Get(url)
	if !ok {
		resp, err := c.client.Get(url)
		if err != nil {
			return locationArea, fmt.Errorf("error when getting url: %s", err)
		}
		info, err := io.ReadAll(resp.Body)
		if err != nil {
			return locationArea, fmt.Errorf("error when io readeing from resp.Body: %s", err)
		}
		data = info
		c.cacheStruct.Add(url, data)
		resp.Body.Close()
	}
	err := json.Unmarshal(data, &locationArea)
	if err != nil {
		return locationArea, fmt.Errorf("error when unmarshaling data into location: %s", err)
	}
	return locationArea, nil
}