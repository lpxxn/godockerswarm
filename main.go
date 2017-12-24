package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		hostName, _ := os.Hostname()
		fmt.Fprintf(w, "HostName: %s", hostName)
	})

	log.Fatal(http.ListenAndServe(":8000", nil))
}
