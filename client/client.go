package client

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
	"gitlab.com/vorticist/logger"
	"os"
)

var (
	fnsUrl = "https://faas-tor1-70ca848e.doserverless.co/api/v1/namespaces/fn-aded6f49-267f-4eb8-81d4-6e6e3875ec90/actions/book/%v?blocking=true&result=true"
)

type Client interface {
	GetEntries() (map[string]string, error)
	SaveEntry(key, value string) (map[string]string, error)
	RemoveEntry(key string) (map[string]string, error)
}

func New() Client {
	godotenv.Load(".env")
	authHeader := fmt.Sprintf("Basic %v", os.Getenv("API_KEY"))
	logger.Infof("API_KEY = %v", authHeader)
	return &client{
		authHeader: authHeader,
	}
}

type client struct {
	authHeader string
}

func (c *client) GetEntries() (map[string]string, error) {
	logger.Info("GetEntries")
	client := resty.New()

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", c.authHeader).
		EnableTrace().
		Post(fmt.Sprintf(fnsUrl, "getEntries"))
	if err != nil {
		logger.Errorf("failed to get entries: %v", err)
		return nil, err
	}

	var result map[string]map[string]string
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		logger.Errorf("failed to unmarshal response: %v", err)
		return nil, err
	}

	return result["entries"], nil
}

func (c *client) SaveEntry(key, value string) (map[string]string, error) {
	client := resty.New()

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", c.authHeader).
		EnableTrace().
		SetBody(map[string]interface{}{
			"key":   key,
			"value": value,
		}).
		Post(fmt.Sprintf(fnsUrl, "saveEntry"))
	if err != nil {
		logger.Errorf("failed to save entry: %v", err)
		return nil, err
	}

	var result map[string]map[string]string
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		logger.Errorf("failed to unmarshal response: %v", err)
		return nil, err
	}

	return result["entries"], nil
}

func (c *client) RemoveEntry(key string) (map[string]string, error) {
	client := resty.New()

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", c.authHeader).
		EnableTrace().
		SetBody(map[string]interface{}{
			"key": key,
		}).
		Post(fmt.Sprintf(fnsUrl, "removeEntry"))
	if err != nil {
		logger.Errorf("failed to save entry: %v", err)
		return nil, err
	}

	var result map[string]map[string]string
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		logger.Errorf("failed to unmarshal response: %v", err)
		return nil, err
	}

	return result["entries"], nil
}
