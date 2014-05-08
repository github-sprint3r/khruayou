package searchnotfound

import (
	"fmt"
	"net/http"
)

<<<<<<< HEAD
func main() {
	http.HandleFunc("/", addHanler)
	http.ListenAndServe(":8080", nil)
=======
func HomeHandler() {
    http.HandleFunc("/", addHanler)
    http.ListenAndServe(":8080", nil)
>>>>>>> b40a424dd6ebf4db89c3d51ee4319a5896976180
}

func addHanler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page!")
}
