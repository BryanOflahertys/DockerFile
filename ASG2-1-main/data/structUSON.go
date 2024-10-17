package data

type Person struct {
	// key nya ini yang akan digunakan untuk decode dan encoding JSON
	Name string `json:"name"`
	Age  int    `json:"age"`
}
