package database

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
	Occupation string `json:"occupation"`
}

var UserList[] User

func Store(user User) {
	UserList = append(UserList, user)
}

func List() []User {
	return UserList
}

func Get(id int) *User {
	for _,user := range(UserList) {
		if user.Id == id {
			return &user
		}
	}
	return nil
}

