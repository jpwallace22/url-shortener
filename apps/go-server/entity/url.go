package entity

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"gorm.io/gorm"
)

type Url struct {
	ID        uint      `gorm:"primaryKey" json:"url_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ShortUrl  string    `json:"short_url"`
	Url       string    `json:"url"`
	QrCode    string    `json:"qr_code"`
	Clicks    int       `gorm:"default:0" json:"clicks"`
	Hash      *string   `json:"hash"`
}

type UrlModel struct {
	DB *gorm.DB
}

func (m *UrlModel) AddUrl(w http.ResponseWriter, r *http.Request) {
	url := Url{ShortUrl: "test", Url: "https://www.google.com"}
	m.DB.Create(&url)
	render.JSON(w, r, url)
}

func (m *UrlModel) UrlRoutes(r chi.Router) {
	r.Post("/", m.AddUrl)
}
