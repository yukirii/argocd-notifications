package notifiers

import (
	"context"
	"fmt"
	"github.com/opsgenie/opsgenie-go-sdk-v2/alert"
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
)

type OpsgenieOptions struct {
	ApiUrl  string            `json:"apiUrl"`
	ApiKeys map[string]string `json:"apiKeys"`
}

type opsgenieNotifier struct {
	opts OpsgenieOptions
}

func NewOpsgenieNotifier(opts OpsgenieOptions) Notifier {
	return &opsgenieNotifier{opts: opts}
}

func (n *opsgenieNotifier) Send(title string, body string, recipient string) error {
	apiKey, ok := n.opts.ApiKeys[recipient]
	if !ok {
		return fmt.Errorf("no API key configured for recipient %s", recipient)
	}
	alertClient, _ := alert.NewClient(&client.Config{
		ApiKey: apiKey,
		OpsGenieAPIURL: client.ApiUrl(n.opts.ApiUrl),
	})
	_, err := alertClient.Create(context.TODO(), &alert.CreateAlertRequest{
		Message: title,
		Description: body,
		Responders: []alert.Responder{
			{
				Type: "team",
				Id:   recipient,
			},
		},
		Source: "Argo CD",
	})
	return err
}
