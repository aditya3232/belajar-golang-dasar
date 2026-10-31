package belajar_golang_dasar

import (
	"fmt"
	"testing"
)

func TestVariable(t *testing.T) {
	// 1. Deklarasi dengan kata kunci `var` dan tipe data eksplisit
	var name string = "Adit Kurniawan"
	fmt.Println("Nama =", name)

	// Mengubah nilai variabel
	name = "Budi Santoso"
	fmt.Println("Nama Baru =", name)

	// 2. Deklarasi `var` tanpa tipe data (Type Inference)
	// Tipe data otomatis ditentukan oleh Go berdasarkan nilai awalnya
	var age = 25
	fmt.Println("Umur =", age)

	// 3. Deklarasi variabel tanpa inisialisasi nilai awal (Zero Value)
	var isStudent bool
	var score float64
	fmt.Println("Zero value bool   =", isStudent) // false
	fmt.Println("Zero value float64=", score)     // 0

	// 4. Deklarasi singkat (Short Declaration) menggunakan `:=`
	// Catatan: Hanya bisa digunakan di dalam fungsi dan hanya untuk inisialisasi awal
	country := "Indonesia"
	fmt.Println("Negara =", country)

	// Mengubah nilai yang dibuat dengan `:=` (gunakan `=` biasa, bukan `:=` lagi)
	country = "Jepang"
	fmt.Println("Negara Baru =", country)

	// 5. Deklarasi Banyak Variabel Sekaligus (Multiple Variable)
	// a. Menggunakan blok var
	var (
		firstName = "Adit"
		lastName  = "Kurniawan"
		city      = "Jakarta"
	)
	fmt.Println("Lengkap =", firstName, lastName, "asal", city)

	// b. Satu baris
	var a, b, c int = 10, 20, 30
	x, y := "Halo", true
	fmt.Println("a, b, c =", a, b, c)
	fmt.Println("x, y    =", x, y)

	// 6. Blank Identifier (_)
	// Digunakan jika ada variabel/nilai yang tidak ingin kita pakai agar tidak kena error "unused variable"
	_ = "Data ini diabaikan"
}

/*
==================================================
PENJELASAN VARIABEL DI GO
==================================================

1. DEFINISI
   - Variabel adalah tempat untuk menyimpan data sementara di memori komputer.
   - Di Go, variabel bersifat "statically typed": tipe data suatu variabel terikat
     dan tidak bisa diubah nilainya ke tipe data lain setelah dideklarasikan.

2. METODE DEKLARASI VARIABEL
   a. Kata Kunci `var` + Tipe Data Eksplisit:
      Contoh: var name string = "Adit"
      Cocok jika ingin menegaskan tipe datanya secara jelas.

   b. Kata Kunci `var` Tanpa Tipe Data:
      Contoh: var age = 25
      Go akan otomatis mendeteksi tipe datanya (type inference).

   c. Short Declaration `:=`:
      Contoh: country := "Indonesia"
      - Singkat dan paling sering digunakan di Go untuk variabel lokal dalam fungsi.
      - Kata kunci `var` dan tipe data tidak perlu ditulis.
      - HANYA bisa digunakan di dalam fungsi (tidak bisa untuk variabel tingkat package/global).
      - Hanya untuk inisialisasi awal; untuk penugasan ulang nilai gunakan operator `=`.

   d. Multiple Declaration:
      - Menggunakan kurung:
        var (
            first = "Adit"
            last  = "Kurniawan"
        )
      - Satu baris:
        var x, y int = 1, 2
        a, b := "foo", "bar"

3. DEFAULT VALUE (ZERO VALUE)
   Jika variabel dideklarasikan tanpa nilai awal, Go akan otomatis mengisinya dengan default value:
   - Bilangan bulat (int, uint)     : 0
   - Bilangan pecahan (float)       : 0.0
   - Boolean (bool)                 : false
   - String (string)                : "" (string kosong)
   - Pointer, slice, map, channel   : nil

4. ATURAN PENTING DI GO
   a. Unused Variable (Error Kompilasi):
      - Di Go, setiap variabel lokal di dalam fungsi HARUS digunakan.
      - Jika variabel dideklarasikan tetapi tidak pernah dipakai, program TIDAK AKAN bisa di-compile.
   b. Blank Identifier (`_`):
      - Digunakan untuk menampung nilai yang sengaja tidak ingin digunakan agar tidak menyebabkan error unused.
   c. Naming Convention:
      - Gunakan camelCase untuk variabel lokal (misal: `firstName`).
      - Huruf awal kapital (PascalCase) menandakan variabel tersebut diekspor (public) ke luar package.

CARA MENJALANKAN FILE INI:
  $ go test -v -run TestVariable
==================================================
*/

