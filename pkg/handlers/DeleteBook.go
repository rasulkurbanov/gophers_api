package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rasulkurbanov/gophers_api/pkg/mocks"
)



func DeleteBook(w http.ResponseWriter, r *http.Request) {
	//Read the dynamic ID parameter
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])


	//Iterate over all mocks.Book array 
	for index, book := range mocks.Books {
		if book.Id == id {
			mocks.Books = append(mocks.Books[:index], mocks.Books[index+1:]...)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("Deleted!")
			break
		}
	}
}