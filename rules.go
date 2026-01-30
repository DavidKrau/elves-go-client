package elves

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// RuleGet- Get specific rule
func (c *Client) RuleGet(id string) (*Attribute, error) {
	url := fmt.Sprintf("%s/v1/api/rule/%s", c.HostName, id)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	body, err := c.RequestResponse200(req)
	if err != nil {
		return nil, err
	}

	attribute := Attribute{}
	err = json.Unmarshal(body, &attribute)
	if err != nil {
		return nil, err
	}

	return &attribute, nil
}

// RuleCreate - Create new rule
func (c *Client) RuleCreate(name string, defaultValue string) (*Attribute, error) {
	url := fmt.Sprintf("%s/v1/api/rule/", c.HostName)
	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	// adding parameter name with variable name
	q.Add("name", name)
	// checking if defaultvalue exist
	if defaultValue != "" {
		// if yes adding parameter with value
		q.Add("default_value", defaultValue)
	}
	// encoding all parameters
	req.URL.RawQuery = q.Encode()

	body, err := c.RequestResponse200(req)
	if err != nil {
		return nil, err
	}

	attribute := Attribute{}
	err = json.Unmarshal(body, &attribute)
	if err != nil {
		return nil, err
	}

	return &attribute, nil
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
