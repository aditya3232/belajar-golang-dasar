package belajar_golang_dasar

import (
	"fmt"
	"strconv"
	"testing"
)

func TestConversion(t *testing.T) {
	// 1. Konversi Antar Tipe Data Number (Sintaks: TipeTujuan(nilai))
	var nilai32 int32 = 100
	var nilai64 int64 = int64(nilai32)     // Konversi aman (upcasting: 32-bit -> 64-bit)
	var nilaiFloat float64 = float64(nilai32) // Konversi int ke float

	fmt.Println("nilai32   =", nilai32)
	fmt.Println("nilai64   =", nilai64)
	fmt.Println("nilaiFloat=", nilaiFloat)

	// 2. Fenomena Integer Overflow (Hati-hati saat downcasting!)
	// int8 hanya mampu menampung -128 s.d. 127
	var angkaBesar int32 = 1000
	var angkaKecil int8 = int8(angkaBesar) // 1000 melebihi 127, nilainya akan memutar (overflow)
	fmt.Println("1000 di-convert ke int8 =", angkaKecil) // Hasilnya -24 (overflow)

	// 3. Konversi Karakter (Byte) ke String
	var name = "Adit"
	var eByte = name[0]         // tipe: byte / uint8 (nilai ASCII: 65)
	var eString = string(eByte) // konversi byte kembali ke string "A"
	fmt.Println("Karakter byte   =", eByte)
	fmt.Println("Karakter string =", eString)

	// 4. Konversi String ke Slice of Byte dan Sebaliknya
	sliceByte := []byte("Halo")
	teksKembali := string(sliceByte)
	fmt.Println("String ke []byte =", sliceByte)   // [72 97 108 111]
	fmt.Println("[]byte ke String =", teksKembali) // Halo

	// 5. Konversi String ke Angka & Sebaliknya (Menggunakan package `strconv`)
	// Di Go, kita TIDAK BISA menulis int("100") secara langsung
	// a. String ke Integer: strconv.Atoi (ASCII to Integer)
	angkaStr := "100"
	angkaInt, err := strconv.Atoi(angkaStr)
	if err == nil {
		fmt.Println("String ke Int (strconv.Atoi) =", angkaInt)
	}

	// b. Integer ke String: strconv.Itoa (Integer to ASCII)
	angka := 2026
	teksAngka := strconv.Itoa(angka)
	fmt.Println("Int ke String (strconv.Itoa) =", teksAngka)
}

/*
==================================================
PENJELASAN KONVERSI TIPE DATA DI GO
==================================================

1. SIFAT STRONGLY TYPED DI GO
   - Di Go, TIDAK ADA konversi implisit (automatic type coercion).
   - Variabel bertipe `int32` tidak bisa dijumlahkan langsung dengan `int64` tanpa melakukan konversi secara eksplisit terlebih dahulu.
   - Format umum konversi tipe dasar:
       TipeTujuan(nilaiAsal)

2. HATI-HATI DENGAN INTEGER OVERFLOW
   - Saat mengonversi tipe integer berukuran besar ke tipe yang lebih kecil (downcasting),
     misal `int32` ke `int8`, pastikan nilainya tidak melebihi kapasitas tipe tujuan.
   - Jika melebihi batas, nilai akan "overflow" (memutar kembali ke nilai minimum),
     sehingga menghasilkan angka yang salah tanpa mengeluarkan pesan error saat kompilasi.

3. KONVERSI BYTE / RUNE KE STRING
   - Karakter dalam string diakses menghasilkan `byte` (`uint8`).
   - Gunakan `string(byteValue)` untuk mengubah nilai byte kembali menjadi karakter string 1 huruf.

4. KONVERSI STRING DENGAN ANGKA (PACKAGE `strconv`)
   - `int("123")` itu ILEGAL / compile error di Go.
   - Untuk mengubah teks berisi angka menjadi bilangan beneran (dan sebaliknya), gunakan package bawaan: `strconv`.
     - `strconv.Atoi(string)` : Mengubah string angka menjadi `int` (mengembalikan 2 nilai: hasil int dan error).
     - `strconv.Itoa(int)`    : Mengubah `int` menjadi representasi teks `string`.
     - `strconv.ParseFloat()` : Mengubah teks menjadi `float64`.
     - `strconv.ParseBool()`  : Mengubah teks "true"/"false" menjadi `bool`.

CARA MENJALANKAN FILE INI:
  $ go test -v -run TestConversion
==================================================
*/

