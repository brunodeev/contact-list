package main

import (
	"net/http"
	"log"
)

func main() {
	go LoadFunc()
	http.HandleFunc("/contacts", GetContacts)
	http.HandleFunc("/contacts/create", CreateContact)
	http.HandleFunc("/contacts/edit", EditContact)
	http.HandleFunc("/contacts/delete", DeleteContact)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
