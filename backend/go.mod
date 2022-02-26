module github.com/mqrc81/zeries

// +heroku goVersion go1.17
// +heroku install ./cmd/...

go 1.17

require (
	github.com/alexedwards/scs/postgresstore v0.0.0-20220216073957-c252878bcf5a
	github.com/alexedwards/scs/v2 v2.5.0
	github.com/cyruzin/golang-tmdb v1.4.3
	github.com/jmoiron/sqlx v1.3.4
	github.com/joho/godotenv v1.4.0
	github.com/labstack/echo/v4 v4.6.3
	github.com/lib/pq v1.10.4
	github.com/nbio/st v0.0.0-20140626010706-e9e8d9816f32
	github.com/spazzymoto/echo-scs-session v1.0.0
)

require (
	github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/labstack/gommon v0.3.1 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.1 // indirect
	golang.org/x/crypto v0.0.0-20220214200702-86341886e292 // indirect
	golang.org/x/net v0.0.0-20220225172249-27dd8689420f // indirect
	golang.org/x/sys v0.0.0-20220224120231-95c6836cb0e7 // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/time v0.0.0-20220224211638-0e9765cccd65 // indirect
)
