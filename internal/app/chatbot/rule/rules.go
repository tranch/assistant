package rule

import (
	"context"
	"fmt"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/chatbot/trigger/tags"
	"github.com/tsundata/assistant/internal/pkg/rulebot"
	"github.com/tsundata/assistant/internal/pkg/util"
	"github.com/tsundata/assistant/internal/pkg/version"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var rules = []Rule{
	{
		Define: "version",
		Help:   `Version info`,
		Parse: func(ctx context.Context, comp rulebot.IComponent, tokens []*Token) []string {
			return []string{version.Info()}
		},
	},
	{
		Define: `menu`,
		Help:   `Show menu`,
		Parse: func(ctx context.Context, comp rulebot.IComponent, tokens []*Token) []string {
			if comp.Middle() == nil {
				return []string{"empty client"}
			}
			reply, err := comp.Middle().GetMenu(ctx, &pb.TextRequest{})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}

			if reply.GetText() == "" {
				return []string{"empty menu"}
			}

			return []string{reply.GetText()}
		},
	},
	{
		Define: `qr [string]`,
		Help:   `Generate QR code`,
		Parse: func(ctx context.Context, comp rulebot.IComponent, tokens []*Token) []string {
			if comp.Middle() == nil {
				return []string{"empty client"}
			}
			if len(tokens) != 2 {
				return []string{"error args"}
			}

			txt := tokens[1].Value
			reply, err := comp.Middle().GetQrUrl(ctx, &pb.TextRequest{
				Text: txt,
			})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}

			return []string{
				reply.GetText(),
			}
		},
	},
	{
		Define: `ut [number]`,
		Help:   `Unix Timestamp`,
		Parse: func(ctx context.Context, comp rulebot.IComponent, tokens []*Token) []string {
			if len(tokens) != 2 {
				return []string{"error args"}
			}

			tt, err := strconv.ParseInt(tokens[1].Value, 10, 64)
			if err != nil {
				return []string{"error call: " + err.Error()}
			}

			t := time.Unix(tt, 0)

			return []string{
				t.String(),
			}
		},
	},
	{
		Define: `rand [number] [number]`,
		Help:   `Unix Timestamp`,
		Parse: func(ctx context.Context, comp rulebot.IComponent, tokens []*Token) []string {
			if len(tokens) != 3 {
				return []string{"error args"}
			}

			minArg := tokens[1].Value
			maxArg := tokens[2].Value
			min, err := strconv.Atoi(minArg)
			if err != nil {
				return []string{"error call: " + err.Error()}
			}
			max, err := strconv.Atoi(maxArg)
			if err != nil {
				return []string{"error call: " + err.Error()}
			}

			rand.Seed(time.Now().Unix())
			t := rand.Intn(max-min) + min

			return []string{
				strconv.Itoa(t),
			}
		},
	},
	{
		Define: `pwd [number]`,
		Help:   `Generate Password`,
		Parse: func(ctx context.Context, comp rulebot.IComponent, tokens []*Token) []string {
			if len(tokens) != 2 {
				return []string{"error args"}
			}

			lenArg := tokens[1].Value
			length, err := strconv.Atoi(lenArg)
			if err != nil {
				return []string{"error call: " + err.Error()}
			}

			pwd := util.GeneratePassword(length, "lowercase|uppercase|numbers")

			return []string{
				pwd,
			}
		},
	},
	{
		Define: "subs list",
		Help:   `List subscribe`,
		Parse: func(ctx context.Context, comp rulebot.IComponent, tokens []*Token) []string {
			if comp.Middle() == nil {
				return []string{"empty client"}
			}
			reply, err := comp.Middle().ListSubscribe(ctx, &pb.SubscribeRequest{})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}

			if reply.GetText() == nil {
				return []string{"empty subscript"}
			}

			return reply.GetText()
		},
	},
	{
		Define: "subs open [string]",
		Help:   `Open subscribe`,
		Parse: func(ctx context.Context, comp rulebot.IComponent, tokens []*Token) []string {
			if comp.Middle() == nil {
				return []string{"empty client"}
			}
			if len(tokens) != 3 {
				return []string{"error args"}
			}

			reply, err := comp.Middle().OpenSubscribe(ctx, &pb.SubscribeRequest{
				Text: tokens[2].Value,
			})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}
			if reply.GetState() {
				return []string{"ok"}
			}

			return []string{"failed"}
		},
	},
	{
		Define: `subs close [string]`,
		Help:   `Close subscribe`,
		Parse: func(ctx context.Context, comp rulebot.IComponent, tokens []*Token) []string {
			if comp.Middle() == nil {
				return []string{"empty client"}
			}
			if len(tokens) != 3 {
				return []string{"error args"}
			}

			reply, err := comp.Middle().CloseSubscribe(ctx, &pb.SubscribeRequest{
				Text: tokens[2].Value,
			})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}
			if reply.GetState() {
				return []string{"ok"}
			}

			return []string{"failed"}
		},
	},
	{
		Define: `view [number]`,
		Help:   `View message`,
		Parse: func(ctx context.Context, comp rulebot.IComponent, tokens []*Token) []string {
			if comp.Message() == nil {
				return []string{"empty client"}
			}
			if len(tokens) != 2 {
				return []string{"error args"}
			}

			id, err := strconv.ParseInt(tokens[1].Value, 10, 64)
			if err != nil {
				return []string{"error args"}
			}
			messageReply, err := comp.Message().Get(ctx, &pb.MessageRequest{Message: &pb.Message{Id: id}})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}

			if messageReply.Message.Id == 0 {
				return []string{"no message"}
			}

			return []string{messageReply.Message.GetText()}
		},
	},
	{
		Define: `run [number]`,
		Help:   `Run message`,
		Parse: func(ctx context.Context, comp rulebot.IComponent, tokens []*Token) []string {
			if comp.Message() == nil {
				return []string{"empty client"}
			}
			if len(tokens) != 2 {
				return []string{"error args"}
			}

			id, err := strconv.ParseInt(tokens[1].Value, 10, 64)
			if err != nil {
				return []string{"error args"}
			}

			reply, err := comp.Message().Run(ctx, &pb.MessageRequest{Message: &pb.Message{Id: id}})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}

			return []string{reply.GetText()}
		},
	},
	{
		Define: `doc`,
		Help:   `Show action docs`,
		Parse: func(ctx context.Context, comp rulebot.IComponent, tokens []*Token) []string {
			if comp.Workflow() == nil {
				return []string{"empty client"}
			}
			reply, err := comp.Workflow().ActionDoc(ctx, &pb.WorkflowRequest{})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}
			res := strings.Builder{}
			res.WriteString("Action:\n")
			res.WriteString(reply.GetText())
			res.WriteString("\n\nTag:\n")
			for k := range tags.Tags() {
				res.WriteString("#")
				res.WriteString(k)
				res.WriteString("\n")
			}
			return []string{res.String()}
		},
	},
	{
		Define: `test`,
		Help:   `Test`,
		Parse: func(ctx context.Context, comp rulebot.IComponent, tokens []*Token) []string {
			_, err := comp.Message().Send(ctx, &pb.MessageRequest{Message: &pb.Message{Text: "test"}})
			if err != nil {
				return []string{err.Error()}
			}
			_, err = comp.Middle().GetSubscribeStatus(ctx, &pb.SubscribeRequest{Text: "example"})
			if err != nil {
				return []string{err.Error()}
			}
			_, err = comp.NLP().Pinyin(ctx, &pb.TextRequest{Text: "测试"})
			if err != nil {
				return []string{err.Error()}
			}
			_, err = comp.Todo().GetTodo(ctx, &pb.TodoRequest{Todo: &pb.Todo{Id: 1}})
			if err != nil {
				return []string{err.Error()}
			}
			_, err = comp.User().GetUser(ctx, &pb.UserRequest{User: &pb.User{Id: 1}})
			if err != nil {
				return []string{err.Error()}
			}
			_, err = comp.Workflow().SyntaxCheck(ctx, &pb.WorkflowRequest{Text: "echo 1"})
			if err != nil {
				return []string{err.Error()}
			}
			return []string{"test done"}
		},
	},
	{
		Define: `stats`,
		Help:   `Stats Info`,
		Parse: func(ctx context.Context, comp rulebot.IComponent, tokens []*Token) []string {
			if comp.Middle() == nil {
				return []string{"empty client"}
			}
			reply, err := comp.Middle().GetStats(ctx, &pb.TextRequest{})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}
			return []string{reply.GetText()}
		},
	},
	{
		Define: `todo [string]`,
		Help:   "Todo something",
		Parse: func(ctx context.Context, comp rulebot.IComponent, tokens []*Token) []string {
			if comp.Todo() == nil {
				return []string{"empty client"}
			}
			if len(tokens) != 2 {
				return []string{"error args"}
			}
			reply, err := comp.Todo().CreateTodo(ctx, &pb.TodoRequest{
				Todo: &pb.Todo{Content: tokens[1].Value},
			})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}
			if !reply.GetState() {
				return []string{"failed"}
			}
			return []string{"success"}
		},
	},
	{
		Define: `role`,
		Help:   "Role info",
		Parse: func(ctx context.Context, comp rulebot.IComponent, tokens []*Token) []string {
			if comp.Middle() == nil {
				return []string{"empty client"}
			}
			reply, err := comp.Middle().GetRoleImageUrl(ctx, &pb.TextRequest{})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}
			return []string{reply.GetText()}
		},
	},
	{
		Define: `pinyin [string]`,
		Help:   "chinese pinyin conversion",
		Parse: func(ctx context.Context, comp rulebot.IComponent, tokens []*Token) []string {
			if comp.NLP() == nil {
				return []string{"empty client"}
			}
			if len(tokens) != 2 {
				return []string{"error args"}
			}
			reply, err := comp.NLP().Pinyin(ctx, &pb.TextRequest{Text: tokens[1].Value})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}
			if len(reply.GetText()) <= 0 {
				return []string{"failed"}
			}
			return []string{strings.Join(reply.GetText(), ", ")}
		},
	},
	{
		Define: `remind [string] [string]`,
		Help:   `Remind something`,
		Parse: func(ctx context.Context, comp rulebot.IComponent, tokens []*Token) []string {
			if len(tokens) != 3 {
				return []string{"error args"}
			}

			arg1 := tokens[1].Value
			arg2 := tokens[2].Value
			fmt.Println(arg1, arg2) // todo remind message

			return []string{}
		},
	},
	{
		Define: `del [number]`,
		Help:   `Delete message`,
		Parse: func(ctx context.Context, comp rulebot.IComponent, tokens []*Token) []string {
			if len(tokens) != 2 {
				return []string{"error args"}
			}

			idStr := tokens[1].Value
			id, err := strconv.ParseInt(idStr, 10, 64)
			if err != nil {
				return []string{"error call: " + err.Error()}
			}

			_, err = comp.Message().Delete(ctx, &pb.MessageRequest{Message: &pb.Message{Id: id}})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}

			return []string{fmt.Sprintf("Deleted %d", id)}
		},
	},
	{
		Define: "cron list",
		Help:   `List cron`,
		Parse: func(ctx context.Context, comp rulebot.IComponent, tokens []*Token) []string {
			if comp.Middle() == nil {
				return []string{"empty client"}
			}
			reply, err := comp.Middle().ListCron(ctx, &pb.CronRequest{})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}

			if reply.GetText() == nil {
				return []string{"empty subscript"}
			}

			return reply.GetText()
		},
	},
	{
		Define: "cron start [string]",
		Help:   `Start cron`,
		Parse: func(ctx context.Context, comp rulebot.IComponent, tokens []*Token) []string {
			if comp.Middle() == nil {
				return []string{"empty client"}
			}
			if len(tokens) != 3 {
				return []string{"error args"}
			}

			reply, err := comp.Middle().StartCron(ctx, &pb.CronRequest{
				Text: tokens[2].Value,
			})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}
			if reply.GetState() {
				return []string{"ok"}
			}

			return []string{"failed"}
		},
	},
	{
		Define: `cron stop [string]`,
		Help:   `Stop cron`,
		Parse: func(ctx context.Context, comp rulebot.IComponent, tokens []*Token) []string {
			if comp.Middle() == nil {
				return []string{"empty client"}
			}
			if len(tokens) != 3 {
				return []string{"error args"}
			}

			reply, err := comp.Middle().StopCron(ctx, &pb.CronRequest{
				Text: tokens[2].Value,
			})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}
			if reply.GetState() {
				return []string{"ok"}
			}

			return []string{"failed"}
		},
	},
}

var Options = []rulebot.Option{
	rulebot.RegisterRuleset(New(rules)),
}
