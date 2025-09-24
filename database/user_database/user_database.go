package userdatabase


type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
	Occupation string `json:"occupation"`
}

var UserList[] User

func StoreUser(user User) {
	UserList = append(UserList, user)
}
func GetUserList() []User {
	return UserList
}
func GetUser(id int) *User {
	for _,user := range(UserList) {
		if user.Id == id {
			return &user
		}
	}
	return nil
}
