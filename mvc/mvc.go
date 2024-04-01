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

func GetKline(symbol string, minute string) (bool, float64, float64, int, []float64, []float64, []float64, []float64, []float64) {

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

	if len(closes) < 500 || closes[10].Float() < 0.0001 {
		return false, 0, 0, 0, []float64{}, []float64{}, []float64{}, []float64{}, []float64{}
	}

	if len(closes) > 500 && closes[10].Float() > 0.0001 {

		day := make([]string, len(dates))
		c := make([]float64, len(closes))
		o := make([]float64, len(opens))
		h := make([]float64, len(highs))

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

		}

		x := len(c)

		diff, dea, _ := talib.Macd(c, 12, 26, 60)
		choose_ma := c[x-1] > talib.Sma(c, 5)[x-1]

		macd1 := 2 * (diff[x-1] - dea[x-1])
		macd2 := 2 * (diff[x-2] - dea[x-2])
		return choose_ma, macd1, macd2, x, c, o, h, diff, dea
	}
	return false, 0, 0, 0, []float64{}, []float64{}, []float64{}, []float64{}, []float64{}
}
func Getprice(symbol string, minute string) {

	choose_ma, macd1, macd2, x, c, o, h, diff, dea := GetKline(symbol, minute)
	if x < 500 {
		return
	}
	//fmt.Println(symbol, minute, x)
	//fmt.Println("symbol--->>>", symbol, "minute--->>>", minute, "choose_ma--->>>", choose_ma, "macd1--->>>", macd1, "macd2--->>>", macd2)
	if c[x-1]/o[x-1] > 1.0015 && c[x-1]/o[x-1] < 1.015 && choose_ma &&
		h[x-1]/c[x-1] < 1.005 && h[x-2]/c[x-2] < 1.005 && diff[x-1] > 0 && dea[x-1] > 0 && macd1 > 0 && macd1/macd2 > 1.5 {
		y := "\n----time--->>" + time.Now().Format("2006-1-2 15:04:02") +
			",symbol----->>>" + symbol +
			",----close1/open1--->>" + strconv.FormatFloat(c[x-1]/o[x-1], 'f', 5, 64) +
			",----close2/open2--->>" + strconv.FormatFloat(c[x-2]/o[x-2], 'f', 5, 64) +
			",----h[x-1]/c[x-1]--->>" + strconv.FormatFloat(h[x-1]/c[x-1], 'f', 5, 64) +
			",----h[x-2]/c[x-2]--->>" + strconv.FormatFloat(h[x-2]/c[x-2], 'f', 5, 64) +
			",----macd1--->>" + strconv.FormatFloat(macd1, 'f', 3, 64) +
			",----macd1/macd2--->>" + strconv.FormatFloat(macd1/macd2, 'f', 3, 64) +
			",----diff--->>" + strconv.FormatFloat(diff[x-1], 'f', 3, 64) +
			",----dea--->>" + strconv.FormatFloat(dea[x-1], 'f', 3, 64) +
			",----minute--->>" + minute
		fmt.Println(y)

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
		write.WriteString(y)

		//Flush将缓存的文件真正写入到文件中
		write.Flush()

		fmt.Println("-------------------------------买入--------------------------------->>>")

		cmd := exec.Command("python", "gorun.py", symbol, minute)
		res, _ := cmd.Output()
		fmt.Println(string(res))
		//SendDingMsg(y)
	}

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
