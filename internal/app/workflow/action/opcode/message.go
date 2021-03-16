package opcode

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/workflow/action/inside"
	"github.com/tsundata/assistant/internal/pkg/utils"
)

type Message struct{}

func NewMessage() *Message {
	return &Message{}
}

func (o *Message) Type() int {
	return TypeOp
}

func (o *Message) Doc() string {
	return "message : (any -> bool)"
}

func (o *Message) Run(ctx *inside.Context, _ []interface{}) (interface{}, error) {
	if ctx.MsgClient == nil {
		return false, nil
	}
	if ctx.Value == nil {
		return false, nil
	}

	var text string
	if str, ok := ctx.Value.(string); ok {
		text = str
	}
	if num, ok := ctx.Value.(int64); ok {
		text = fmt.Sprintf("%d", num)
	}
	if objects, ok := ctx.Value.(map[string]interface{}); ok {
		b, err := json.Marshal(objects)
		if err != nil {
			return false, nil
		}
		text = utils.ByteToString(b)
	}
	if arrays, ok := ctx.Value.([]interface{}); ok {
		b, err := json.Marshal(arrays)
		if err != nil {
			return false, nil
		}
		text = utils.ByteToString(b)
	}

	state, err := ctx.MsgClient.Send(context.Background(), &pb.MessageRequest{Text: text})
	if err != nil {
		return false, err
	}
	ctx.SetValue(state.GetState())
	return state.GetState(), nil
}