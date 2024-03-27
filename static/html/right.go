package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

// parseCSV parses the provided CSV file and returns a slice of Candlesticks
func parseCSV(filePath string) ([][5]interface{}, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var candlesticks [][5]interface{}
	for _, record := range records[1:] { // Skipping the header line
		// Create a Candlestick struct from the record and append it to the slice
		date := record[0]
		open := parseToFloat(record[2])
		high := parseToFloat(record[3])
		low := parseToFloat(record[4])
		close := parseToFloat(record[1])
		candlestick := [5]interface{}{date, open, close, low, high}
		candlesticks = append(candlesticks, candlestick) // Use the correct variable name
	}
	return candlesticks, nil
}

// parseToFloat takes a string and converts it to a float64
func parseToFloat(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}
func main() {
	const symbol = "ETH-USDT-SWAP"
	filePath := fmt.Sprintf("../../datas/old_data/%s/%s-15min.csv", symbol, symbol)
	candlesticks, _ := parseCSV(filePath)
	fmt.Println(candlesticks)
}
