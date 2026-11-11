package elves

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// FAAPath is one entry of a Santa FAA rule's Paths. It uses Santa's key names.
type FAAPath struct {
	Path     string `json:"Path"`
	IsPrefix bool   `json:"IsPrefix"`
}

// FAAProcess is one entry of a Santa FAA rule's Processes. Santa matches a
// process only if every key that is set matches.
type FAAProcess struct {
	SigningID         string `json:"SigningID,omitempty"`
	TeamID            string `json:"TeamID,omitempty"`
	PlatformBinary    *bool  `json:"PlatformBinary,omitempty"`
	CDHash            string `json:"CDHash,omitempty"`
	CertificateSha256 string `json:"CertificateSha256,omitempty"`
	BinaryPath        string `json:"BinaryPath,omitempty"`
}

// FAARule is one WatchItems entry of the Santa File Access Authorization
// policy that Elves exports into the Santa MDM profile.
type FAARule struct {
	ID       int64  `json:"id,omitempty"`
	Name     string `json:"name"`
	RuleType string `json:"rule_type"`

	Paths     []FAAPath    `json:"paths"`
	Processes []FAAProcess `json:"processes"`

	AllowReadAccess bool `json:"allow_read_access"`
	AuditOnly       bool `json:"audit_only"`
	// audit, enforce or disabled. Elves sets enforce on new Paths* rules that
	// leave it empty.
	ParentDirectoryProtection string `json:"parent_directory_protection"`
	EnableSilentMode          bool   `json:"enable_silent_mode"`
	EnableSilentTTYMode       bool   `json:"enable_silent_tty_mode"`
	BlockMessage              string `json:"block_message"`
	EventDetailURL            string `json:"event_detail_url"`
	EventDetailText           string `json:"event_detail_text"`

	// Shown to users on the FAA event page, not exported to Santa
	Description string `json:"description"`

	// Likely mistakes that Santa accepts. Set by Elves on read and write, ignored on write.
	Warnings []string `json:"warnings,omitempty"`
}

// FAASettings holds the FAA policy's own keys. Elves bumps Version on every
// rule change, and on a settings update that leaves it unchanged.
type FAASettings struct {
	Version         string `json:"version"`
	EventDetailURL  string `json:"event_detail_url"`
	EventDetailText string `json:"event_detail_text"`
}

// FAARules - List all FAA rules
func (c *Client) FAARules() ([]FAARule, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/v1/api/faa/rules/", c.HostName), nil)
	if err != nil {
		return nil, err
	}
	body, err := c.RequestResponse200(req)
	if err != nil {
		return nil, err
	}
	resp := struct {
		Results []FAARule `json:"results"`
	}{}
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}
	return resp.Results, nil
}

// FAARuleGet - Get a specific FAA rule
func (c *Client) FAARuleGet(id string) (*FAARule, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/v1/api/faa/rules/%s", c.HostName, id), nil)
	if err != nil {
		return nil, err
	}
	body, err := c.RequestResponse200(req)
	if err != nil {
		return nil, err
	}
	rule := FAARule{}
	if err := json.Unmarshal(body, &rule); err != nil {
		return nil, err
	}
	return &rule, nil
}

// FAARuleCreate - Create a new FAA rule. Names must be unique.
func (c *Client) FAARuleCreate(rule *FAARule) (*FAARule, error) {
	jsonBody, err := json.Marshal(rule)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/v1/api/faa/rules/", c.HostName), bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	body, err := c.RequestResponse201(req)
	if err != nil {
		return nil, err
	}
	created := FAARule{}
	if err := json.Unmarshal(body, &created); err != nil {
		return nil, err
	}
	return &created, nil
}

// FAARuleUpdate - Replace an FAA rule
func (c *Client) FAARuleUpdate(id string, rule *FAARule) (*FAARule, error) {
	jsonBody, err := json.Marshal(rule)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/v1/api/faa/rules/%s", c.HostName, id), bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	body, err := c.RequestResponse200(req)
	if err != nil {
		return nil, err
	}
	updated := FAARule{}
	if err := json.Unmarshal(body, &updated); err != nil {
		return nil, err
	}
	return &updated, nil
}

// FAARuleDelete - Delete an FAA rule
func (c *Client) FAARuleDelete(id string) error {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/v1/api/faa/rules/%s", c.HostName, id), nil)
	if err != nil {
		return err
	}
	_, err = c.RequestResponse204(req)
	return err
}

// FAASettingsGet - Get the FAA policy settings
func (c *Client) FAASettingsGet() (*FAASettings, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/v1/api/faa/settings", c.HostName), nil)
	if err != nil {
		return nil, err
	}
	body, err := c.RequestResponse200(req)
	if err != nil {
		return nil, err
	}
	settings := FAASettings{}
	if err := json.Unmarshal(body, &settings); err != nil {
		return nil, err
	}
	return &settings, nil
}

// FAASettingsUpdate - Replace the FAA policy settings. If Version is the
// current version, Elves bumps it; the returned settings have the saved version.
func (c *Client) FAASettingsUpdate(settings *FAASettings) (*FAASettings, error) {
	jsonBody, err := json.Marshal(settings)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/v1/api/faa/settings", c.HostName), bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	body, err := c.RequestResponse200(req)
	if err != nil {
		return nil, err
	}
	saved := FAASettings{}
	if err := json.Unmarshal(body, &saved); err != nil {
		return nil, err
	}
	return &saved, nil
}
