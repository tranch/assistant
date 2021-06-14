package task

import (
	"github.com/RichardKnop/machinery/v2"
	"github.com/google/wire"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/task/service"
	"github.com/tsundata/assistant/internal/app/task/work"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/logger"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc"
	"google.golang.org/grpc"
)

func NewApp(c *config.AppConfig, bus *event.Bus, logger *logger.Logger, rs *rpc.Server, q *machinery.Server, client *rpc.Client) (*app.Application, error) {
	// service
	s := service.NewTask(q)
	err := rs.Register(func(gs *grpc.Server) error {
		pb.RegisterTaskServer(gs, s)
		return nil
	})
	if err != nil {
		return nil, err
	}

	a, err := app.New(c, logger, app.RPCServerOption(rs))
	if err != nil {
		return nil, err
	}

	// worker
	go func() {
		workflowTask := work.NewWorkflowTask(bus, client)
		err = q.RegisterTasks(map[string]interface{}{
			"run":  workflowTask.Run,
		})
		if err != nil {
			logger.Error(err)
			return
		}

		worker := q.NewWorker(c.Name, 0)
		worker.SetErrorHandler(func(err error) {
			logger.Error(err)
		})
		err = worker.Launch()
		if err != nil {
			logger.Error(err)
			return
		}
	}()

	return a, nil
}

var ProviderSet = wire.NewSet(NewApp)
