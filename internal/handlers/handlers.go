package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

const limit = 10 << 20 // 10 MB

func Index(w http.ResponseWriter, r *http.Request) {
	indexPage := "index.html"
	f, err := os.ReadFile(indexPage)
	if err != nil {
		http.Error(w, "error reading index.html", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Content-Length", fmt.Sprint(len(f)))
	w.WriteHeader(http.StatusOK)
	w.Write(f)
}

func Upload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "not post method", http.StatusMethodNotAllowed)
		return
	}

	r.ParseMultipartForm(limit)
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "error getting file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	data := make([]byte, handler.Size)
	_, err = file.Read(data)
	if err != nil {
		http.Error(w, "error reading file", http.StatusInternalServerError)
		return
	}

	convertedData := service.Convert(string(data))

	uploadsFolder := "uploads"
	err = os.Mkdir(uploadsFolder, 0755)
	if err != nil && !errors.Is(err, os.ErrExist) {
		http.Error(w, "error creating dir", http.StatusInternalServerError)
		return
	}

	root, err := os.OpenRoot(uploadsFolder)
	if err != nil {
		http.Error(w, "error opening dir", http.StatusInternalServerError)
		return
	}
	defer root.Close()

	newFileName := time.Now().UTC().Format("01.02.2006 15:04:05") + filepath.Ext(handler.Filename)
	dst, err := root.Create(newFileName)
	if err != nil {
		http.Error(w, "error creating file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	_, err = dst.WriteString(convertedData)
	if err != nil {
		http.Error(w, "error saving file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Content-Length", fmt.Sprint(len(convertedData)))
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, convertedData)
}
