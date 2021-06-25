// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package collection

import (
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/logger"
	"github.com/tsundata/assistant/internal/pkg/middleware/consul"
	redis2 "github.com/tsundata/assistant/internal/pkg/middleware/redis"
)

// Injectors from wire.go:

func CreateRedisClient(id string) (*redis.Client, error) {
	client, err := consul.New()
	if err != nil {
		return nil, err
	}
	appConfig := config.NewConfig(id, client)
	redisClient, err := redis2.New(appConfig)
	if err != nil {
		return nil, err
	}
	return redisClient, nil
}

// wire.go:

var testProviderSet = wire.NewSet(logger.ProviderSet, config.ProviderSet, consul.ProviderSet, redis2.ProviderSet)