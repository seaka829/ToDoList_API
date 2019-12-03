package main

import (
	"./db"
	"./router"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db.Init()
	defer db.Close()

	router.Init()
}
