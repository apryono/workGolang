package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tugas/GoResto/connect"
	"tugas/GoResto/model"

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
	}
	defer db.Close()

	rows, err := db.Query("Select menu_id,nama,harga from menu")
	if err != nil {
		fmt.Printf("[MenuHandler.getAllMenu.Query] Error when query data with error : %v \n", err.Error())
	}

	for rows.Next() {
		if err := rows.Scan(&menu.ID, &menu.Name, &menu.Price); err != nil {
			data := model.ResponseWrapper{
				Success: false,
				Message: "Scanning is no valid",
			}
			menuJSON, err := json.Marshal(data)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Oops! Something went wrong."))
				fmt.Printf("[GetAllMeMenuHandler.getAllMenu.Scan.Marshal] Error when rows scan with error : %v\n", err.Error())
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(menuJSON)
			return
		}
		result = append(result, menu)
	}

	defer rows.Close()

	data := model.ResponseWrapper{
		Success: true,
		Message: "Success",
		Data:    result,
	}
	menuJSON, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Oops! Something went wrong."))
		fmt.Printf("[GetAllMeMenuHandler.getAllMenu.Data.Marshal] Error when rows Marshal with error : %v\n", err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(menuJSON)
	return

}

// CreateMenuHandler for call router menu handler
func CreateMenuHandler(r *mux.Router) {
	menu := MenuHandler{}

	r.HandleFunc("/menu", menu.getAllMenu).Methods(http.MethodGet)
}
