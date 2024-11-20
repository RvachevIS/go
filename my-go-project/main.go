package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/items", getItems)
	http.HandleFunc("/items/create", createItem)
	http.HandleFunc("/items/update", updateItem)
	http.HandleFunc("/items/delete", deleteItem)
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
