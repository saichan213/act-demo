package main

import (
	"os"
	"testing"
)

func TestFileOutput(t *testing.T) {
	// 手動でファイルを生成する関数を呼び出して出力
	main()

	// ファイルの内容を確認
	content, err := os.ReadFile("test.txt")
	if err != nil {
		t.Fatalf("ファイルの読み込みに失敗: %v", err)
	}
	defer os.Remove("test.txt") // テスト終了時にファイルを削除

	expected := "test"
	if string(content) != expected {
		t.Errorf("期待される出力: %s, 実際の出力: %s", expected, string(content))
	}
}
