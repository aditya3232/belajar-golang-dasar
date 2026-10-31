package belajar_golang_dasar

import (
	"fmt"
	"testing"
)

// 1. Fungsi Cleanup / Logging untuk Defer
func logging() {
	fmt.Println("Selesai memanggil function (Logging via Defer)")
}

func runApplication(value int) {
	// `defer` menjadwalkan eksekusi `logging()` tepat sebelum `runApplication` selesai
	defer logging()

	fmt.Println("Aplikasi sedang berjalan...")
	result := 100 / value
	fmt.Println("Hasil perhitungan =", result)
}

// 2. Demonstrasi Urutan Eksekusi Multiple Defer (LIFO: Last In, First Out)
func testMultipleDefer() {
	fmt.Println("--- Mulai Multiple Defer ---")
	defer fmt.Println("Defer 1 (Didaftarkan pertama)")
	defer fmt.Println("Defer 2 (Didaftarkan kedua)")
	defer fmt.Println("Defer 3 (Didaftarkan ketiga)")
	fmt.Println("--- Akhir Badan Fungsi ---")
	// Urutan eksekusi output: Defer 3 -> Defer 2 -> Defer 1
}

// 3. Fungsi Recovery untuk Menangkap Panic
func endAppWithRecover() {
	// `recover()` wajib dipanggil di dalam fungsi defer
	message := recover()
	if message != nil {
		fmt.Println("TERJADI PANIC DITANGKAP OLEH RECOVER:", message)
		fmt.Println("Aplikasi tetap aman dan tidak mengalami crash!")
	}
	fmt.Println("Aplikasi selesai dihentikan secara rapi.")
}

// 4. Fungsi yang Memperagakan Panic
func connectDatabase(isValid bool) {
	// Daftarkan fungsi recovery sebelum potensi panic terjadi
	defer endAppWithRecover()

	if !isValid {
		// Memicu panic: menghentikan alur normal dan melempar pesan fatal
		panic("GAGAL TERHUBUNG KE DATABASE!")
	}

	fmt.Println("Berhasil terhubung ke Database!")
}

func TestDeferPanicRecover(t *testing.T) {
	// ==========================================
	// 1. Pengujian Defer Dasar
	// ==========================================
	fmt.Println("=== 1. Pengujian Defer ===")
	runApplication(10)

	// ==========================================
	// 2. Pengujian Urutan Multiple Defer (LIFO)
	// ==========================================
	fmt.Println("\n=== 2. Pengujian Multiple Defer (LIFO) ===")
	testMultipleDefer()

	// ==========================================
	// 3. Pengujian Panic & Recover
	// ==========================================
	fmt.Println("\n=== 3. Pengujian Panic & Recover ===")
	// Kasus Normal (tanpa panic)
	connectDatabase(true)

	fmt.Println()

	// Kasus Panic (panic terjadi, ditangkap oleh recover, program tidak mati)
	connectDatabase(false)

	fmt.Println("\nEksekusi kode setelah connectDatabase(false) tetap berjalan lancar!")
}

/*
==================================================
PENJELASAN DEFER, PANIC, DAN RECOVER DI GO
==================================================

1. DEFER
   - Kata kunci `defer` digunakan untuk menjadwalkan pemanggilan sebuah fungsi
     agar dieksekusi TEPAT SEBELUM fungsi pembungkusnya selesai (return).
   - Defer AKAN SELALU DIJALANKAN, baik fungsi selesai secara normal maupun berhenti karena panic.
   - Sifat Multiple Defer: Mengikuti prinsip LIFO (Last In, First Out) atau sistem Stack.
     Fungsi defer yang didaftarkan paling akhir akan dieksekusi paling pertama saat fungsi selesai.
   - KASUS PENGGUNAAN UTAMA:
     - Resource cleanup (menutup file: `file.Close()`, menutup koneksi database, membuka kunci mutex: `mu.Unlock()`).

2. PANIC
   - `panic` adalah fungsi bawaan yang digunakan untuk menghentikan alur program secara tiba-tiba
     ketika terjadi kesalahan fatal yang tidak dapat dipulihkan (unrecoverable error).
   - Saat panic dipanggil:
     a. Eksekusi alur biasa langsung terhenti.
     b. Seluruh fungsi `defer` yang sudah terdaftar tetap akan dieksekusi secara berurutan (stack unwinding).
     c. Jika tidak ditangkap oleh `recover`, program akan CRASH dan mencetak pesan panic beserta stack trace.
   - CATATAN BEST PRACTICE DI GO:
     - Jangan gunakan `panic` untuk error bisnis sehari-hari (seperti input salah atau data tidak ketemu).
       Gunakan mekanisme pengembalian `error` biasa.
     - `panic` biasanya hanya digunakan saat inisialisasi aplikasi (misal konfigurasi env penting tidak ditemukan).

3. RECOVER
   - `recover` adalah fungsi bawaan untuk menangkap data panic dan memulihkan alur program agar tidak crash.
   - SYARAT MUTLAK RECOVER:
     `recover()` HANYA BERFUNGSI JIKA DIPANGGIL DI DALAM FUNGSI `defer`!
     Jika dipanggil di luar defer, `recover()` hanya akan mengembalikan `nil` tanpa efek apapun.
   - Nilai balik `recover()` adalah nilai/pesan yang dioper ke `panic(...)`.
     Jika tidak ada panic yang terjadi, `recover()` mengembalikan `nil`.

CARA MENJALANKAN FILE INI:
  $ go test -v -run TestDeferPanicRecover
==================================================
*/
