package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Student struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Student
}

var students []Student

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/student")
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

func getAllStudent(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {

		var students Student
		var result []Student
		var response Response

		db, err := connect()
		handleError(err)
		defer db.Close()

		rows, err := db.Query("Select id,name,gender from tb_student")
		handleError(err)

		for rows.Next() {
			if err := rows.Scan(&students.ID, &students.Name, &students.Gender); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Oops! Something went wrong."))
				return
			}

			result = append(result, students)

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

func getStudentByID(w http.ResponseWriter, req *http.Request) {
	var students Student
	var result []Student
	var response Response

	db, err := connect()
	handleError(err)
	defer db.Close()

	// mendapatkan id dengan fitur mux
	vars := mux.Vars(req)
	fmt.Println(vars["id"])

	// id, _ := strconv.Atoi(req.FormValue("id"))
	id, _ := strconv.Atoi(vars["id"])

	rows, err := db.Query("Select id,name,gender from tb_student where id = ?", id)
	handleError(err)

	for rows.Next() {
		if err := rows.Scan(&students.ID, &students.Name, &students.Gender); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Oops! Something went wrong."))
			return
		}

		result = append(result, students)

	}

	response.Status = 1
	response.Message = "Success"
	response.Data = result

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func insertStudent(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		db, err := connect()
		handleError(err)
		defer db.Close()

		tx, err := db.Begin()
		handleError(err)

		var result []Student
		var response Response

		reqBody, err := ioutil.ReadAll(req.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Opps. Something went wrong. Please contact admin"))
			fmt.Println(err.Error())
			tx.Rollback()
			return
		}
		var std Student
		err = json.Unmarshal(reqBody, &std)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Opps. Something went wrong. Please contact admin"))
			fmt.Println(err.Error())
			tx.Rollback()
			return
		}

		result = append(result, std)

		for i := 0; i < len(result); i++ {
			stmt, err := db.Prepare("insert into tb_student (name,gender) values (?,?)")
			handleError(err)
			stmt.Exec(result[i].Name, result[i].Gender)
		}

		err = tx.Commit()
		handleError(err)

		response.Status = 1
		response.Message = "Success"
		w.Write([]byte("Insert data to database"))

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

func updateStudent(w http.ResponseWriter, req *http.Request) {
	if req.Method == "PUT" {
		db, err := connect()
		handleError(err)
		defer db.Close()

		tx, err := db.Begin()
		handleError(err)

		var result []Student
		var response Response

		id, _ := strconv.Atoi(req.FormValue("id"))

		reqBody, err := ioutil.ReadAll(req.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Opps. Something went wrong. Please contact admin"))
			fmt.Println(err.Error())
			tx.Rollback()
			return
		}
		var std Student
		err = json.Unmarshal(reqBody, &std)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Opps. Something went wrong. Please contact admin"))
			fmt.Println(err.Error())
			tx.Rollback()
			return
		}

		result = append(result, std)
		for i := 0; i < len(result); i++ {
			stmt, err := db.Prepare("UPDATE tb_student SET name = ?, gender = ? where id = ?")
			handleError(err)
			stmt.Exec(result[i].Name, result[i].Gender, id)
		}

		err = tx.Commit()
		handleError(err)

		response.Status = 1
		response.Message = "Success"
		w.Write([]byte("Update data to database"))

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)

	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Oops!!.Method tidak di Support"))
		return
	}
}

func deleteStudent(w http.ResponseWriter, req *http.Request) {
	if req.Method == "DELETE" {
		db, err := connect()
		handleError(err)
		defer db.Close()

		tx, err := db.Begin()
		handleError(err)

		var response Response

		id, _ := strconv.Atoi(req.FormValue("id"))

		_, err = db.Exec("DELETE from tb_student where id = ?", id)
		handleError(err)

		err = tx.Commit()
		handleError(err)

		response.Status = 1
		response.Message = "Success"
		w.Write([]byte("Delete data from database"))

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Oops!!.Method tidak di Support"))
		return
	}
}

func studentHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		getStudentByID(w, req)
	} else if req.Method == "POST" {
		insertStudent(w, req)
	} else if req.Method == "PUT" {
		updateStudent(w, req)
	} else if req.Method == "DELETE" {
		deleteStudent(w, req)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method tidak di Support"))
		return
	}
}

func main() {
	port := "1234"
	router := mux.NewRouter()
	router.HandleFunc("/students", getAllStudent)
	router.HandleFunc("/student/{id}", studentHandler) // menggunakan path variable
	router.HandleFunc("/student", studentHandler)
	http.Handle("/", router)
	fmt.Println("Starting Web Server At Port : " + port)
	fmt.Println(http.ListenAndServe(":"+port, router))
}
