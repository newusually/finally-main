package mvc

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"github.com/markcheno/go-talib"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func Getsymbols() []gjson.Result {
	time.Sleep(time.Millisecond)

	// 获取当前时间 或者使用 time.Date(year, month, ...)
	t := time.Now()
	timeStamp := t.Unix()
	client := &http.Client{

		Transport: &http.Transport{

			TLSNextProto: map[string]func(string, *tls.Conn) http.RoundTripper{},
		},
	}
	req, err := http.NewRequest("GET", "https://www.okx.com/priapi/v5/public/simpleProduct?t="+strconv.Itoa(int(timeStamp))+"&instType=SWAP", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("authority", "www.okx.com")
	req.Header.Set("timeout", "10000")
	req.Header.Set("x-cdn", "https://static.okx.com")
	req.Header.Set("devid", "8ccf140e-e4ab-4a46-8582-738445cad57c")
	req.Header.Set("accept-language", "zh-CN")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.87 Safari/537.36 SE 2.X MetaSr 1.0")
	req.Header.Set("accept", "application/json")
	req.Header.Set("x-utc", "8")
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("app-type", "web")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("referer", "https://www.okx.com/trade-swap/btc-usdt-swap")
	req.Header.Set("cookie", "locale=zh_CN; defaultLocale=zh_CN; _gcl_au=1.1.2108489782."+strconv.Itoa(int(timeStamp))+"; _ga=GA1.2.1119126875."+strconv.Itoa(int(timeStamp))+"; _gid=GA1.2.1241734301."+strconv.Itoa(int(timeStamp))+"; amp_56bf9d=4KsK9IqNpGzx7Sx3l_9DPR...1g2ehgo0j.1g2ej7b4i.4.0.4")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)

	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)

	}
	src := string(bodyText)
	//fmt.Println(src)
	//"instType":"SWAP","instId":"GODS-USDT-SWAP","last":"0.7141","lastSz":"13",
	instId := gjson.Get(src, "data.#.instId").Array()
	//fmt.Println(instId)
	return instId
}

func GetIsBuy(symbol string, minute string) (bool, string, string, string, string, string) {

	time.Sleep(time.Millisecond * 10)

	// 获取当前时间 或者使用 time.Date(year, month, ...)
	t := time.Now()
	timeStamp := t.Unix()
	client := &http.Client{

		Transport: &http.Transport{

			TLSNextProto: map[string]func(string, *tls.Conn) http.RoundTripper{},
		},
	}
	req, err := http.NewRequest("GET", "https://www.okx.com/priapi/v5/rubik/stat/taker-volume?instType=CONTRACTS&period="+minute+"&ccy="+symbol+"&t="+strconv.Itoa(int(timeStamp)), nil)
	if err != nil {
		panic(err)

	}
	req.Header.Set("authority", "www.okx.com")
	req.Header.Set("timeout", "10000")
	req.Header.Set("x-cdn", "https://www.okx.com")
	req.Header.Set("devid", "a5ceb850-4efb-4a3f-baff-21da4fce8858")
	req.Header.Set("accept-language", "zh-CN")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.87 Safari/537.36 SE 2.X MetaSr 1.0")
	req.Header.Set("accept", "application/json")
	req.Header.Set("x-utc", "8")
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("app-type", "web")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("referer", "https://www.okx.com/zh-hans/markets/data/contracts")
	req.Header.Set("cookie", "locale=zh_CN; defaultLocale=zh_CN; _gcl_au=1.1.1520807996."+strconv.Itoa(int(timeStamp))+"; _ga=GA1.2.1752370991."+strconv.Itoa(int(timeStamp))+"; _gid=GA1.2.650161560."+strconv.Itoa(int(timeStamp))+"; amp_56bf9d=y9J2I5hN4sKjIiyZROsSAs...1g1isehgq.1g1isehgs.2.0.2; _gat_UA-35324627-3=1")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)

	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	src := string(bodyText)

	dates := gjson.Get(src, "data.#.0").Array()

	sales := gjson.Get(src, "data.#.1").Array()
	buys := gjson.Get(src, "data.#.2").Array()

	day := make([]string, len(dates))
	sale := make([]float64, len(sales))
	buy := make([]float64, len(buys))

	for i := 0; i < len(dates); i++ {

		a := dates[len(dates)-i-1].Str
		e, _ := strconv.ParseInt(a, 10, 64)
		day[i] = time.Unix(0, e*int64(time.Millisecond)).Format("2006-01-02 15:04:05")

		b := sales[len(sales)-i-1].Str
		f, _ := strconv.ParseFloat(b, 64)
		sale[i] = f

		d := buys[len(buys)-i-1].Str
		g, _ := strconv.ParseFloat(d, 64)
		buy[i] = g
	}
	x := len(day)
	if x > 30 {
		var sumBuy1, sumSale1 float64
		var sumBuy2, sumSale2 float64
		var sumBuy3, sumSale3 float64

		sumBuy1 = buy[x-1]
		sumSale1 = sale[x-1]
		sumBuy2 = buy[x-2]
		sumSale2 = sale[x-2]
		sumBuy3 = buy[x-3]
		sumSale3 = sale[x-3]

		buySaleStr1 := fmt.Sprintf("%.5f", sumBuy1/sumSale1)
		buySale1, _ := strconv.ParseFloat(buySaleStr1, 64)

		buySaleStr2 := fmt.Sprintf("%.5f", sumBuy2/sumSale2)
		buySale2, _ := strconv.ParseFloat(buySaleStr2, 64)

		buySaleStr3 := fmt.Sprintf("%.5f", sumBuy3/sumSale3)
		buySale3, _ := strconv.ParseFloat(buySaleStr3, 64)

		// Parse the date string to time.Time
		layout := "2006-01-02 15:04:05"
		ts, _ := time.Parse(layout, day[x-1])

		// Get the current time
		now := time.Now()

		// Calculate the time 15 minutes before and after the current time
		fifteenMinutesBefore := now.Add(-(15) * time.Minute)
		fifteenMinutesAfter := now.Add((8 * 60) * time.Minute)

		// Check if t is between fifteenMinutesBefore and fifteenMinutesAfter
		if ts.After(fifteenMinutesBefore) && ts.Before(fifteenMinutesAfter) {
			if buySale1 > 0.5 && buySale1 < 0.7 && buySale1 > buySale2 && buySale2 > buySale3 {
				return true, fmt.Sprintf("%.5f", buySale1), fmt.Sprintf("%.5f", buySale2), fmt.Sprintf("%.5f", buySale3), fmt.Sprintf("%.5f", buySale1/buySale2), fmt.Sprintf("%.5f", buySale2/buySale3)
			} else {
				return false, fmt.Sprintf("%.5f", buySale1), fmt.Sprintf("%.5f", buySale2), fmt.Sprintf("%.5f", buySale3), fmt.Sprintf("%.5f", buySale1/buySale2), fmt.Sprintf("%.5f", buySale2/buySale3)
			}

		}

	}
	return false, "0", "0", "0", "0", "0"
}

