package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("httpreq")

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	http.ListenAndServe(":8111", mux)
}
