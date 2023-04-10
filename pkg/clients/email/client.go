package email

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"go.uber.org/zap"
)

type IEmailClient interface {
	SendEmail(ctx context.Context, target string) (string, error)
}

type client struct {
	log zap.Logger
	url string
}

func NewClient(log zap.Logger, url string) IEmailClient {
	return &client{
		log: log,
		url: url,
	}
}

func (c *client) SendEmail(ctx context.Context, target string) (string, error) {
	target = strings.Replace(target, "@", "%40", -1)
	url := fmt.Sprintf("%s?email=%s", c.url, target)
	resp, err := http.Post(url, "application/json", nil)
	if err != nil {
		return "", err
	}

	// read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.log.Error("error reading response body", zap.Error(err))
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		c.log.Error("error sending email", zap.Error(err), zap.Any("status", resp.StatusCode), zap.Any("body", string(body)), zap.Any("url", url))
		return "", fmt.Errorf("error sending email")
	}

	return string(body), nil
}
