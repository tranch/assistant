package main

import (
	"flag"
	"github.com/tsundata/assistant/internal/app/spider"
	"github.com/tsundata/assistant/internal/app/spider/rpcclients"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/etcd"
	"github.com/tsundata/assistant/internal/pkg/jaeger"
	"github.com/tsundata/assistant/internal/pkg/logger"
	"github.com/tsundata/assistant/internal/pkg/redis"
	"github.com/tsundata/assistant/internal/pkg/transports/rpc"
)

func CreateApp(cf string) (*app.Application, error) {
	viper, err := config.New(cf)
	if err != nil {
		return nil, err
	}
	log := logger.NewLogger()

	t, err := jaeger.NewConfiguration(viper, log)
	if err != nil {
		return nil, err
	}
	j, err := jaeger.New(t)
	if err != nil {
		return nil, err
	}

	redisOption, err := redis.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	rdb, err := redis.New(redisOption)
	if err != nil {
		return nil, err
	}

	etcdOption, err := etcd.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	e, err := etcd.New(etcdOption)
	if err != nil {
		return nil, err
	}

	clientOptions, err := rpc.NewClientOptions(viper, j)
	if err != nil {
		return nil, err
	}
	client, err := rpc.NewClient(clientOptions, e)
	if err != nil {
		return nil, err
	}
	msgClient, err := rpcclients.NewMessageClient(client)
	if err != nil {
		return nil, err
	}
	subClient, err := rpcclients.NewSubscribeClient(client)
	if err != nil {
		return nil, err
	}
	midClient, err := rpcclients.NewMiddleClient(client)
	if err != nil {
		return nil, err
	}

	appOptions, err := spider.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	application, err := spider.NewApp(appOptions, rdb, log, msgClient, midClient, subClient)
	if err != nil {
		return nil, err
	}
	return application, nil
}

var configFile = flag.String("f", "spider.yml", "set config file which will loading")

func main() {
	flag.Parse()

	a, err := CreateApp(*configFile)
	if err != nil {
		panic(err)
	}

	if err := a.Start(); err != nil {
		panic(err)
	}

	a.AwaitSignal()
}