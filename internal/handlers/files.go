package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi"
	"gitlab.com/upn-belajar-go/configs"
	"gitlab.com/upn-belajar-go/transport/http/middleware"
)

type FileHandler struct {
	Config *configs.Config
}

func ProvideFileHandler(conf *configs.Config) FileHandler {
	return FileHandler{
		Config: conf,
	}
}

func (h *FileHandler) Router(r chi.Router, middleware *middleware.JWT) {
	r.Route("/files", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Get("/", h.ReadFile)
		})
	})
}

func (h *FileHandler) ReadFile(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Query().Get("path")
	dir := h.Config.App.File.Dir
	fileLocation := filepath.Join(dir, filename)
	img, err := os.Open(fileLocation)

	if err != nil {
		http.Error(w, "File Not Found", http.StatusInternalServerError)
	}
	defer img.Close()
	// w.Header().Set("Content-Type", "image/png") // <-- set the content-type header
	io.Copy(w, img)
}
