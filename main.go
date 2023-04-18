package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"net/http"
	"os"
	"runtime/debug"
	"strconv"
	"strings"
	"time"
	"github.com/gopalrg310/BitBurst/handler"
	"github.com/gopalrg310/BitBurst/utils"
)



var Log = logrus.New()

func main() {
	/*	db, err := sql.Open("postgres", "postgres://postgres:@localhost:5432/bitburstasses?sslmode=disable")
		if err != nil {
			log.Fatalf("Error opening database: %v", err)
		}
		defer db.Close()*/
	// ensure to change values as needed.
	databaseUrl := "postgres://postgres:password@postgres:5432/bitburstasses"

	// this returns connection pool
	userService:=handler.NewUserService()
	// to close DB pool
	dbPool,ctx:=utils.ConnectDB(0,databaseUrl)
	defer dbPool.Close(ctx)
	userService.Db := dbPool

	r := mux.NewRouter()
	r.HandleFunc("/users/{uid}/add", userService.AddTransactionHandler).Methods("POST")
	r.HandleFunc("/users/{uid}/balance", userService.BalanceHandler).Methods("GET")
	r.HandleFunc("/users/{uid}/history", userService.HistoryHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}
