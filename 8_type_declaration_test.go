package belajar_golang_dasar

import (
	"fmt"
	"testing"
)

// Deklarasi tipe data baru (Type Definition)
// Format: type NamaTipeBaru TipeDasar
type NoKTP string
type Married bool
type Umur int

// Deklarasi Type Alias (menggunakan tanda sama dengan =)
// Format: type NamaAlias = TipeDasar
type Teks = string

func TestTypeDeclaration(t *testing.T) {
	// 1. Menggunakan Tipe Data Baru yang Sudah Dideklarasikan
	var ktpAdit NoKTP = "3201012345670001"
	var sudahMenikah Married = false
	var usia Umur = 25

	fmt.Println("No KTP  =", ktpAdit)
	fmt.Println("Menikah =", sudahMenikah)
	fmt.Println("Usia    =", usia)

	// 2. Konversi Antara Tipe Data Asal dan Tipe Data Baru
	// Karena NoKTP adalah tipe data yang BERBEDA dari string (meski tipe dasarnya string),
	// kita WAJIB melakukan konversi eksplisit jika ingin memindahkannya ke variabel string biasa.
	var ktpString string = string(ktpAdit)
	fmt.Println("KTP sebagai string biasa =", ktpString)

	var nomorContoh string = "1122334455660002"
	var ktpBaru NoKTP = NoKTP(nomorContoh)
	fmt.Println("String ke tipe NoKTP     =", ktpBaru)

	// 3. Perbedaan dengan Type Alias (`=`)
	// Type Alias hanyalah nama lain (sinonim) dari tipe yang sudah ada, sehingga kompatibel 100% tanpa konversi.
	var deskripsi Teks = "Ini menggunakan tipe alias"
	var stringBiasa string = deskripsi // Bisa langsung di-assign tanpa string(deskripsi)
	fmt.Println("Type Alias =", stringBiasa)
}

/*
==================================================
PENJELASAN TYPE DECLARATION DI GO
==================================================

1. DEFINISI
   - Type Declaration adalah kemampuan di Go untuk membuat tipe data baru dari tipe data yang sudah ada (underlying type).
   - Menggunakan kata kunci: `type`

2. MANFAAT TYPE DECLARATION
   a. Semantic Meaning (Domain Modeling):
      - Kode menjadi jauh lebih ekspresif dan mudah dibaca.
      - Contoh: Mengetahui bahwa suatu variabel adalah `NoKTP` jauh lebih informatif daripada sekadar `string`.
   b. Type Safety (Keamanan Tipe):
      - Mencegah kesalahan passing argumen secara tidak sengaja.
      - Misal jika ada `type UserID int` dan `type OrderID int`, compiler akan mencegah UserID dimasukkan ke fungsi yang membutuhkan OrderID, meskipun keduanya sama-sama berbasis integer.
   c. Dapat Diberi Method Sendiri:
      - Tipe data baru yang kita buat dapat ditambahkan fungsi/method khusus (akan dipelajari pada materi Struct & Method).

3. PERBEDAAN TYPE DEFINITION VS TYPE ALIAS
   a. Type Definition (`type NamaBaru TipeAsal`):
      - Membuat tipe data BARU yang independen.
      - Bersifat tidak kompatibel secara langsung dengan tipe asalnya (wajib eksplisit casting: `string(ktp)`).
   b. Type Alias (`type NamaAlias = TipeAsal`):
      - Menggunakan tanda sama dengan (`=`).
      - Hanyalah nama alternatif (alias/sinonim) untuk tipe data yang sama.
      - Tidak perlu konversi dan 100% kompatibel langsung.
      - Contoh bawaan Go: `byte` adalah alias dari `uint8`, dan `rune` adalah alias dari `int32`.

CARA MENJALANKAN FILE INI:
  $ go test -v -run TestTypeDeclaration
==================================================
*/

