package pokeapi

import (
	"fmt"
	"encoding/json"
	"io"
)

func (c *Client) PokemonInfo(url string) (Pokemon, error) {
	var Pokemon Pokemon
	var data []byte
	data, ok := c.cacheStruct.Get(url)
	if !ok {
		resp, err := c.client.Get(url)
		if err != nil {
			return Pokemon, fmt.Errorf("error when getting url: %s", err)
		}
		info, err := io.ReadAll(resp.Body)
		if err != nil {
			return Pokemon, fmt.Errorf("error when io readeing from resp.Body: %s", err)
		}
		data = info
		c.cacheStruct.Add(url, data)
		resp.Body.Close()
	}
	err := json.Unmarshal(data, &Pokemon)
	if err != nil {
		return Pokemon, fmt.Errorf("error when unmarshaling data into location: %s", err)
	}
	return Pokemon, nil
}