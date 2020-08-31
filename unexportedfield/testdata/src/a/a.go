package a

type Person struct {
	name     string `json:"name"` // want "name field is unexported though it has a json tag"
	age      int    `json:"age"`  // want "age field is unexported though it has a json tag"
	Hoge     string `json:"hoge"`
	aaa, BBB string
}
