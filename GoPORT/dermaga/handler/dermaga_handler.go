package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"tugas/GoPORT/dermaga"
	"tugas/GoPORT/model"
	"tugas/GoPORT/response"

	"github.com/gorilla/mux"
)

// DermagaHandler struct
type DermagaHandler struct {
	dermagaUsecase dermaga.DermagaUsecase
}

// CreateDermagaHandler router for Dermaga will call for function main
func CreateDermagaHandler(r *mux.Router, dermagaUsecase dermaga.DermagaUsecase) {
	dock := DermagaHandler{dermagaUsecase}
	r.HandleFunc("/dock", dock.getAllDermaga).Methods(http.MethodGet)
	r.HandleFunc("/dock", dock.insertDermaga).Methods(http.MethodPost)
	r.HandleFunc("/dock/{id}", dock.getDermagaByID).Methods(http.MethodGet)
	r.HandleFunc("/dock/{id}", dock.updateDataDermaga).Methods(http.MethodPut)
	r.HandleFunc("/dock/{id}", dock.deleteDataByID).Methods(http.MethodDelete)
}

func (d *DermagaHandler) getAllDermaga(w http.ResponseWriter, r *http.Request) {
	dock, err := d.dermagaUsecase.GetAllDermaga()
	if dock != nil {
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				response.RespondWithError(w, "There is no data ")
				fmt.Printf("[DermagaHandler.CreateDermagaHandler.getAllKapal.ErrNoRows] Error with >> %v\n", sql.ErrNoRows)
				return
			}
			response.RespondWithError(w, "Data have problem!")
			fmt.Printf("[DermagaHandler.CreateDermagaHandler.getAllKapal.GetAllDermaga] Error with : %v", err)
			return
		}
		response.RespondWithSuccess(w, "Success", dock)
		return
	}
	response.RespondWithError(w, "Data tidak tersedia")
	return
}

func (d *DermagaHandler) getDermagaByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	strID := vars["id"]

	id, err := strconv.Atoi(strID)
	if err != nil {
		response.RespondWithError(w, "ID harus angka ya!")
		fmt.Printf("[DermagaHandler.CreateDermagaHandler.getDermagaByID.Atoi] Error with : %v", err)
		return
	}
	dock, err := d.dermagaUsecase.GetByID(id)

	if dock != nil {
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				response.RespondWithError(w, "There is no data by id")
				fmt.Printf("[DermagaHandler.CreateDermagaHandler.getDermagaByID.GetByID.ErrNoRows] Error with >> %v\n", sql.ErrNoRows)
				return
			}
			response.RespondWithError(w, "Data have problem!")
			fmt.Printf("[DermagaHandler.CreateDermagaHandler.getDermagaByID.GetByID] Error with : %v", err)
			return
		}
		response.RespondWithSuccess(w, "Success", dock)
		return
	}
	response.RespondWithError(w, "data ID yang anda input tidak tersedia")
	return
}

func (d *DermagaHandler) insertDermaga(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.RespondWithError(w, "Ooops, Something went wrong")
		fmt.Println("[DermagaHandler.insertDermagaByID.ReadAll] error when reading request body : " + err.Error())
		return
	}
	dock := model.Dermaga{}
	err = json.Unmarshal(reqBody, &dock)
	if err != nil {
		response.RespondWithError(w, "Ooops, Something went wrong")
		fmt.Println("[DermagaHandler.insertDermagaByID.Unmarshal] error when Unmarshal request body : " + err.Error())
		return
	}

	validation(dock, w)

	err = d.dermagaUsecase.InsertData(&dock)
	if err != nil {
		response.RespondWithError(w, "Oops! Something went wrong, dock code can't be same code")
		fmt.Printf("[DermagaHandler.insertDermagaByID.dermagaUsecase.InsertData] error when call insert service : %v", err.Error())
		return
	}
	response.RespondWithSuccess(w, "Insert Data Success", dock)
	return
}

func validation(data model.Dermaga, w http.ResponseWriter) {
	if data.Kode == "" {
		response.RespondWithError(w, "Kode must be Exist")
		return
	}

	if data.Status == "" {
		response.RespondWithError(w, "Status must be Exist")
		return
	}

}

func (d *DermagaHandler) updateDataDermaga(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	strID := vars["id"]

	id, err := strconv.Atoi(strID)
	if err != nil {
		response.RespondWithError(w, "ID dock harus angka ya!")
		fmt.Printf("[KapalHandler.CreateKapalHandler.updateDataKapal.Atoi] Error with : %v", err)
		return
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.RespondWithError(w, "Ooops, Something went wrong")
		fmt.Println("[DermagaHandler.insertDermagaByID.ReadAll] error when reading request body : " + err.Error())
		return
	}
	dock := model.Dermaga{}
	err = json.Unmarshal(reqBody, &dock)
	if err != nil {
		response.RespondWithError(w, "Ooops, Something went wrong")
		fmt.Println("[DermagaHandler.insertDermagaByID.Unmarshal] error when Unmarshal request body : " + err.Error())
		return
	}

	validation(dock, w)

	err = d.dermagaUsecase.UpdateData(id, &dock)
	if err != nil {
		response.RespondWithError(w, "Oops! Something went wrong, ship code can't be same code")
		fmt.Printf("[DermagaHandler.insertDermagaByID.UpdateData] error when call insert service : %v", err.Error())
		return
	}

	response.RespondWithSuccess(w, "Update Data Success", dock)
	return

}

func (d *DermagaHandler) deleteDataByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	strID := vars["id"]

	id, err := strconv.Atoi(strID)
	if err != nil {
		response.RespondWithError(w, "ID harus angka ya!")
		fmt.Printf("[DermagaHandler.CreateDermagaHandler.deleteDataByID.Atoi] Error with : %v", err)
		return
	}

	err = d.dermagaUsecase.DeleteByID(id)
	if err != nil {
		response.RespondWithError(w, "Sorry, ID not valid")
		fmt.Printf("[DermagaHandler.CreateDermagaHandler.deleteDataByID.dermagaUsecase] Error when checking id with : %v", err)
		return
	}

	response.RespondWithSuccess(w, "Delete data success", nil)
	return
}
