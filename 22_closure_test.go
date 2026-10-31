package belajar_golang_dasar

import (
	"fmt"
	"testing"
)

// 1. Function Factory (Membuat fungsi dengan state privat terisolasi)
// Fungsi ini mengembalikan anonymous function yang "mengingat" variabel `count`
func createCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// 2. Custom Filter Generator (Membuat validator dinamis menggunakan closure)
func makeScoreValidator(minScore int) func(int) bool {
	return func(score int) bool {
		// Mengakses `minScore` dari fungsi pembungkusnya
		return score >= minScore
	}
}

func TestClosure(t *testing.T) {
	// ==========================================
	// CONTOH 1: Closure Mengubah Variabel Scope Luar
	// ==========================================
	counter := 0
	name := "Adit"

	// Anonymous function ini membentuk closure karena mengakses & mengubah variabel di luar dirinya
	increment := func() {
		fmt.Println("Mengeksekusi increment...")
		counter++ // HATI-HATI: variabel `counter` di luar ikut termutasi!
		name = "Adit Dimodifikasi"
	}

	increment()
	increment()

	fmt.Println("Nilai counter luar sekarang =", counter) // 2
	fmt.Println("Nilai name luar sekarang    =", name)    // Adit Dimodifikasi

	// ==========================================
	// CONTOH 2: State Encapsulation (Private Counter)
	// ==========================================
	fmt.Println("\n--- State Encapsulation ---")
	counterA := createCounter()
	counterB := createCounter()

	// counterA dan counterB memiliki memori `count` masing-masing yang terpisah & privat
	fmt.Println("counterA:", counterA()) // 1
	fmt.Println("counterA:", counterA()) // 2

	fmt.Println("counterB:", counterB()) // 1 (tidak terpengaruh oleh counterA)
	fmt.Println("counterA:", counterA()) // 3

	// ==========================================
	// CONTOH 3: Kapan Pakai (Dynamic Filter)
	// ==========================================
	fmt.Println("\n--- Dynamic Filter Generator ---")
	lulusSD := makeScoreValidator(60)
	lulusKuliah := makeScoreValidator(75)

	nilaiMurid := 70
	fmt.Println("Nilai 70 lulus SD?     =", lulusSD(nilaiMurid))     // true (>= 60)
	fmt.Println("Nilai 70 lulus Kuliah? =", lulusKuliah(nilaiMurid)) // false (< 75)

	// ==========================================
	// CONTOH 4: Kapan Harus Hati-hati (Capture by Reference)
	// ==========================================
	fmt.Println("\n--- Jebakan: Capture by Reference ---")
	var operations []func()

	angka := 10
	// Closure menangkap REFERENSI (pointer) dari variabel `angka`, bukan nilainya saat itu
	operations = append(operations, func() {
		fmt.Println("Nilai angka dalam closure:", angka)
	})

	angka = 999 // Nilai di luar diubah setelah fungsi ditambahkan ke slice

	// Ketika fungsi dipanggil, ia membaca nilai terakhir dari `angka`
	operations[0]() // Mencetak 999, BUKAN 10!
}

/*
==================================================
PENJELASAN CLOSURE DI GO
==================================================

1. DEFINISI CLOSURE
   - Closure adalah kemampuan sebuah fungsi (biasanya anonymous function) untuk "mengingat"
     dan berinteraksi dengan variabel-variabel yang berada di luar lingkup tubuhnya sendiri (lexical scope).
   - Fungsi tersebut "menutup" (closes over) variabel di sekitarnya sehingga variabel tersebut tetap hidup
     di memori selama fungsi closure masih ada.

2. KAPAN SEBAIKNYA MENGGUNAKAN CLOSURE? (USE CASES)
   a. State Encapsulation (Enkapsulasi Ringan):
      - Ingin memiliki variabel privat yang hanya bisa diakses dan dimanipulasi oleh fungsi tertentu
        tanpa perlu membuat struct dan method lengkap (seperti contoh `createCounter`).
   b. Function Factory / Generator:
      - Membuat fungsi baru yang telah terkonfigurasi dengan parameter tertentu (seperti `makeScoreValidator`).
   c. Middleware & Decorator:
      - Menyelipkan konteks data luar ke dalam fungsi handler (sangat umum di web server seperti `net/http` atau `Gin`).

3. KAPAN HARUS HATI-HATI DENGAN CLOSURE? (PITFALLS & DANGERS)
   a. Efek Samping Tak Sengaja (Side Effects):
      - Jika Anda membuat anonymous function dan tidak sengaja mengubah variabel di luar fungsi,
        variabel tersebut akan berubah di seluruh program dan berpotensi menimbulkan bug yang sulit dilacak.
      - TIPS: Jika hanya ingin menggunakan nilainya tanpa mengubah variabel luar, oper variabel tersebut
        sebagai parameter argumen fungsi.

   b. Capture by Reference (Bukan Snapshot):
      - Closure mengikat variabel ke lokasinya di memori (by reference), bukan menyalin nilainya (*snapshot*).
      - Jika variabel di luar berubah sebelum closure dieksekusi, closure akan membaca nilai yang BARU (seperti di Contoh 4).

   c. Potensi Memory Leak (Memori Tertahan):
      - Variabel luar yang terikat dengan closure akan dipindahkan compiler dari Stack ke Heap (Escape Analysis).
      - Selama fungsi closure masih disimpan atau direferensikan, variabel luar tersebut TIDAK BISA dibersihkan
        oleh Garbage Collector, meskipun fungsi pembungkus aslinya sudah selesai dijalankan.

CARA MENJALANKAN FILE INI:
  $ go test -v -run TestClosure
==================================================
*/
