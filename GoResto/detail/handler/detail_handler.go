package handler

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"tugas/GoResto/connect"
	"tugas/GoResto/model"
	res "tugas/GoResto/response"
	success "tugas/GoResto/response"

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

	err = db.QueryRow("Select tr_id,meja_id,status_transaksi from transactions where meja_id = ? and status_transaksi = 'uncheck'", id).Scan(&trx.TransID, &trx.MejaID, &trx.StatusTR)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			res.RespondWithError(w, "Maaf! meja yang ada input tidak tersedia")
			fmt.Printf("[DetailHandler.getDetailOrder.QueryRow.ErrNoRows] with error :")
			return
		}
		res.RespondWithError(w, "Oops! Something went wrong")
		fmt.Printf("[DetailHandler.getDetailOrder.QueryRow] Error when rowsDT QueryRow data with error : %v \n", err.Error())
		return
	}

	status := "uncheck"
	query := "select sum(qty*harga) from menu join orders join transactions on menu.menu_id = orders.menu_id and transactions.tr_id = orders.transaksi_id where meja_id = ? and status_transaksi = ?"
	err = db.QueryRow(query, id, status).Scan(&trx.GranTot)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			res.RespondWithError(w, "Oops! meja tidak tersedia")
			return
		}
		res.RespondWithError(w, "Oops! Something went wrong")
		fmt.Printf("[DetailHandler.getDetailOrder.Query.status] Error when status query data with error : %v \n", err.Error())
		return
	}

	qry := "select nama, qty, harga, (qty*harga) from menu join orders join transactions on menu.menu_id = orders.menu_id and transactions.tr_id = orders.transaksi_id where meja_id = ? and status_transaksi = ? group by nama"
	rowsDT, err := db.Query(qry, id, status)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			res.RespondWithError(w, "Oops! Data tidak tersedia berdasarkan no meja tersebut")
			return
		}
		res.RespondWithError(w, "Oops! Something went wrong")
		fmt.Printf("[DetailHandler.getDetailOrder.Query.rowsDT] Error when rowsDT query data with error : %v \n", err.Error())
		return
	}

	for rowsDT.Next() {
		if err = rowsDT.Scan(&detail.Menu, &detail.Qty, &detail.Price, &detail.Total); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				res.RespondWithError(w, "Maaf, Data tidak tersedia")
				return
			}
			res.RespondWithError(w, "Oops! Something went wrong")
			fmt.Printf("[DetailHandler.getDetailOrder.Query.rowsDT] Error with : %v\n ", err)
			return
		}
		getDetail = append(getDetail, detail)
		fmt.Println(getDetail)
	}
	defer rowsDT.Close()

	if trx.MejaID != 0 {
		trx = model.Transaction{
			TransID:  trx.TransID,
			MejaID:   id,
			StatusTR: "Belum dibayar",
			Detail:   getDetail,
			GranTot:  trx.GranTot,
		}

		success.RespondWithSuccess(w, "Success", trx)

		_, err = db.Exec("update transactions set status_transaksi = ? where meja_id = ?", "check", id)
		if err != nil {
			fmt.Printf("[DetailHandler.getDetailOrder.Exec] Error when update data transactions with >> %v\n ", err)
			return
		}

		_, err = db.Exec("update meja set status = ? where meja_id = ?", "close", id)
		if err != nil {
			fmt.Printf("[DetailHandler.getDetailOrder.Exec] Error when update data meja with >> %v\n", err)
			return
		}
		return

	}
	res.RespondWithError(w, "Sorry, Nothing Data Transaction")
	return

}

// CreateDetailHandler calling for routing
func CreateDetailHandler(r *mux.Router) {
	detail := DetailHandler{}
	r.HandleFunc("/detail/{id}", detail.getDetailOrder).Methods(http.MethodGet)
}
