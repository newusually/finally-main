package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type ServerTimeResponse struct {
	Code string `json:"code"`
	Data []struct {
		Ts string `json:"ts"`
	} `json:"data"`
	Msg string `json:"msg"`
}

func getServerTime() (time.Time, error) {
	url := "https://www.okx.com/api/v5/public/time"
	resp, err := http.Get(url)
	if err != nil {
		return time.Time{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return time.Time{}, err
	}

	var serverTimeResponse ServerTimeResponse
	err = json.Unmarshal(body, &serverTimeResponse)
	if err != nil {
		return time.Time{}, err
	}

	// Convert the timestamp from string to int64
	ts, err := strconv.ParseInt(serverTimeResponse.Data[0].Ts, 10, 64)
	if err != nil {
		return time.Time{}, err
	}

	// Convert the timestamp from milliseconds to seconds and create a time.Time object
	serverTime := time.Unix(0, ts*int64(time.Millisecond))
	fmt.Println(serverTime)

	return serverTime, nil
}

func main() {

	url := "https://www.okx.com/api/v5/account/balance?ccy=USDT"
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("OK-ACCESS-KEY", "118c7c11-bfc4-4d09-9811-eb08a86336ba")
	req.Header.Add("OK-ACCESS-SIGN", "9gvnDUubt7GLvQeA8NqKfuRD6rmN5eh21yKxGqkY5pI=")
	req.Header.Add("OK-ACCESS-TIMESTAMP", "0")
	req.Header.Add("OK-ACCESS-PASSPHRASE", "USUA51tomato58")
	req.Header.Add("x-simulated-trading", "0")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error on response.\n[ERRO] -", err)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
