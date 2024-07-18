package users_crud_api

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func getUser(db *sql.DB, userId int) (*User, error) {
	var err error
	query := "SELECT * FROM users WHERE ID = ?"
	row := db.QueryRow(query, userId)
	user := &User{}
	err = row.Scan(&user.id, &user.Name, &user.Email)
	if err != nil {
		log.Fatal("Couldn't read the user from db there isn't such user")
		return nil, err
	}
	return user, nil
}
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	var msg = "[READ]"
	//	Open (connect) the database
	db, err := sql.Open(DbDriver, DbUser+":"+DbPass+"@/"+DbName)
	if err != nil {
		log.Fatal("Error during opening the database, couldn't open the db", err, msg)
	}
	defer func() {
		err = db.Close()
		if err != nil {
			log.Fatal(err, "Cannot close the db", msg)
		}
	}()
	// Getting the {id} parameter
	id := mux.Vars(r)
	userId, err := strconv.Atoi(id["id"])
	//	Fetching the user from the db
	user, err := getUser(db, userId)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	// Convert the user object to JSON and send it in the response
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		log.Fatal("Cannot Encode the values", err)
	}
}
