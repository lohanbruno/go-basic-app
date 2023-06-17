package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/lohanbruno/go-bookstore/pkg/models"
	"net/http"
	"strconv"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	bookParams := mux.Vars(r)
	
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, _ := models.GetBookById(id)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}



func UpdateBook(w http.ResponseWriter, r *http.Request) {}

func DeleteBook(w http.ResponseWriter, r *http.Request) {}
