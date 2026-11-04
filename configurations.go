package elves

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Configs - List all configurations
func (c *Client) Configs() ([]Configuration, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/v1/api/config/", c.HostName), nil)
	if err != nil {
		return nil, err
	}
	body, err := c.RequestResponse200(req)
	if err != nil {
		return nil, err
	}
	configs := []Configuration{}
	if err := json.Unmarshal(body, &configs); err != nil {
		return nil, err
	}
	return configs, nil
}

// ConfigGet - Get a specific configuration
func (c *Client) ConfigGet(id string) (*Configuration, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/v1/api/config/%s", c.HostName, id), nil)
	if err != nil {
		return nil, err
	}
	body, err := c.RequestResponse200(req)
	if err != nil {
		return nil, err
	}
	config := Configuration{}
	if err := json.Unmarshal(body, &config); err != nil {
		return nil, err
	}
	return &config, nil
}

// ConfigCreate - Create a new configuration
func (c *Client) ConfigCreate(config *Configuration) (*Configuration, error) {
	jsonBody, err := json.Marshal(config)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/v1/api/config/", c.HostName), bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	body, err := c.RequestResponse201(req)
	if err != nil {
		return nil, err
	}
	created := Configuration{}
	if err := json.Unmarshal(body, &created); err != nil {
		return nil, err
	}
	return &created, nil
}

// ConfigUpdate - Replace a configuration. The default configuration (ID 1)
// can't be changed.
func (c *Client) ConfigUpdate(id string, config *Configuration) error {
	jsonBody, err := json.Marshal(config)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPatch, fmt.Sprintf("%s/v1/api/config/%s", c.HostName, id), bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	_, err = c.RequestResponse200(req)
	return err
}

// ConfigDelete - Delete a configuration. Its Macs move to the default
// configuration (ID 1), which can't be deleted.
func (c *Client) ConfigDelete(id string) error {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/v1/api/config/%s", c.HostName, id), nil)
	if err != nil {
		return err
	}
	_, err = c.RequestResponse204(req)
	return err
}
