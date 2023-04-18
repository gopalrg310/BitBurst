package main

import (
	"github.com/gorilla/mux"

	"log"
	"net/http"

	"github.com/gopalrg310/bitburst/handler"
	"github.com/gopalrg310/bitburst/utils"
)

func main() {
	// ensure to change values as needed.
	databaseUrl := "postgres://postgres:password@postgres:5432/bitburstasses"

	// this returns connection pool
	userService := handler.NewUserService()

	dbPool, _ := utils.ConnectDB(0, databaseUrl)
	// to close DB pool
	defer dbPool.Close()
	userService.Db = dbPool

	r := mux.NewRouter()
	r.HandleFunc("/users/{uid}/add", userService.AddTransactionHandler).Methods("POST")
	r.HandleFunc("/users/{uid}/balance", userService.BalanceHandler).Methods("GET")
	r.HandleFunc("/users/{uid}/history", userService.HistoryHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}
