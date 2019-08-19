package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Define struct of contacts
type Contact struct {
	First string `json:"first"`
	Last  string `json:"last'`
	Email string `"json: email"`
}

// Declare global Contacts array (simulation database)
type Contacts []Contact

func allContacts(w http.ResponseWriter, r *http.Request) {
	contacts := Contacts{
		Contact{First: "Test FirstName:", Last: "Test LastName:", Email: "Test Email:"},
		Contact{First: "Kyle", Last: "Gallagher", Email: "kyle.gallagher@gmial.com"},
		Contact{First: "Dan", Last: "Green", Email: "<3u"},
	}

	fmt.Println("Endpoint Hit: All")
	fmt.Fprint(w, "All Contacts:\n")
	json.NewEncoder(w).Encode(contacts)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: homepage")
	fmt.Fprintf(w, "Welcome to the HomePage")

}

// Register endpoints here
func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/all", allContacts)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()
}
