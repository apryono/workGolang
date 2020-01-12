package transactions

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Trans struct {
	OrderID    int    `json:"Order ID"`
	TransID    int    `json:"Transaction ID"`
	MejaID     int    `json:"Nomor Meja"`
	Status     string `json:"Status"`
	MenuID     int    `json:"Menu ID"`
	Menu       string `json:"Menu"`
	Qty        int    `json:"Qty"`
	Note       string `json:"Note"`
	Price      int    `json:"Harga"`
	Total      int    `json:"Total"`
	GrandTotal int    `json:"Grand Total"`
}

var trx = Trans{}
var result []Trans

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

func GetTransactionHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	db, err := connect()
	if err != nil {
		w.Write([]byte("Oops, something error"))
		fmt.Printf("[GetTransactionHandler] Error when connect database with error : %v\n ", err.Error())
		return
	}
	defer db.Close()

	err = db.QueryRow("Select count(tr_id) from transactions").Scan(&trx.TransID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Opps. Something went wrong. Please contact admin"))
		fmt.Printf("[GetTransactionHandler] Error when scan with error : %v\n ", err.Error())
		return
	}
	if trx.TransID >= id {
		err = db.QueryRow("select meja_id, nama, qty, harga, (qty * harga) from menu join orders join transactions on menu.menu_id = orders.menu_id and transactions.tr_id = orders.transaksi_id group by nama").Scan(&trx.MejaID, &trx.Menu, &trx.Qty, &trx.Price, &trx.Total)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Opps. Something went wrong. Please contact admin"))
			fmt.Printf("[GetTransactionHandler] Error when Query Select join with error : %v\n ", err.Error())
			return
		}

		err = db.QueryRow("select sum(qty * harga) as grandTotal from menu join orders join transactions on menu.menu_id = orders.menu_id and transactions.tr_id = orders.transaksi_id where tr_id = ?", id).Scan(&trx.GrandTotal)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Opps. Something went wrong. Please contact admin"))
			fmt.Printf("[GetTransactionHandler] Error when Query Grand Total with error : %v\n ", err.Error())
			return
		}

		result = append(result, trx)
		json, err := json.Marshal(result)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Opps. Something went wrong. Please contact admin"))
			fmt.Printf("[GetTransactionHandler] Error when json.Marshal with error : %v\n ", err.Error())
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(json)

		_, err = db.Exec("UPDATE meja set status = 'open' where meja_id = ?", &trx.MejaID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Opps. Something went wrong. Please contact admin"))
			fmt.Printf("[GetTransactionHandler] Error when update with error : %v\n ", err.Error())
			return
		}
	} else {
		w.Write([]byte("Sorry!!!, Data Not Found "))
	}

}
