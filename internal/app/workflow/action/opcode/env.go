package opcode

import (
	"context"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/workflow/action/inside"
)

type Env struct{}

func NewEnv() *Env {
	return &Env{}
}

func (o *Env) Type() int {
	return TypeCond
}

func (o *Env) Doc() string {
	return "env [string] : (nil -> string)"
}

func (o *Env) Run(ctx context.Context, comp *inside.Component, params []interface{}) (interface{}, error) {
	if len(params) != 1 {
		return nil, nil
	}

	if comp.Middle == nil {
		return nil, nil
	}

	if text, ok := params[0].(string); ok {
		reply, err := comp.Middle.GetSetting(ctx, &pb.TextRequest{Text: text})
		if err != nil {
			return nil, err
		}
		comp.SetValue(reply.GetValue())
		return reply.GetValue(), nil
	}

	return nil, nil
}
