package belajar_golang_dasar

import (
	"fmt"
	"testing"
)

func TestIfExpression(t *testing.T) {
	// 1. Percabangan if Sederhana
	name := "Adit"

	if name == "Adit" {
		fmt.Println("Halo Adit!")
	}

	// 2. if - else if - else
	nilai := 80

	if nilai >= 90 {
		fmt.Println("Predikat: A (Sangat Baik)")
	} else if nilai >= 75 {
		fmt.Println("Predikat: B (Baik)")
	} else if nilai >= 60 {
		fmt.Println("Predikat: C (Cukup)")
	} else {
		fmt.Println("Predikat: D (Perlu Remedial)")
	}

	// 3. If dengan Short Statement (Init Statement)
	// Sintaks: if variable := ekspresi; kondisi { ... }
	// Variabel `length` dideklarasikan langsung di header if
	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang, karakter:", length)
	} else {
		fmt.Println("Nama sudah sesuai, karakter:", length)
	}

	// 4. Pembuktian Scope Variabel pada Short Statement
	// Variabel `length` di atas HANYA hidup di dalam blok `if` dan `else` tersebut.
	// Jika kita mencoba memanggil `fmt.Println(length)` di luar blok ini, akan terjadi compile error:
	// "undefined: length"

	// 5. Contoh Penggunaan Nyata If Short Statement: Pengecekan Error / Map
	user := map[string]string{
		"username": "adit_dev",
		"role":     "admin",
	}

	if role, exists := user["role"]; exists && role == "admin" {
		fmt.Println("Akses diizinkan: Pengguna adalah Admin")
	} else {
		fmt.Println("Akses ditolak!")
	}
}

/*
==================================================
PENJELASAN IF EXPRESSION DI GO
==================================================

1. DEFINISI & SINTAKS DASAR
   - Digunakan untuk percabangan logika berdasarkan kondisi bertipe `bool`.
   - Di Go, kondisi TIDAK MEMERLUKAN tanda kurung buka-tutup `()`:
     - Benar : if nilai > 80 { ... }
     - Gaya C: if (nilai > 80) { ... } (tidak disarankan / akan otomatis dihilangkan oleh gofmt).

2. ATURAN WAJIB KURUNG KURAWAL `{}`
   - Di Go, blok kode di dalam `if`, `else if`, dan `else` WAJIB dibungkus tanda kurung kurawal `{}`
     meskipun isinya hanya terdiri dari satu baris instruksi.
   - Tidak ada sintaks satu baris tanpa kurung seperti di bahasa C/Java/JS:
     // ERROR di Go:
     if x > 0
         fmt.Println(x)

3. ATURAN PENEMPATAN `else` & `else if`
   - Kata kunci `else` atau `else if` HARUS diletakkan pada baris yang sama dengan kurung kurawal tutup `}` dari blok sebelumnya:
     // BENAR:
     } else {
     }

     // SALAH (Compile Error!):
     }
     else {
     }
   - Alasan: Go compiler otomatis menyisipkan titik koma (semicolon `;`) di akhir baris setelah `}`,
     sehingga baris `else` yang terpisah akan dianggap perintah baru tanpa pasangan `if`.

4. IF DENGAN SHORT STATEMENT (FITUR UNGGULAN GO)
   - Go mendukung pembuatan variabel lokal tepat sebelum pengecekan kondisi:
       if statementSingkat; kondisi {
           // blok kode
       }
   - Manfaat:
     a. Scope isolation: Variabel tersebut hanya hidup di dalam blok `if`, `else if`, dan `else` tersebut.
     b. Mencegah pencemaran scope fungsi (scope pollution) dengan variabel sementara yang tidak dibutuhkan lagi di baris selanjutnya.
     c. Pola yang sangat umum digunakan untuk error handling di Go:
        if err := doSomething(); err != nil {
            // tangani error
        }

5. TIDAK ADA TERNARY OPERATOR DI GO
   - Go sengaja TIDAK MENYEDIAKAN operator ternary (`kondisi ? benar : salah`).
   - Tim desainer Go berpendapat ternary sering kali membuat kode sulit dibaca saat kompleks,
     sehingga percabangan selalu diwajibkan menggunakan struktur `if-else` yang jelas dan eksplisit.

CARA MENJALANKAN FILE INI:
  $ go test -v -run TestIfExpression
==================================================
*/

