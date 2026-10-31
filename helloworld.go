package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}

/*
==================================================
PANDUAN EKSEKUSI & BUILD DI GO
==================================================

1. MENJALANKAN APP (Development)
   - Digunakan untuk mencoba kode secara langsung tanpa menghasilkan file binary fisik:
     $ go run helloworld.go

2. MELAKUKAN BUILD (Production / Deployment)
   - Mengompilasi program menjadi satu file binary executable mandiri:

     a. Build default (nama binary mengikuti nama module di go.mod):
        $ go build

     b. Build dengan nama output khusus (-o):
        $ go build -o myapp helloworld.go

     c. Menjalankan file binary yang sudah di-build:
        $ ./myapp
        (atau ./belajar-golang-dasar jika menggunakan build default)

3. CROSS-COMPILATION
   - Go memungkinkan build binary untuk OS dan arsitektur lain tanpa perlu berpindah mesin,
     cukup dengan mengatur environment variable GOOS dan GOARCH:

     a. Target Windows (64-bit):
        $ GOOS=windows GOARCH=amd64 go build -o app.exe helloworld.go

     b. Target Linux (64-bit):
        $ GOOS=linux GOARCH=amd64 go build -o app-linux helloworld.go

     c. Target macOS (Apple Silicon / M-series ARM64):
        $ GOOS=darwin GOARCH=arm64 go build -o app-mac helloworld.go
==================================================
*/
