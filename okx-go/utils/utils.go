package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"finally-main/okx-go/consts"
	"fmt"
	"strings"
	"time"
)

func Sign(message, secretKey string) string {
	mac := hmac.New(sha256.New, []byte(secretKey))
	mac.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}

func PreHash(timestamp, method, requestPath, body string) string {
	return fmt.Sprintf("%s%s%s%s", timestamp, strings.ToUpper(method), requestPath, body)
}

func GetHeader(apiKey, sign, timestamp, passphrase, flag string) map[string]string {
	header := make(map[string]string)
	header[consts.CONTENT_TYPE] = consts.APPLICATION_JSON
	header[consts.OK_ACCESS_KEY] = apiKey
	header[consts.OK_ACCESS_SIGN] = sign
	header[consts.OK_ACCESS_TIMESTAMP] = timestamp
	header[consts.OK_ACCESS_PASSPHRASE] = passphrase
	header["x-simulated-trading"] = flag
	return header
}

func ParseParamsToStr(params map[string]string) string {
	var strParams []string
	for key, value := range params {
		strParams = append(strParams, fmt.Sprintf("%s=%s", key, value))
	}
	return "?" + strings.Join(strParams, "&")
}

func GetTimestamp() string {
	return time.Now().UTC().Format(time.RFC3339)
}
