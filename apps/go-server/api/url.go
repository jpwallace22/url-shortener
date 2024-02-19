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
	Password  *string   `gorm:"-" json:"password" validate:"omitempty,max=72,min=8"`
	Hash      []byte    `json:"-"`
}

type UrlModel struct {
	DB        *gorm.DB
	Validator *validator.Validate
}

// POST /urls/shorten
func (m *UrlModel) shorten(w http.ResponseWriter, r *http.Request) {
	createHandler(w, r, func() *JSONError {
		var url Url
		if err := json.NewDecoder(r.Body).Decode(&url); err != nil {
			return &JSONError{Code: http.StatusBadRequest}
		}

		if err := m.Validator.Struct(url); err != nil {
			errors := TranslateErrors(err)
			return &JSONError{Code: http.StatusBadRequest, Message: errors}
		}

		if url.Password != nil {
			hash, err := bcrypt.GenerateFromPassword([]byte(*url.Password), bcrypt.MinCost)
			if err != nil {
				return &JSONError{Code: http.StatusInternalServerError}
			}
			url.Hash = hash
		}

		pngBytes, err := qrcode.Encode("https://example.org", qrcode.Medium, 256)
		if err != nil {
			return &JSONError{Code: http.StatusInternalServerError}
		}

		url.QrCode = "data:image/png;base64," + base64.StdEncoding.EncodeToString(pngBytes)
		url.ShortUrl, url.ID = generateShortURL()

		m.DB.Create(&url)
		render.JSON(w, r, url)
		return nil
	})
}

type verifyUrlRequest struct {
	ID       string `json:"url_id" validate:"required"`
	Password string `json:"password" validate:"required,max=72,min=8"`
}

// POST /urls/verify
func (m *UrlModel) verify(w http.ResponseWriter, r *http.Request) {
	createHandler(w, r, func() *JSONError {
		var url Url
		var body verifyUrlRequest
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			return &JSONError{Code: 400}
		}

		if err := m.Validator.Struct(body); err != nil {
			errors := TranslateErrors(err)
			return &JSONError{Code: 400, Message: errors}
		}

		result := m.DB.First(&url, "id = ?", body.ID)
		if result.Error != nil {
			return &JSONError{Code: 404}
		}

		if err := bcrypt.CompareHashAndPassword(url.Hash, []byte(body.Password)); err != nil {
			return &JSONError{Code: 404}
		}

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

	appUrl := os.Getenv("PUBLIC_API_URL")
	shortUrl := fmt.Sprintf("%s/%s", appUrl, path)
	return shortUrl, path
}
