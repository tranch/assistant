package message

import (
	"github.com/spf13/viper"
	"github.com/tsundata/assistant/internal/app/message/service"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/transports/rpc"
)

type Options struct {
	Name string
	v    *viper.Viper
}

func NewOptions(v *viper.Viper) (*Options, error) {
	var err error
	o := new(Options)
	o.v = v

	return o, err
}

func NewApp(o *Options, rs *rpc.Server) (*app.Application, error) {
	var slack service.Slack
	slack.V = o.v
	rs.Register(&slack)

	a, err := app.New(o.Name, app.RpcServerOption(rs))

	if err != nil {
		return nil, err
	}

	return a, nil
}
