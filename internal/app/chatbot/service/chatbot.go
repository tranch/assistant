package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/influxdata/cron"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/chatbot/repository"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/log"
	"github.com/tsundata/assistant/internal/pkg/robot"
	"github.com/tsundata/assistant/internal/pkg/robot/action"
	"github.com/tsundata/assistant/internal/pkg/robot/action/opcode"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
	"github.com/tsundata/assistant/internal/pkg/robot/rulebot"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc/exception"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc/md"
	"github.com/tsundata/assistant/internal/pkg/util"
	"gorm.io/gorm"
	"strings"
	"time"
)

type Chatbot struct {
	logger  log.Logger
	bus     event.Bus
	rdb     *redis.Client
	bot     *rulebot.RuleBot
	repo    repository.ChatbotRepository
	message pb.MessageSvcClient
	middle  pb.MiddleSvcClient
	comp    component.Component
}

func NewChatbot(
	logger log.Logger,
	bus event.Bus,
	rdb *redis.Client,
	repo repository.ChatbotRepository,
	message pb.MessageSvcClient,
	middle pb.MiddleSvcClient,
	bot *rulebot.RuleBot,
	comp component.Component) *Chatbot {
	return &Chatbot{
		logger:  logger,
		bus:     bus,
		rdb:     rdb,
		bot:     bot,
		repo:    repo,
		message: message,
		middle:  middle,
		comp:    comp,
	}
}

func (s *Chatbot) Handle(ctx context.Context, payload *pb.ChatbotRequest) (*pb.ChatbotReply, error) {
	reply, err := s.message.Get(ctx, &pb.MessageRequest{Message: &pb.Message{Id: payload.MessageId, Uuid: payload.MessageUuid}})
	if err != nil {
		return nil, err
	}

	r := robot.NewRobot()

	// touch updated
	err = s.repo.TouchGroupUpdatedAt(ctx, reply.Message.GroupId)
	if err != nil {
		return nil, err
	}

	// group bots
	groupBots, err := s.repo.ListGroupBot(ctx, reply.Message.GetGroupId())
	if err != nil {
		return nil, err
	}
	groupBotsMap := make(map[string]struct{})
	for _, bot := range groupBots {
		groupBotsMap[bot.Identifier] = struct{}{}
	}

	// help
	outMessages, err := r.Help(groupBots, reply.Message.GetText())
	if err != nil {
		return nil, err
	}

	if len(outMessages) == 0 {
		// bot settings
		groupSetting, err := s.repo.GetGroupSetting(ctx, reply.Message.GetGroupId())
		if err != nil {
			return nil, err
		}
		groupBotSetting, err := s.repo.GetGroupBotSettingByGroup(ctx, reply.Message.GetGroupId())
		if err != nil {
			return nil, err
		}

		fmt.Println(groupSetting)    // todo
		fmt.Println(groupBotSetting) // todo

		// trigger
		err = r.ProcessTrigger(ctx, s.comp, reply.Message)
		if err != nil {
			return nil, err
		}

		// lexer
		tokens, objects, tags, commands, err := r.ParseText(reply.Message)
		if err != nil {
			return nil, err
		}

		// filter bots
		inBots, err := s.repo.GetBotsByText(ctx, objects)
		if err != nil {
			return nil, err
		}
		for k, item := range inBots {
			if _, ok := groupBotsMap[item.Identifier]; !ok {
				delete(inBots, k)
			}
		}

		if len(commands) > 0 {
			var commandsBots []*pb.Bot
			if len(objects) == 0 {
				commandsBots = groupBots
			} else {
				for k := range inBots {
					commandsBots = append(commandsBots, inBots[k])
				}
			}
			for _, item := range commandsBots {
				for _, commandText := range commands {
					outMessages, err = r.ProcessCommand(ctx, s.comp, item, commandText)
					if err != nil {
						return nil, err
					}
				}
			}
		}

		// tags
		if len(tags) > 0 {
			for _, tag := range tags {
				_, err = s.middle.SaveModelTag(ctx, &pb.ModelTagRequest{
					Model: &pb.ModelTag{
						Service: enum.Message,
						Model:   util.ModelName(pb.Message{}),
						ModelId: reply.Message.GetId(),
					},
					Tag: tag,
				})
				if err != nil {
					s.logger.Error(err)
				}
			}
		}

		if len(outMessages) == 0 {
			// run
			outMessages, err = r.ProcessWorkflow(ctx, s.comp, tokens, inBots)
			if err != nil {
				return nil, err
			}
		}
	}

	// send message
	for botId, messages := range outMessages {
		for _, text := range messages {
			outMessage := &pb.Message{
				GroupId:      reply.Message.GetGroupId(),
				UserId:       reply.Message.GetUserId(),
				Sender:       botId,
				SenderType:   enum.MessageBotType,
				Receiver:     reply.Message.GetUserId(),
				ReceiverType: enum.MessageUserType,
				Type:         enum.MessageTypeText,
				Text:         text,
				Status:       0,
				Direction:    enum.MessageIncomingDirection,
				SendTime:     util.Format(time.Now().Unix()),
			}
			_, err = s.message.Save(ctx, &pb.MessageRequest{Message: outMessage})
			if err != nil {
				return nil, err
			}
			err = s.bus.Publish(ctx, enum.Message, event.MessageChannelSubject, outMessage)
			if err != nil {
				return nil, err
			}
		}
	}

	return &pb.ChatbotReply{
		State: true,
	}, nil
}

