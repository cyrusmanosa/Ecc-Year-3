package main

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	getData, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=Tokyo,jp&units=metric&appid=[appid]") // [appid]は実際の値に置換して使用
	if err != nil {
		fmt.Println(err)
		return
	}
	defer getData.Body.Close()

	jsonData, err := ioutil.ReadAll(getData.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	js, _ := simplejson.NewJson(jsonData)
	temp, _ := js.GetPath("main", "temp").Float64() // キーmain.tempの値をfloat64型として取得
	utc, _ := js.GetPath("dt").Int64()              // キーdtの値をint64型として取得
	date := time.Unix(utc, 0)                       // dtから取得した値をtime.Unix関数によりtime.Time型に変換

	fmt.Println("気温:", temp, "℃")
	fmt.Println("取得時刻:", date)
}
