package controller

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/gateway/health"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/log"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc/md"
	"github.com/tsundata/assistant/internal/pkg/vendors/newrelic"
	"github.com/tsundata/assistant/internal/pkg/version"
	"go.uber.org/zap"
	"google.golang.org/grpc/health/grpc_health_v1"
	"strings"
	"time"
)

type GatewayController struct {
	opt    *config.AppConfig
	rdb    *redis.Client
	logger log.Logger
	nr     *newrelic.App
	bus    event.Bus

	messageSvc   pb.MessageSvcClient
	middleSvc    pb.MiddleSvcClient
	workflowSvc  pb.WorkflowSvcClient
	userSvc      pb.UserSvcClient
	chatbotSvc   pb.ChatbotSvcClient
	healthClient *health.Client
}

func NewGatewayController(
	opt *config.AppConfig,
	rdb *redis.Client,
	logger log.Logger,
	nr *newrelic.App,
	bus event.Bus,
	messageSvc pb.MessageSvcClient,
	middleSvc pb.MiddleSvcClient,
	workflowSvc pb.WorkflowSvcClient,
	chatbotSvc pb.ChatbotSvcClient,
	userSvc pb.UserSvcClient,
	healthClient *health.Client) *GatewayController {
	return &GatewayController{
		opt:          opt,
		rdb:          rdb,
		logger:       logger,
		nr:           nr,
		bus:          bus,
		messageSvc:   messageSvc,
		middleSvc:    middleSvc,
		workflowSvc:  workflowSvc,
		userSvc:      userSvc,
		chatbotSvc:   chatbotSvc,
		healthClient: healthClient,
	}
}

func (gc *GatewayController) Index(c *fiber.Ctx) error {
	return c.SendString(fmt.Sprintf("Gateway %s", version.Version))
}

func (gc *GatewayController) GetChart(c *fiber.Ctx) error {
	var in pb.ChartData
	err := c.QueryParser(&in)
	if err != nil {
		return err
	}

	reply, err := gc.middleSvc.GetChartData(md.Outgoing(c), &pb.ChartDataRequest{ChartData: &in})
	if err != nil {
		return err
	}
	return c.JSON(reply)
}

func (gc *GatewayController) WebhookTrigger(c *fiber.Ctx) error {
	var in pb.TriggerRequest
	err := c.BodyParser(&in)
	if err != nil {
		return err
	}

	reply, err := gc.workflowSvc.WebhookTrigger(md.Outgoing(c), &in)
	if err != nil {
		return err
	}
	return c.JSON(reply)
}

func (gc *GatewayController) GetPage(c *fiber.Ctx) error {
	var in pb.Page
	err := c.QueryParser(&in)
	if err != nil {
		return err
	}

	reply, err := gc.middleSvc.GetPage(md.Outgoing(c), &pb.PageRequest{Page: &in})
	if err != nil {
		return err
	}
	return c.JSON(reply)
}

func (gc *GatewayController) Authorization(c *fiber.Ctx) error {
	var in pb.LoginRequest
	err := c.BodyParser(&in)
	if err != nil {
		return err
	}

	reply, err := gc.userSvc.Login(md.Outgoing(c), &in)
	if err != nil {
		return err
	}
	return c.JSON(reply)
}

func (gc *GatewayController) StoreAppOAuth(c *fiber.Ctx) error {
	var in pb.AppRequest
	err := c.BodyParser(&in)
	if err != nil {
		return err
	}

	reply, err := gc.middleSvc.StoreAppOAuth(md.Outgoing(c), &in)
	if err != nil {
		return err
	}
	return c.JSON(reply)
}

// GetApps godoc
// @Summary Get Apps
// @Description get apps
// @ID get-apps
// @Accept json
// @Produce json
// @Success 200 {object} pb.TextRequest
// @Router /apps [get]
func (gc *GatewayController) GetApps(c *fiber.Ctx) error {
	var in pb.TextRequest
	err := c.QueryParser(&in)
	if err != nil {
		return err
	}

	reply, err := gc.middleSvc.GetApps(md.Outgoing(c), &in)
	if err != nil {
		return err
	}
	return c.JSON(reply)
}

func (gc *GatewayController) GetMessages(c *fiber.Ctx) error {
	var in pb.GetMessagesRequest
	err := c.QueryParser(&in)
	if err != nil {
		return err
	}

	reply, err := gc.messageSvc.ListByGroup(md.Outgoing(c), &in)
	if err != nil {
		return err
	}
	return c.JSON(reply)
}

func (gc *GatewayController) GetMaskingCredentials(c *fiber.Ctx) error {
	var in pb.TextRequest
	err := c.QueryParser(&in)
	if err != nil {
		return err
	}

	reply, err := gc.middleSvc.GetMaskingCredentials(md.Outgoing(c), &in)
	if err != nil {
		return err
	}
	return c.JSON(reply)
}

