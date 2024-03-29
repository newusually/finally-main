package account

import (
	"context"
	"fmt"
	binance_connector "github.com/binance/binance-connector-go"
)

func GetAccount(apiKey string, secretKey string, baseURL string) string {

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// Binance Account Information (USER_DATA) - GET /api/v3/account
	accountInformation, err := client.NewGetAccountService().Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return ""
	}
	accountInfo := binance_connector.PrettyPrint(accountInformation)

	return accountInfo
}
