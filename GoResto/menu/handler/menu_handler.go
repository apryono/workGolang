package handler

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"tugas/GoResto/connect"
	"tugas/GoResto/model"
	"tugas/GoResto/response"

	"github.com/gorilla/mux"
)

// MenuHandler for struct
type MenuHandler struct {
}

func (m *MenuHandler) getAllMenu(w http.ResponseWriter, r *http.Request) {
	var menu = model.Menu{}
	var result = []model.Menu{}

	db, err := connect.ConnectHandler()
	if err != nil {
		fmt.Printf("[MenuHandler.getAllMenu] Error when connecting to database with error : %v \n", err.Error())
		return
	}
	defer db.Close()

	rows, err := db.Query("Select menu_id,nama,harga from menu")
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			response.RespondWithError(w, "Theer is no data")
			return
		}
		response.RespondWithError(w, "Oops!. Something went wrong")
		fmt.Printf("[MenuHandler.getAllMenu.Query] Error when query data with error : %v \n", err.Error())
		return
	}

	for rows.Next() {
		if err := rows.Scan(&menu.ID, &menu.Name, &menu.Price); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				response.RespondWithError(w, "There is no data")
				return
			}
			response.RespondWithError(w, "Oops! Something went wrong.")
			fmt.Printf("[MenuHandler.getAllMenu.Next] Error when scanning with >> %v\n", err)
			return
		}
		result = append(result, menu)
	}

	defer rows.Close()
	response.RespondWithSuccess(w, "Success", result)
	return

}

// CreateMenuHandler for call router menu handler
func CreateMenuHandler(r *mux.Router) {
	menu := MenuHandler{}

	r.HandleFunc("/menu", menu.getAllMenu).Methods(http.MethodGet)
}
