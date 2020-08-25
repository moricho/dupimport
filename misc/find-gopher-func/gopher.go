package main

import "fmt"

type GOPHER struct {
	Gopher string `json:"gopher"`
}

func sub() {
	const gopher = "GOPHER"
	gogopher := Gopher()
	gogopher.Gopher = gopher
	fmt.Println(gogopher)
}

func Gopher() (gopher *GOPHER) {
	gopher = &GOPHER{Gopher: "gopher"}
	return
}
