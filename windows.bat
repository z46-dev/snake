@echo off

echo Building wasm...

go env -w GOOS=js GOARCH=wasm
go build -o ./public/main.wasm ./public/main.go

echo Starting main.go

go env -w GOOS=windows GOARCH=amd64
go run main.go