package api

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"github.com/skip2/go-qrcode"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Url struct {
	ID        string    `gorm:"primaryKey" json:"url_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ShortUrl  string    `json:"short_url"`
	Url       string    `json:"url" validate:"required,url"`
	QrCode    string    `json:"qr_code"`
	Clicks    *int      `gorm:"default:0" json:"clicks"`
	Password  string    `gorm:"-" json:"password"`
	Hash      []byte    `json:"-"`
}

type UrlModel struct {
	DB        *gorm.DB
	Validator *validator.Validate
}

// POST /url/shorten
func (m *UrlModel) shorten(w http.ResponseWriter, r *http.Request) {
	createHandler(w, r, func() *JSONError {
		var url Url
		if err := json.NewDecoder(r.Body).Decode(&url); err != nil {
			println("FAILED")
			return &JSONError{400, "Bad Request"}
		}

		if err := m.Validator.Struct(url); err != nil {
			errors := TranslateErrors(err)
			return &JSONError{400, errors}
		}

		if url.Password != "" {
			hash, err := bcrypt.GenerateFromPassword([]byte(url.Password), 1)
			if err != nil {
				return &JSONError{500, "Internal Error"}
			}
			url.Hash = hash
		}

		pngBytes, err := qrcode.Encode("https://example.org", qrcode.Medium, 256)
		if err != nil {
			return &JSONError{500, "Internal Error"}
		}

		url.QrCode = "data:image/png;base64," + base64.StdEncoding.EncodeToString(pngBytes)
		url.ShortUrl, url.ID = generateShortURL()

		m.DB.Create(&url)
		render.JSON(w, r, url)
		return nil
	})
}

/*
Returns:
  - A pointer to the short_url using PUBLIC_APP_URL.
  - The generated path used as the url_id param.
*/
func generateShortURL() (string, string) {
	const (
		pathLength = 6
		charset    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_"
	)

	randBytes := make([]byte, 6)
	_, err := rand.Read(randBytes)
	if err != nil {
		panic(err)
	}

	path := base64.RawURLEncoding.EncodeToString(randBytes)

	path = strings.Map(func(r rune) rune {
		if strings.ContainsRune(charset, r) {
			return r
		}
		return -1
	}, path)

	if len(path) > pathLength {
		path = path[:pathLength]
	}

	appUrl := os.Getenv("PUBLIC_APP_URL")
	shortUrl := fmt.Sprintf("%s/%s", appUrl, path)
	return shortUrl, path
}
