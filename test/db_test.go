package main

import (
    "database/sql"
    "fmt"
    "github.com/joho/godotenv"
    "log"
    "os"
    _ "github.com/lib/pq"
    "time"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require",
        os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // データベース接続のテスト
    err = db.Ping()
    if err != nil {
        log.Fatal("Cannot connect to the database:", err)
    }

    // テストデータの挿入
    _, err = db.Exec("INSERT INTO files (filename, filepath, upload_time) VALUES ($1, $2, $3)",
        "testfile.txt", "/testpath", time.Now())
    if err != nil {
        log.Fatal("Insertion error:", err)
    }

    // データの取得
    rows, err := db.Query("SELECT filename, filepath, upload_time FROM files")
    if err != nil {
        log.Fatal("Query error:", err)
    }
    defer rows.Close()

    for rows.Next() {
        var filename, filepath string
        var uploadTime time.Time
        err = rows.Scan(&filename, &filepath, &uploadTime)
        if err != nil {
            log.Fatal("Scan error:", err)
        }
        fmt.Println("Filename:", filename, "Filepath:", filepath, "Upload time:", uploadTime)
    }
}
