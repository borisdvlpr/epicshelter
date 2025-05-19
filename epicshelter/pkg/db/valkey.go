package cache

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/valkey-io/valkey-go"

	"github.com/borisdvlpr/epicshelter/pkg/config"
)

type Client struct {
	client valkey.Client
	ttl    int
}

func NewClient(cfg *config.Config) (*Client, error) {
	client, err := valkey.NewClient(valkey.ClientOption{
		InitAddress: []string{cfg.Host},
		Password:    cfg.Password,
		SelectDB:    cfg.Db,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to create Valkey client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	timeout := client.B().Ping().Build()
	if err := client.Do(ctx, timeout).Error(); err != nil {
		return nil, fmt.Errorf("failed to connect to Valkey: %v", err)
	}

	log.Printf("Connected to Valkey.")
	return &Client{client: client, ttl: cfg.TTL}, nil
}

func (c *Client) Get(ctx context.Context, key string) ([]byte, error) {
	cmd := c.client.B().Get().Key(key).Build()
	return c.client.Do(ctx, cmd).AsBytes()
}

func (c *Client) Set(ctx context.Context, key string, value []byte) error {
	cmd := c.client.B().Set().Key(key).Value(string(value)).Px(time.Duration(c.ttl) * time.Second).Build()
	return c.client.Do(ctx, cmd).Error()
}

func (c *Client) Exists(ctx context.Context, key string) (bool, error) {
	cmd := c.client.B().Exists().Key(key).Build()
	return c.client.Do(ctx, cmd).AsBool()
}

func (c *Client) Close() {
	c.client.Close()
}