func GetKline(symbol string, minute string) (bool, float64, float64, float64, int, []float64, float64, float64) {

	//fmt.Println(symboldemo, symbol)
	time.Sleep(time.Millisecond * 10)

	// 获取当前时间 或者使用 time.Date(year, month, ...)
	t := time.Now()
	timeStamp := t.Unix()
	client := &http.Client{

		Transport: &http.Transport{

			TLSNextProto: map[string]func(string, *tls.Conn) http.RoundTripper{},
		},
	}
	req, err := http.NewRequest("GET", "https://www.okx.com/priapi/v5/market/candles?instId="+symbol+"&bar="+minute+"&after=&limit=1400&t="+strconv.Itoa(int(timeStamp)), nil)
	if err != nil {
		panic(err)

	}
	req.Header.Set("authority", "www.okx.com")
	req.Header.Set("timeout", "10000")
	req.Header.Set("x-cdn", "https://static.okx.com")
	req.Header.Set("devid", "a5ceb850-4efb-4a3f-baff-21da4fce8858")
	req.Header.Set("accept-language", "zh-CN")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.87 Safari/537.36 SE 2.X MetaSr 1.0")
	req.Header.Set("accept", "application/json")
	req.Header.Set("x-utc", "8")
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("app-type", "web")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("referer", "https://www.okx.com/trade-swap/btc-usdt-swap")
	req.Header.Set("cookie", "locale=zh_CN; defaultLocale=zh_CN; _gcl_au=1.1.1520807996."+strconv.Itoa(int(timeStamp))+"; _ga=GA1.2.1752370991."+strconv.Itoa(int(timeStamp))+"; _gid=GA1.2.650161560."+strconv.Itoa(int(timeStamp))+"; amp_56bf9d=y9J2I5hN4sKjIiyZROsSAs...1g1isehgq.1g1isehgs.2.0.2; _gat_UA-35324627-3=1")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)

	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	src := string(bodyText)
	//"date", "open", "high", "low", "close", "p", "vol"
	closes := gjson.Get(src, "data.#.4").Array()
	dates := gjson.Get(src, "data.#.0").Array()
	opens := gjson.Get(src, "data.#.1").Array()
	highs := gjson.Get(src, "data.#.2").Array()
	vols := gjson.Get(src, "data.#.6").Array()

	if len(closes) < 500 || closes[10].Float() < 0.0001 {
		return false, 0, 0, 0, 0, []float64{}, 0, 0
	}

	if len(closes) > 500 && closes[10].Float() > 0.0001 {

		day := make([]string, len(dates))
		c := make([]float64, len(closes))
		o := make([]float64, len(opens))
		h := make([]float64, len(highs))
		v := make([]float64, len(vols))

		for i := 0; i < len(closes); i++ {

			a := dates[len(dates)-i-1].Str
			e, _ := strconv.ParseInt(a, 10, 64)
			day[i] = time.Unix(0, e*int64(time.Millisecond)).Format("2006-01-02 15:04:05")

			b := closes[len(closes)-i-1].Str
			f, _ := strconv.ParseFloat(b, 64)
			c[i] = f

			d := opens[len(opens)-i-1].Str
			g, _ := strconv.ParseFloat(d, 64)
			o[i] = g

			p := highs[len(highs)-i-1].Str
			rs, _ := strconv.ParseFloat(p, 64)
			h[i] = rs

			vv := vols[len(vols)-i-1].Str
			vv1, _ := strconv.ParseFloat(vv, 64)
			v[i] = vv1

		}

		x := len(c)

		diff, dea, _ := talib.Macd(c, 12, 26, 60)
		choose_ma := c[x-1] > talib.Sma(c, 5)[x-1]

		macd1 := 2 * (diff[x-1] - dea[x-1])
		macd2 := 2 * (diff[x-2] - dea[x-2])
		macd3 := 2 * (diff[x-3] - dea[x-3])

		vol1 := v[x-1]
		vol2 := v[x-2]
		return choose_ma, macd1, macd2, macd3, x, c, vol1, vol2
	}
	return false, 0, 0, 0, 0, []float64{}, 0, 0
}
func Getprice(symbol string, minute string) {

	_, macd1, macd2, macd3, x, c, vol1, vol2 := GetKline(symbol, minute)
	if x < 500 {
		return
	}

	if vol1 > vol2 {
		return

	}

	//布林曲线

	upper, middle, _ := talib.BBands(c, 26, 2, 2, 0)

	// Slope斜率
	slope := talib.LinearRegSlope(talib.Sma(c, 5), 5)
	cosa45 := slope[len(slope)-1]
	backCosa45 := slope[len(slope)-2]
	if cosa45 < 0 || backCosa45 < 0 {
		return
	}
	//fmt.Println("symbol:--->>>", symbol, "----cosa45/backCosa45--->>>", cosa45/backCosa45, "----upper/middle--->>>", upper[x-1]/middle[x-1], minute)
	if macd1 > 0 && macd1/macd2 > 1.03 && macd1/macd2 < 2 && macd2 > macd3 &&
		cosa45/backCosa45 > 1.1 && cosa45/backCosa45 < 2 &&
		upper[x-1]/middle[x-1] > 1.01 && upper[x-1]/middle[x-1] < 2 {
		y := "\n----time--->>" + time.Now().Format("2006-1-2 15:04:02") +
			",symbol----->>>" + symbol +
			",----cosa45/backCosa45--->>" + strconv.FormatFloat(cosa45/backCosa45, 'f', 5, 64) +
			",----upper/middle--->>" + strconv.FormatFloat(upper[x-1]/middle[x-1], 'f', 5, 64) +
			",----macd1/macd2--->>" + strconv.FormatFloat(macd1/macd2, 'f', 3, 64) +
			",----vol1/vol2--->>" + strconv.FormatFloat(vol1/vol2, 'f', 3, 64) +
			",----minute--->>" + minute
		fmt.Println(y)

		//SendDingMsg(y)
	}

}

