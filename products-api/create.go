package products_api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func createBook(db *sql.DB, name string, price string, numPages int) error {
	query := "INSERT INTO books values (?, ?, ?)"
	_, err := db.Exec(query, name, price, numPages)
	if err != nil {
		return err
	}
	return nil
}
func CreateBookHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open(DbDriver, DbUser+":"+DbPass+"@/"+DbName)
	if err != nil {
		panic(err)
	}
	defer func() {
		err = db.Close()
		if err != nil {
			log.Fatal("Cannot close the DB")
		}
	}()
	// Parse JSON data from the request body
	var book Book
	err = json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		log.Fatal(err)
	}
	err = createBook(db, book.name, book.price, book.numPages)
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusCreated)
	_, err = fmt.Fprintln(w, "User created successfully")
	if err != nil {
		log.Print("Cannot write the http header [CREATE]")
	}
}
