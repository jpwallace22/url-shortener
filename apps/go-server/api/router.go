package api

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	_ "github.com/joho/godotenv/autoload"
)

type Context struct {
	Urls *UrlModel
}

func BuildRouter(ctx *Context) *chi.Mux {
	r := chi.NewRouter()
	r.Use(
		cors.Handler(cors.Options{
			AllowedOrigins: []string{"*"},
		}),
		middleware.Logger,
		middleware.AllowContentType("application/json"),
		render.SetContentType(render.ContentTypeJSON),
		StripJSONFieldsMiddleware([]string{"password", "hash"}),
		middleware.CleanPath,
		middleware.RedirectSlashes,
		middleware.Recoverer,
	)
	r.Route("/urls", ctx.Urls.urlRoutes)
	r.Get("/{id}", ctx.Urls.redirectToUrl)

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

func (m *UrlModel) urlRoutes(r chi.Router) {
	r.Post("/shorten", m.shorten)
	r.Post("/verify", m.verify)
}

func (m *UrlModel) redirectToUrl(w http.ResponseWriter, r *http.Request) {
	createHandler(w, r, func() *JSONError {
		var url Url

		urlId := chi.URLParam(r, "id")
		result := m.DB.First(&url, "id = ?", urlId)
		if result.Error != nil {
			return &JSONError{Code: http.StatusNotFound}
		}

		if url.Hash != nil {
			verificationPage := fmt.Sprintf("%s/verify?id=%s", os.Getenv("PUBLIC_APP_URL"), url.ID)
			http.Redirect(w, r, verificationPage, http.StatusFound)
		}

		http.Redirect(w, r, url.Url, http.StatusFound)
		return nil
	})
}

func createHandler(w http.ResponseWriter, r *http.Request, f func() *JSONError) {
	if err := f(); err != nil {
		err.WriteResponse(w)
	}
}
