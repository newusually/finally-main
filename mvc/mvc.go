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

func GetIsBuy(symbol string, minute string) bool {

	time.Sleep(time.Millisecond * 10)
	x, close, vol1, vol2, diff, dea := GetKline(symbol, minute)

	if x < 500 {
		return false
	} else {

		macd1 := 2 * (diff[x-1] - dea[x-1]) * (10000 / close)
		macd2 := 2 * (diff[x-2] - dea[x-2]) * (10000 / close)

		if macd1 > macd2 && macd1 > 40 && macd1 < 60 && macd1/macd2 > 1.1 && macd1/macd2 < 1.5 && vol1 > vol2 {
			log := "\ntime--->>" + time.Now().Format("2006-1-2 15:04:02") +
				",symbol--->>" + symbol +
				",macd1--->>" + fmt.Sprintf("%.5f", macd1) +
				",macd2--->>" + fmt.Sprintf("%.5f", macd2) +
				",macd1/macd2--->>" + fmt.Sprintf("%.5f", macd1/macd2) +
				",close--->>" + fmt.Sprintf("%.5f", close) +
				",vol1--->>" + fmt.Sprintf("%.5f", vol1) +
				",vol2--->>" + fmt.Sprintf("%.5f", vol2) +
				",diff--->>" + fmt.Sprintf("%.5f", diff[x-1]) +
				",dea--->>" + fmt.Sprintf("%.5f", dea[x-1]) +
				",minute--->>" + minute
			fmt.Println(log)
			GetWriter(log)
			return true
		} else {
			return false
		}
	}

}

func GetKline(symbol string, minute string) (int, float64, float64, float64, []float64, []float64) {

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
		return 0, 0, 0, 0, []float64{}, []float64{}
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

		vol1 := v[x-1]
		vol2 := v[x-2]
		return x, c[x-1], vol1, vol2, diff, dea
	}
	return 0, 0, 0, 0, []float64{}, []float64{}
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
