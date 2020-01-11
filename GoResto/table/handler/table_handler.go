package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"tugas/GoResto/connect"
	"tugas/GoResto/model"

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
		fmt.Printf("[TableHandler.getAllTabel.Query] Error when query data with error : %v \n", err.Error())
	}

	for rows.Next() {
		if err := rows.Scan(&table.MejaID, &table.Status); err != nil {
			data := model.ResponseWrapper{
				Success: false,
				Message: "Scanning is no valid",
			}
			tableJSON, err := json.Marshal(data)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Oops! Something went wrong."))
				fmt.Printf("[TableHandler.getAllTabel.Scan.Marshal] Error when rows scan with error : %v\n", err.Error())
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(tableJSON)
			return
		}
		result = append(result, table)
	}
	defer rows.Close()

	data := model.ResponseWrapper{
		Success: true,
		Message: "Success",
		Data:    result,
	}
	tableJSON, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Oops! Something went wrong."))
		fmt.Printf("[TableHandler.getAllTabel.Data.Marshal] Error when rows Marshal with error : %v\n", err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(tableJSON)
	return
}

func (t TableHandler) getTabelByID(w http.ResponseWriter, r *http.Request) {
	table := model.Table{}

	db, err := connect.ConnectHandler()
	if err != nil {
		fmt.Printf("[TableHandler.getTabelByID.ConnectHandler] Error when connecting to database with error : %v \n", err.Error())
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		fmt.Printf("[TableHandler.getTabelByID.Begin] Error when connecting to begin with error : %v \n", err.Error())
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		data := model.ResponseWrapper{
			Success: false,
			Message: "ID harus angka",
		}
		tableJSON, err := json.Marshal(data)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Ooops, Something Went Wrong"))
			fmt.Printf("[TableHandler.getTabelByID.Marshal.vars] Error when do json Marshalling for error handling : %v \n", err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(tableJSON)
		return
	}

	err = db.QueryRow("select status from meja where meja_id = ?", id).Scan(&table.Status)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Printf("[QueryRow] Data Kosong : %v \n", err)
			w.Write([]byte("Data untuk id ini tidak tersedia"))
			return
		}
		fmt.Printf("[TableHandler.getTabelByID.QueryRow] Error when connecting to QueryRow with error : %v \n", err.Error())
		return
	}

	if table.Status == "close" {
		_, err = db.Exec("UPDATE meja set status = 'open' where meja_id = ?", id)
		if err != nil {
			w.Write([]byte("Oops! Something went wrong"))
			fmt.Printf("[TableHandler.getTabelByID.Exec] Error when updated with error : %v\n ", err.Error())
			return
		}

		table := model.Table{
			MejaID: id,
			Status: "open",
		}

		data := model.ResponseWrapper{
			Success: true,
			Message: "Successfully Booked !",
			Data:    table,
		}

		tableJSON, err := json.Marshal(data)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Ooops, Something Went Wrong"))
			fmt.Printf("[TableHandler.getTabelByID.Marshal.Status] Error when do json Marshalling for error handling : %v \n", err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(tableJSON)
		return
	}
	data := model.ResponseWrapper{
		Success: false,
		Message: "Sorry!, The table you choose has booked",
	}
	tableJSON, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Ooops, Something Went Wrong"))
		fmt.Printf("[TableHandler.getTabelByID.Marshal.Status] Error when do json Marshalling for error handling : %v \n", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(tableJSON)

	err = tx.Commit()
	if err != nil {
		fmt.Printf("[TableHandler.getTabelByID.Commit] Error when connecting to begin with error : %v \n", err.Error())
	}

}

// CreateTableHandler for calling router
func CreateTableHandler(r *mux.Router) {
	table := TableHandler{}
	r.HandleFunc("/table", table.getAllTabel).Methods(http.MethodGet)
	r.HandleFunc("/table/{id}", table.getTabelByID).Methods(http.MethodPut)

}
