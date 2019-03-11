package main

import (
	"gorm/db"
	"gorm/server"
)

func main() {
	db.Init()
	server.Init()
}
