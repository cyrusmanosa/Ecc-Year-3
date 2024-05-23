package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"time"
// )

// // Raspberry Pi側のプログラムと同じデータ型を定義
// type TempData struct {
// 	Temp      float32
// 	EntryTime time.Time
// }

// func main() {
// 	getData, err := http.Get("http://[Raspberry PiのIPアドレス]:4000/json") // JSON取得用アドレスを指定
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer getData.Body.Close() // HTTP接続のクローズ

// 	// HTTPレスポンスのBody（JSONデータ）を取得する
// 	jsonData, err := ioutil.ReadAll(getData.Body)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	// JSONデータを構造体に変換する
// 	var tempData []TempData
// 	err = json.Unmarshal(jsonData, &tempData)
// 	if err != nil || len(jsonData) == 0 {
// 		fmt.Println(err)
// 		return
// 	}

// 	// 取得したデータのうち、最大値と最少値を検索する
// 	max := tempData[0]
// 	min := tempData[0]
// 	for _, t := range tempData {
// 		if max.Temp < t.Temp {
// 			max = t
// 		}
// 		if min.Temp > t.Temp {
// 			min = t
// 		}
// 	}
// 	fmt.Println("最大値のデータは", max)
// 	fmt.Println("最小値のデータは", min)
// }