func (s *Chatbot) Register(ctx context.Context, request *pb.BotRequest) (*pb.StateReply, error) {
	bot, err := s.repo.GetByIdentifier(ctx, request.Bot.Identifier)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if bot.Id > 0 {
		request.Bot.Id = bot.Id
		err = s.repo.Update(ctx, request.Bot)
		if err != nil {
			return nil, err
		}
		return &pb.StateReply{State: true}, nil
	}
	request.Bot.Uuid = util.UUID()
	_, err = s.repo.Create(ctx, request.Bot)
	if err != nil {
		return nil, err
	}
	return &pb.StateReply{State: true}, nil
}

func (s *Chatbot) GetBot(ctx context.Context, payload *pb.BotRequest) (*pb.BotReply, error) {
	bot, err := s.repo.GetGroupBot(ctx, payload.GroupUuid, payload.BotUuid)
	if err != nil {
		return nil, err
	}
	return &pb.BotReply{Bot: &bot}, nil
}

func (s *Chatbot) GetBots(ctx context.Context, payload *pb.BotsRequest) (*pb.BotsReply, error) {
	bots, err := s.repo.GetBotsByGroupUuid(ctx, payload.GroupUuid)
	if err != nil {
		return nil, err
	}
	return &pb.BotsReply{Bots: bots}, nil
}

func (s *Chatbot) UpdateBotSetting(_ context.Context, _ *pb.BotSettingRequest) (*pb.StateReply, error) {
	return &pb.StateReply{State: true}, nil
}

func (s *Chatbot) GetGroups(ctx context.Context, _ *pb.GroupRequest) (*pb.GetGroupsReply, error) {
	id, _ := md.FromIncoming(ctx)

	// default group
	const defaultGroupName = "System"
	_, err := s.repo.GetGroupByName(ctx, id, defaultGroupName)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		_, err = s.repo.CreateGroup(ctx, &pb.Group{
			Uuid:      util.UUID(),
			UserId:    id,
			Name:      defaultGroupName,
			Avatar:    "",
			CreatedAt: time.Now().Unix(),
			UpdatedAt: time.Now().Unix(),
		})
		if err != nil {
			return nil, err
		}
		// todo group --- system bot
	}

	// bot avatar
	bots, err := s.repo.GetBotsByUser(ctx, id)
	if err != nil {
		return nil, err
	}

	botAvatar := make(map[int64][]string)
	for _, bot := range bots {
		botAvatar[bot.GroupId] = append(botAvatar[bot.GroupId], bot.Avatar)
	}

	groups, err := s.repo.ListGroup(ctx, id)
	if err != nil {
		return nil, err
	}
	var res []*pb.GroupItem
	for _, group := range groups {
		var avatars []string
		if v, ok := botAvatar[group.Id]; ok {
			avatars = v
		}
		// last message
		last, err := s.message.LastByGroup(ctx, &pb.LastByGroupRequest{GroupId: group.Id})
		if err != nil {
			return nil, err
		}
		res = append(res, &pb.GroupItem{
			Sequence:    group.Sequence,
			Type:        group.Type,
			Uuid:        group.Uuid,
			Name:        group.Name,
			Avatar:      group.Avatar,
			UnreadCount: 0, // todo
			LastMessage: &pb.LastMessage{
				LastSender: last.Message.GetSenderName(),
				Content:    last.Message.GetText(),
			},
			BotAvatar: avatars,
		})
	}
	return &pb.GetGroupsReply{Groups: res}, nil
}

