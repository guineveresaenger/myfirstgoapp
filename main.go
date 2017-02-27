package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

func index(w http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		log.Printf("%v", err)
	}
	t.Execute(w, "./templates/index.html")

}

func colors(w http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")
	log.Println(name)

	// s := `./templates/colors.html`
	// t := template.Must(template.New("test").Parse(s))
	t, _ := template.ParseFiles("./templates/colors.html")
	t.Execute(w, name)
}

func GetPersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}

	}
	json.NewEncoder(w).Encode(&Person{})
}

func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var person Person
	_ = json.NewDecoder(req.Body).Decode(&person) // ?? what does this line do exactly?
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range people {
		if item.ID == params["id"] {
			// the following line slices right before current index, and right after, skipping the current index.
			// Read up on the ... in Go.
			people = append(people[:index], people[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(people)
}

func main() {
	router := mux.NewRouter()
	people = append(people, Person{ID: "1", Firstname: "Guin", Lastname: "Awesome", Address: &Address{City: "Seattle", State: "WA"}})
	people = append(people, Person{ID: "2", Firstname: "Brendan", Lastname: "Batman"})

	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("DELETE")
	router.HandleFunc("/", index).Methods("GET")
	router.HandleFunc("/colors", colors).Methods("GET")

	log.Fatal(http.ListenAndServe(":12345", router))

}
