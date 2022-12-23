package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  string `json:"age"`
}

func HandleFunc(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		w.Write([]byte(`{"method":"GET"}`))
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

func main() {
	u := []User{
		{Id: "1", Name: "peach", Age: "26"},
		{Id: "2", Name: "peachvon", Age: "29"},
	}
	b, err := json.Marshal(u)
	if err != nil {
		return
	}
	fmt.Printf("%s", b)
	http.HandleFunc("/", HandleFunc)
	log.Fatal(http.ListenAndServe(":2565", nil))
	fmt.Println("Hello !")
}
