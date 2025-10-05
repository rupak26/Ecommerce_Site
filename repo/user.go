package repo

import ( 
  "github.com/jmoiron/sqlx"
  "strings"
  "fmt"
)

type UserRepo interface {
	Create(u User) (*User , error)
	Get() ([]User, error)
	Find(email , password string) (*User , error)
	GetById(usrID int) (*User, error)
	UpdateUser(user User) (*User , error) 
	PatchUser(id int, req UpdateUserReq) (*User, error)
	Delete(id int) (error)
}

type UpdateUserReq struct {
	Name *string         `json:"name"`
	Age  *int            `json:"age"`
	Email *string        `json:"email"`
	Password *string     `json:"password"`
	Occupation *string   `json:"occupation"`
}

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

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) Create(usr User) (*User , error) {
    query := `
		INSERT INTO users (
		  name, 
		  age, 
		  email, 
		  password, 
		  occupation
		) VALUES (
		  :name , 
		  :age , 
		  :email , 
		  :password , 
		  :occupation
		)
		RETURNING id;
	`

	var id int
	rows , err := r.db.NamedQuery(query, usr) 
	if err != nil {
		return nil , err
	}
	defer rows.Close()
	if rows.Next(){
		rows.Scan(&id)
	}
	
	usr.Id = id
	return &usr , nil
}



func (r *userRepo) GetById(usrID int) (*User, error) {
	var user User
	query := `SELECT * FROM users WHERE id = $1`
	err := r.db.Get(&user, query, usrID)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) Find(email, password string) (*User, error) {
	var user User
	query := `SELECT * FROM users WHERE email = $1 AND password = $2`
	err := r.db.Get(&user, query, email, password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) Get() ([]User, error) {
	var users []User
	query := `SELECT * FROM users`
	err := r.db.Select(&users, query)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepo) UpdateUser(user User) (*User , error) {
	query := `
		UPDATE users
		SET 
			name = :name,
			age = :age,
			email = :email,
			password = :password,
			occupation = :occupation
		WHERE id = :id
		RETURNING id, name, age, email, password, occupation;
	`

	rows, err := r.db.NamedQuery(query, user)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.StructScan(&user)
		if err != nil {
			return nil, err
		}
		return &user, nil
	}

	return nil, nil
} 

func (r *userRepo) PatchUser(id int, req UpdateUserReq) (*User, error) {
	query := "UPDATE users SET "
	params := map[string]interface{}{"id": id}
	setClauses := []string{}

	if req.Name != nil {
		setClauses = append(setClauses, "name = :name")
		params["name"] = *req.Name
	}
	if req.Age != nil {
		setClauses = append(setClauses, "age = :age")
		params["age"] = *req.Age
	}
	if req.Email != nil {
		setClauses = append(setClauses, "email = :email")
		params["email"] = *req.Email
	}
	if req.Password != nil {
		setClauses = append(setClauses, "password = :password")
		params["password"] = *req.Password
	}
	if req.Occupation != nil {
		setClauses = append(setClauses, "occupation = :occupation")
		params["occupation"] = *req.Occupation
	}

	if len(setClauses) == 0 {
		return nil, fmt.Errorf("no fields provided to update")
	}

	query += strings.Join(setClauses, ", ") + " WHERE id = :id RETURNING id, name, age, email, password, occupation;"
    
	rows, err := r.db.NamedQuery(query, params)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var user User
	if rows.Next() {
		if err := rows.StructScan(&user); err != nil {
			return nil, err
		}
		return &user, nil
	}
	return nil, fmt.Errorf("user with id %d not found", id)
}

func (r userRepo) Delete(id int) error {
	query := `DELETE from 
	         users 
	         WHERE id=$1`
	_,err := r.db.Exec(query , id)
	if err != nil {
		return err
	}
	return nil
}