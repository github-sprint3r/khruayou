package main_item

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
)

import "database/sql"
import _ "github.com/go-sql-driver/mysql"

type MenuItem struct {
	id      int64
	name_th string
	name_en string
	cat_id  int64
	price   int64
	unit    string
}

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	title := r.URL.Path[len("/backend/menu_item"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "menu_item", p)
}

func JsonHandler(w http.ResponseWriter, r *http.Request) {

	category_id := r.URL.Query().Get("id")

	if category_id == "" {
		category_id = "1"
	}
	db, err := sql.Open("mysql", "root:1q2w3e4r@tcp(119.59.97.11:3306)/KhruaYou?charset=utf8")
	defer db.Close()

	if err != nil {
		fmt.Printf(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Printf(err.Error())
	}

	statementQuery, err := db.Prepare("SELECT id,name_th,name_en, price FROM menu_item WHERE cat_id=" + category_id + " ORDER BY name_th")
	if err != nil {
		panic(err.Error())
	}
	defer statementQuery.Close()

	rows, err := statementQuery.Query()
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	/*
	   for rows.Next() {
	       var id int
	       var name_th string
	       var name_en string
	       var price int
	       rows.Scan(&id, &name_th, &name_en, &price)
	       fmt.Printf("CID=%d, Namethai=%s, Nameeng=%s, %d \n", id, name_th, name_en, price)
	   }
	*/
	s := "{ \"data\": ["

	for rows.Next() {
		var id int
		var name_th string
		var name_en string
		var price int
		rows.Scan(&id, &name_th, &name_en, &price)
		//s += fmt.Sprintf("{\"id\": \"%d\", \"name_th\":\"%s\",\"name_en\":\"%s\",\"price\":\"%d\"},",id, name_th, name_en,price);
		s += fmt.Sprintf("{\"id\": \"%d\", \"name\":\"%s <br> %s\",\"price\":\"%d\"},", id, name_th, name_en, price)
	}

	s += "]}"

	s = strings.Replace(s, ",]", "]", -1)

	fmt.Fprintf(w, s)

}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, p)
}

func main() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("/Users/seksan/Workspace/khruayou/assets"))))
	http.HandleFunc("/backend/menu_item", HomeHandler)
	http.HandleFunc("/backend/menu_item/json", JsonHandler)
	http.ListenAndServe(":8080", nil)
}
