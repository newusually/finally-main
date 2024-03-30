package client

import (
	"bytes"
	"encoding/json"
	"finally-main/okx-go/consts"
	"finally-main/okx-go/utils"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	APIKey        string
	APISecretKey  string
	Passphrase    string
	UseServerTime bool
	Flag          string
}

func NewClient(apiKey, apiSecretKey, passphrase string, useServerTime bool, flag string) *Client {
	return &Client{
		APIKey:        apiKey,
		APISecretKey:  apiSecretKey,
		Passphrase:    passphrase,
		UseServerTime: useServerTime,
		Flag:          flag,
	}
}

func (c *Client) Request(method, requestPath string, params map[string]string) (map[string]interface{}, error) {
	if method == consts.GET {
		requestPath = requestPath + utils.ParseParamsToStr(params)
	}

	url := consts.API_URL + requestPath

	timestamp := utils.GetTimestamp()

	if c.UseServerTime {
		timestamp = c.GetServerTimestamp()
	}

	body := ""
	if method == consts.POST {
		jsonBody, _ := json.Marshal(params)
		body = string(jsonBody)
	}

	sign := utils.Sign(utils.PreHash(timestamp, method, requestPath, body), c.APISecretKey)
	header := utils.GetHeader(c.APIKey, sign, timestamp, c.Passphrase, c.Flag)

	var response map[string]interface{}
	fmt.Println("url:", url, sign, header)

	req, err := http.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	if err != nil {
		return response, err
	}

	for key, value := range header {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(bodyBytes, &response)

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return response, fmt.Errorf("non-2xx HTTP code: %d", resp.StatusCode)
	}

	return response, nil
}

func (c *Client) GetServerTimestamp() string {
	url := consts.API_URL + consts.SERVER_TIMESTAMP_URL
	resp, err := http.Get(url)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	var response map[string]interface{}
	json.Unmarshal(bodyBytes, &response)

	return response["ts"].(string)
}
