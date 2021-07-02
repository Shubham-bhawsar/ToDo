package main

import (
	"fmt"
	"strings"
)

func main() {
	// r := router.Router()
	// // fs := http.FileServer(http.Dir("build"))
	// // http.Handle("/", fs)
	// fmt.Println("Starting server on the port 8080...")

	// log.Fatal(http.ListenAndServe(":8080", r))
	str1 := "Hii"
	str2 := "Hii"
	fmt.Println(&str1, &str2)
	fmt.Println(strings.Compare(str1, str2))
}
