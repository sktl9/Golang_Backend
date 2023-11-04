package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Замените PostController на реализацию вашего контроллера
// Например, это могут быть функции, которые принимают http.ResponseWriter и *http.Request

func Router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/posts", PostControllerCreate).Methods("POST")
	r.HandleFunc("/posts", PostControllerGetAll).Methods("GET")
	r.HandleFunc("/update", PostControllerUpdateOne).Methods("PUT")
	r.HandleFunc("/posts/{id}", PostControllerGetOne).Methods("GET")
	r.HandleFunc("/posts", PostControllerUpdate).Methods("PUT")
	r.HandleFunc("/posts/{id}", PostControllerDelete).Methods("DELETE")

	return r
}

func PostControllerCreate(w http.ResponseWriter, r *http.Request) {
	// Ваш код для создания поста
}

func PostControllerGetAll(w http.ResponseWriter, r *http.Request) {
	// Ваш код для получения всех постов
}

func PostControllerUpdateOne(w http.ResponseWriter, r *http.Request) {
	// Ваш код для обновления одного поста
}

func PostControllerGetOne(w http.ResponseWriter, r *http.Request) {
	// Ваш код для получения одного поста
}

func PostControllerUpdate(w http.ResponseWriter, r *http.Request) {
	// Ваш код для обновления поста
}

func PostControllerDelete(w http.ResponseWriter, r *http.Request) {
	// Ваш код для удаления поста
}
