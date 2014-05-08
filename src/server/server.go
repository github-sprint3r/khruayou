package main

import (
	//	"menu_category"
	//	"menu_item"
	//	"memu_recommend"
	"menu_search"
	"net/http"
)

func makeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r)
	}
}

func main() {
	http.Handle("/assets/", http.FileServer(http.Dir(".")))
	// http.Handle("/assets", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))

	//	http.Handle("/assets/", http.FileServer(http.Dir(".")))
	//http.HandleFunc("/menu_category", makeHandler(menu_category.HomeHandler))
	//http.HandleFunc("/menu_item", makeHandler(menu_item.HomeHandler))
	//http.HandleFunc("/memu_recommend", makeHandler(memu_recommend.HomeHandler))

	// ค้นหารายการอาหาร เว็บ
	http.HandleFunc("/menu_search", makeHandler(menu_search.HomeHandler))
	// ค้นหารายการอาหาร JSON API
	http.HandleFunc("/api", makeHandler(menu_search.SearchAPIHandle))

	http.ListenAndServe(":8080", nil)
}
