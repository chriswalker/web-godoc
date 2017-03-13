package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	port    = flag.String("p", "8000", "Port for the doc server to listen on")
	docRoot = flag.String("docroot", "", "Top-level directory to serve static content from")
)

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

	// Search requests handled as normal
	http.ListenAndServe(":"+*port, nil)
}
