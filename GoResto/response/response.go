package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tugas/GoResto/model"
)

// RespondWithError call for error
func RespondWithError(w http.ResponseWriter, msg string) {
	w.Header().Set("Content-Type", "application/json")
	error := model.Error{
		Status:  false,
		Message: msg,
	}
	err := json.NewEncoder(w).Encode(error)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Printf("[DetailHandler.getDetailOrder.RespondWithError] Error when rows RespondWithError with error : %v\n", err.Error())
		return
	}
	return
}

// RespondWithSuccess call for success
func RespondWithSuccess(w http.ResponseWriter, msg string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	success := model.Success{
		Status:  true,
		Message: msg,
		Data:    data,
	}
	err := json.NewEncoder(w).Encode(success)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Printf("[DetailHandler.getDetailOrder.RespondWithSuccess] Error when Marshal with error : %v\n", err.Error())
		return
	}
	return
}
