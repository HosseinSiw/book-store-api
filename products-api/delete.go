package products_api

//
//import (
//	"database/sql"
//	"github.com/gorilla/mux"
//	"log"
//	"net/http"
//	"strconv"
//)
//
//func deleteProduct(db *sql.DB, productId int) error {
//	query := "DELETE FROM books WHERE ID = ?"
//	_, err := db.Exec(query, productId)
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
//func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
//	db, err := sql.Open(DbDriver, DbUser+":"+DbPass+"@/"+DbName)
//	if err != nil {
//		log.Fatal("Error during opening the database, couldn't open the db", err)
//	}
//	defer func() {
//		err = db.Close()
//		if err != nil {
//			log.Fatal(err, "Cannot close the db")
//		}
//	}()
//	id := mux.Vars(r)["id"]
//	productId, err := strconv.Atoi(id)
//	if err != nil {
//		log.Fatal(err)
//	}
//	err = deleteProduct(db, productId)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//
//}
