DSN="host=localhost port=5432 user=postgres password=password dbname=concurrency sslmode=disable timezone=UTC connect_timeout=5"
BINARY_NAME=myapp.exe
REDIS="127.0.0.1:6379"

## build: builds all binaries
build:
	@echo Building...
	@go build -o ${BINARY_NAME} ./cmd/web
	@echo Build completed!

run: build
	@echo Starting...
	powershell -noprofile -command $$env:DSN='${DSN}'; $$env:REDIS='${REDIS}'; ./${BINARY_NAME}
	@echo Started!

start: run
