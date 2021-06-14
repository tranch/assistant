package config

import (
	"errors"
	"fmt"
	"github.com/google/wire"
	"github.com/hashicorp/consul/api"
	"github.com/tsundata/assistant/internal/pkg/util"
	"gopkg.in/yaml.v2"
	"os"
	"strings"
)

type AppConfig struct {
	kv *api.KV

	Name string `json:"name"`

	Http    Http    `json:"http"`
	Rpc     Rpc     `json:"rpc"`
	Web     Web     `json:"web"`
	Gateway Gateway `json:"gateway"`
	Plugin  Plugin  `json:"plugin"`
	Storage Storage `json:"storage"`

	Mysql    Mysql    `json:"mysql"`
	Redis    Redis    `json:"redis"`
	Influx   Influx   `json:"influx"`
	Rabbitmq Rabbitmq `json:"rabbitmq"`
	Jaeger   Jaeger   `json:"jaeger"`
	Nats     Nats     `json:"nats"`

	Slack    Slack    `json:"slack"`
	Rollbar  Rollbar  `json:"rollbar"`
	Telegram Telegram `json:"telegram"`
}

func NewConfig(consul *api.Client) *AppConfig {
	kv := consul.KV()
	configNamespace := os.Getenv("CONFIG_NAMESPACE")
	if configNamespace == "" {
		panic("config namespace error")
	}
	pair, _, err := kv.Get(fmt.Sprintf("config/%s", configNamespace), nil)
	if err != nil {
		panic(err)
	}

	// Watch todo

	var xc AppConfig
	err = yaml.Unmarshal(pair.Value, &xc)
	if err != nil {
		panic(err)
	}
	xc.kv = kv

	return &xc
}

func (c *AppConfig) GetConfig(key string) (string, error) {
	result, _, err := c.kv.Get(fmt.Sprintf("config/%s", key), nil)
	if err != nil {
		return "", err
	}
	if result != nil {
		return util.ByteToString(result.Value), nil
	}
	return "", errors.New("result error")
}

func (c *AppConfig) GetSetting(key string) (string, error) {
	result, _, err := c.kv.Get(fmt.Sprintf("setting/%s", key), nil)
	if err != nil {
		return "", err
	}
	if result != nil {
		return util.ByteToString(result.Value), nil
	}
	return "", errors.New("result error")
}

func (c *AppConfig) SetSetting(key, value string) error {
	_, err := c.kv.Put(&api.KVPair{
		Key:   fmt.Sprintf("setting/%s", key),
		Value: util.StringToByte(value),
	}, nil)
	return err
}

func (c *AppConfig) GetSettings() (map[string]string, error) {
	kvs, _, err := c.kv.List("setting", nil)
	if err != nil {
		return nil, err
	}
	result := make(map[string]string)
	for _, ev := range kvs {
		result[strings.ReplaceAll(ev.Key, "setting/", "")] = util.ByteToString(ev.Value)
	}
	return result, nil
}

var ProviderSet = wire.NewSet(NewConfig)
