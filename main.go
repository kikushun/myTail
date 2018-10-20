package main

import (
	"flag"
	"fmt"
	"github.com/serverhorror/rog-go/reverse" // ファイル末尾からスキャンできる
	"os"
)

func main() {

	// オプション
	var n, start int
	flag.IntVar(&n, "n", 10, "出力する行数")
	flag.IntVar(&start, "start", 1, "スタート位置")
	flag.Parse()

	// 必須
	args := flag.Args()
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "ファイルパスを指定してください。")
		os.Exit(1)
	}

	// コマンド引数で渡された各ファイルをmyTailする
	for _, filePath := range args {
		myTail(filePath, &n, &start)
	}
}

// 引数で渡されたファイルを末尾の[start]行目から[n]行出力
func myTail(filePath string, n, start *int) error {

	// ファイルパスから対象のファイルを開く
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "ファイルが見つかりませんでした。", filePath)
		return fmt.Errorf("ファイルが見つかりませんでした。[%s]", err)
	}
	defer file.Close()

	// 終了位置
	endLine := *n + *start

	// 末尾から１行ずつスキャン
	scanner := reverse.NewScanner(file)
	for i := 1; scanner.Scan(); i++ {

		// 終了位置に達したら終了
		if i >= endLine {
			break
		}

		// スタート位置に達するまでcontinue
		if i < *start {
			continue
		}

		// 出力(LIFO)
		defer fmt.Println(scanner.Text())
	}

	defer fmt.Println("***************", file.Name(), "***************")

	return nil
}
