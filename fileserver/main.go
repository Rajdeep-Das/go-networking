package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("FileServer Stared on Port : 8080")
	http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))
}
