package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	_ "github.com/joho/godotenv/autoload"
	"github.com/jpwallace22/link-shortener/api"
	"github.com/jpwallace22/link-shortener/db"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	db, err := db.Init()
	if err != nil {
		log.Fatal("Could not connect to database")
	}

	// singleton validator because it caches results
	validator := validator.New(validator.WithRequiredStructEnabled())
	ctx := api.Context{
		Urls: &api.UrlModel{
			DB:        db,
			Validator: validator,
		},
	}

	router := api.BuildRouter(&ctx)

	port := fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))
	h2s := &http2.Server{}
	server := &http.Server{
		Addr:    port,
		Handler: h2c.NewHandler(router, h2s),
	}

	fmt.Printf("Server is running on http://localhost%s\n", port)
	log.Fatal(server.ListenAndServe())
}
