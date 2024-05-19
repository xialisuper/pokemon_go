package api

import (
	"net/http"
	"pokemoncli/cache"
	"time"
)

type Client struct {
	httpClient *http.Client
	cache      cache.Cache
}

// NewClient creates a new Client with the given timeout and cache interval.
func NewClient(timeout, cacheInterval time.Duration) Client {

	return Client{
		httpClient: &http.Client{
			Timeout: timeout,
		},
		cache: cache.NewCache(cacheInterval),
	}
}
