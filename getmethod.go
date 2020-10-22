package main

import (
	"fmt"
	"net/http"
)

func FromData(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "naveena", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "template/index.html")

	case "POST":
		var nav = r.ParseForm()
		if nav != nil {
			fmt.Fprintln(w, "naveena1", nav)
			return
		}
		var name = r.FormValue("name")
		var age = r.FormValue("age")
		var gender = r.FormValue("gender")
		var place = r.FormValue("place")

		fmt.Fprintln(w, "name =", name)
		fmt.Fprintln(w, "age =", age)
		fmt.Fprintln(w, "gender =", gender)
		fmt.Fprintln(w, "place =", place)

	default:
		fmt.Fprintln(w, "naveena", "22", "female", "andhrapradesh")

	}
}
func main() {
	http.HandleFunc("/", FromData)

	fmt.Println("Server is running on port 8080....")
	http.ListenAndServe(":8080", nil)
}
