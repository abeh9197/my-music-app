package handlers

import (
	"fmt"
	"io"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	// 10MBのファイルサイズ制限を設定
	const MaxUploadSize = 10 << 20 // 10 MB
	r.Body = http.MaxBytesReader(w, r.Body, MaxUploadSize)
	if err := r.ParseMultipartForm(MaxUploadSize); err != nil {
		http.Error(w, fmt.Sprintf("The uploaded file is too big: %s. Maximum size is %d MB.", err.Error(), MaxUploadSize>>20), http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, fmt.Sprintf("Could not retrieve the file: %s", err.Error()), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// MIMEタイプの検証
	buf := make([]byte, 512)
	if _, err := file.Read(buf); err != nil {
		http.Error(w, fmt.Sprintf("Error reading file: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	filetype := mime.TypeByExtension(filepath.Ext(handler.Filename))
	if filetype == "" {
		filetype = http.DetectContentType(buf)
	}
	if !strings.HasPrefix(filetype, "audio/") {
		http.Error(w, "The provided file format is not allowed. Please upload an audio file.", http.StatusBadRequest)
		return
	}

	file.Seek(0, io.SeekStart)

	dir := "./uploads/"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, os.ModePerm)
	}
	filePath := filepath.Join(dir, handler.Filename)

	dst, err := os.Create(filePath)
	if err != nil {
		http.Error(w, fmt.Sprintf("Could not save file: %s", err.Error()), http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	if _, err = io.Copy(dst, file); err != nil {
		http.Error(w, fmt.Sprintf("Error saving file: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "File uploaded successfully: %s", filePath)
}
