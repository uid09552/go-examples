package main

import (
	"database/sql"
	"log"
	"fmt"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(user, password, dbname, connectionString string) {
	conn := fmt.Sprint(connectionString)
	var err error
	a.DB, err = sql.Open("postgres", conn)
	if err != nil {
		log.Fatal(err)
	}
	a.Router = mux.NewRouter()
}

func (a *App) Run(addr string) int {
	return 0
}
