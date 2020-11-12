package main

import (
	"flag"
	"github.com/tsundata/assistant/internal/app/gateway"
	"github.com/tsundata/assistant/internal/app/gateway/controllers"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/logger"
	"github.com/tsundata/assistant/internal/pkg/transports/http"
	"github.com/tsundata/assistant/internal/pkg/transports/rpc"
)

func CreateApp(cf string) (*app.Application, error) {
	viper, err := config.New(cf)
	if err != nil {
		return nil, err
	}
	httpOptions, err := http.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	clientOptions, err := rpc.NewClientOptions(viper)
	if err != nil {
		return nil, err
	}
	subClient, err := rpc.NewClient(clientOptions, "subscribe", "Subscribe")
	if err != nil {
		return nil, err
	}
	msgClient, err := rpc.NewClient(clientOptions, "message", "Message")
	if err != nil {
		return nil, err
	}
	gatewayOptions, err := gateway.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	log := logger.NewLogger()
	gatewayController := controllers.NewGatewayController(gatewayOptions, log, subClient, msgClient)
	initControllers := controllers.CreateInitControllersFn(gatewayController)
	router := http.NewRouter(httpOptions, initControllers)
	server, err := http.New(httpOptions, router)
	if err != nil {
		return nil, err
	}
	application, err := gateway.NewApp(gatewayOptions, log, server)
	if err != nil {
		return nil, err
	}
	return application, nil
}

var configFile = flag.String("f", "gateway.yml", "set config file which will loading")

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
