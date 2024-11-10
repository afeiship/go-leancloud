package leancloud

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

// LeanCloudClient is a struct to hold necessary configurations and an HTTP client
type LeanCloudClient struct {
	lcId    string
	lcKey   string
	baseURL string
	client  *http.Client
}

// New creates a new instance of LeanCloudClient
func New() *LeanCloudClient {
	return &LeanCloudClient{
		lcId:    os.Getenv("LEANCLOUD_ID"),
		lcKey:   os.Getenv("LEANCLOUD_KEY"),
		baseURL: "https://" + os.Getenv("LEANCLOUD_HOST") + "/1.1/classes/options",
		client:  &http.Client{},
	}
}

type LcResult struct {
	Key       string `json:"key"`
	Value     any    `json:"value"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	ObjectId  string `json:"objectId"`
}

// Get retrieves a key-value pair from LeanCloud
func (c *LeanCloudClient) Get(key string) (LcResult, error) {
	resurl := fmt.Sprintf("%s/%s", c.baseURL, key)
	req, err := http.NewRequest("GET", resurl, nil)
	if err != nil {
		return LcResult{}, err
	}

	// Add headers
	req.Header.Add("X-LC-Id", c.lcId)
	req.Header.Add("X-LC-Key", c.lcKey)
	req.Header.Add("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return LcResult{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return LcResult{}, errors.New("failed to retrieve data, status code: " + resp.Status)
	}

	// Parse response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return LcResult{}, err
	}

	var jsonBody LcResult
	if err := json.Unmarshal(body, &jsonBody); err != nil {
		return LcResult{}, err
	}

	return jsonBody, nil
}

// Value fetches only the value associated with a given key
func (c *LeanCloudClient) Value(key string) (any, error) {
	result, err := c.Get(key)
	if err != nil {
		return nil, err
	}
	return result.Value, nil
}