func (s *Chatbot) CreateGroup(ctx context.Context, payload *pb.GroupRequest) (*pb.StateReply, error) {
	_, err := s.repo.CreateGroup(ctx, payload.Group)
	if err != nil {
		return nil, err
	}
	return &pb.StateReply{State: true}, nil
}

func (s *Chatbot) GetGroup(ctx context.Context, payload *pb.GroupRequest) (*pb.GetGroupReply, error) {
	id, _ := md.FromIncoming(ctx)
	group, err := s.repo.GetGroupByUUID(ctx, payload.Group.GetUuid())
	if err != nil {
		return nil, err
	}
	if group.UserId != id {
		return nil, exception.ErrGrpcUnauthenticated
	}
	return &pb.GetGroupReply{Group: &pb.GroupItem{
		Id:       group.Id,
		Sequence: group.Sequence,
		Type:     group.Type,
		Uuid:     group.Uuid,
		Name:     group.Name,
		Avatar:   group.Avatar,
	}}, nil
}

func (s *Chatbot) CreateGroupBot(ctx context.Context, payload *pb.GroupBotRequest) (*pb.StateReply, error) {
	err := s.repo.CreateGroupBot(ctx, payload.GroupId, payload.Bot)
	if err != nil {
		return nil, err
	}
	return &pb.StateReply{State: true}, nil
}

func (s *Chatbot) DeleteGroupBot(ctx context.Context, payload *pb.GroupBotRequest) (*pb.StateReply, error) {
	err := s.repo.DeleteGroupBot(ctx, payload.GroupId, payload.Bot.GetId())
	if err != nil {
		return nil, err
	}
	return &pb.StateReply{State: true}, nil
}

func (s *Chatbot) UpdateGroupBotSetting(ctx context.Context, payload *pb.BotSettingRequest) (*pb.StateReply, error) {
	group, err := s.repo.GetGroupByUUID(ctx, payload.GroupUuid)
	if err != nil {
		return nil, err
	}
	bot, err := s.repo.GetByUUID(ctx, payload.BotUuid)
	if err != nil {
		return nil, err
	}
	err = s.repo.UpdateGroupBotSetting(ctx, group.Id, bot.Id, payload.Kvs)
	if err != nil {
		return nil, err
	}
	return &pb.StateReply{State: true}, nil
}

func (s *Chatbot) UpdateGroupSetting(ctx context.Context, payload *pb.GroupSettingRequest) (*pb.StateReply, error) {
	group, err := s.repo.GetGroupByUUID(ctx, payload.GroupUuid)
	if err != nil {
		return nil, err
	}
	err = s.repo.UpdateGroupSetting(ctx, group.Id, payload.Kvs)
	if err != nil {
		return nil, err
	}
	return &pb.StateReply{State: true}, nil
}

func (s *Chatbot) DeleteGroup(ctx context.Context, payload *pb.GroupRequest) (*pb.StateReply, error) {
	err := s.repo.DeleteGroup(ctx, payload.Group.GetId())
	if err != nil {
		return nil, err
	}
	return &pb.StateReply{State: true}, nil
}

func (s *Chatbot) UpdateGroup(ctx context.Context, payload *pb.GroupRequest) (*pb.StateReply, error) {
	err := s.repo.UpdateGroup(ctx, payload.Group)
	if err != nil {
		return nil, err
	}
	return &pb.StateReply{State: true}, nil
}

