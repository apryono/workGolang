package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"tugas/GoResto/connect"
	"tugas/GoResto/model"

	"github.com/gorilla/mux"
)

// OrderHandler struct
type OrderHandler struct {
}

func (o OrderHandler) insertOrderHandler(w http.ResponseWriter, r *http.Request) {
	data := []model.Order{}

	db, err := connect.ConnectHandler()
	if err != nil {
		fmt.Printf("[OrderHandler.insertOrderHandler] Error when connecting to database with error : %v \n", err.Error())
	}
	defer db.Close()

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Opps. Something went wrong. Please contact admin"))
		fmt.Printf("[OrderHandler.insertOrderHandler.ReadAll] Error when request body with error : %v\n ", err.Error())
		return
	}

	err = json.Unmarshal(reqBody, &data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Opps. Something went wrong. Please contact admin"))
		fmt.Printf("[OrderHandler.insertOrderHandler.Unmarshal] Error when unmarshal with error : %v\n ", err.Error())
		return
	}

	insertDB(data[0].MejaID, data[0].Pesanan, w)

}

func insertDB(ID int, data []model.DetailOrder, w http.ResponseWriter) {
	menu := model.Table{}

	db, err := connect.ConnectHandler()
	if err != nil {
		fmt.Printf("[OrderHandler.insertOrderHandler.insertDB.ConnectHandler] Error when connecting to database with error : %v \n", err.Error())
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		fmt.Printf("[OrderHandler.insertOrderHandler.insertDB.Begin] Error when Begin to database with error : %v \n", err.Error())
	}
	status := "uncheck"
	res, err := db.Exec("INSERT into transactions (meja_id,tgl_transaksi,status_transaksi) VALUES (?,current_date(),?)", ID, status)
	if err != nil {
		fmt.Printf("[OrderHandler.insertOrderHandler.insertDB.Exec] Error when insert with error : %v\n ", err.Error())
		tx.Rollback()
		return
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		fmt.Printf("[OrderHandler.insertOrderHandler.insertDB.LastInsertId] Error when last insert with error : %v\n ", err.Error())
		tx.Rollback()
		return
	}

	rows, err := db.Query("SELECT status from meja where meja_id = ?", ID)
	if err != nil {
		fmt.Printf("[OrderHandler.insertOrderHandler.insertDB.Query] Error when Query with error : %v\n ", err.Error())
		tx.Rollback()
		return
	}

	for rows.Next() {
		err = rows.Scan(&menu.Status)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				w.Write([]byte("Meja yang anda pilih tidak tersedia"))
				tx.Rollback()
				return
			}
			fmt.Printf("[OrderHandler.insertOrderHandler.insertDB.rows.Next.Scan] Error when scan with error : %v\n ", err.Error())
			tx.Rollback()
			return
		}
	}
	defer rows.Close()

	if menu.Status == "open" {
		for i := 0; i < len(data); i++ {
			_, err = tx.Exec("INSERT INTO orders (transaksi_id,menu_id,qty,notes) VALUES (?,?,?,?)", lastID, data[i].MenuID, data[i].Qty, data[i].Notes)
			if err != nil {
				if errors.Is(err, sql.ErrNoRows) {
					w.Write([]byte("Oops! Menu yang anda pilih tidak ada."))
					tx.Rollback()
				}
				fmt.Printf("[OrderHandler.insertOrderHandler.insertDB.Status] Error when insert transaksi with error : %v\n ", err.Error())
				tx.Rollback()
				return
			}
		}

		err = tx.Commit()
		if err != nil {
			fmt.Printf("[OrderHandler.insertOrderHandler.insertDB.Commit] Error when connecting to begin with error : %v \n", err.Error())
		}

		table := model.Table{
			MejaID: ID,
			Status: "Open",
		}

		data := model.ResponseWrapper{
			Success: true,
			Message: "Success Insert Data",
			Data:    table,
		}

		orderJSON, err := json.Marshal(data)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Ooops, Something Went Wrong"))
			fmt.Printf("[OrderHandler.insertOrderHandler.insertDB.Marshal] Error when do json Marshalling for error handling : %v \n", err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(orderJSON)

	} else {
		table := model.Table{
			MejaID: ID,
			Status: "Close",
		}

		data := model.ResponseWrapper{
			Success: false,
			Message: "Table You Chooses Is Still Closed",
			Data:    table,
		}

		orderJSON, err := json.Marshal(data)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Ooops, Something Went Wrong"))
			fmt.Printf("[OrderHandler.insertOrderHandler.insertDB.Marshal] Error when do json Marshalling for error handling : %v \n", err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(orderJSON)
	}
}

// CreateOrderHandler calling for routing
func CreateOrderHandler(r *mux.Router) {
	order := OrderHandler{}
	r.HandleFunc("/order", order.insertOrderHandler).Methods(http.MethodPost)
}
