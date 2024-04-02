package runtime

import (
	"finally-main/mvc"
	"fmt"
	"regexp"
)

func Run1() {
	Run("1m")
}

func Run3() {
	Run("3m")
}

func Run5() {
	Run("5m")
}

func Run15() {
	Run("15m")
}

func Run1H() {
	Run("1H")
}

func Run2H() {
	Run("2H")
}

func Run4H() {
	Run("4H")
}

func Savecsvfinal() {
	//mvc.Savecsv("1m")
	//mvc.Savecsv("3m")
	//mvc.Savecsv("5m")
	mvc.Savecsv("15m")

}

func Run6H() {
	Run("6H")

}

func Run12H() {
	mvc.Savecsv("1H")
	mvc.Savecsv("2H")
	mvc.Savecsv("4H")
	mvc.Savecsv("6H")
	mvc.Savecsv("12H")
	mvc.Savecsv("1D")
}

func Run(minute string) {
	defer func() {
		if r := recover(); r != nil {
			// 处理异常
			fmt.Println("Exception caught:", r)
		}
	}()

	_, _, macdEth15m, _, _, _, _, _, _ := mvc.GetKline("ETH-USDT-SWAP", "15m")

	if macdEth15m > 0 {
		symbollist := mvc.Getsymbols()

		for i := 0; i < len(symbollist); i++ {
			symbol := symbollist[i].String()
			re := regexp.MustCompile(`^(\w+)-`)
			match := re.FindStringSubmatch(symbol)

			if len(match) > 1 {
				symbol = match[1] + "-USDT-SWAP"
			} else {
				fmt.Println("No match found")
			}

			choose := symbol != "USTC-USDT-SWAP" && symbol != "USDC-USDT-SWAP" &&
				symbol != "BTC-USDT-SWAP" && symbol != "ETH-USDT-SWAP" && symbol != "ETC-USDT-SWAP" &&
				symbol != "BCH-USDT-SWAP" && symbol != "DOGE-USDT-SWAP" && symbol != "SOL-USDT-SWAP" &&
				symbol != "XRP-USDT-SWAP" && symbol != "AVAX-USDT-SWAP" && symbol != "BSV-USDT-SWAP" &&
				symbol != "OP-USDT-SWAP" && symbol != "LTC-USDT-SWAP" && symbol != "ADA-USDT-SWAP" &&
				symbol != "LINK-USDT-SWAP" && symbol != "TRX-USDT-SWAP" && symbol != "MKR-USDT-SWAP"

			if choose {
				mvc.Getprice(symbol, minute)
			}

		}
	}

}
