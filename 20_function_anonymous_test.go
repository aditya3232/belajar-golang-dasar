package belajar_golang_dasar

import (
	"fmt"
	"strings"
	"testing"
)

// Tipe data fungsi untuk blacklist user
type Blacklist func(string) bool

// Fungsi yang memproses registrasi user dengan validasi blacklist
func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Printf("User '%s' diblokir! Tidak dapat mendaftar.\n", name)
	} else {
		fmt.Printf("Selamat datang, '%s'! Registrasi berhasil.\n", name)
	}
}

// Fungsi pembantu untuk memfilter slice angka
func filterAngka(numbers []int, condition func(int) bool) []int {
	var result []int
	for _, num := range numbers {
		if condition(num) {
			result = append(result, num)
		}
	}
	return result
}

func TestFunctionAnonymous(t *testing.T) {
	// 1. Anonymous Function Disimpan ke dalam Variabel
	// Fungsi dibuat tanpa nama dan langsung dimasukkan ke variabel `blacklistRoot`
	blacklistRoot := func(name string) bool {
		return strings.ToLower(name) == "root" || strings.ToLower(name) == "admin"
	}

	registerUser("Adit", blacklistRoot)  // Berhasil
	registerUser("admin", blacklistRoot) // Diblokir

	// 2. Anonymous Function Langsung Ditulis sebagai Argumen (Inline Function)
	// Kita tidak perlu membuat fungsi terpisah jika aturannya hanya dipakai sekali di sini
	fmt.Println("--- Anonymous Function Inline ---")
	registerUser("anjing", func(name string) bool {
		return strings.ToLower(name) == "anjing" || strings.ToLower(name) == "kasar"
	})

	registerUser("Budi Santoso", func(name string) bool {
		return strings.ToLower(name) == "anjing" || strings.ToLower(name) == "kasar"
	})

	// 3. Contoh Nyata: Filter Koleksi Menggunakan Inline Anonymous Function
	// Sangat ringkas karena tidak perlu mendefinisikan fungsi isOdd / isMultipleOfFive di luar
	angkaList := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 15, 20}

	// a. Mengambil angka kelipatan 5 langsung dengan anonymous function
	kelipatan5 := filterAngka(angkaList, func(n int) bool {
		return n%5 == 0
	})
	fmt.Println("Kelipatan 5 =", kelipatan5) // [5 10 15 20]

	// b. Mengambil angka ganjil yang lebih besar dari 5
	ganjilBesar := filterAngka(angkaList, func(n int) bool {
		return n%2 != 0 && n > 5
	})
	fmt.Println("Ganjil > 5  =", ganjilBesar) // [7 9 15]

	// 4. Immediately Invoked Function Expression (IIFE)
	// Anonymous function yang didefinisikan sekaligus LANGSUNG DIJALANKAN seketika
	// Perhatikan tanda kurung `(10, 20)` di baris penutup kurung kurawal
	hasilHitung := func(a, b int) int {
		return (a + b) * 2
	}(10, 20) // Langsung dieksekusi dengan argumen 10 dan 20

	fmt.Println("Hasil IIFE (10 + 20) * 2 =", hasilHitung) // 60
}

/*
==================================================
PENJELASAN ANONYMOUS FUNCTION DI GO
==================================================

1. DEFINISI
   - Anonymous Function (dikenal juga sebagai Function Literal atau Lambda) adalah fungsi yang dibuat tanpa nama.
   - Di Go, fungsi dapat dibuat di mana saja, termasuk di dalam blok kode fungsi lain.

2. KEUNTUNGAN PENGGUNAAN
   a. Mengurangi Scope Pollution:
      - Jika sebuah logika fungsi hanya dibutuhkan di satu tempat tertentu (sekali pakai / throwaway),
        kita tidak perlu mengotori package level dengan membuat fungsi bernama terpisah.
   b. Lebih Ringkas & Deklaratif:
      - Logika fungsi langsung terlihat di tempat fungsi tersebut dipanggil (inline),
        seperti saat melakukan filter data pada contoh di atas.
   c. Sangat Sering Digunakan untuk Concurrency (Goroutine):
      - Di Go, sering kali goroutine dijalankan dengan anonymous function:
        `go func() { ... }()`

3. CARA PENGGUNAAN
   a. Disimpan ke variabel:
      `myFunc := func() { ... }`
      `myFunc()`

   b. Sebagai argumen langsung (Inline callback):
      `doSomething(func(x int) bool { return x > 0 })`

   c. Langsung dieksekusi (IIFE):
      `func() { ... }()`

CARA MENJALANKAN FILE INI:
  $ go test -v -run TestFunctionAnonymous
==================================================
*/

