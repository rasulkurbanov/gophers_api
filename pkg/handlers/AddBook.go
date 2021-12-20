package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"

	"github.com/rasulkurbanov/gophers_api/pkg/mocks"
	"github.com/rasulkurbanov/gophers_api/pkg/models"
)


func AddBook(w http.ResponseWriter, r *http.Request) {
	// Read request body	
	defer r.Body.Close()	
	body, err := ioutil.ReadAll(r.Body)

	//If there is an error, log it 
	if err != nil {
		log.Fatalln(err)
	}

	var book models.Book
	json.Unmarshal(body, &book)

	//Append a new book to Book mocks
	book.Id = rand.Intn(100)
	mocks.Books = append(mocks.Books, book)

	//Send a 201 created message response 
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")
}

