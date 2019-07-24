package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Student struct {
	Result  string `json:"result"`
	Name    string `json:"name"`
	RollNum int    `json:rollnum`
	Branch  string `json:"branch"`
}

var StudentData []Student

func main() {
	fmt.Println("REST 2.0")
	StudentData = []Student{
		{Result: "failed", Name: "MeghNa", RollNum: 121, Branch: "Arts"},
		{Result: "Failed", Name: "NeHa", RollNum: 100, Branch: "Science"},
		{Result: "passed", Name: "ArunPurus", RollNum: 100, Branch: "Arts"},
		{Result: "passed", Name: "Sannni", RollNum: 100, Branch: "Science"},
	}
	myrouter := mux.NewRouter()
	myrouter.HandleFunc("/studentlist", Getstudenetdata)
	myrouter.HandleFunc("/getstudent/{branch}", getSciencestudent)
	myrouter.HandleFunc("/addstudent", AddStudent).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", myrouter))
	fmt.Println("code completed")
}

func Getstudenetdata(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(StudentData); err != nil {
		panic(err)
	}
}

func getSciencestudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["branch"]
	for _, s := range StudentData {
		if s.Branch == key {
			json.NewEncoder(w).Encode(s)
		}
	}
}

func AddStudent(w http.ResponseWriter, r *http.Request) {
	resBody, _ := ioutil.ReadAll(r.Body)
	fmt.Println(resBody)
	var s Student
	json.Unmarshal(resBody, &s)
	StudentData = append(StudentData, s)
	json.NewEncoder(w).Encode(s)
}