func (gc *GatewayController) GetCredential(c *fiber.Ctx) error {
	var in pb.CredentialRequest
	err := c.QueryParser(&in)
	if err != nil {
		return err
	}

	reply, err := gc.middleSvc.GetCredential(md.Outgoing(c), &in)
	if err != nil {
		return err
	}
	return c.JSON(reply)
}

func (gc *GatewayController) CreateCredential(c *fiber.Ctx) error {
	var in pb.KVsRequest
	err := c.BodyParser(&in)
	if err != nil {
		return err
	}

	reply, err := gc.middleSvc.CreateCredential(md.Outgoing(c), &in)
	if err != nil {
		return err
	}
	return c.JSON(reply)
}

func (gc *GatewayController) GetSettings(c *fiber.Ctx) error {
	var in pb.TextRequest
	err := c.QueryParser(&in)
	if err != nil {
		return err
	}

	reply, err := gc.middleSvc.GetSettings(md.Outgoing(c), &in)
	if err != nil {
		return err
	}
	return c.JSON(reply)
}

func (gc *GatewayController) CreateSetting(c *fiber.Ctx) error {
	var in pb.KVRequest
	err := c.BodyParser(&in)
	if err != nil {
		return err
	}

	reply, err := gc.middleSvc.CreateSetting(md.Outgoing(c), &in)
	if err != nil {
		return err
	}
	return c.JSON(reply)
}

func (gc *GatewayController) GetActionMessages(c *fiber.Ctx) error {
	var in pb.TextRequest
	err := c.QueryParser(&in)
	if err != nil {
		return err
	}

	reply, err := gc.messageSvc.GetActionMessages(md.Outgoing(c), &in)
	if err != nil {
		return err
	}
	return c.JSON(reply)
}

func (gc *GatewayController) CreateActionMessage(c *fiber.Ctx) error {
	var in pb.TextRequest
	err := c.BodyParser(&in)
	if err != nil {
		return err
	}

	reply, err := gc.messageSvc.CreateActionMessage(md.Outgoing(c), &in)
	if err != nil {
		return err
	}
	return c.JSON(reply)
}

func (gc *GatewayController) DeleteWorkflowMessage(c *fiber.Ctx) error {
	var in pb.Message
	err := c.BodyParser(&in)
	if err != nil {
		return err
	}

	reply, err := gc.messageSvc.DeleteWorkflowMessage(md.Outgoing(c), &pb.MessageRequest{Message: &in})
	if err != nil {
		return err
	}
	return c.JSON(reply)
}

func (gc *GatewayController) RunMessage(c *fiber.Ctx) error {
	var in pb.Message
	err := c.BodyParser(&in)
	if err != nil {
		return err
	}

	reply, err := gc.messageSvc.Run(md.Outgoing(c), &pb.MessageRequest{Message: &in})
	if err != nil {
		return err
	}
	return c.JSON(reply)
}

func (gc *GatewayController) SendMessage(c *fiber.Ctx) error {
	var in pb.Message
	err := c.BodyParser(&in)
	if err != nil {
		return err
	}

	reply, err := gc.messageSvc.Send(md.Outgoing(c), &pb.MessageRequest{Message: &in})
	if err != nil {
		return err
	}
	return c.JSON(reply)
}

func (gc *GatewayController) GetRoleImage(c *fiber.Ctx) error {
	var in pb.RoleRequest
	err := c.QueryParser(&in)
	if err != nil {
		return err
	}

	reply, err := gc.userSvc.GetRoleImage(md.Outgoing(c), &in)
	if err != nil {
		return err
	}
	return c.JSON(reply)
}

func (gc *GatewayController) Health(c *fiber.Ctx) error {
	str := strings.Builder{}
	gc.healthClient.Status.Range(func(key, value interface{}) bool {
		if service, ok := key.(string); ok {
			str.WriteString(service)
			str.WriteString(": ")
		}
		if status, ok := value.(grpc_health_v1.HealthCheckResponse_ServingStatus); ok {
			str.WriteString(strings.ToLower(grpc_health_v1.HealthCheckResponse_ServingStatus_name[int32(status)]))
			str.WriteString("\n")
		}
		return true
	})
	return c.SendString(str.String())
}

func (gc *GatewayController) GetGroups(c *fiber.Ctx) error {
	var in pb.Group
	err := c.QueryParser(&in)
	if err != nil {
		return err
	}

	reply, err := gc.chatbotSvc.GetGroups(md.Outgoing(c), &pb.GroupRequest{Group: &in})
	if err != nil {
		return err
	}
	return c.JSON(reply)
}

func (gc *GatewayController) GetGroup(c *fiber.Ctx) error {
	var in pb.Group
	err := c.QueryParser(&in)
	if err != nil {
		return err
	}

	reply, err := gc.chatbotSvc.GetGroup(md.Outgoing(c), &pb.GroupRequest{Group: &in})
	if err != nil {
		return err
	}
	return c.JSON(reply)
}

