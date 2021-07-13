package opcode

import (
	"context"
	"errors"
	"github.com/go-resty/resty/v2"
	"github.com/tsundata/assistant/internal/app/workflow/action/inside"
	"github.com/tsundata/assistant/internal/pkg/util"
	"time"
)

type Get struct{}

func NewGet() *Get {
	return &Get{}
}

func (o *Get) Type() int {
	return TypeOp
}

func (o *Get) Doc() string {
	return "get [any] : (nil -> any)"
}

func (o *Get) Run(_ context.Context, comp *inside.Component, params []interface{}) (interface{}, error) {
	if len(params) != 1 {
		return nil, errors.New("error params")
	}
	if text, ok := params[0].(string); ok {
		if util.IsUrl(text) {
			client := resty.New()
			client.SetTimeout(time.Minute)
			resp, err := client.R().Get(text)
			if err != nil {
				return nil, err
			}
			result := util.ByteToString(resp.Body())
			comp.SetValue(result)
			return result, nil
		}
	}
	return nil, nil
}
