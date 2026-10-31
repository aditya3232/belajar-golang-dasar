package belajar_golang_dasar

import (
	"fmt"
	"testing"
)

func TestBoolean(t *testing.T) {
	// 1. Nilai Boolean Dasar
	fmt.Println("Benar =", true)
	fmt.Println("Salah =", false)

	// 2. Operasi Logika Boolean
	fmt.Println("true && false =", true && false) // AND: bernilai true jika kedua kondisi true
	fmt.Println("true || false =", true || false) // OR : bernilai true jika salah satu kondisi true
	fmt.Println("!true         =", !true)         // NOT: membalikkan nilai boolean (negasi)

	// 3. Hasil Operasi Perbandingan
	nilaiUjian := 85
	lulus := nilaiUjian >= 75
	fmt.Println("Apakah lulus? =", lulus)
}

/*
==================================================
PENJELASAN TIPE DATA BOOLEAN DI GO
==================================================

1. DEFINISI & NILAI
   - Tipe data boolean di Go direpresentasikan dengan kata kunci: `bool`
   - Hanya memiliki 2 kemungkinan nilai:
     - `true`  (benar)
     - `false` (salah)

2. DEFAULT VALUE (ZERO VALUE)
   - Zero value untuk tipe `bool` di Go adalah `false`.
   - Contoh: `var aktif bool` (nilai awalnya otomatis `false`).

3. OPERATOR LOGIKA BOOLEAN
   - `&&` (AND) : Menghasilkan `true` HANYA JIKA kedua operan bernilai `true`.
   - `||` (OR)  : Menghasilkan `true` JIKA SALAH SATU operan bernilai `true`.
   - `!`  (NOT) : Membalikkan nilai boolean (`!true` menjadi `false`, `!false` menjadi `true`).

4. KARAKTERISTIK PENTING (STRICT TYPING DI GO)
   - Go tidak memiliki konsep "truthy" atau "falsy" seperti JavaScript atau Python.
   - Angka `1` tidak sama dengan `true`, dan `0` tidak sama dengan `false`.
   - Boolean tidak dapat di-convert/cast secara langsung ke tipe number (int) atau sebaliknya.
   - Kondisi pada percabangan (`if`, `for`) WAJIB menghasilkan ekspresi bertipe `bool`.

CARA MENJALANKAN FILE INI:
  $ go test -v -run TestBoolean
==================================================
*/

