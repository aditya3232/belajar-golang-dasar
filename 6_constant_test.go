package belajar_golang_dasar

import (
	"fmt"
	"testing"
)

func TestConstant(t *testing.T) {
	// 1. Deklarasi Constant Tunggal
	// Nilai wajib langsung diisi dan tidak dapat diubah lagi
	const firstName string = "Adit"
	const lastName = "Kurniawan" // Type inference otomatis mendeteksi sebagai string
	const pi = 3.14

	fmt.Println("Nama Depan    =", firstName)
	fmt.Println("Nama Belakang =", lastName)
	fmt.Println("Nilai Pi      =", pi)

	// Jika baris di bawah ini diaktifkan, akan terjadi compile error:
	// firstName = "Budi" // ERROR: cannot assign to firstName

	// 2. Deklarasi Multiple Constant Sekaligus
	const (
		appName = "Belajar Golang Dasar"
		version = "1.0.0"
		author  = "Adit"
	)
	fmt.Println("App     =", appName)
	fmt.Println("Version =", version)
	fmt.Println("Author  =", author)

	// 3. Penggunaan `iota` (Auto-increment Constant / Enum)
	// iota otomatis dimulai dari 0 dan bertambah 1 untuk setiap baris berikutnya dalam blok const
	const (
		Senin = iota // 0
		Selasa       // 1
		Rabu         // 2
		Kamis        // 3
		Jumat        // 4
	)
	fmt.Println("Hari Senin =", Senin)
	fmt.Println("Hari Selasa=", Selasa)
	fmt.Println("Hari Rabu  =", Rabu)

	// 4. Constant yang tidak digunakan TIDAK menyebabkan error kompilasi
	// (Berbeda dengan variabel 'var' lokal yang wajib dipakai)
	const unusedConst = "Saya tidak dipakai tapi tidak akan error"
}

/*
==================================================
PENJELASAN CONSTANT (KONSTANTA) DI GO
==================================================

1. DEFINISI
   - Constant adalah nilai yang dideklarasikan dengan kata kunci `const`.
   - Sekali nilai diberikan ke constant, nilainya TIDAK BISA DIUBAH lagi sepanjang program berjalan (bersifat immutable / tetap).

2. ATURAN PENULISAN CONSTANT
   a. Wajib Langsung Diinisialisasi:
      - Berbeda dengan `var` yang memiliki zero value jika tidak diisi, `const` WAJIB langsung diberikan nilai saat dibuat.
      - Contoh valid  : const phi = 3.14
      - Contoh invalid: const phi // Compile error!

   b. Nilai Tidak Bisa Ditugaskan Ulang (Re-assign):
      - Percobaan mengganti nilai constant (misal `pi = 3.14159`) akan menghasilkan compile error.

   c. Unused Constant Diperbolehkan:
      - Di Go, variabel lokal (`var` atau `:=`) yang tidak dipakai akan menimbulkan error kompilasi.
      - Namun, `const` yang tidak pernah dipakai TIDAK AKAN menyebabkan error kompilasi.

3. TYPED VS UNTYPED CONSTANT
   - Typed Constant   : const x int = 10 (tipe datanya dikunci sebagai int).
   - Untyped Constant : const x = 10 (memiliki presisi tinggi dan fleksibel menyesuaikan tipe data saat digunakan dalam ekspresi).

4. KATA KUNCI `iota`
   - `iota` adalah fitur khusus di Go untuk membuat konstanta berurutan otomatis (sering dipakai sebagai enum).
   - Nilai `iota` di-reset kembali ke 0 setiap kali blok `const (...)` baru dibuka.
   - Baris berikutnya cukup menuliskan nama konstantanya saja tanpa perlu mengulang kata kunci `iota`.

CARA MENJALANKAN FILE INI:
  $ go test -v -run TestConstant
==================================================
*/

