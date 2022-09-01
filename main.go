package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	cfg := getConf()

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Heartbeat("/ping"))
	r.Use(middleware.RequestLogger(&middleware.DefaultLogFormatter{Logger: log.New(os.Stdout, "", log.LstdFlags)}))
	r.Use(middleware.StripSlashes)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", indexRoute)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Hmmmm 🌚 Are you lost?", http.StatusMethodNotAllowed)
	})

	fmt.Println("Server is running on port " + cfg.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), r))
}
