package elves

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// RuleGet- Get specific rule
func (c *Client) RuleGet(id string) (*Rule, error) {
	url := fmt.Sprintf("%s/v1/api/rule/%s", c.HostName, id)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	body, err := c.RequestResponse200(req)
	if err != nil {
		return nil, err
	}

	rule := Rule{}
	err = json.Unmarshal(body, &rule)
	if err != nil {
		return nil, err
	}

	return &rule, nil
}

// RuleCreate - Create new rule
func (c *Client) RuleCreate(rule *Rule) (*Rule, error) {
	url := fmt.Sprintf("%s/v1/api/rule/", c.HostName)

	// Marshal the Rule struct to JSON
	jsonBody, err := json.Marshal(rule)
	if err != nil {
		return nil, err
	}

	// Create a new request with the JSON body
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	body, err := c.RequestResponse201(req)
	if err != nil {
		return nil, err
	}

	createdRule := Rule{}
	err = json.Unmarshal(body, &createdRule)
	if err != nil {
		return nil, err
	}

	return &createdRule, nil
}

// RuleUpdate - Updates an rule
func (c *Client) RuleUpdate(id string, rule *Rule) error {
	url := fmt.Sprintf("%s/v1/api/rule/%s", c.HostName, id)

	// Marshal the Rule struct to JSON
	jsonBody, err := json.Marshal(rule)
	if err != nil {
		return err
	}

	// Create a new request with the JSON body
	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")

	body, err := c.RequestResponse200(req)
	if err != nil {
		return err
	}

	updatedRule := Rule{}
	err = json.Unmarshal(body, &updatedRule)
	if err != nil {
		return err
	}

	return nil
}

// RuleDelete - Deletes an Rule
func (c *Client) RuleDelete(id string) error {
	url := fmt.Sprintf("%s/v1/api/rule/%s", c.HostName, id)
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	body, err := c.RequestResponse204(req)

	if err != nil {
		return err
	}

	if string(body) != "" {
		return errors.New(string(body))
	}

	return nil
}
