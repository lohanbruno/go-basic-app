package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lohanbruno/go-bookstore/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods(http.MethodPost)
	router.HandleFunc("/book/", controllers.GetBooks).Methods(http.MethodGet)
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods(http.MethodGet)
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods(http.MethodPut)
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods(http.MethodDelete)
}
