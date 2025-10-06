package domain

type User struct {
	Id   int            `json:"id"`
	Name string         `json:"name"`
	Age  int            `json:"age"`
	Email string        `json:"email"`
	Password string     `json:"password"`
	Occupation string   `json:"occupation"`
}

type UpdateUserReq struct {
	Name *string         `json:"name"`
	Age  *int            `json:"age"`
	Email *string        `json:"email"`
	Password *string     `json:"password"`
	Occupation *string   `json:"occupation"`
}
