package registry

import (
	"errors"
	"github.com/spf13/viper"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/transports/rpc"
	"log"
)

type Options struct {
	Name string
}

func NewOptions(v *viper.Viper) (*Options, error) {
	var err error
	o := new(Options)

	if err = v.UnmarshalKey("app", o); err != nil {
		return nil, errors.New("unmarshal app option error")
	}

	log.Println("load application options success")

	return o, err
}

func NewApp(o *Options, rs *rpc.Registry) (*app.Application, error) {
	a, err := app.New(o.Name, app.RegistryServerOption(rs))

	if err != nil {
		return nil, err
	}

	return a, nil
}
