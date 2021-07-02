package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type student struct {
	Fullname      string
	CourseOfStudy string
	Email         string
	Year          int
	Prom          float32
}

//type Studets []Student

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	student1 := student{
		Fullname:      "Willam",
		CourseOfStudy: "Software Engineer",
		Email:         "williamhuchim@gmail.com",
		Year:          3,
		Prom:          9.3,
	}

	jsonStd, _ := json.Marshal(student1)

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonStd)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", HomeHandler).Methods("GET")
	http.ListenAndServe(":3000", r)
}
