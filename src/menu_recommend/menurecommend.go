package menu_recommend

import (
	"html/template"
	"net/http"
)

func MenuRecommendHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("menu_recommend/menu_recommend.html")
	t.Execute(w, map[string]interface{}{
		"project_name": "Menu Recommend",
	})
}
