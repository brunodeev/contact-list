package main

type Contact struct {
	ID 	string	`json:"id"`
	Created string 	`json:"created"`
	Name 	string 	`json:"name"`
	Phones 	[]Phone `json:"phones"`
}
