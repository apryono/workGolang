package routes

import (
	"net/http"
	"tugas/RestoranModular/controllers/menus"
	"tugas/RestoranModular/controllers/order"
	"tugas/RestoranModular/controllers/tables"
	"tugas/RestoranModular/controllers/transactions"

	"github.com/gorilla/mux"
)

/* untuk di panggil di fungsi main main*/
func MenusHandler(r *mux.Router) {
	r.StrictSlash(true)
	r.HandleFunc("/", menus.GetAllMenu).Methods(http.MethodGet)
}

func TablesHandler(r *mux.Router) {
	r.StrictSlash(true)
	r.HandleFunc("/", tables.GetAllStatusMeja).Methods(http.MethodGet)
	r.HandleFunc("/{id}", tables.GetStatusMejaByID).Methods(http.MethodPut)
}

func OrderHandler(r *mux.Router) {
	r.StrictSlash(true)
	r.HandleFunc("/", order.InsertOrderHandler).Methods(http.MethodPost)
}

func TransactionHandler(r *mux.Router) {
	r.StrictSlash(true)
	r.HandleFunc("/{id}", transactions.GetTransactionHandler).Methods(http.MethodPost)
}
