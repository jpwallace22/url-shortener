package api

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type responseWriterInterceptor struct {
	http.ResponseWriter
	body              bytes.Buffer
	blacklistedFields []string
}

/*
Creates a middleware that removes specified JSON fields from the response body.
The middleware intercepts the HTTP response and manipulates the JSON payload to exclude any fields listed
in the `blacklistedFields` parameter. This is useful for filtering out sensitive or unnecessary data from
API responses before they reach the client.

Parameters:
- blacklistedFields []string: A list of JSON field names to be removed from the response.

Returns:
- A middleware function that wraps an http.Handler to filter the response.
*/
func StripJSONFieldsMiddleware(blacklistedFields []string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			wrappedWriter := &responseWriterInterceptor{ResponseWriter: w, blacklistedFields: blacklistedFields}
			next.ServeHTTP(wrappedWriter, r)
			wrappedWriter.stripResponse()
		})
	}
}

func (w *responseWriterInterceptor) Write(data []byte) (int, error) {
	return w.body.Write(data)
}
func (w *responseWriterInterceptor) stripResponse() {
	var data interface{} // Use interface{} to accommodate arrays at the top level as well.
	if err := json.Unmarshal(w.body.Bytes(), &data); err != nil {
		w.ResponseWriter.Write(w.body.Bytes())
		return
	}

	stripFields(data, w.blacklistedFields)

	modifiedData, _ := json.Marshal(data)
	w.ResponseWriter.Write(modifiedData)
}

// stripFields recursively removes specified fields from the data.
func stripFields(data interface{}, blacklistedFields []string) {
	switch d := data.(type) {
	case map[string]interface{}:
		for _, field := range blacklistedFields {
			delete(d, field)
		}
		for _, value := range d {
			stripFields(value, blacklistedFields)
		}

	case []interface{}:
		for _, value := range d {
			stripFields(value, blacklistedFields)
		}
	}
}
