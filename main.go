package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/pedrocmart/crud-go/api"
)

func main() {
	api.Run()
}
