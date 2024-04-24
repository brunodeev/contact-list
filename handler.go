package main

import (
	"time"
	"net/http"
	"encoding/json"
	"github.com/google/uuid"
)

func GetContacts(res http.ResponseWriter, req *http.Request) {
	result, _ := json.Marshal(ContactList)
	res.Write(result)
}

func CreateContact(res http.ResponseWriter, req *http.Request) {
	var contact Contact
	_ = json.NewDecoder(req.Body).Decode(&contact)
	contact.ID = uuid.New().String()
	tempTime := time.Now()
	contact.Created = tempTime.Format("02-01-2006 15:04:05")
	
	ContactList = append(ContactList, contact)
	res.Write([]byte("Cadastrado com sucesso!"))
}

func EditContact(res http.ResponseWriter, req *http.Request){
	query := req.URL.Query().Get("id")
	index := -1
	var contact Contact
	
	for i, contact := range ContactList {
		if contact.ID == query {
			index = i
		}
	}
	
	if index == -1 {
		res.Write([]byte("Falha ao atualizar!"))
		return
	}
	
	_ = json.NewDecoder(req.Body).Decode(&contact)
	
	if contact.Name != "" {
		ContactList[index].Name = contact.Name
	}
	if contact.Phones != nil {
		ContactList[index].Phones = contact.Phones
	}
	
	tempTime := time.Now()
	ContactList[index].Created = tempTime.Format("02-01-2006 15:04:05")
	
	res.Write([]byte("Atualizado com sucesso!"))
}

func DeleteContact(res http.ResponseWriter, req *http.Request){
	query := req.URL.Query().Get("id")
	var index int
	
	for i, contact := range ContactList {
		if contact.ID == query {
			index = i
		}
	}
	ContactList = append(ContactList[:index], ContactList[index+1:]...)
	res.Write([]byte("Removido com sucesso!"))
}



