func (s *Chatbot) GetGroupBotSetting(ctx context.Context, request *pb.BotSettingRequest) (*pb.BotSettingReply, error) {
	kv, err := s.repo.GetGroupBotSettingByUuid(ctx, request.GroupUuid, request.BotUuid)
	if err != nil {
		return nil, err
	}
	return &pb.BotSettingReply{Kvs: kv}, nil
}

func (s *Chatbot) GetGroupSetting(ctx context.Context, request *pb.GroupSettingRequest) (*pb.GroupSettingReply, error) {
	kv, err := s.repo.GetGroupSettingByUuid(ctx, request.GroupUuid)
	if err != nil {
		return nil, err
	}
	return &pb.GroupSettingReply{Kvs: kv}, nil
}

func (s *Chatbot) GetGroupId(ctx context.Context, payload *pb.UuidRequest) (*pb.IdReply, error) {
	group, err := s.repo.GetGroupByUUID(ctx, payload.Uuid)
	if err != nil {
		return nil, err
	}
	return &pb.IdReply{Id: group.Id}, nil
}

func (s *Chatbot) SyntaxCheck(_ context.Context, payload *pb.WorkflowRequest) (*pb.StateReply, error) {
	switch payload.Type {
	case enum.MessageTypeAction:
		if payload.GetText() == "" {
			return nil, errors.New("empty action")
		}
		p, err := action.NewParser(action.NewLexer([]rune(payload.GetText())))
		if err != nil {
			return &pb.StateReply{State: false}, err
		}
		tree, err := p.Parse()
		if err != nil {
			return &pb.StateReply{State: false}, err
		}

		symbolTable := action.NewSemanticAnalyzer()
		err = symbolTable.Visit(tree)
		if err != nil {
			return &pb.StateReply{State: false}, err
		}

		return &pb.StateReply{State: true}, nil
	default:
		return &pb.StateReply{State: false}, nil
	}
}

func (s *Chatbot) RunAction(ctx context.Context, payload *pb.WorkflowRequest) (*pb.WorkflowReply, error) {
	if payload.GetText() == "" {
		return nil, errors.New("empty action")
	}
	p, err := action.NewParser(action.NewLexer([]rune(payload.GetText())))
	if err != nil {
		return nil, err
	}
	tree, err := p.Parse()
	if err != nil {
		return nil, err
	}

	symbolTable := action.NewSemanticAnalyzer()
	err = symbolTable.Visit(tree)
	if err != nil {
		return nil, err
	}

	i := action.NewInterpreter(ctx, tree)
	i.SetComponent(s.bus, s.rdb, s.message, s.middle, s.logger)
	_, err = i.Interpret()
	if err != nil {
		return nil, err
	}

	var result string
	if i.Comp.Debug {
		result = fmt.Sprintf("Tracing\n-------\n %s", i.Stdout())
	}

	return &pb.WorkflowReply{
		Text: result,
	}, nil
}

func (s *Chatbot) WebhookTrigger(ctx context.Context, payload *pb.TriggerRequest) (*pb.WorkflowReply, error) {
	trigger, err := s.repo.GetTriggerByFlag(ctx, payload.Trigger.GetType(), payload.Trigger.GetFlag())
	if err != nil {
		return nil, err
	}

	// Authorization
	if trigger.Secret != "" && payload.Trigger.GetSecret() != trigger.Secret {
		return nil, errors.New("error secret")
	}

	if trigger.MessageId <= 0 {
		return nil, errors.New("error trigger")
	}

	// publish event
	err = s.bus.Publish(ctx, enum.Chatbot, event.WorkflowRunSubject, pb.Message{
		Id: trigger.MessageId,
	})
	if err != nil {
		return nil, err
	}

	return &pb.WorkflowReply{}, nil
}

