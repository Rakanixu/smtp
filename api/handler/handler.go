package handler

import (
	"encoding/json"
	smtp "github.com/Rakanixu/smtp/srv/proto/smtp"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/errors"
	api "github.com/micro/micro/api/proto"
	"golang.org/x/net/context"
	"net/http"
)

// Smtp struct
type Smtp struct{}

// Send (mail) API handler
func (s *Smtp) Send(ctx context.Context, req *api.Request, rsp *api.Response) error {
	if len(req.Body) <= 0 {
		return errors.BadRequest("go.micro.api.smtp", "Required fields")
	}

	var input map[string]interface{}
	json.Unmarshal([]byte(req.Body), &input)

	var recipient = make([]string, len(input["recipient"].([]interface{})))
	for k, v := range input["recipient"].([]interface{}) {
		recipient[k] = v.(string)
	}

	sendReq := client.NewRequest(
		"go.micro.srv.smtp",
		"SMTP.Send",
		&smtp.SendRequest{
			Recipient: recipient,
			Subject:   input["subject"].(string),
			Body:      input["body"].(string),
		},
	)
	sendRsp := &smtp.SendResponse{}

	if err := client.Call(ctx, sendReq, sendRsp); err != nil {
		return errors.InternalServerError("go.micro.api.smtp", err.Error())
	}

	rsp.StatusCode = http.StatusOK
	rsp.Body = `{}`

	return nil
}
