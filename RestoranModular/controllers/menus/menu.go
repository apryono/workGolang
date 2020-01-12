package menus

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Menu struct {
	ID    int    `json:"ID"`
	Name  string `json:"Nama"`
	Price int    `json:"Harga"`
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Menu
}

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

func InputMenusHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		db, err := connect()
		handleError(err)
		defer db.Close()

		tx, err := db.Begin()
		handleError(err)

		var response Response

		reqBody, err := ioutil.ReadAll(req.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Opps. Something went wrong. Please contact admin"))
			fmt.Printf("[InputMenusHandler] Error when request body with error : %v\n", err.Error())
			tx.Rollback()
			return
		}
		var menu Menu
		err = json.Unmarshal(reqBody, &menu)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Opps. Something went wrong. Please contact admin"))
			fmt.Printf("[InputMenusHandler] Error when json.Unmarshal with error : %v\n", err.Error())
			tx.Rollback()
			return
		}

		_, err = db.Exec("insert into menu (nama,harga) values (?,?)", menu.Name, menu.Price)
		handleError(err)

		err = tx.Commit()
		handleError(err)

		response.Status = 1
		response.Message = "Success"
		w.Write([]byte("Insert Menu to database"))

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

func GetAllMenu(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {

		var menus Menu
		var result []Menu
		var response Response

		db, err := connect()
		handleError(err)
		defer db.Close()

		rows, err := db.Query("Select menu_id,nama,harga from menu")
		handleError(err)

		for rows.Next() {
			if err := rows.Scan(&menus.ID, &menus.Name, &menus.Price); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Oops! Something went wrong."))
				fmt.Printf("[GetAllMenu] Error when rows scan with error : %v\n", err.Error())
				return
			}

			result = append(result, menus)

		}

		response.Status = 1
		response.Message = "Success"
		response.Data = result

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Oops!! Method tidak di Support"))
		return
	}

}
