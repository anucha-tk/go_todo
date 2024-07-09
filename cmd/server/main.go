package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	port := ":3000"
	flag.StringVar(&port, "port", port, "port to listen on")
	flag.Parse()

	router := routes()

	server := &http.Server{
		Addr:              "localhost" + port,
		Handler:           router,
		ReadHeaderTimeout: time.Second * 5,
	}

	fmt.Printf("Listening on http://localhost%s\n", port)

	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
