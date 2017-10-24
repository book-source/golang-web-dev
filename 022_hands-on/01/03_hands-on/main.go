package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "index")
}

func dog(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "dog")
	tpl, err := template.ParseFiles("dog.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	tpl.Execute(w, " dogName")

}

func me(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "larry")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal(err)
	}
}
