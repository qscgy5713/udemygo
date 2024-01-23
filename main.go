package main

import (
	"fmt"
	"strings"

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

	// 聲明變量
	var i0 = 19
	var i1 int8 = 0b10011 // 二進制 19
	var i2 int16 = 0o23   // 八進制 19
	var i3 int32 = 0x13   // 16 進制 19
	fmt.Printf("10 %d\n", i0)
	fmt.Printf("10 type %T\n", i0)
	fmt.Printf("2 %d\n", i1)
	fmt.Printf("2 type %T\n", i1)
	fmt.Printf("8 %d\n", i2)
	fmt.Printf("8 type %T\n", i2)
	fmt.Printf("16 %d\n", i3)
	fmt.Printf("16 type %T\n", i3)

	var f0 *float64
	ff0 := 1.333
	f0 = &ff0
	fmt.Printf("%.3f\n", *f0)

	// 複數類型
	// complex128()(64位實數和虛數)
	// complex64()(32位實數和虛數)
	var com1 complex128 = complex(2, -3)
	var com2 complex64 = complex(9, 2)
	fmt.Println(com1, com2)

	// 類型轉換
	sum := 17
	count := 5
	var mean *float32
	mm := float32(sum) / float32(count)
	mean = &mm
	fmt.Printf("mean: %.3f\n", *mean)
	var long1 int64 = 9999999999
	var int1 int32 = int32(long1)
	fmt.Println(int1)
	float1 := 3.456
	var int2 = int32(float1)
	fmt.Println(int2)

	// string
	s1 := "hello world"
	s2 := s1 + "fuck"
	s3 := "\u5927\u5bb6\u597d"
	s4 := "呵呵"
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Printf("s1 len %d\n", len(s1))
	fmt.Printf("%s\n", s1)
	fmt.Printf("%c\n", s1[0])

	text := "go programming"
	subText := "go"
	res := strings.Contains(text, subText)
	fmt.Println(res)

	text2 := "cat"
	t3 := strings.Replace(text2, "t", "r", 1)
	fmt.Println(t3)

	text3 := " I love Golang"
	t4 := strings.ToUpper(text3)
	fmt.Println(t4)
	t4 = strings.ToLower(text3)
	fmt.Println(t4)
	t5 := strings.Split(text3, " ")
	fmt.Printf("%v\n", t5)
}
