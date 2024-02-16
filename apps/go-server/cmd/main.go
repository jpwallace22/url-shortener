package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	_ "github.com/joho/godotenv/autoload"
	"github.com/jpwallace22/link-shortener/db"
	"github.com/jpwallace22/link-shortener/entity"
)

type Context struct {
	urls *entity.UrlModel
}

func buildRoutes(ctx Context) *chi.Mux {
	r := chi.NewRouter()
	r.Use(render.SetContentType(
		render.ContentTypeJSON),
		middleware.Logger,
		middleware.RedirectSlashes,
		middleware.Recoverer,
	)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, map[string]interface{}{
			"hello": "world",
		})
	})

	r.Route("/url", ctx.urls.UrlRoutes)

	// Walk all available routes and log them on initialization
	runner := func(method string, route string, handler http.Handler, middleware ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s", method, route)
		return nil
	}
	if err := chi.Walk(r, runner); err != nil {
		log.Panicf("Failure on walk check -- Logging err: %s\n", err.Error())
	}

	return r
}

func main() {
	db, err := db.Init()
	if err != nil {
		log.Fatal("Could not connect to database")
	}

	ctx := Context{urls: &entity.UrlModel{DB: db}}
	router := buildRoutes(ctx)

	port := fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))
	fmt.Printf("Server is running on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}
