package main

import (
	"fmt"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("POST /v1/users", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("created new user"))
	})

	router.HandleFunc("GET /v1/users", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("all users"))
	})

	router.HandleFunc("PATCH /v1/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		w.Write([]byte(fmt.Sprintf("updated user %s", id)))
	})

	router.HandleFunc("DELETE /v1/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		w.Write([]byte(fmt.Sprintf("deleted user %s", id)))
	})

	http.ListenAndServe(":8080", router)
}
