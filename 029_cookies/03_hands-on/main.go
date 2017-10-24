package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", count)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	err := http.ListenAndServe(":8082", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func count(w http.ResponseWriter, r *http.Request) {
	//
	// get current count var in cookie
	// cur count is not set
	access, err := r.Cookie("access")
	if err != nil {
		log.Println("count is not set", err)
		access = &http.Cookie{Name: "access", Value: "0"}
		// 加上这句以后则cookie设置不成功！！！！
		// fmt.Fprintln(w, "new value is set")
	}
	//
	//
	// set count = count+1 in new cookie
	accessCountValue, err := strconv.Atoi(access.Value)
	if err != nil {
		log.Fatalln("conv to int err")
	}
	accessCountValue++
	access.Value = strconv.Itoa(accessCountValue)
	http.SetCookie(w, access)
	fmt.Fprintln(w, "new value is set", access.Value)
	// io.WriteString(w, access.Value)
}
