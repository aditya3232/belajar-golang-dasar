package belajar_golang_dasar

import (
	"fmt"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	fmt.Println("Hello World")
}

/*
==================================================
PANDUAN EKSEKUSI & BUILD DI GO
==================================================

1. MENJALANKAN TEST
   - Menjalankan test ini secara spesifik:
     $ go test -v -run TestHelloWorld

   - Menjalankan semua test di direktori:
     $ go test -v ./...

2. MENGENAL CARA KERJA PROGRAM GO BIASA (Jika Menggunakan package main):
   - Menjalankan langsung (Development):
     $ go run <nama_file>.go

   - Membangun executable binary (Production):
     $ go build -o myapp <nama_file>.go

3. CROSS-COMPILATION
   - Environment variable GOOS dan GOARCH:
     $ GOOS=windows GOARCH=amd64 go build -o app.exe main.go
     $ GOOS=linux GOARCH=amd64 go build -o app-linux main.go
     $ GOOS=darwin GOARCH=arm64 go build -o app-mac main.go
==================================================
*/

