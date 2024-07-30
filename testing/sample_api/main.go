package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	iName := os.Getenv("INSTANCE_NAME")
	router := http.NewServeMux()

	router.HandleFunc("POST /v1/users", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("created new user %s", iName)))
	})

	router.HandleFunc("GET /v1/users", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("all users for %s", iName)))
	})

	router.HandleFunc("PATCH /v1/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		w.Write([]byte(fmt.Sprintf("updated user %s %s", id, iName)))
	})

	router.HandleFunc("DELETE /v1/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		w.Write([]byte(fmt.Sprintf("deleted user %s %s", id, iName)))
	})

	http.ListenAndServe(":8080", router)
}
