module main

go 1.19

replace Hangman => ../hangman-classic

require (
	Hangman v0.0.0-00010101000000-000000000000
	github.com/gorilla/sessions v1.2.1
	github.com/satori/go.uuid v1.2.0
	golang.org/x/crypto v0.5.0
	gorm.io/driver/sqlite v1.4.4
	gorm.io/gorm v1.24.3
)

require (
	github.com/gorilla/securecookie v1.1.1 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.2.0 // indirect
	github.com/jinzhu/gorm v1.9.16 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/lib/pq v1.1.1 // indirect
	github.com/mattn/go-sqlite3 v1.14.15 // indirect
	golang.org/x/net v0.5.0 // indirect
	golang.org/x/text v0.6.0 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gorm.io/driver/postgres v1.4.6 // indirect
)
