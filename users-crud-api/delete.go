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

func deleteUser(db *sql.DB, userID int) error {
	var err error
	query := "DELETE FROM users WHERE ID = ?"
	_, err = db.Exec(query, userID)
	if err != nil {
		log.Fatal("Couldn't delete the user")
		return err
	}
	return nil
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open(DbDriver, DbUser+":"+DbPass+"@/"+DbName)
	if err != nil {
		panic(err)
	}

	vars := mux.Vars(r)
	id := vars["id"]
	userID, _ := strconv.Atoi(id)
	user := deleteUser(db, userID)
	if err != nil {
		http.Error(w, "no such user, review the id", http.StatusBadGateway)
		return
	}
	_, err = fmt.Fprintln(w, "User deleted successfully")
	if err != nil {
		log.Fatal(err)
	}
	// Convert the user object to JSON and send it in the response
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		log.Fatal(err)
	}
}
