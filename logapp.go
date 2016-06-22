package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		f, err := os.OpenFile("logapp.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
		if err != nil {
		    log.Fatal("error opening file: %v", err)
		}
		defer f.Close()
		log.SetOutput(f)
		dump, err := httputil.DumpRequest(r, true)
		log.Printf("%s\n\n", dump)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
