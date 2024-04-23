package main

import (
	"net/http"
	"log"
	"fmt"
)

func main() {
	fmt.Println("Rodando...")
	http.HandleFunc("/contacts", GetContacts)
	http.HandleFunc("/contacts/create", CreateContact)
	http.HandleFunc("/contacts/edit", EditContact)
	http.HandleFunc("/contacts/delete", DeleteContact)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
