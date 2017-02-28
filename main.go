package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// type Person struct {
// 	ID        string   `json:"id,omitempty"`
// 	Firstname string   `json:"firstname,omitempty"`
// 	Lastname  string   `json:"lastname,omitempty"`
// 	Address   *Address `json:"address,omitempty"`
// }
//
// type Address struct {
// 	City  string `json:"city,omitempty"`
// 	State string `json:"state,omitempty"`
// }
//
// var people []Person

type Color struct {
	ID   string
	Name string
}

var colors []Color

var green = Color{ID: "4", Name: "green"}

func index(w http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		log.Printf("%v", err)
	}
	t.Execute(w, "./templates/index.html")

}

func favoriteColor(w http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")
	log.Println(name)
	t, err := template.ParseFiles("./templates/favoriteColor.html")
	if err != nil {
		log.Printf("%v", err)
	}
	// color := generateRandomColor()

	t.Execute(w, colors)
}

func generateRandomColor() Color {
	var color = Color{ID: "5", Name: "purple"}
	return color
}

// func GetPersonEndpoint(w http.ResponseWriter, req *http.Request) {
// 	params := mux.Vars(req)
// 	for _, item := range people {
// 		if item.ID == params["id"] {
// 			json.NewEncoder(w).Encode(item)
// 			return
// 		}
//
// 	}
// 	json.NewEncoder(w).Encode(&Person{})
// }
//
// func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
// 	json.NewEncoder(w).Encode(people)
// }
//
// func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request) {
// 	params := mux.Vars(req)
// 	var person Person
// 	_ = json.NewDecoder(req.Body).Decode(&person) // ?? what does this line do exactly?
// 	person.ID = params["id"]
// 	people = append(people, person)
// 	json.NewEncoder(w).Encode(people)
// }
//
// func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request) {
// 	params := mux.Vars(req)
// 	for index, item := range people {
// 		if item.ID == params["id"] {
// 			// the following line slices right before current index, and right after, skipping the current index.
// 			// Read up on the ... in Go.
// 			people = append(people[:index], people[index+1:]...)
// 			break
// 		}
// 	}
// 	json.NewEncoder(w).Encode(people)
// }

func main() {
	router := mux.NewRouter()
	colors = append(colors, Color{ID: "1", Name: "blue"})
	colors = append(colors, Color{ID: "2", Name: "red"})
	colors = append(colors, Color{ID: "3", Name: "yellow"})
	// people = append(people, Person{ID: "1", Firstname: "Guin", Lastname: "Awesome", Address: &Address{City: "Seattle", State: "WA"}})
	// people = append(people, Person{ID: "2", Firstname: "Brendan", Lastname: "Batman"})

	// router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	// router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")
	// router.HandleFunc("/people/{id}", CreatePersonEndpoint).Methods("POST")
	// router.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("DELETE")
	router.HandleFunc("/", index).Methods("GET")
	router.HandleFunc("/favoriteColor", favoriteColor).Methods("GET")

	log.Fatal(http.ListenAndServe(":12345", router))

}
