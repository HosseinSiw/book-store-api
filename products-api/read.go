package products_api

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func getBook(db *sql.DB, bookId int) (*Book, error) {
	var err error
	var book = &Book{}
	query := "SELECT * FROM books WHERE Id = ?"
	row := db.QueryRow(query, bookId)
	err = row.Scan(&book.name, &book.price, &book.numPages)
	if err != nil {
		log.Fatal("Couldn't read the book from db there isn't such user")
		return nil, err
	}
	return book, nil
}

func GetBookHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open(DbDriver, DbUser+":"+DbPass+"@/"+DbName)
	if err != nil {
		panic(err)
	}
	defer func() {
		err = db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	vars := mux.Vars(r)
	id := vars["id"]
	bookId, err := strconv.Atoi(id)
	book, err := getBook(db, bookId)
	if err != nil {
		http.Error(w, "Cannot retrive the book", http.StatusBadGateway)
	}
	//	Encode the book into json format
	w.Header().Set("Content-type", "application/json")
	err = json.NewEncoder(w).Encode(book)
	if err != nil {
		log.Fatal(err)
	}
}
