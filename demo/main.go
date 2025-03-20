package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.Create("test.txt") // test.txtというファイルを作成
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() // mainが終了する前にファイルを閉じる

	_, err = file.WriteString("test") // "test"と書き込む
	if err != nil {
		log.Fatal(err)
	}
}
