module github.com/mqrc81/zeries

// +heroku goVersion go1.17
// +heroku install ./cmd/...

go 1.17

require (
	github.com/alexedwards/scs/postgresstore v0.0.0-20220216073957-c252878bcf5a
	github.com/alexedwards/scs/v2 v2.5.0
	github.com/cyruzin/golang-tmdb v1.4.3
	github.com/go-chi/chi/v5 v5.0.7
	github.com/jmoiron/sqlx v1.3.4
	github.com/joho/godotenv v1.4.0
	github.com/lib/pq v1.10.4
)

require (
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
)
