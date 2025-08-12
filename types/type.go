package types


type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
	Occupation string `json:"occupation"`
}