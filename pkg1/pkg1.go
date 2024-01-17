package pkg1 //自定義包

var Money = 2000 //聲明變量 首字大寫為公開使用的
var xyz = 10     // 私有

// 首字大寫 為公開
func Add(a, b int) int {
	return a + b
}

// 私有
func add(a, b int) int {
	return a + b
}
