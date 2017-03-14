package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

var (
	port    = flag.String("p", "8000", "Port for the doc server to listen on")
	docRoot = flag.String("docroot", "", "Top-level directory to serve static content from")
)

func search(w http.ResponseWriter, r *http.Request) {
	// Get query
	q := r.URL.Query().Get("q")
	fmt.Println(q)
	// Shell out
	args := strings.Split(q, " ")
	output, err := exec.Command("godoc", args...).Output()
	if err != nil {
		log.Fatal(err)
		// Pretty bad if we can't run it, but we'll check it's presence along
		// with a correct $GOPATH etc when refactoring
		os.Exit(1)
	}
	fmt.Println(string(output))
	// Write response to client
	fmt.Fprintf(w, string(output))
}

func main() {
	flag.Parse()

	if len(*docRoot) == 0 {
		fmt.Println("-docroot flag missing, cannot start server")
		os.Exit(1)
	}
	log.Println("Serving static content from:", *docRoot)

	/// Static content served via the Go fileserver
	fs := http.FileServer(http.Dir(*docRoot))
	http.Handle("/", fs)
	http.Handle("/css", fs)
	http.Handle("/scripts", fs)
	http.Handle("/images", fs)

	http.HandleFunc("/search", search)

	// Search requests handled as normal
	http.ListenAndServe(":"+*port, nil)
}
