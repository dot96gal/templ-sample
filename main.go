package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dot96gal/templ-sample/handlers"
	"github.com/dot96gal/templ-sample/storage"
)

func main() {
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("./assets"))
	mux.Handle("GET /assets/", http.StripPrefix("/assets", fs))

	state := storage.NewState()
	handlePostPath := "/"
	indexHandler := handlers.NewIndexHandler(state, handlePostPath)
	mux.HandleFunc("GET /", indexHandler.Get)
	mux.HandleFunc(fmt.Sprintf("POST %s", handlePostPath), indexHandler.Post)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, os.Interrupt, os.Kill)
	defer stop()

	srv := &http.Server{
		Addr:    "0.0.0.0:3000",
		Handler: mux,
	}

	log.Println("Listening on :3000")
	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	<-ctx.Done()
	log.Println("Receive shutdown signal")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Println("Gracefull shutdown...")
	err := srv.Shutdown(ctx)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Complete shutdown")
}
