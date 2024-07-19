package users_crud_api

type User struct {
	id       int
	Name     string
	Email    string
	Password string
}

func (*User) checkPassword() {

}
