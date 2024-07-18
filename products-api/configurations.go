package products_api

const (
	DbDriver = "mysql"
	DbUser   = "root"
	DbPass   = "0150112378MySQL"
	DbName   = "goCrudApp"
)

type Book struct {
	id       int
	name     string
	price    string
	numPages int
}
