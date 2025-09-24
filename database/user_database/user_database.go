package userdatabase


type User struct {
	Id   int            `json:"id"`
	Name string         `json:"name"`
	Age  int            `json:"age"`
	Email string        `json:"email"`
	Password string     `json:"password"`
	Occupation string   `json:"occupation"`
}

type ReqLogin struct {
	Email string      `json:"email"`
	Password string   `json:"password"`
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

func Find(email , password string) *User {
    for _,user := range(UserList) {
		if user.Email == email && user.Password == password {
			return &user
		}
	}
	return nil 
}
