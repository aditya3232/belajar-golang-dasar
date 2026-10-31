package belajar_golang_dasar

import (
	"fmt"
	"testing"
)

func TestNumber(t *testing.T) {
	// 1. Bilangan Bulat (Integer)
	fmt.Println("Satu =", 1)
	fmt.Println("Dua =", 2)

	// 2. Bilangan Desimal (Floating Point)
	fmt.Println("Tiga koma lima =", 3.5)

	// 3. Format Penulisan Angka Lainnya (Number Literal)
	fmt.Println("Biner (0b1010)   =", 0b1010)    // 10
	fmt.Println("Oktal (0o12)     =", 0o12)      // 10
	fmt.Println("Heksa (0xFF)     =", 0xFF)      // 255
	fmt.Println("Underscore       =", 1_000_000) // 1000000 (pemisah agar mudah dibaca)
}

/*
==================================================
PENJELASAN TIPE DATA NUMBER DI GO
==================================================

1. INTEGER (Bilangan Bulat)
   a. Signed Integer (bisa bernilai negatif dan positif):
      - int8   : -128 s.d. 127 (8-bit / 1 byte)
      - int16  : -32.768 s.d. 32.767 (16-bit / 2 bytes)
      - int32  : -2.147.483.648 s.d. 2.147.483.647 (32-bit / 4 bytes)
      - int64  : -9.223.372.036.854.775.808 s.d. 9.223.372.036.854.775.807 (64-bit / 8 bytes)

   b. Unsigned Integer (hanya bernilai positif & nol):
      - uint8  : 0 s.d. 255 (8-bit)
      - uint16 : 0 s.d. 65.535 (16-bit)
      - uint32 : 0 s.d. 4.294.967.295 (32-bit)
      - uint64 : 0 s.d. 18.446.744.073.709.551.615 (64-bit)

   c. Sesuai Arsitektur OS/Mesin:
      - int    : Minimal 32-bit (di OS 64-bit menjadi int64) -> Tipe integer default & paling umum
      - uint   : Minimal 32-bit (di OS 64-bit menjadi uint64)

   d. Tipe Data Alias:
      - byte   : Alias untuk uint8
      - rune   : Alias untuk int32 (biasanya untuk representasi karakter Unicode)

2. FLOATING POINT (Bilangan Pecahan / Desimal)
   - float32 : Presisi floating-point 32-bit
   - float64 : Presisi floating-point 64-bit (tipe default saat membuat angka pecahan)

3. COMPLEX (Bilangan Kompleks / Imajiner)
   - complex64  : Bagian riil dan imajiner masing-masing float32
   - complex128 : Bagian riil dan imajiner masing-masing float64

4. FORMAT NUMBER LITERAL (Go 1.13+)
   - Biner        : Diawali `0b` atau `0B` (contoh: 0b1010 = 10)
   - Oktal        : Diawali `0o` atau `0O` (contoh: 0o12 = 10)
   - Heksadesimal : Diawali `0x` atau `0X` (contoh: 0xFF = 255)
   - Underscore _ : Pemisah visual digit agar mudah dibaca (contoh: 1_000_000)

CARA MENJALANKAN FILE INI:
  $ go test -v -run TestNumber
==================================================
*/

