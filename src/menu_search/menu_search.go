package menu_search

import "html"
import "strings"
import "html/template"
import "database/sql"
import "fmt"
import "net/http"
import _ "github.com/go-sql-driver/mysql"

type MenuItem struct {
	id      int64
	name_th string
	name_en string
	cat_id  int64
	price   int64
	unit    string
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("html/menu_search/menusearch.html")
	t.Execute(w, "")
}

func SearchAPIHandle(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:1q2w3e4r@tcp(119.59.97.11:3306)/KhruaYou")
	defer db.Close()

	if err != nil {
		fmt.Printf(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Printf(err.Error())
	}

	statementQuery, err := db.Prepare("SELECT * FROM menu_item")
	if err != nil {
		panic(err.Error())
	}
	defer statementQuery.Close()

	rows, err := statementQuery.Query()
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	//list := make([]*Category, 0)

	s := "["

	for rows.Next() {
		cat := new(MenuItem)
		rows.Scan(&cat.id, &cat.name_th, &cat.name_en, &cat.cat_id, &cat.price, &cat.unit)
		name_th := html.EscapeString(cat.name_th)
		name_en := html.EscapeString(strings.Trim(cat.name_en, " "))
		unit := html.EscapeString(cat.unit)
		s += fmt.Sprintf("{\"id\": \"%d\", \"name_th\":\"%s\",\"name_en\":\"%s\",\"cat_id\":\"%d\",\"price\":\"%d\",\"unit\":\"%s\"},", cat.id, name_th, name_en, cat.cat_id, cat.price, unit)
	}

	s += "]"

	s = strings.Replace(s, ",]", "]", -1)

	fmt.Fprintf(w, s)

}
