package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/anucha-tk/go_todo/internal/database"
)

func main() {
	port := ":3000"
	flag.StringVar(&port, "port", port, "port to listen on")
	flag.Parse()

	// database
	if err := database.Init(); err != nil {
		log.Fatalf("cannot connect database: %v", err)
	}

	router := routes()

	server := &http.Server{
		Addr:              "localhost" + port,
		Handler:           router,
		ReadHeaderTimeout: time.Second * 5,
	}

	log.Printf("Listening on http://localhost%s\n", port)

	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