func (s *Chatbot) CronTrigger(ctx context.Context, _ *pb.TriggerRequest) (*pb.WorkflowReply, error) {
	triggers, err := s.repo.ListTriggersByType(ctx, "cron")
	if err != nil {
		return nil, err
	}

	for _, trigger := range triggers {
		var lastTime time.Time
		key := fmt.Sprintf("workflow:cron:%d:time", trigger.MessageId)
		t := s.rdb.Get(ctx, key).Val()
		if t == "" {
			lastTime = time.Time{}
		} else {
			lastTime, err = time.ParseInLocation("2006-01-02 15:04:05", t, time.Local)
			if err != nil {
				return nil, err
			}
		}

		p, err := cron.ParseUTC(trigger.When)
		if err != nil {
			return nil, err
		}
		nextTime, err := p.Next(lastTime)
		if err != nil {
			return nil, err
		}

		now := time.Now()
		if nextTime.Before(now) {
			// time
			s.rdb.Set(ctx, key, now.Format("2006-01-02 15:04:05"), 0)

			// publish event
			err = s.bus.Publish(ctx, enum.Chatbot, event.WorkflowRunSubject, pb.Message{Id: trigger.MessageId})
			if err != nil {
				return nil, err
			}
		}
	}

	return &pb.WorkflowReply{}, nil
}

func (s *Chatbot) CreateTrigger(ctx context.Context, payload *pb.TriggerRequest) (*pb.StateReply, error) {
	var trigger pb.Trigger
	trigger.Type = payload.Trigger.GetType()
	trigger.Kind = payload.Trigger.GetKind()
	trigger.MessageId = payload.Trigger.GetMessageId()

	switch payload.Trigger.GetKind() {
	case enum.MessageTypeAction:
		if payload.Info.GetMessageText() == "" {
			return nil, errors.New("empty action")
		}
		p, err := action.NewParser(action.NewLexer([]rune(payload.Info.GetMessageText())))
		if err != nil {
			return nil, err
		}
		tree, err := p.Parse()
		if err != nil {
			return nil, err
		}

		symbolTable := action.NewSemanticAnalyzer()
		err = symbolTable.Visit(tree)
		if err != nil {
			return nil, err
		}

		if symbolTable.Cron == nil && symbolTable.Webhook == nil {
			return &pb.StateReply{State: false}, nil
		}

		if symbolTable.Cron != nil {
			trigger.Type = "cron"
			trigger.When = symbolTable.Cron.When

			// store
			_, err = s.repo.CreateTrigger(ctx, &trigger)
			if err != nil {
				return nil, err
			}
		}

		if symbolTable.Webhook != nil {
			trigger.Type = "webhook"
			trigger.Flag = symbolTable.Webhook.Flag
			trigger.Secret = symbolTable.Webhook.Secret

			find, err := s.repo.GetTriggerByFlag(ctx, trigger.Type, trigger.Flag)
			if err != nil {
				return nil, err
			}

			if find.Id > 0 {
				return nil, errors.New("exist flag: " + trigger.Flag)
			}

			// store
			_, err = s.repo.CreateTrigger(ctx, &trigger)
			if err != nil {
				return nil, err
			}
		}

		return &pb.StateReply{State: true}, nil
	default:
		return &pb.StateReply{State: false}, nil
	}
}

func (s *Chatbot) DeleteTrigger(ctx context.Context, payload *pb.TriggerRequest) (*pb.StateReply, error) {
	err := s.repo.DeleteTriggerByMessageID(ctx, payload.Trigger.GetMessageId())
	if err != nil {
		return &pb.StateReply{State: false}, err
	}

	return &pb.StateReply{State: true}, nil
}

func (s *Chatbot) ActionDoc(_ context.Context, payload *pb.WorkflowRequest) (*pb.WorkflowReply, error) {
	var docs string
	if payload.GetText() == "" {
		docs = strings.Join(opcode.Docs(), "\n")
	} else {
		docs = opcode.Doc(payload.GetText())
	}
	return &pb.WorkflowReply{
		Text: docs,
	}, nil
}

func (s *Chatbot) ListWebhook(ctx context.Context, _ *pb.WorkflowRequest) (*pb.WebhooksReply, error) {
	triggers, err := s.repo.ListTriggersByType(ctx, "webhook")
	if err != nil {
		return nil, err
	}
	var flags []string
	for _, item := range triggers {
		flags = append(flags, item.Flag)
	}
	return &pb.WebhooksReply{Flag: flags}, nil
}
