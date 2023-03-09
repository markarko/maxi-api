package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/markarko/maxi-api/handlers"
)

func main() {
	handlers.InitMongodb()

	l := log.New(os.Stdout, "maxi-api ", log.LstdFlags)
	ph := handlers.NewProducts(l)
	akh := handlers.NewApiKeyHandler(l)

	sm := http.NewServeMux()
	sm.HandleFunc("/products", ph.ServeHTTP)
	sm.HandleFunc("/api-key", akh.ServeHTTP)

	s := &http.Server{
		Addr:         "localhost:8080",
		Handler:      sm,
		ErrorLog:     l,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown", sig)
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
