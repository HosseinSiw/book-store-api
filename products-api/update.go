package products_api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func updateProduct(db *sql.DB, id int, name string, price string, numPages int) error {
	query := "UPDATE book SET (?, ?, ?) WHERE id = ?"
	_, err := db.Exec(query, name, price, numPages, id)
	if err != nil {
		log.Fatal("Cannot Update the product")
		return err
	}
	return nil
}
func UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open(DbDriver, DbUser+":"+DbPass+"@/"+DbName)
	if err != nil {
		log.Fatal("Error during opening the database, couldn't open the db", err)
	}
	defer func() {
		err = db.Close()
		if err != nil {
			log.Fatal(err, "Cannot close the db")
		}
	}()
	var product = &Book{}
	err = json.NewEncoder(w).Encode(&product)
	if err != nil {
		log.Fatal(err)
	}
	args := mux.Vars(r)
	id, err := strconv.Atoi(args["id"])
	if err != nil {
		log.Fatal(err)
	}
	err = updateProduct(db, id, product.name, product.price, product.numPages)
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusOK)
	fmt.Println("Product updated successfully")
}
