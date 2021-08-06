package main

/*
	WEBPLANNER:
	An application with the following features:
		- Submit entry per day with goals, and events
		- Log and store every entry and organize into list
		- extract data to do math functions like mean, additon, subtract
		- modify entries
		- auto calculate transportation time b/w 2 locations
		- send reminders? option

*/

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type Page struct {
	Title string
	Body  []byte
}

func main() {
	//router
	r := newRouter()

	//r.HandleFunc("/hello", handler).Methods("GET")
	http.ListenAndServe(":8080", r)
}

func (p *Page) savePage() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &Page{Title: title, Body: body}, err
}

//e.g.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "First Server")
}

func viewHandler(w http.ResponseWriter, req *http.Request) {
	title := req.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", handler).Methods("GET")
	return r
}
