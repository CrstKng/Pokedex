package pokeapi

import (
	"net/http"
	"time"
	"github.com/CrstKng/pokedexcli/internal/pokecache"
)

type Client struct {
	client *http.Client
	cacheStruct *pokecache.Cache
}

func NewClient(cacheInterval, timeout time.Duration) *Client{
	return &Client{
		cacheStruct : pokecache.NewCache(cacheInterval),
		client : &http.Client{
			Timeout: timeout,
		},
	}
}