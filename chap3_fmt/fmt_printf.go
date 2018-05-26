package main

import "fmt"

func main() {

	/*
	 * 基本的な使い方
	 */
	fmt.Println("## 基本的な使い方")

	// 基本
	fmt.Printf("数値=%d\n", 5)

	// 数値用書式
	fmt.Printf("10進=%d 2進=%b 8進=%o 16進=%x\n", 17, 17, 17, 17)

	// パラメータ不足
	fmt.Printf("%d%d%d\n", 1, 2)

	// パラメータ過剰
	fmt.Printf("%d%d%d", 1, 2, 3, 4)

	/*
	 * %v %#v %+v %T
	 */
	fmt.Println("\n\n## %v %#v %+v %T")

	// %v
	fmt.Printf("数値=%v 文字列%v 配列=%v\n", 5, "hoge", [...]int{1, 2, 3})

	// %#v
	fmt.Printf("数値=%#v 文字列%#v 配列=%#v\n", 5, "hoge", [...]int{1, 2, 3})

	// %+v
	fmt.Printf("数値=%+v 文字列%+v 配列=%+v\n", 5, "hoge", [...]int{1, 2, 3})

	// %T
	fmt.Printf("数値=%T 文字列%T 配列=%T\n", 5, "hoge", [...]int{1, 2, 3})
}
