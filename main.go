package main

import (
	"database/sql"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
)

type Color struct {
	ID   int
	Name string
}

var colors []Color

//assign global database variable
var db *sql.DB

//run this at start??
func init() {
	var err error
	db, err = sql.Open("mysql", "root@tcp(ae80bf560ffd111e6ac3f06ff6ffae64-2060924781.us-west-2.elb.amazonaws.com:3306)/myfirstgoapp")
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	// defer db.Close()
}

func index(w http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFiles("/templates/index.html")
	if err != nil {
		log.Printf("%v", err)
	}
	t.Execute(w, colors)

}

func favoriteColor(w http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")
	log.Println(name)
	log.Println(colors)
	t, err := template.ParseFiles("/templates/favoriteColor.html")
	if err != nil {
		log.Printf("%v", err)
	}
	color := generateRandomColor(len(name))
	t.Execute(w, color)
}

func generateRandomColor(seedNum int) Color {
	// using the name length, generate a number that matches an index in colors
	seed := int64(seedNum)
	rand.Seed(time.Now().Unix() / seed)
	n := rand.Intn(len(colors))
	// match color
	for index, color := range colors {
		if n == index {
			return color
		}
	}
	return Color{}
}

func newColorForm(w http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFiles("/templates/newColorForm.html")
	if err != nil {
		log.Printf("%v", err)
	}
	t.Execute(w, "/templates/newColorForm.html")
}

func addNewColor(w http.ResponseWriter, req *http.Request) {
	log.Println("addnewcolor is running")

	// get color name
	newName := req.URL.Query().Get("newColor")

	// if field was empty
	if newName == "" {
		http.Redirect(w, req, "/newColorForm", http.StatusFound)
		return
	}

	log.Println(newName)

	_, err := db.Exec("INSERT INTO colors (id, name) VALUES(?, ?)", 0, newName)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	log.Println("color", newName, "inserted")
	//
	// rowsAffected, err := result.RowsAffected()
	// if err != nil {
	//   http.Error(w, http.StatusText(500), 500)
	//   return
	// }
	//
	// log.Println("")

	newColor := Color{Name: newName}

	colors = append(colors, newColor)

	http.Redirect(w, req, "/", http.StatusFound)
}

func deleteColor(w http.ResponseWriter, req *http.Request) {
	log.Println("delete color running")

	name := req.URL.Query().Get("deleteColor")
	log.Println(name)

	_, err := db.Exec("DELETE FROM colors WHERE name=?", name)
	if err != nil {

		http.Error(w, http.StatusText(500), 500)
		return
	}
	log.Println("color", name, "deleted")

	// remove color from colors
	for index, color := range colors {
		if color.Name == name {
			colors = append(colors[:index], colors[index+1:]...)
			break
		}
	}

	http.Redirect(w, req, "/", http.StatusFound)
}

func main() {
	log.Println("main is running")
	defer db.Close()

	// get rows
	rows, err := db.Query("SELECT * FROM colors")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	//rows is a funny database type, so we scan them. Scan will assign the values found in the row to a
	//new empty Color by field.
	for rows.Next() {
		var color = Color{}
		err := rows.Scan(&color.ID, &color.Name)
		if err != nil {
			log.Fatal(err)
		}
		// append each color ot colors slice
		colors = append(colors, color)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	//check output in terminal
	for _, color := range colors {
		log.Println(color.Name)
	}

	router := mux.NewRouter()

	router.HandleFunc("/", index).Methods("GET")
	router.HandleFunc("/favoriteColor", favoriteColor).Methods("GET")
	router.HandleFunc("/addNewColor", addNewColor).Methods("GET")
	router.HandleFunc("/newColorForm", newColorForm).Methods("GET")
	router.HandleFunc("/deleteColor", deleteColor).Methods("GET")

	log.Fatal(http.ListenAndServe(":12345", router))

}
