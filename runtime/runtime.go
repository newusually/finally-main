package runtime

import (
	"finally-main/mvc"
	"fmt"
	"strings"
	"time"
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
	for _, v := range symbollist {
		symbol := strings.Replace(v.String(), "-USDT-SWAP", "", -1)

		isbuy, buysale1, buysale2, buysale3, buysale1_2, buysale2_3 := mvc.GetIsBuy(symbol, minute)
		if isbuy {
			y := "\n----time--->>" + time.Now().Format("2006-1-2 15:04:02") +
				",symbol----->>>" + symbol + "-USDT-SWAP" +
				",----buysale1--->>" + buysale1 +
				",----buysale2--->>" + buysale2 +
				",----buysale3--->>" + buysale3 +
				",----buysale1_2--->>" + buysale1_2 +
				",----buysale2_3--->>" + buysale2_3 +
				",----minute--->>" + minute
			fmt.Println(y)
			mvc.GetWriter(y)
			mvc.Buy(symbol+"-USDT-SWAP", minute)
		}
	}

}
