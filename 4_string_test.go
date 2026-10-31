package belajar_golang_dasar

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	// 1. String Dasar
	fmt.Println("Nama Depan  :", "Adit")
	fmt.Println("Nama Belakang:", "Kurniawan")

	// 2. Fungsi bawaan: len() untuk menghitung jumlah byte/panjang string
	fmt.Println("Panjang string 'Adit' =", len("Adit")) // 4

	// 3. Mengakses byte karakter berdasarkan index (dimulai dari 0)
	// Hasil akses index adalah representasi byte (uint8/ASCII), bukan string langsung
	fmt.Println("Karakter byte ke-0 dari 'Adit'  =", "Adit"[0])         // 65 (ASCII untuk 'A')
	fmt.Println("Karakter string ke-0 dari 'Adit' =", string("Adit"[0])) // 'A' (dikonversi kembali ke string)

	// 4. Penggabungan string (concatenation)
	fullName := "Adit" + " " + "Kurniawan"
	fmt.Println("Gabungan string               =", fullName)

	// 5. Raw String Literal (menggunakan backtick `)
	// Mempertahankan baris baru (newline) dan kutip tanpa perlu escape character
	rawString := `Ini adalah baris satu
Ini adalah baris dua
Bisa memuat "tanda kutip" dengan bebas`
	fmt.Println(rawString)
}

/*
==================================================
PENJELASAN TIPE DATA STRING DI GO
==================================================

1. DEFINISI
   - Tipe data string di Go direpresentasikan dengan kata kunci: `string`
   - String adalah kumpulan karakter yang di-encode menggunakan format UTF-8.

2. PENULISAN STRING
   a. Double Quotes ("..."):
      - String standar satu baris.
      - Mendukung escape sequence seperti `\n` (baris baru), `\t` (tab), `\"` (kutip).
   b. Backticks (`...`):
      - Disebut "Raw String Literal".
      - Mendukung banyak baris (multiline) dan mempertahankan spasi/newline persis seperti yang diketik.
      - Tidak memproses escape sequence.

3. DEFAULT VALUE (ZERO VALUE)
   - Zero value untuk tipe `string` adalah string kosong: `""`.

4. FUNGSI & OPERASI BAWAAN STRING
   a. len("string"):
      - Menghitung jumlah BYTE dalam string (bukan selalu jumlah karakter).
      - Untuk karakter ASCII umum (A-Z, 0-9), 1 karakter = 1 byte.
      - Karakter non-ASCII (seperti huruf beraksen, kanji, atau emoji) menggunakan lebih dari 1 byte (contoh UTF-8).

   b. Mengambil Karakter ("string"[index]):
      - Indeks dimulai dari 0.
      - Mengembalikan nilai BYTE (`uint8`) dari karakter tersebut, bukan string.
      - Untuk menampilkannya kembali sebagai karakter string, gunakan casting: `string(str[index])`.

   c. Penggabungan String:
      - Menggunakan operator `+` (contoh: "Halo" + " " + "Dunia").

5. SIFAT IMMUTABLE (TIDAK DAPAT DIUBAH)
   - String di Go bersifat immutable (tidak bisa dimutasi).
   - Kode seperti `str[0] = 'B'` akan menghasilkan compile error.
   - Jika ingin memodifikasi isi string, buat string baru atau gunakan slice of byte `[]byte(str)` / `strings.Builder`.

CARA MENJALANKAN FILE INI:
  $ go test -v -run TestString
==================================================
*/

