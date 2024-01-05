package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// ファイルの解析と最大サイズの設定
    // 10MBのファイルサイズ制限を設定
	r.ParseMultipartForm(10 << 20)

	// フォームからファイルを取得
	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()

	// 保存するディレクトリとファイル名を指定
    // ここではカレントディレクトリのuploadsフォルダに保存します
	dir := "./uploads/"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, os.ModePerm) // ディレクトリがない場合は作成
	}
	filePath := filepath.Join(dir, handler.Filename)

	// ファイルをサーバー上に保存
	dst, err := os.Create(filePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// ファイルをコピー
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "File uploaded successfully: %s", filePath)
}