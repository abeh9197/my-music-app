# my-music-app
## バックエンドのAPI起動
cd back
go run cmd/my-music-app/main.go

# upload request sample.
curl -F 'file=@README.md' http://localhost:8080/upload

psql -h {hostname} -p 5432 -U {user_name} -d {db_name}
