package menu_category

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"net/http"
)

/*func homeHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("xxx.html")
	t.Execute(w, map[string]interface{}{
		"project_name": "MY DATA",
	})
}*/

/*func addcatalogHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("Add_Catalog_New.html")
	t.Execute(w, map[string]interface{}{
		"project_name": "MY DATA",
	})
}*/

type Management struct {
	CID      int
	NameThai string
	NameEng  string
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	var CID int
	var NameThai string
	var NameEng string
	t, _ := template.ParseFiles("menu_catalog.html")

	db, err := sql.Open("mysql", "root:1q2w3e4r@tcp(119.59.97.11:3306)/KhruaYou")
	//statementQuery, err := db.Prepare("SELECT CID, Namethai, Nameeng FROM management")
	statementQuery, err := db.Prepare("SELECT id, name_th, name_en FROM category")
	if err != nil {
		panic(err.Error())
	}
	defer statementQuery.Close()

	rows, err := statementQuery.Query()
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	//var s = make([]*Management, 0, 100)
	var cat = make([]Management, 0, 100)

	for rows.Next() {
		rows.Scan(&CID, &NameThai, &NameEng)
		data := Management{CID, NameThai, NameEng}
		cat = append(cat, data)
		//fmt.Println(data.NameEng)
	}

	t.Execute(w, map[string]interface{}{
		"Results": cat,
	})
}

/*func makeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r)
	}
}*/

func main() {
	http.Handle("/assets/", http.FileServer(http.Dir(".")))
	//http.HandleFunc("/", makeHandler(homeHandler))
	http.HandleFunc("/", makeHandler(HomeHandler))
	//http.HandleFunc("/Add_Catalog_New", makeHandler(addcatalogHandler))
	http.ListenAndServe(":8080", nil)
}