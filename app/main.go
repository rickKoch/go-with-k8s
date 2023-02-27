package main

import (
	"fmt"
	"net/http"
	"os"
)

var hit = 0

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hit++
		fmt.Printf("there was one more hit. Hit num: %d\n", hit)
		fmt.Fprintf(w, "Welcome!!!")
	})

	// listen to port
	if err := http.ListenAndServe(":5555", nil); err != nil {
		fmt.Printf("There's been an error: %v", err)
		os.Exit(1)
	}
}
