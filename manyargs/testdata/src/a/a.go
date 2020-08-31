package a

func hoge(a, b, c, d, e, f int, str string) int { // want "hoge: too many args"
	return a + b + c + d + e + f
}

func fuga(a, b, c int) int {
	return a + b + c
}
