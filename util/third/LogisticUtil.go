package third

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"github.com/tidwall/gjson"
	"io"
	"net/http"
	"net/url"
	"strings"
)

//快递鸟

var kdUrl string
var uid string
var ApiKey string

func InitKd() {
	kdUrl = viper.GetString("kd.url")
	uid = viper.GetString("kd.uid")
	ApiKey = viper.GetString("kd.apiKey")
	//GetLogistic("EMS", "1212405644905")
}
func GetLogistic(code, num string) []Trace {
	// 组装应用级参数
	RequestData := fmt.Sprintf(`{"ShipperCode":"%s","LogisticCode":"%s"}`, code, num)

	sign := Sign(RequestData)
	// 组装系统级参数
	params := map[string]string{
		"RequestType": "8001",
		"EBusinessID": uid,
		"DataType":    "2",
		"RequestData": RequestData,
		"DataSign":    sign,
	}
	result := gjson.Parse(post(kdUrl, params))
	fmt.Println(result)
	if result.Get("Success").Bool() {
		var list []Trace
		json.Unmarshal([]byte(result.Get("Traces").String()), &list)
		return list
	}
	return nil
}

func Sign(n string) (data string) {
	str := n + ApiKey
	w := md5.New()
	io.WriteString(w, str)
	sign := base64.StdEncoding.EncodeToString([]byte(hex.EncodeToString(w.Sum(nil))))
	return sign
}

func post(u string, params map[string]string) string {
	var values []string
	for k, v := range params {
		values = append(values, fmt.Sprintf("%s=%s", k, url.QueryEscape(v)))
	}
	fmt.Println(values)
	resp, err := http.Post(u, "application/x-www-form-urlencoded", strings.NewReader(strings.Join(values, "&")))
	if err != nil || resp.StatusCode != 200 {
		fmt.Println(err.Error())
	}
	contentBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	return string(contentBytes)
}

type Trace struct {
	AcceptStation string `json:"acceptStation"`
	AcceptTime    string `json:"acceptTime"`
	Location      string `json:"location"`
}
