package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// @Component{
// 	templateUrl: "./index.html"
// }

type BIO struct {
	NAME   string `jason:"name"`
	AGE    string `jason:"age"`
	GENDER string `jason:"gender"`
	PHONE  string `jason:"number"`
}

var bio []BIO

func getbio(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bio)
	//http.ServeFile(w, r, "template/index.html")
}

func getdata(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //get param
	//loops throgh book and find id
	for _, item := range bio {
		if item.NAME == params["name"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(bio)

}

func createBio(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var Bio BIO
	_ = json.NewDecoder(r.Body).Decode(&Bio)
	Bio.NAME = strconv.Itoa(rand.Intn(1000))
	bio = append(bio, Bio)
	json.NewEncoder(w).Encode(Bio)
}

// delete bio
func deletedata(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range bio {
		if item.NAME == params["name"] {
			bio = append(bio[:index], bio[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(bio)
}

func main() {

	//init router
	r := mux.NewRouter()

	//Mock data - implement data base
	bio = append(bio, BIO{NAME: "naveena", AGE: "22", GENDER: "female", PHONE: "123456"})
	bio = append(bio, BIO{NAME: "babu", AGE: "24", GENDER: "male", PHONE: "147852"})
	//route handlers / endpoints
	r.HandleFunc("/api/bio", getbio).Methods("GET")
	r.HandleFunc("/api/bio/{name}", getdata).Methods("GET")
	r.HandleFunc("/api/bio", createBio).Methods("POST")
	r.HandleFunc("/api/bio/{name}", deletedata).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8001", r))
}
