/*
ここはパッケージコメントの最初になるから見出しではないよー

Hで始まり単一行かつ句読点なしかつ前が見出しではないのでこれは見出し

段落段落
段落段落
段落段落

次の段落
次の段落

    整形済みテキスt

次のやつはリンクになるはず。
https://golang.org/
*/
package sample

func Hoge(str string) string {
	str2 := str + "HAHAHA"
	return str2
	// Output: aaa
}

//Article は1つの記事を表します。
type Article struct {
	// 記事のタイトル
	Title string

	// 記事本文
	Body string
}

type Test struct {
	A string
}
