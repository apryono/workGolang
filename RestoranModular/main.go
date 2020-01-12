package main

import (
	"fmt"
	"net/http"
	"tugas/RestoranModular/routes"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	port := "8080"

	r := mux.NewRouter()
	rMenu := r.PathPrefix("/menu").Subrouter()
	rTable := r.PathPrefix("/table").Subrouter()
	rOrder := r.PathPrefix("/order").Subrouter()
	rTrans := r.PathPrefix("/transaction").Subrouter()
	routes.MenusHandler(rMenu)
	routes.TablesHandler(rTable)
	routes.OrderHandler(rOrder)
	routes.TransactionHandler(rTrans)

	fmt.Println("Starting server at port : " + port)
	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		fmt.Printf("[main] ListenAndServe Error with issues: %v\n", err.Error())
	}
}
