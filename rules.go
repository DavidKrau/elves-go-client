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

	attribute := Rule{}
	err = json.Unmarshal(body, &attribute)
	if err != nil {
		return nil, err
	}

	return &attribute, nil
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
	body, err := c.RequestResponse200(req)
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
func (c *Client) RuleUpdate(id string, defaultValue string) error {
	url := fmt.Sprintf("%s/v1/api/rule/%s", c.HostName, id)
	req, err := http.NewRequest(http.MethodPatch, url, nil)
	if err != nil {
		return err
	}

	q := req.URL.Query()
	// checking if defaultvalue exist
	if defaultValue != "" {
		// if yes adding parameter with value
		q.Add("default_value", defaultValue)
	} else {
		q.Add("default_value", "")
	}
	// encoding all parameters
	req.URL.RawQuery = q.Encode()

	body, err := c.RequestResponse200(req)
	if err != nil {
		return err
	}

	if string(body) != "" {
		return errors.New(string(body))
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

	body, err := c.RequestResponse200(req)

	if err != nil {
		return err
	}

	if string(body) != "" {
		return errors.New(string(body))
	}

	return nil
}
