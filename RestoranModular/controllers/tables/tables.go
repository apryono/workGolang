package tables

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Table struct {
	NomorMeja int    `json:"Nomor Meja"`
	Status    string `json:"Status"`
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Table
}

var response Response

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/restoran")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func handleError(err error) {
	if err != nil {
		fmt.Println("Oopss !! Something went wrong.!")
		fmt.Println(err.Error())
		return
	}

}

func GetAllStatusMeja(w http.ResponseWriter, req *http.Request) {

	var tables Table
	var result []Table

	db, err := connect()
	handleError(err)
	defer db.Close()

	rows, err := db.Query("Select meja_id,status from meja")
	if err != nil {
		w.Write([]byte("Oooops, something error"))
		fmt.Printf("[GetAllStatusMeja] Error when Query database with error : %v\n ", err.Error())
	}

	for rows.Next() {
		if err := rows.Scan(&tables.NomorMeja, &tables.Status); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Oops! Something went wrong."))
			fmt.Printf("[GetAllStatusMeja] Error when rows scan with error : %v\n", err.Error())
			return
		}

		result = append(result, tables)

	}

	response.Status = 1
	response.Message = "Success"
	response.Data = result

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func GetStatusMejaByID(w http.ResponseWriter, req *http.Request) {

	var tables Table

	db, err := connect()
	handleError(err)
	defer db.Close()

	tx, err := db.Begin()
	handleError(err)

	vars := mux.Vars(req)
	id, _ := strconv.Atoi(vars["id"])

	err = db.QueryRow("SELECT status from meja where meja_id = ?", id).Scan(&tables.Status)
	handleError(err)
	if tables.Status == "close" {
		_, err = db.Exec("UPDATE meja set status = 'open' where meja_id = ?", id)
		if err != nil {
			w.Write([]byte("Oops! Something went wrong"))
			fmt.Printf("[GetStatusMejaByID] Error when updated with error : %v\n ", err.Error())
			return
		}

		w.Write([]byte("Successfully Booked. !"))
	} else {
		w.Write([]byte("Sorry!, The table you choose has booked"))
	}

	err = tx.Commit()
	handleError(err)

}
