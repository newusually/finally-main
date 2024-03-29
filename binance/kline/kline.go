package kline

import (
	"context"
	"encoding/csv"
	"fmt"
	"os"
	"time"

	binance_connector "github.com/binance/binance-connector-go"
)

func getKlines(symbol string) []*binance_connector.KlinesResponse {
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient("", "", baseURL)

	// Klines
	klines, err := client.NewKlinesService().
		Symbol(symbol).Interval("15m").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return klines
}

func SaveKlines(symbol string) {

	// Create a CSV file
	file, err := os.Create("../datas/other/" + symbol + ".csv")
	if err != nil {
		fmt.Println("Cannot create file", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	writer.Write([]string{"date", "open", "high", "low", "close", "volume"})

	klines := getKlines(symbol)

	// Write data
	for _, kline := range klines {
		// Convert Unix timestamp to Beijing time
		openTime := time.Unix(int64(kline.OpenTime/1000), 0).In(time.FixedZone("CST", 8*3600)).Format("2006-01-02 15:04:05")

		writer.Write([]string{
			openTime,
			kline.Open,
			kline.High,
			kline.Low,
			kline.Close,
			kline.Volume,
		})
	}
}
