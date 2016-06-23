package handler

import (
	proto "github.com/Rakanixu/smtp/srv/proto/smtp"
	smtp "github.com/Rakanixu/smtp/srv/smtp"
	"github.com/micro/go-micro/errors"
	"golang.org/x/net/context"
)

// SMTP struct
type SMTP struct{}

// Send srv handler
func (s *SMTP) Send(ctx context.Context, req *proto.SendRequest, rsp *proto.SendResponse) error {
	if len(req.Recipient) <= 0 {
		return errors.BadRequest("go.micro.srv.smtp.Smtp.Send", "Recipient required")
	}

	if err := smtp.Send(req.Recipient, req.Subject, req.Body); err != nil {
		return errors.InternalServerError("go.micro.srv.smtp.Smtp.Send", err.Error())
	}

	return nil
}
