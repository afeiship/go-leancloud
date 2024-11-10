package leancloud

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

// func Set(key string, value interface{}) {}

var lcId = os.Getenv("LEANCLOUD_ID")
var lcKey = os.Getenv("LEANCLOUD_KEY")
var lcHost = os.Getenv("LEANCLOUD_HOST")
var apiBaseURL = "https://" + lcHost + "/1.1/classes/options"

type LcResult struct {
	Key       string `json:"key"`
	Value     any    `json:"value"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	ObjectId  string `json:"objectId"`
}

func Get(key string) LcResult {
	client := &http.Client{}
	resURL := apiBaseURL + "/" + key
	req, err := http.NewRequest("GET", resURL, nil)
	if err != nil {
		log.Fatal(err)
	}

	// add headers
	req.Header.Add("X-LC-Id", lcId)
	req.Header.Add("X-LC-Key", lcKey)
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	jsonBody := LcResult{}

	err = json.Unmarshal(body, &jsonBody)
	if err != nil {
		log.Fatal(err)
	}

	return jsonBody
}

// get Value
func Value(key string) any {
	result := Get(key)
	return result.Value
}
