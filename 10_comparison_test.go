package belajar_golang_dasar

import (
	"fmt"
	"testing"
)

func TestComparison(t *testing.T) {
	// 1. Perbandingan Angka (Number)
	angka1 := 100
	angka2 := 200

	fmt.Println("100 > 200  =", angka1 > angka2)  // false (lebih besar)
	fmt.Println("100 < 200  =", angka1 < angka2)  // true  (lebih kecil)
	fmt.Println("100 >= 200 =", angka1 >= angka2) // false (lebih besar atau sama dengan)
	fmt.Println("100 <= 200 =", angka1 <= angka2) // true  (lebih kecil atau sama dengan)
	fmt.Println("100 == 200 =", angka1 == angka2) // false (sama dengan)
	fmt.Println("100 != 200 =", angka1 != angka2) // true  (tidak sama dengan)

	// 2. Perbandingan Teks (String)
	// Di Go, string dibandingkan secara leksikografis (urutan abjad / nilai byte ASCII)
	nama1 := "Adit"
	nama2 := "Adit"
	nama3 := "Budi"

	fmt.Println("'Adit' == 'Adit' =", nama1 == nama2) // true
	fmt.Println("'Adit' != 'Budi' =", nama1 != nama3) // true
	fmt.Println("'Adit' < 'Budi'  =", nama1 < nama3)  // true (karena byte 'A' bernilai 65 < 'B' bernilai 66)

	// 3. Perbandingan Boolean
	// Boolean hanya mendukung operator `==` dan `!=` (tidak mendukung < atau >)
	statusA := true
	statusB := false
	fmt.Println("true == false =", statusA == statusB) // false
	fmt.Println("true != false =", statusA != statusB) // true

	// 4. Strict Type Checking (Kedua operan WAJIB memiliki tipe yang sama persis)
	var nilaiInt32 int32 = 100
	var nilaiInt64 int64 = 100

	// Baris di bawah ini akan menyebabkan COMPILE ERROR jika tidak di-cast:
	// fmt.Println(nilaiInt32 == nilaiInt64) // ERROR: mismatched types int32 and int64

	// Solusi: konversi salah satu tipe data agar identik
	sama := int64(nilaiInt32) == nilaiInt64
	fmt.Println("Hasil setelah konversi tipe =", sama) // true
}

/*
==================================================
PENJELASAN OPERASI PERBANDINGAN DI GO
==================================================

1. DEFINISI
   - Operasi perbandingan digunakan untuk membandingkan dua buah nilai.
   - Hasil dari operasi perbandingan SELALU menghasilkan nilai bertipe boolean (`bool`), yaitu `true` atau `false`.

2. DAFTAR OPERATOR PERBANDINGAN
   - `>`  : Lebih dari
   - `<`  : Kurang dari
   - `>=` : Lebih dari atau sama dengan
   - `<=` : Kurang dari atau sama dengan
   - `==` : Sama dengan
   - `!=` : Tidak sama dengan

3. ATURAN KETAT TIPE DATA (TYPE SAFETY)
   - Di Go, perbandingan HANYA bisa dilakukan antara dua nilai yang memiliki TIPE DATA YANG SAMA PERSIS.
   - Anda tidak dapat membandingkan `int` dengan `int32`, `int` dengan `float64`, atau `string` dengan `int`.
   - Jika ingin membandingkan nilai yang berbeda tipe data, lakukan konversi tipe data (type casting) secara eksplisit terlebih dahulu.

4. PERBANDINGAN STRING SECARA LEKSIKOGRAFIS
   - Go mendukung operator `<`, `<=`, `>`, `>=` pada string.
   - Karakter dibandingkan huruf demi huruf berdasarkan nilai byte-nya (urutan kode ASCII/UTF-8).
   - Huruf kapital memiliki nilai byte lebih kecil daripada huruf kecil (contoh: 'A' = 65, 'a' = 97, sehingga `"Apple" < "apple"` bernilai `true`).

5. BATASAN OPERATOR PADA BOOLEAN
   - Tipe data `bool` hanya bisa dibandingkan menggunakan operator `==` dan `!=`.
   - Operator relasional seperti `>`, `<`, `>=`, `<=` tidak berlaku untuk boolean dan akan menimbulkan compile error.

CARA MENJALANKAN FILE INI:
  $ go test -v -run TestComparison
==================================================
*/

