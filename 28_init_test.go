package belajar_golang_dasar

import (
	"fmt"
	"testing"
)

// Variabel level package yang akan diinisialisasi oleh fungsi init()
var (
	databaseConnection string
	initLogs           []string
)

// 1. Fungsi init() Pertama
// Otomatis dieksekusi oleh Go runtime saat package di-load, SEBELUM fungsi test dijalankan
func init() {
	databaseConnection = "MySQL Connected (diinisialisasi oleh init)"
	initLogs = append(initLogs, "1. init() pertama dijalankan")
}

// 2. Fungsi init() Kedua
// Sebuah package / file di Go diperbolehkan memiliki LEBIH DARI SATU fungsi init()
func init() {
	initLogs = append(initLogs, "2. init() kedua dijalankan")
}

func TestInit(t *testing.T) {
	fmt.Println("=== Pengujian Fungsi init() di Go ===")

	// 1. Membuktikan variabel sudah terisi sebelum TestInit() mulai berjalan
	fmt.Println("Status Database Connection :", databaseConnection)

	// 2. Membuktikan urutan eksekusi fungsi init()
	fmt.Println("\nUrutan Log Eksekusi init():")
	for _, log := range initLogs {
		fmt.Println(" -", log)
	}

	fmt.Println("\n3. Fungsi TestInit() baru dieksekusi sekarang!")
}

/*
==================================================
PENJELASAN FUNGSI init() DI GO
==================================================

1. DEFINISI FUNGSI init()
   - `init()` adalah fungsi bawaan khusus di Go yang otomatis dieksekusi saat suatu package pertama kali dimuat (load).
   - Fungsi ini dieksekusi SEBELUM fungsi `main()` (pada aplikasi biasa) atau fungsi `TestXxx()` (pada pengujian).
   - Format: `func init()` (tidak menerima parameter dan tidak mengembalikan nilai).

2. KARAKTERISTIK UNIK FUNGSI init()
   a. Tidak Bisa Dipanggil Manual:
      - Kita tidak bisa memanggil `init()` secara eksplisit dari kode lain; pemanggilan diatur 100% oleh Go runtime.
   b. Boleh Lebih dari Satu (Multiple init):
      - Dalam satu file atau package, Anda bisa mendeklarasikan banyak fungsi `init()`.
      - Semuanya akan dieksekusi berurutan sesuai urutan deklarasi.
   c. Urutan Inisialisasi Program Go:
      1. Inisialisasi variabel package-level (misal: `var a = 10`).
      2. Eksekusi seluruh fungsi `init()` dari package dependency yang di-import.
      3. Eksekusi fungsi `init()` di package utama.
      4. Eksekusi fungsi `main()` atau unit test.

3. BLANK IDENTIFIER IMPORT (`_`)
   - Terkadang kita ingin mengimpor package HANYA untuk menjalankan fungsi `init()` yang ada di dalamnya,
     tanpa memanggil fungsi/struct apapun dari package tersebut.
   - Gunakan blank identifier (`_`):
       import _ "github.com/lib/pq" // Driver PostgreSQL
       import _ "github.com/go-sql-driver/mysql" // Driver MySQL
   - Driver database mendaftarkan dirinya ke package `database/sql` di dalam fungsi `init()` masing-masing.

4. KAPAN HARUS PAKAI & KAPAN HARUS DIHINDARI?
   - Kapan Tepat Digunakan:
     - Registrasi driver database (seperti contoh di atas).
     - Menyiapkan variabel konfigurasi default / cache awal.
   - Kapan Harus DIHINDARI (Bad Practice):
     - Hindari proses I/O berat atau koneksi jaringan lambat di dalam `init()` (seperti ping API remote).
     - Hindari logika yang berpotensi error fatal, karena `init()` tidak bisa mengembalikan `error`.
       Jika koneksi DB gagal di dalam `init()`, program hanya bisa panic, dan ini menyulitkan pembuatan unit test (sulit di-mock).

CARA MENJALANKAN FILE INI:
  $ go test -v -run TestInit
==================================================
*/