func GetWriter(log string) {
	filePath := "..\\datas\\log\\buylog.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	//及时关闭file句柄
	defer file.Close()
	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)

	write.WriteString("\n")
	write.WriteString(log)

	//Flush将缓存的文件真正写入到文件中
	write.Flush()

}

func Buy(symbol string, minute string) {
	fmt.Println("-------------------------------买入--------------------------------->>>")

	cmd := exec.Command("python", "gorun.py", symbol, minute)
	res, _ := cmd.Output()
	fmt.Println(string(res))
}

func Getcashbal() {
	cmd := exec.Command("python", "cash.py")
	res, _ := cmd.Output()
	fmt.Println(string(res))
}

func Getcashhistory() {
	cmd := exec.Command("python", "cashhistory.py")
	res, _ := cmd.Output()
	fmt.Println(string(res))
}

func GetuplRatio() {
	cmd := exec.Command("python", "getuplRatio.py")
	res, _ := cmd.Output()
	fmt.Println(string(res))
}

func Savecsv(minute string) {
	cmd := exec.Command("python", "savecsv.py", minute)
	res, _ := cmd.Output()
	fmt.Println(string(res))
}

func SendDingMsg(msg string) {
	//请求地址模板
	webHook := `https://oapi.dingtalk.com/robot/send?access_token=f8195c9e4ad6da4427d67e80dffed5d07ecaca1d1e79462fb5c0a9c6b12e90f2`
	content := `{"msgtype": "text",
        "text": {"content": "` + msg + `"}
    }`
	//创建一个请求
	req, err := http.NewRequest("POST", webHook, strings.NewReader(content))
	if err != nil {
		// handle error
	}

	client := &http.Client{}
	//设置请求头
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	//发送请求
	resp, err := client.Do(req)
	//关闭请求
	defer resp.Body.Close()

	if err != nil {
		// handle error
	}
}
