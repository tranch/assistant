// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package rulebot

import (
	"github.com/google/wire"
	"github.com/tsundata/assistant/internal/app/chatbot/bot/finance/service"
	repository3 "github.com/tsundata/assistant/internal/app/chatbot/bot/okr/repository"
	service4 "github.com/tsundata/assistant/internal/app/chatbot/bot/okr/service"
	repository2 "github.com/tsundata/assistant/internal/app/chatbot/bot/system/repository"
	service3 "github.com/tsundata/assistant/internal/app/chatbot/bot/system/service"
	"github.com/tsundata/assistant/internal/app/chatbot/bot/todo/repository"
	service2 "github.com/tsundata/assistant/internal/app/chatbot/bot/todo/service"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/global"
	"github.com/tsundata/assistant/internal/pkg/log"
	"github.com/tsundata/assistant/internal/pkg/middleware/etcd"
	"github.com/tsundata/assistant/internal/pkg/middleware/jaeger"
	"github.com/tsundata/assistant/internal/pkg/middleware/mysql"
	"github.com/tsundata/assistant/internal/pkg/middleware/rabbitmq"
	"github.com/tsundata/assistant/internal/pkg/middleware/redis"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc/rpcclient"
	"github.com/tsundata/assistant/internal/pkg/vendors/newrelic"
	"github.com/tsundata/assistant/internal/pkg/vendors/rollbar"
)

// Injectors from wire.go:

func CreateRuleBot(id string) (*RuleBot, error) {
	client, err := etcd.New()
	if err != nil {
		return nil, err
	}
	appConfig := config.NewConfig(id, client)
	connection, err := rabbitmq.New(appConfig)
	if err != nil {
		return nil, err
	}
	rollbarRollbar := rollbar.New(appConfig)
	logger := log.NewZapLogger(rollbarRollbar)
	logLogger := log.NewAppLogger(logger)
	bus := event.NewRabbitmqBus(connection, logLogger)
	app, err := newrelic.New(appConfig, logger)
	if err != nil {
		return nil, err
	}
	redisClient, err := redis.New(appConfig, app)
	if err != nil {
		return nil, err
	}
	configuration, err := jaeger.NewConfiguration(appConfig, logLogger)
	if err != nil {
		return nil, err
	}
	tracer, err := jaeger.New(configuration)
	if err != nil {
		return nil, err
	}
	clientOptions, err := rpc.NewClientOptions(tracer)
	if err != nil {
		return nil, err
	}
	rpcClient, err := rpc.NewClient(clientOptions, appConfig, logLogger)
	if err != nil {
		return nil, err
	}
	messageSvcClient, err := rpcclient.NewMessageClient(rpcClient)
	if err != nil {
		return nil, err
	}
	chatbotSvcClient, err := rpcclient.NewChatbotClient(rpcClient)
	if err != nil {
		return nil, err
	}
	middleSvcClient, err := rpcclient.NewMiddleClient(rpcClient)
	if err != nil {
		return nil, err
	}
	storageSvcClient, err := rpcclient.NewStorageClient(rpcClient)
	if err != nil {
		return nil, err
	}
	userSvcClient, err := rpcclient.NewUserClient(rpcClient)
	if err != nil {
		return nil, err
	}
	financeSvcServer := service.NewFinance()
	locker := global.NewLocker(client)
	idSvcClient, err := rpcclient.NewIdClient(rpcClient)
	if err != nil {
		return nil, err
	}
	globalID := global.NewID(appConfig, idSvcClient)
	conn, err := mysql.New(appConfig)
	if err != nil {
		return nil, err
	}
	todoRepository := repository.NewMysqlTodoRepository(locker, globalID, conn)
	todoSvcServer := service2.NewTodo(bus, logLogger, todoRepository)
	systemRepository := repository2.NewMysqlSystemRepository(logLogger, globalID, conn)
	systemSvcServer := service3.NewSystem(systemRepository, logLogger, locker)
	okrRepository := repository3.NewMysqlOkrRepository(locker, globalID, conn)
	okrSvcServer := service4.NewOkr(bus, okrRepository, middleSvcClient)
	componentComponent := component.NewComponent(appConfig, bus, redisClient, logLogger, messageSvcClient, chatbotSvcClient, middleSvcClient, storageSvcClient, userSvcClient, financeSvcServer, todoSvcServer, systemSvcServer, okrSvcServer)
	ruleBot := New(componentComponent)
	return ruleBot, nil
}

// wire.go:

var testProviderSet = wire.NewSet(log.ProviderSet, config.ProviderSet, etcd.ProviderSet, ProviderSet, rpcclient.ProviderSet, redis.ProviderSet, rollbar.ProviderSet, rpc.ProviderSet, jaeger.ProviderSet, newrelic.ProviderSet, component.ProviderSet, event.ProviderSet, rabbitmq.ProviderSet, global.ProviderSet, mysql.ProviderSet, service4.ProviderSet, repository.ProviderSet, service2.ProviderSet, repository2.ProviderSet, service3.ProviderSet, repository3.ProviderSet, service.ProviderSet)
