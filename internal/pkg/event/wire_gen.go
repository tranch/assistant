// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package event

import (
	"github.com/google/wire"
	"github.com/rabbitmq/amqp091-go"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/log"
	"github.com/tsundata/assistant/internal/pkg/middleware/etcd"
	"github.com/tsundata/assistant/internal/pkg/middleware/rabbitmq"
)

// Injectors from wire.go:

func CreateRabbitmq(id string) (*amqp091.Connection, error) {
	client, err := etcd.New()
	if err != nil {
		return nil, err
	}
	appConfig := config.NewConfig(id, client)
	connection, err := rabbitmq.New(appConfig)
	if err != nil {
		return nil, err
	}
	return connection, nil
}

// wire.go:

var testProviderSet = wire.NewSet(log.ProviderSet, config.ProviderSet, etcd.ProviderSet, ProviderSet, rabbitmq.ProviderSet)
