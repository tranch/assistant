package trigger

import (
	"context"
	"fmt"
	"github.com/tsundata/assistant/api/model"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/chatbot/trigger/ctx"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/util"
	"github.com/tsundata/assistant/internal/pkg/vendors/email"
	"regexp"
	"strconv"
	"strings"
)

type Email struct {
	text  string
	email []string
}

func NewEmail() *Email {
	return &Email{}
}

func (t *Email) Cond(text string) bool {
	re := regexp.MustCompile(`(?m)` + util.EmailRegex)
	ts := re.FindAllString(text, -1)

	if len(ts) == 0 {
		return false
	}

	t.text = text
	for _, item := range ts {
		t.text = strings.ReplaceAll(t.text, item, "")
		t.text = strings.ReplaceAll(t.text, "<mailto:|>", "") // special
		t.email = append(t.email, item)
	}

	t.email = clear(t.email)

	return true
}

func (t *Email) Handle(ctx context.Context, comp *ctx.Component) {
	reply, err := comp.Middle.GetCredential(ctx, &pb.CredentialRequest{Type: email.ID})
	if err != nil {
		return
	}
	host := ""
	port := ""
	username := ""
	password := ""
	for _, item := range reply.GetContent() {
		if item.Key == email.Host {
			host = item.Value
		}
		if item.Key == email.Port {
			port = item.Value
		}
		if item.Key == email.Username {
			username = item.Value
		}
		if item.Key == email.Password {
			password = item.Value
		}
	}
	if host == "" || port == "" || username == "" || password == "" {
		return
	}

	p, err := strconv.Atoi(port)
	if err != nil {
		return
	}

	config := email.Config{
		Host:     host,
		Port:     p,
		Username: username,
		Password: password,
	}

	for _, mail := range t.email {
		err = email.SendEmail(&config, mail, t.text, fmt.Sprintf(`%s <br>
---------------- <br>
Sended by Assistant
`, t.text))
		if err != nil {
			comp.Logger.Error(err)
			return
		}
		err := comp.Bus.Publish(ctx, event.SendMessageSubject, model.Message{Text: fmt.Sprintf("Sended to Mail: %s", mail)})
		if err != nil {
			return
		}
	}
}
