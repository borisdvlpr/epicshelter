package service

import (
	"context"
	"fmt"

	cache "github.com/borisdvlpr/epicshelter/pkg/db"
)

var ErrCacheConn = fmt.Errorf("unable to connect to server")
var ErrCacheMiss = fmt.Errorf("requested key not found")

type ApiService struct {
	cache *cache.Client
}

func NewApiService(cache *cache.Client) *ApiService {
	return &ApiService{
		cache: cache,
	}
}

func (s *ApiService) GetCache(ctx context.Context, key string) ([]byte, error) {
	if s.cache == nil {
		return nil, ErrCacheConn
	}

	data, err := s.cache.Get(ctx, key)
	if err != nil {
		return nil, ErrCacheMiss
	}

	return data, nil
}

func (s *ApiService) SetCache(ctx context.Context, key string, data []byte) error {
	if s.cache == nil {
		return ErrCacheConn
	}

	return s.cache.Set(ctx, key, data)
}
