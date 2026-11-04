package elves

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// ScimGroups - List all SCIM groups
func (c *Client) ScimGroups() ([]ScimGroup, error) {
	var groups []ScimGroup
	for page := 1; ; page++ {
		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/v1/api/scimgroups/?page=%d", c.HostName, page), nil)
		if err != nil {
			return nil, err
		}
		body, err := c.RequestResponse200(req)
		if err != nil {
			return nil, err
		}
		var resp struct {
			MoreResults bool        `json:"more_results"`
			Results     []ScimGroup `json:"results"`
		}
		if err := json.Unmarshal(body, &resp); err != nil {
			return nil, err
		}
		groups = append(groups, resp.Results...)
		if !resp.MoreResults {
			return groups, nil
		}
	}
}

// ScimGroupRules - List the rules assigned to a SCIM group
func (c *Client) ScimGroupRules(groupID int64) ([]AssignedRule, error) {
	return c.assignedRules(fmt.Sprintf("%s/v1/api/scimgroups/%d/rule/", c.HostName, groupID))
}

// ScimGroupRuleAssign - Include or exclude a rule for a SCIM group. Assigning
// an already assigned rule changes its action.
func (c *Client) ScimGroupRuleAssign(groupID, ruleID int64, action string) error {
	return c.assignRule(fmt.Sprintf("%s/v1/api/scimgroups/%d/rule/%d", c.HostName, groupID, ruleID), action)
}

// ScimGroupRuleUnassign - Remove a rule from a SCIM group
func (c *Client) ScimGroupRuleUnassign(groupID, ruleID int64) error {
	return c.unassignRule(fmt.Sprintf("%s/v1/api/scimgroups/%d/rule/%d", c.HostName, groupID, ruleID))
}

// ClientRules - List the rules assigned to a client (by serial number)
func (c *Client) ClientRules(serial string) ([]AssignedRule, error) {
	return c.assignedRules(fmt.Sprintf("%s/v1/api/client/%s/rule/", c.HostName, url.PathEscape(serial)))
}

// ClientRuleAssign - Include or exclude a rule for a client. Assigning an
// already assigned rule changes its action.
func (c *Client) ClientRuleAssign(serial string, ruleID int64, action string) error {
	return c.assignRule(fmt.Sprintf("%s/v1/api/client/%s/rule/%d", c.HostName, url.PathEscape(serial), ruleID), action)
}

// ClientRuleUnassign - Remove a rule from a client
func (c *Client) ClientRuleUnassign(serial string, ruleID int64) error {
	return c.unassignRule(fmt.Sprintf("%s/v1/api/client/%s/rule/%d", c.HostName, url.PathEscape(serial), ruleID))
}

// ScimUserByEmail - Find a SCIM user by email or user name, ignoring case.
// It returns nil when there is none.
func (c *Client) ScimUserByEmail(email string) (*ScimUser, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/v1/api/scimusers/?email=%s", c.HostName, url.QueryEscape(email)), nil)
	if err != nil {
		return nil, err
	}
	body, err := c.RequestResponse200(req)
	if err != nil {
		return nil, err
	}
	var resp struct {
		Results []ScimUser `json:"results"`
	}
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}
	switch len(resp.Results) {
	case 0:
		return nil, nil
	case 1:
		return &resp.Results[0], nil
	default:
		return nil, fmt.Errorf("more than one SCIM user matches %q", email)
	}
}

// RuleScimUsers - List the SCIM users a rule is assigned to
func (c *Client) RuleScimUsers(ruleID int64) ([]RuleScimUser, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/v1/api/rule/%d/scimusers", c.HostName, ruleID), nil)
	if err != nil {
		return nil, err
	}
	body, err := c.RequestResponse200(req)
	if err != nil {
		return nil, err
	}
	users := []RuleScimUser{}
	if err := json.Unmarshal(body, &users); err != nil {
		return nil, err
	}
	return users, nil
}

// ScimUserRules - List the rules assigned to a SCIM user
func (c *Client) ScimUserRules(userID int64) ([]AssignedRule, error) {
	return c.assignedRules(fmt.Sprintf("%s/v1/api/scimusers/%d/rule/", c.HostName, userID))
}

// ScimUserRuleAssign - Include or exclude a rule for a SCIM user. Assigning an
// already assigned rule changes its action.
func (c *Client) ScimUserRuleAssign(userID, ruleID int64, action string) error {
	return c.assignRule(fmt.Sprintf("%s/v1/api/scimusers/%d/rule/%d", c.HostName, userID, ruleID), action)
}

// ScimUserRuleUnassign - Remove a rule from a SCIM user
func (c *Client) ScimUserRuleUnassign(userID, ruleID int64) error {
	return c.unassignRule(fmt.Sprintf("%s/v1/api/scimusers/%d/rule/%d", c.HostName, userID, ruleID))
}

// RuleConfigs - List the configurations a rule is assigned to
func (c *Client) RuleConfigs(ruleID int64) ([]RuleConfig, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/v1/api/rule/%d/configs", c.HostName, ruleID), nil)
	if err != nil {
		return nil, err
	}
	body, err := c.RequestResponse200(req)
	if err != nil {
		return nil, err
	}
	configs := []RuleConfig{}
	if err := json.Unmarshal(body, &configs); err != nil {
		return nil, err
	}
	return configs, nil
}

// ConfigRules - List the rules assigned to a configuration
func (c *Client) ConfigRules(configID int64) ([]AssignedRule, error) {
	return c.assignedRules(fmt.Sprintf("%s/v1/api/config/%d/rule/", c.HostName, configID))
}

// ConfigRuleAssign - Include or exclude a rule for a configuration. Assigning
// an already assigned rule changes its action.
func (c *Client) ConfigRuleAssign(configID, ruleID int64, action string) error {
	return c.assignRule(fmt.Sprintf("%s/v1/api/config/%d/rule/%d", c.HostName, configID, ruleID), action)
}

// ConfigRuleUnassign - Remove a rule from a configuration
func (c *Client) ConfigRuleUnassign(configID, ruleID int64) error {
	return c.unassignRule(fmt.Sprintf("%s/v1/api/config/%d/rule/%d", c.HostName, configID, ruleID))
}

func (c *Client) assignedRules(endpoint string) ([]AssignedRule, error) {
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}
	body, err := c.RequestResponse200(req)
	if err != nil {
		return nil, err
	}
	rules := []AssignedRule{}
	if err := json.Unmarshal(body, &rules); err != nil {
		return nil, err
	}
	return rules, nil
}

func (c *Client) assignRule(endpoint, action string) error {
	jsonBody, err := json.Marshal(map[string]string{"action": action})
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	_, err = c.RequestResponse201(req)
	return err
}

func (c *Client) unassignRule(endpoint string) error {
	req, err := http.NewRequest(http.MethodDelete, endpoint, nil)
	if err != nil {
		return err
	}
	_, err = c.RequestResponse204(req)
	return err
}
