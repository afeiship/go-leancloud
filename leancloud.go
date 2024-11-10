package leancloud

import (
	"bytes"
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
	Value     string `json:"value"`
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

	// try to unmarshal value
	var value string
	if err := json.Unmarshal([]byte(jsonBody.Value), &value); err == nil {
		jsonBody.Value = value
	}
	return jsonBody, nil
}

// Value fetches only the value associated with a given key
func (c *LeanCloudClient) Value(key string) (string, error) {
	result, err := c.Get(key)
	if err != nil {
		return "", err
	}
	return result.Value, nil
}

// set
func (c *LeanCloudClient) Set(key string, value any) error {
	resurl := fmt.Sprintf("%s/%s", c.baseURL, key)

	// 将 value 转换为 JSON 字符串
	jsonValue, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("failed to marshal value to JSON string: %w", err)
	}

	body := map[string]any{
		"value": string(jsonValue),
	}

	jsonData, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	req, err := http.NewRequest("PUT", resurl, bytes.NewReader(jsonData))
	if err != nil {
		return err
	}

	// Add headers
	req.Header.Add("X-LC-Id", c.lcId)
	req.Header.Add("X-LC-Key", c.lcKey)
	req.Header.Add("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return errors.New("failed to put data, status code: " + resp.Status)
	}

	return nil
}
