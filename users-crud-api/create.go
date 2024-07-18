package users_crud_api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	DbDriver = "mysql"
	DbUser   = "root"
	DbPass   = "0150112378MySQL"
	DbName   = "goCrudApp"
)

type User struct {
	id    int
	Name  string
	Email string
}

func CreateUser(db *sql.DB, name string, email string) error {
	command := "INSERT INTO users values (?, ?)"
	_, err := db.Exec(command, name, email)
	if err != nil {
		return err
	}
	return nil
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	db, err := sql.Open(DbDriver, DbUser+":"+DbPass+"@/"+DbName)
	if err != nil {
		log.Fatal("Cannot open the data base", err.Error())
	}
	defer func() {
		err = db.Close()
		if err != nil {
			log.Fatal("Cannot close the database")
		}
	}()
	// Parse JSON data from the request body
	var user User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatal("Cannot decode the content into a user type")
	}
	err = CreateUser(db, user.Name, user.Email)
	if err != nil {
		log.Fatal("Cannot Create a new user")
	}
	w.WriteHeader(http.StatusCreated)
	_, err = fmt.Fprintln(w, "User created successfully")
	if err != nil {
		log.Print("Cannot write the http header [CREATE]")
	}
}
