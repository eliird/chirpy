package main

import (
	"log"
	"net/http"
)

func handleHome(w http.ResponseWriter, req *http.Request) {
	return
}

func main() {
	const port = "8080"
	const filepathRoot = "."

	//create the mux
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(filepathRoot)))

	//create a server struct
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())
}
