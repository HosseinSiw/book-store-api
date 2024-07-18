package users_crud_api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func UpdateUser(db *sql.DB, userId int, name string, email string) error {
	var (
		err   error
		query = "UPDATE users SET name = ?, email = ? WHERE id = ?"
	)
	_, err = db.Exec(query, name, email, userId)
	if err != nil {
		fmt.Println("Error during the [Updating]")
		return err
	}
	return nil
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	db, err := sql.Open(DbDriver, DbUser+":"+DbPass+"@/"+DbName)
	if err != nil {
		fmt.Println("[UPDATE]")
		panic(err.Error())
	}
	defer func() {
		err = db.Close()
		if err != nil {
			log.Fatal("Cannot close the DB ", err)
		}
	}()
	// Getting the id and converting into integer
	vars := mux.Vars(r)
	id := vars["id"]
	userId, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal("Error during the converting", err)
	}
	var user User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}
	// Call the GetUser function to fetch the user data from the database
	err = UpdateUser(db, userId, user.Name, user.Email)
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusOK)
	fmt.Println("User updated successfully")
}
