package api

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	_ "github.com/joho/godotenv/autoload"
)

type Context struct {
	Urls *UrlModel
}

func BuildRouter(ctx *Context) *chi.Mux {
	r := chi.NewRouter()
	r.Use(
		middleware.Logger,
		middleware.AllowContentType("application/json"),
		render.SetContentType(render.ContentTypeJSON),
		StripJSONFieldsMiddleware([]string{"password", "hash"}),
		middleware.CleanPath,
		middleware.RedirectSlashes,
		middleware.Recoverer,
	)
	r.Route("/url", ctx.Urls.urlRoutes)
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
}

func (m *UrlModel) redirectToUrl(w http.ResponseWriter, r *http.Request) {
	createHandler(w, r, func() *JSONError {
		var url Url
		urlId := chi.URLParam(r, "id")
		result := m.DB.First(&url, "id = ?", urlId)
		if result.Error != nil {
			return &JSONError{404, "Not Found"}
		}

		http.Redirect(w, r, url.Url, http.StatusFound)
		return nil
	})
}

func createHandler(w http.ResponseWriter, r *http.Request, f func() *JSONError) {
	if err := f(); err != nil {
		sendJSONError(w, err.Message, err.Code)
	}
}
