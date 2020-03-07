package render

import (
	"net/http"
)

// Render http reponse render.
type Render interface {
	// Render render it to http response writer.
	Render(http.ResponseWriter) error
	// WriteContentType write content-type to http response writer.
	WriteContentType(w http.ResponseWriter)
}

var (
	_ Render = JSON{}
)

func writeContentType(w http.ResponseWriter, value []string) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = value
	}
}
