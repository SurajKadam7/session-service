module github.com/SurajKadam7/session-service

go 1.19

require (
	github.com/SurajKadam7/msg-info-service v0.0.0-20230926170229-62ca112898f3
	github.com/go-kit/kit v0.13.0
	github.com/go-kit/log v0.2.1
	github.com/gorilla/mux v1.8.0
	github.com/redis/go-redis/v9 v9.2.0
)

require (
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/go-logfmt/logfmt v0.5.1 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.4.3 // indirect
	github.com/jackc/puddle/v2 v2.2.1 // indirect
	golang.org/x/crypto v0.9.0 // indirect
	golang.org/x/sync v0.1.0 // indirect
	golang.org/x/text v0.9.0 // indirect
)

replace github.com/suraj.kadam7/msg-info-srv => github.com/SurajKadam7/msg-info-service v0.0.0-20230926141126-bd0f52f5b15b
