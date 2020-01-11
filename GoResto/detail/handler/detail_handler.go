package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"tugas/GoResto/connect"
	"tugas/GoResto/model"

	"github.com/gorilla/mux"
)

// DetailHandler struct
type DetailHandler struct {
}

func (d DetailHandler) getDetailOrder(w http.ResponseWriter, r *http.Request) {

	var trx = model.Transaction{}
	// var getTrx = []model.Transaction{}
	var detail = model.Detail{}
	var getDetail = []model.Detail{}

	db, err := connect.ConnectHandler()
	if err != nil {
		fmt.Printf("[DetailHandler.getDetailOrder.ConnectHandler] Error when connecting to database with error : %v \n", err.Error())
	}
	defer db.Close()

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	err = db.QueryRow("Select tr_id,meja_id,status_transaksi from transactions where meja_id = ?", id).Scan(&trx.TransID, &trx.MejaID, &trx.StatusTR)
	if err != nil {
		fmt.Printf("[DetailHandler.getDetailOrder.QueryRow] Error when rowsDT QueryRow data with error : %v \n", err.Error())
	}

	status := "uncheck"
	query := "select sum(qty*harga) from menu join orders join transactions on menu.menu_id = orders.menu_id and transactions.tr_id = orders.transaksi_id where meja_id = ? and status_transaksi = ?"
	err = db.QueryRow(query, id, status).Scan(&trx.GranTot)

	qry := "select nama, qty, harga, (qty*harga) from menu join orders join transactions on menu.menu_id = orders.menu_id and transactions.tr_id = orders.transaksi_id where meja_id = ? and status_transaksi = ? group by nama"
	rowsDT, err := db.Query(qry, id, status)
	if err != nil {
		fmt.Printf("[DetailHandler.getDetailOrder.Query.rowsDT] Error when rowsDT query data with error : %v \n", err.Error())
	}

	for rowsDT.Next() {
		if err = rowsDT.Scan(&detail.Menu, &detail.Qty, &detail.Price, &detail.Total); err != nil {
			data := model.ResponseWrapper{
				Success: false,
				Message: "Scanning is no valid",
			}
			menuJSON, err := json.Marshal(data)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Oops! Something went wrong."))
				fmt.Printf("[DetailHandler.getDetailOrder.Scan.Marshal.rowsDT] Error when rowsTR scan with error : %v\n", err.Error())
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(menuJSON)
			return
		}
		getDetail = append(getDetail, detail)
	}
	defer rowsDT.Close()

	if trx.MejaID == id && trx.StatusTR == "uncheck" {
		trx = model.Transaction{
			TransID:  trx.TransID,
			MejaID:   id,
			StatusTR: "Belum dibayar",
			GranTot:  trx.GranTot,
			Detail:   getDetail,
		}

		data := model.ResponseWrapper{
			Success: true,
			Message: "Success",
			Data:    trx,
		}
		tableJSON, err := json.Marshal(data)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Oops! Something went wrong."))
			fmt.Printf("[DetailHandler.getDetailOrder.Data.Marshal] Error when rows Marshal with error : %v\n", err.Error())
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(tableJSON)

		_, err = db.Exec("update transactions set status_transaksi = ? where meja_id = ?", "check", id)
		if err != nil {
			fmt.Printf("[DetailHandler.getDetailOrder.Exec] Error when update data transactions with >> %v\n ", err)
			return
		}
		return

	}
	data := model.ResponseWrapper{
		Success: false,
		Message: "Sorry, Nothing Data Transaction",
	}
	tableJSON, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Oops! Something went wrong."))
		fmt.Printf("[DetailHandler.getDetailOrder.Data.Marshal] Error when rows Marshal with error : %v\n", err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(tableJSON)
	return

}

// CreateDetailHandler calling for routing
func CreateDetailHandler(r *mux.Router) {
	detail := DetailHandler{}
	r.HandleFunc("/detail/{id}", detail.getDetailOrder).Methods(http.MethodGet)
}
