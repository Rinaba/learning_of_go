package main

import (
	"fmt"
//	"time"
)

func main () {
	var data int
	go func() {
		data++
	} ()
	//以下のSLEEP処理を有効にするとif文の評価前にdataを加算する
	//期待値は "表示なし"
	//time.Sleep(1*time.Second)
	if data == 0 {
		//以下のSLEEP処理を有効にするとif文評価後にdataを加算する
		//期待値は "the value is 1."
		//time.Sleep(1*time.Second)
		fmt.Printf("the value is %v.\n", data)
	}
}
