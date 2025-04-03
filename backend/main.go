package main

import (
	"fmt"
	"net/http"
)

func main() {
	// env := os.Getenv("BECKEND_URL")
	// fmt.Println(env)

	ConnectDatabase()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Online backend")
	})

	fmt.Println("Start Server on 8080")
	http.ListenAndServe(":8080", nil)
}
