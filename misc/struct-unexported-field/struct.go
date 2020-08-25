package main

type Person struct {
	// タグが付いているのに非公開なのでNG
	name string `json:"name"`
	age  int    `json:"age"`
	Hoge string `json:"hoge"`
}
