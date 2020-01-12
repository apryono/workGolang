package orders

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Menu struct {
	MenuID int    `json:"MenuID"`
	Qty    string `json:"Qty"`
	Notes  string `json:"Notes"`
}

type Order struct {
	MejaID  int    `json:"MejaID"`
	Pesanan []Menu `json:"Pesanan"`
}

type Meja struct {
	MejaID int    `json:"MejaID"`
	Status string `json:"Status"`
}

var ListMenu []Menu

var ListMeja []Meja

// type Error struct {
// 	Message string `json: "message"`
// }

// var jsonErr Error

// func respondWithError(w http.ResponseWriter, status int, error Error) {
// 	w.Header().Set("Content-Type", "application/json")
// 	responseJSON(w, status, error)
// }

// func responseJSON(w http.ResponseWriter, status int, data interface{}) {
// 	w.WriteHeader(status)
// 	json.NewEncoder(w).Encode(data)
// }

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

func PostOrderHandler(w http.ResponseWriter, r *http.Request) {
	var data []Order

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Opps. Something went wrong. Please contact admin"))
		fmt.Printf("[GetOrderHandler] Error when request body with error : %v\n ", err.Error())
		return
	}

	err = json.Unmarshal(reqBody, &data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Opps. Something went wrong. Please contact admin"))
		fmt.Printf("[GetOrderHandler] Error when unmarshal with error : %v\n ", err.Error())
		return
	}

	InsertDB(data[0].MejaID, data[0].Pesanan)

}

func InsertDB(ID int, data []Menu) {
	var meja Meja

	db, err := connect()
	handleError(err)
	defer db.Close()

	tx, err := db.Begin()
	handleError(err)

	res, err := db.Exec("INSERT into transactions (meja_id,tgl_transaksi) VALUES (?,current_date())", ID)
	if err != nil {
		fmt.Printf("[InsertDB] Error when insert with error : %v\n ", err.Error())
		tx.Rollback()
		return
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		fmt.Printf("[InsertDB] Error when last insert with error : %v\n ", err.Error())
		tx.Rollback()
		return
	}

	rows, err := db.Query("SELECT status from meja where meja_id = ?", ID)
	if err != nil {
		fmt.Printf("[InsertDB] Error when scan with error : %v\n ", err.Error())
	}

	for rows.Next() {
		err = rows.Scan(&meja.Status)
		if err != nil {
			fmt.Printf("[InsertDB] Error when scan with error : %v\n ", err.Error())
			tx.Rollback()
			return
		}
	}

	if meja.Status == "open" {

		for i := 0; i < len(data); i++ {
			_, err = tx.Exec("INSERT INTO orders (transaksi_id,menu_id,qty,notes) VALUES (?,?,?,?)", lastID, data[i].MenuID, data[i].Qty, data[i].Notes)
			if err != nil {
				fmt.Printf("[InsertDB] Error when insert transaksi with error : %v\n ", err.Error())
				tx.Rollback()
				return
			}
		}

		err = tx.Commit()
		handleError(err)

		fmt.Println("Data Berhasil Di tambahkan")

	} else {
		fmt.Println("Status Meja Closed")
	}

}
