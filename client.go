package elves

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Client holds all the information required to connect to a server
type Client struct {
	HostName     string
	APIKeyName   string
	APIKeySecret string
	httpClient   *http.Client
}

func NewClient(hostname string, apikeyname string, apikeysecret string) *Client {
	return &Client{
		HostName:     hostname,
		APIKeyName:   apikeyname,
		APIKeySecret: apikeysecret,
		httpClient: &http.Client{
			Timeout: 60 * time.Second,
		},
	}
}

// 200 response
func (c *Client) RequestResponse200(req *http.Request) ([]byte, error) {
	req.SetBasicAuth(c.APIKeyName, c.APIKeySecret)
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		resBody := new(bytes.Buffer)
		_, err = resBody.ReadFrom(res.Body)
		if err != nil {
			return nil, fmt.Errorf("got a non 200 status code: %v", res.StatusCode)
		}
		return nil, fmt.Errorf("got a non 200 status code: %v - %s - %s", res.StatusCode, req.URL, resBody.String())
	}

	return body, nil
}

// 201 response
func (c *Client) RequestResponse201(req *http.Request) ([]byte, error) {
	req.SetBasicAuth(c.APIKeyName, c.APIKeySecret)
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusCreated {
		resBody := new(bytes.Buffer)
		_, err = resBody.ReadFrom(res.Body)
		if err != nil {
			return nil, fmt.Errorf("got a non 201 status code: %v", res.StatusCode)
		}
		return nil, fmt.Errorf("got a non 201 status code: %v - %s - %s", res.StatusCode, req.URL, resBody.String())
	}

	return body, nil
}

// 204 response
func (c *Client) RequestResponse204(req *http.Request) ([]byte, error) {
	req.SetBasicAuth(c.APIKeyName, c.APIKeySecret)
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusNoContent {
		resBody := new(bytes.Buffer)
		_, err = resBody.ReadFrom(res.Body)
		if err != nil {
			return nil, fmt.Errorf("got a non 204 status code: %v", res.StatusCode)
		}
		return nil, fmt.Errorf("got a non 204 status code: %v - %s", res.StatusCode, resBody.String())
	}

	return body, nil
}
