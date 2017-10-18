// package mainではmain関数を定義することでエントリポイントになる
package main

// importすると、packageのメンバが使用可能になる
// フォーマットI/O C言語のprintfおよびscanf的なやつ
import "fmt"

func main() {
    // 大文字のメンバはimportで利用可能になる
	fmt.Println("Hello ihysk!")
}
