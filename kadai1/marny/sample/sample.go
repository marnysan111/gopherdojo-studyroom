// Copyright
// +build darwin

package sample

func Hoge(str string) string {
	str2 := str + "HAHAHA"

	return str2
}

//Article は1つの記事を表します。
type Article struct {
	// 記事のタイトル
	Title string

	// 記事本文
	Body string
}
