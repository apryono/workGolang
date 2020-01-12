package handler

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"tugas/GoResto/connect"
	"tugas/GoResto/model"
	"tugas/GoResto/response"

	"github.com/gorilla/mux"
)

// TableHandler struct
type TableHandler struct {
}

func (t TableHandler) getAllTabel(w http.ResponseWriter, r *http.Request) {
	var table = model.Table{}
	var result = []model.Table{}

	db, err := connect.ConnectHandler()
	if err != nil {
		fmt.Printf("[TableHandler.getAllTabel] Error when connecting to database with error : %v \n", err.Error())
	}
	defer db.Close()

	rows, err := db.Query("Select meja_id,status from meja")
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			response.RespondWithError(w, "There is no data")
			return
		}
		response.RespondWithError(w, "Oops! Something went wrong")
		fmt.Printf("[TableHandler.getAllTabel.Query] Error when query data with error : %v \n", err.Error())
		return
	}

	for rows.Next() {
		if err := rows.Scan(&table.MejaID, &table.Status); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				response.RespondWithError(w, "There is no data")
				return
			}
			response.RespondWithError(w, "Oops! Something went wrong")
			fmt.Printf("[TableHandler.getAllTabel.Query.rows.Next] Error when query data with error : %v \n", err.Error())
			return
		}
		result = append(result, table)
	}
	defer rows.Close()

	response.RespondWithSuccess(w, "Success", result)
	return
}

func (t TableHandler) getTabelByID(w http.ResponseWriter, r *http.Request) {
	table := model.Table{}

	db, err := connect.ConnectHandler()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			response.RespondWithError(w, "There is no data")
			return
		}
		response.RespondWithError(w, "Oops! Something went wrong")
		fmt.Printf("[TableHandler.getTabelByID.ConnectHandler] Error when connecting to database with error : %v \n", err.Error())
		return
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		fmt.Printf("[TableHandler.getTabelByID.Begin] Error when connecting to begin with error : %v \n", err.Error())
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response.RespondWithError(w, "ID harus angka!")
		return
	}

	err = db.QueryRow("select status from meja where meja_id = ?", id).Scan(&table.Status)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			response.RespondWithError(w, "There is no data")
			fmt.Printf("[QueryRow] Data Kosong : %v \n", err)
			return
		}
		response.RespondWithError(w, "Oops! Something went wrong")
		fmt.Printf("[TableHandler.getTabelByID.QueryRow] Error when connecting to QueryRow with error : %v \n", err.Error())
		return
	}

	if table.Status == "close" {
		_, err = db.Exec("UPDATE meja set status = 'open' where meja_id = ?", id)
		if err != nil {
			response.RespondWithError(w, "Oops! Something went wrong")
			fmt.Printf("[TableHandler.getTabelByID.Exec] Error when updated with error : %v\n ", err.Error())
			return
		}

		table := model.Table{
			MejaID: id,
			Status: "open",
		}

		response.RespondWithSuccess(w, "Successfully Booked !", table)
		return
	}
	response.RespondWithError(w, "Sorry!, The table you choose has booked")

	err = tx.Commit()
	if err != nil {
		fmt.Printf("[TableHandler.getTabelByID.Commit] Error when connecting to begin with error : %v \n", err.Error())
		return
	}

}

// CreateTableHandler for calling router
func CreateTableHandler(r *mux.Router) {
	table := TableHandler{}
	r.HandleFunc("/table", table.getAllTabel).Methods(http.MethodGet)
	r.HandleFunc("/table/{id}", table.getTabelByID).Methods(http.MethodPut)

}
