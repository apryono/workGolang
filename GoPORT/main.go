package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	dock "tugas/GoPORT/dermaga/handler"
	dockpo "tugas/GoPORT/dermaga/repo"
	dockcase "tugas/GoPORT/dermaga/usecase"
	kapal "tugas/GoPORT/kapal/handler"
	kapalpo "tugas/GoPORT/kapal/repo"
	kapalcase "tugas/GoPORT/kapal/usecase"
)

func main() {
	port := "8080"
	conStr := "root:@tcp(localhost:3306)/ship_harbor" // string untuk memanggil database name

	db, err := sql.Open("mysql", conStr)
	if err != nil {
		log.Fatal("Error When Connect to DB " + conStr + " : " + err.Error()) // log fatal digunakan karena tidak terhubung ke database
	}
	defer db.Close()

	r := mux.NewRouter().StrictSlash(true) // memanggil library mux router

	kapalRepo := kapalpo.CreateKapalRepoMysql(db)           // calling repository query
	kapalUsecase := kapalcase.CreateKapalUsecase(kapalRepo) // calling usecase validation
	dockRepo := dockpo.CreateDermagaRepoMysql(db)
	dockUsecase := dockcase.CreateDermagaUsecase(dockRepo)

	/* ======== End Call Repo And UseCase ========== */

	kapal.CreateKapalHandler(r, kapalUsecase) // calling handler to run the router
	dock.CreateDermagaHandler(r, dockUsecase) // calling handler to run the router

	fmt.Println("Starting Web Server at port : " + port)
	err = http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatal(err)
	}
}
