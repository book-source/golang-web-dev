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
	count, err := r.Cookie("count")
	if err != nil {
		log.Println("count is not set", err)
		count = &http.Cookie{Name: "count", Value: "0"}
		// 加上这句以后则cookie设置不成功！！！！
		// fmt.Fprintln(w, "new value is set")
	}
	//
	//
	// set count = count+1 in new cookie
	countCountValue, err := strconv.Atoi(count.Value)
	if err != nil {
		log.Fatalln("conv to int err")
	}
	countCountValue++
	count.Value = strconv.Itoa(countCountValue)
	http.SetCookie(w, count)
	fmt.Fprintln(w, "new value is set", count.Value)
	// io.WriteString(w, count.Value)
}
