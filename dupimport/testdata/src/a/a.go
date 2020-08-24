package a

import (
	fmt1 "fmt"
	fmt2 "fmt" // want "fmt package is duplicated"
)

func main() {
	fmt1.Println("fmt1")
	fmt2.Println("fmt2")
}
