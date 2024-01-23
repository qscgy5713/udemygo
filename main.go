package main

import (
	"fmt"

	"github.com/qscgy5713/udemygo/pkg1"
	"github.com/qscgy5713/udemygo/pkg2"
)

func main() {
	chapter1()
}

func chapter1() {
	// 變量聲明
	var str1 = "hello world"
	// 變量聲明沒有初始化
	var variable1 int
	var variable2 string
	// 變量同類型聲明
	var variable3, variable4 int = 3, 4
	// 變量不同同類型聲明
	var v1, v2, v3 = 1, "aa", 4.4
	fmt.Println(str1, variable1, variable2, variable3, variable4, v1, v2, v3)

	// 短變量聲明
	a1 := 1
	a2 := "ddd"
	a3 := 4.2
	fmt.Println(a1, a2, a3)
	fmt.Printf("%.3f", a3)

	// 常量聲明 不可修改性
	const pi = 3.14
	const na = "test"
	fmt.Println(pi, na)

	// 格式輸出
	b1 := 64
	fmt.Printf("十進制: %d\n", b1)
	fmt.Printf("八進制: %o\n", b1)
	fmt.Printf("十六進制: %x\n", b1)
	a := pkg1.Add(1, 3)
	fmt.Println(a)
	fmt.Println(pkg2.Money)
}
