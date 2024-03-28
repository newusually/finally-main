package userinfo

import (
	"encoding/json"
	"fmt"
	"os"
)

type ApiKeys struct {
	APIKey    string `json:"api_key"`
	SecretKey string `json:"secret_key"`
}

func GetUserinfo() (string, string) {
	file, err := os.Open("../datas/api_bn.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return "", ""
	}
	defer file.Close()

	var keys ApiKeys
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&keys)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return "", ""
	}

	fmt.Println("API Key:", keys.APIKey)
	fmt.Println("Secret Key:", keys.SecretKey)
	return keys.APIKey, keys.SecretKey
}
