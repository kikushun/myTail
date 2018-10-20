package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestMyTail(t *testing.T) {

	tmp, _ := ioutil.TempFile(".", "tmpTest")
	str :=
		`aaaaaa
	bbbbb
	ccccc
	ddddd
	eeeee
	fffff
	ggggg`

	ioutil.WriteFile(tmp.Name(), []byte(str), 777)

	// 成功テスト
	n, start := 5, 1
	err := myTail(tmp.Name(), &n, &start)
	if err != nil {
		t.Fatal(err)
	}

	// 成功テスト
	n, start = 10, 2
	err = myTail(tmp.Name(), &n, &start)
	if err != nil {
		t.Fatal(err)
	}

	// 失敗テスト
	os.Remove(tmp.Name())
	err = myTail(tmp.Name(), &n, &start)
	if err == nil {
		t.Fatal("存在しないファイルなのにエラーにならず")
	}

}
