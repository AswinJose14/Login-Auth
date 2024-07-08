package main

import (
	"fmt"

	"github.com/AswinJose14/Login-Auth/config"
	"github.com/AswinJose14/Login-Auth/db"
	"github.com/AswinJose14/Login-Auth/server"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		panic("Unable to load config")
	}
	fmt.Println("Test is starting")
	db, err := db.Init()
	if err != nil {
		fmt.Printf("unable to connect to DB, %v", err)
	}
	server.Start(db, config)
}
