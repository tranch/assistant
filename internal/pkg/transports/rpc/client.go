package rpc

import (
	"context"
	"time"

	"github.com/tsundata/rpc/xclient"
)

type ClientOptions struct {
	Wait     time.Duration
	Tag      string
	Registry string
}

func NewClientOptions() (*ClientOptions, error) {
	var (
		err error
		o   = new(ClientOptions)
	)
	return o, err
}

type ClientOptional func(o *ClientOptions)

func WithTimeout(d time.Duration) ClientOptional {
	return func(o *ClientOptions) {
		o.Wait = d
	}
}

type Client struct {
	o *ClientOptions
}

func NewClient(o *ClientOptions) (*Client, error) {
	return &Client{
		o: o,
	}, nil
}

func (c *Client) Dial(service string, options ...ClientOptional) (*xclient.XClient, error) {
	d := xclient.NewRegistryDiscovery(c.o.Registry, 0)
	xc := xclient.NewXClient(d, xclient.RoundRobinSelect, nil)
	// TODO
	defer func() { _ = xc.Close() }()
	var reply interface{}
	err := xc.Call(context.Background(), service, nil, &reply)
	return xc, err
}