func (gc *GatewayController) GetMessage(c *fiber.Ctx) error {
	var in pb.Message
	err := c.QueryParser(&in)
	if err != nil {
		return err
	}

	reply, err := gc.messageSvc.Get(md.Outgoing(c), &pb.MessageRequest{Message: &in})
	if err != nil {
		return err
	}
	return c.JSON(reply)
}

func (gc *GatewayController) GetUser(c *fiber.Ctx) error {
	var in pb.User
	err := c.QueryParser(&in)
	if err != nil {
		return err
	}

	reply, err := gc.userSvc.GetUser(md.Outgoing(c), &pb.UserRequest{User: &in})
	if err != nil {
		return err
	}
	return c.JSON(reply)
}

func (gc *GatewayController) GetGroupSetting(c *fiber.Ctx) error {
	var in pb.GroupSettingRequest
	err := c.QueryParser(&in)
	if err != nil {
		return err
	}

	reply, err := gc.chatbotSvc.GetGroupSetting(md.Outgoing(c), &in)
	if err != nil {
		return err
	}
	return c.JSON(reply)
}

func (gc *GatewayController) UpdateGroupSetting(c *fiber.Ctx) error {
	var in pb.GroupSettingRequest
	err := c.BodyParser(&in)
	if err != nil {
		return err
	}

	reply, err := gc.chatbotSvc.UpdateGroupSetting(md.Outgoing(c), &in)
	if err != nil {
		return err
	}
	return c.JSON(reply)
}

func (gc *GatewayController) GetGroupBots(c *fiber.Ctx) error {
	var in pb.BotsRequest
	err := c.QueryParser(&in)
	if err != nil {
		return err
	}

	reply, err := gc.chatbotSvc.GetBots(md.Outgoing(c), &in)
	if err != nil {
		return err
	}
	return c.JSON(reply)
}

func (gc *GatewayController) GetGroupBot(c *fiber.Ctx) error {
	var in pb.BotRequest
	err := c.QueryParser(&in)
	if err != nil {
		return err
	}

	reply, err := gc.chatbotSvc.GetBot(md.Outgoing(c), &in)
	if err != nil {
		return err
	}
	return c.JSON(reply)
}

func (gc *GatewayController) GetGroupBotSetting(c *fiber.Ctx) error {
	var in pb.BotSettingRequest
	err := c.QueryParser(&in)
	if err != nil {
		return err
	}

	reply, err := gc.chatbotSvc.GetGroupBotSetting(md.Outgoing(c), &in)
	if err != nil {
		return err
	}
	return c.JSON(reply)
}

func (gc *GatewayController) UpdateGroupBotSetting(c *fiber.Ctx) error {
	var in pb.BotSettingRequest
	err := c.BodyParser(&in)
	if err != nil {
		return err
	}

	reply, err := gc.chatbotSvc.UpdateGroupBotSetting(md.Outgoing(c), &in)
	if err != nil {
		return err
	}
	return c.JSON(reply)
}

func (gc *GatewayController) GetInboxes(c *fiber.Ctx) error {
	var in pb.InboxRequest
	err := c.QueryParser(&in)
	if err != nil {
		return err
	}

	reply, err := gc.messageSvc.ListInbox(md.Outgoing(c), &in)
	if err != nil {
		return err
	}
	return c.JSON(reply)
}

func (gc *GatewayController) Notify(conn *websocket.Conn, userId int64) {
	gc.logger.Info("[Notify] listening...", zap.Any("user", userId))

	go func() {
		for {
			_, data, err := conn.ReadMessage()
			if err != nil {
				gc.logger.Warn(err.Error())
				break
			}
			var inbox pb.Inbox
			err = json.Unmarshal(data, &inbox)
			if err != nil {
				gc.logger.Error(err)
				continue
			}
			_, err = gc.messageSvc.MarkReadInbox(md.BuildAuthContext(userId), &pb.InboxRequest{InboxId: inbox.Id})
			if err != nil {
				gc.logger.Error(err)
				continue
			}
		}
	}()

	t := time.NewTicker(10 * time.Second)
	for {
		<-t.C
		err := conn.WriteMessage(websocket.PingMessage, []byte{})
		if err != nil {
			t.Stop()
			gc.logger.Warn(err.Error())
			break
		}

		inbox, err := gc.messageSvc.LastInbox(md.BuildAuthContext(userId), &pb.InboxRequest{})
		if err != nil {
			gc.logger.Error(err)
			continue
		}

		if len(inbox.Inbox) > 0 && inbox.Inbox[0].Id > 0 {
			data, err := json.Marshal(inbox.Inbox[0])
			if err != nil {
				gc.logger.Error(err)
				continue
			}
			err = conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				t.Stop()
				gc.logger.Warn(err.Error())
				break
			}
			_, err = gc.messageSvc.MarkSendInbox(md.BuildAuthContext(userId), &pb.InboxRequest{InboxId: inbox.Inbox[0].Id})
			if err != nil {
				gc.logger.Error(err)
				continue
			}
		}
	}

	gc.logger.Info("[Notify] end", zap.Any("user", userId))
}
