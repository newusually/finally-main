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
	symbollist := mvc.Getsymbols()
	//fmt.Println(symbollist)

	uniqueSymbols := make(map[string]bool)
	re := regexp.MustCompile(`-(USDT)-SWAP$`)

	for _, v := range symbollist {
		symbol := re.ReplaceAllString(v.String(), "")

		if _, exists := uniqueSymbols[symbol]; !exists {
			uniqueSymbols[symbol] = true

			isbuy := mvc.GetIsBuy(symbol+"-USDT-SWAP", minute)
			if isbuy {
				mvc.Buy(symbol+"-USDT-SWAP", minute)
			}

		}
	}

}
