package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"tugas/GoPORT/kapal"
	"tugas/GoPORT/model"
	"tugas/GoPORT/response"
)

// KapalHandler struct
type KapalHandler struct {
	kapalUsecase kapal.KapalUsecase
}

// CreateKapalHandler router for kapal will call for function main
func CreateKapalHandler(r *mux.Router, kapalUsecase kapal.KapalUsecase) {
	kapal := KapalHandler{kapalUsecase}
	r.HandleFunc("/kapal", kapal.getAllKapal).Methods(http.MethodGet)
	r.HandleFunc("/kapal", kapal.insertDataKapal).Methods(http.MethodPost)
	r.HandleFunc("/kapal/{id}", kapal.getKapalByID).Methods(http.MethodGet)
	r.HandleFunc("/kapal/{id}", kapal.updateDataKapal).Methods(http.MethodPut)
	r.HandleFunc("/kapal/{id}", kapal.deleteKapalByID).Methods(http.MethodDelete)
}

func (k *KapalHandler) getAllKapal(w http.ResponseWriter, r *http.Request) {
	kapal, err := k.kapalUsecase.GetAllKapal()
	if kapal != nil {
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				response.RespondWithError(w, "There is no data ")
				fmt.Printf("[KapalHandler.CreateKapalHandler.getAllKapal.ErrNoRows] Error with >> %v\n", sql.ErrNoRows)
				return
			}
			response.RespondWithError(w, "Data have problem!")
			fmt.Printf("[KapalHandler.CreateKapalHandler.getAllKapal] Error with : %v", err)
			return
		}
		response.RespondWithSuccess(w, "Success", kapal)
		return
	}
	response.RespondWithError(w, "Data tidak tersedia")
	return
}

func (k *KapalHandler) getKapalByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	strID := vars["id"]

	id, err := strconv.Atoi(strID)
	if err != nil {
		response.RespondWithError(w, "ID harus angka ya!")
		fmt.Printf("[KapalHandler.CreateKapalHandler.getKapalByID.Atoi] Error with : %v", err)
		return
	}
	kapal, err := k.kapalUsecase.GetByID(id)

	if kapal != nil {
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				response.RespondWithError(w, "There is no data by id")
				fmt.Printf("[KapalHandler.CreateKapalHandler.GetByID.ErrNoRows] Error with >> %v\n", sql.ErrNoRows)
				return
			}
			response.RespondWithError(w, "Data have problem!")
			fmt.Printf("[KapalHandler.CreateKapalHandler.getKapalByID.GetByID] Error with : %v", err)
			return
		}
		response.RespondWithSuccess(w, "Success", kapal)
		return
	}
	response.RespondWithError(w, "data ID yang anda input tidak tersedia")
	return

}

func (k *KapalHandler) deleteKapalByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	strID := vars["id"]

	id, err := strconv.Atoi(strID)
	if err != nil {
		response.RespondWithError(w, "ID harus angka ya!")
		fmt.Printf("[KapalHandler.CreateKapalHandler.deleteKapalByID.Atoi] Error with : %v", err)
		return
	}

	err = k.kapalUsecase.DeleteByID(id)
	if err != nil {
		response.RespondWithError(w, "Sorry, ID not valid")
		fmt.Printf("[KapalHandler.CreateKapalHandler.deleteKapalByID.kapalUsecase] Error when checking id with : %v", err)
		return
	}

	response.RespondWithSuccess(w, "Delete data success", nil)
	return

}

func (k *KapalHandler) insertDataKapal(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.RespondWithError(w, "Ooops, Something went wrong")
		fmt.Println("[KapalHandler.insertDataKapal.ReadAll] error when reading request body : " + err.Error())
		return
	}
	kapal := model.Kapal{}
	err = json.Unmarshal(reqBody, &kapal)
	if err != nil {
		response.RespondWithError(w, "Ooops, Something went wrong")
		fmt.Println("[KapalHandler.insertDataKapal.Unmarshal] error when Unmarshal request body : " + err.Error())
		return
	}

	validation(kapal, w)

	err = k.kapalUsecase.InsertData(&kapal)
	if err != nil {
		response.RespondWithError(w, "Oops! Something went wrong, ship code can't be same code")
		fmt.Printf("[KapalHandler.insertDataKapal.kapalUsecase.InsertData] error when call insert service : %v", err.Error())
		return
	}
	response.RespondWithSuccess(w, "Insert Data Success", kapal)
	return

}

func validation(data model.Kapal, w http.ResponseWriter) {
	if data.Kode == "" {
		response.RespondWithError(w, "Kode must be Exist")
		return
	}

	if data.Muatan == 0 {
		response.RespondWithError(w, "Muatan must be Exist")
		return
	}

	if data.Status == "" {
		response.RespondWithError(w, "Status must be Exist")
		return
	}

	if data.IsDelete == 0 {
		data.IsDelete = 0
		return
	}

}

func (k *KapalHandler) updateDataKapal(w http.ResponseWriter, r *http.Request) {

	kapal := model.Kapal{}
	vars := mux.Vars(r)
	strID := vars["id"]

	id, err := strconv.Atoi(strID)
	if err != nil {
		response.RespondWithError(w, "ID kapal harus angka ya!")
		fmt.Printf("[KapalHandler.CreateKapalHandler.updateDataKapal.Atoi] Error with : %v", err)
		return
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.RespondWithError(w, "Ooops, Something went wrong")
		fmt.Println("[KapalHandler.updateDataKapal.ReadAll] error when reading request body : " + err.Error())
		return
	}

	kapal = model.Kapal{}
	err = json.Unmarshal(reqBody, &kapal)
	if err != nil {
		response.RespondWithError(w, "Ooops, Something went wrong")
		fmt.Println("[KapalHandler.updateDataKapal.Unmarshal] error when Unmarshal request body : " + err.Error())
		return
	}

	validation(kapal, w)

	err = k.kapalUsecase.UpdateData(id, &kapal)
	if err != nil {
		response.RespondWithError(w, "Oops! Something went wrong, ship code can't be same code")
		fmt.Printf("[KapalHandler.updateDataKapal.kapalUsecase.InsertData] error when call insert service : %v", err.Error())
		return
	}

	response.RespondWithSuccess(w, "Update Data Success", kapal)
	return

}
