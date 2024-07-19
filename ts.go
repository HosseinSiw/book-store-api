package main

import (
	"fmt"
	"github.com/HosseinSiw/book-store-api/users-crud-api"
)

func main() {
	data := "samplePassword"
	//shifts := []int{1, 2, 3, 4, 5, 6, 7, 7, 8, 9, 10}
	//for _, shift := range shifts {
	//	fmt.Println(users_crud_api.CaesarEncrypt(data, shift))
	//}
	fmt.Println(users_crud_api.CaesarEncrypt(data, 10))
}
