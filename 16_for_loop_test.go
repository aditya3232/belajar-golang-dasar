package belajar_golang_dasar

import (
	"fmt"
	"testing"
)

func TestForLoop(t *testing.T) {
	// 1. For Loop Standar (Init statement; Kondisi; Post statement)
	fmt.Println("--- 1. For Loop Standar ---")
	for i := 1; i <= 3; i++ {
		fmt.Println("Perulangan ke:", i)
	}

	// 2. For Loop Hanya dengan Kondisi (Berfungsi seperti 'while' di bahasa lain)
	fmt.Println("--- 2. For seperti 'while' ---")
	counter := 1
	for counter <= 3 {
		fmt.Println("Counter:", counter)
		counter++
	}

	// 3. For Tanpa Kondisi (Infinite Loop) dengan Kata Kunci `break`
	fmt.Println("--- 3. For Infinite Loop + break ---")
	angka := 1
	for {
		fmt.Println("Angka:", angka)
		if angka == 3 {
			break // Berhenti dan keluar dari perulangan
		}
		angka++
	}

	// 4. Penggunaan Kata Kunci `continue`
	// `continue` akan melewati (skip) sisa kode di iterasi saat ini dan melompat ke iterasi berikutnya
	fmt.Println("--- 4. For dengan continue (Cetak ganjil saja) ---")
	for i := 1; i <= 6; i++ {
		if i%2 == 0 {
			continue // Lewati angka genap
		}
		fmt.Println("Angka ganjil:", i)
	}

	// 5. For Range pada Slice / Array
	fmt.Println("--- 5. For Range pada Slice ---")
	names := []string{"Adit", "Budi", "Joko"}
	for index, name := range names {
		fmt.Printf("Index %d: %s\n", index, name)
	}

	// Mengabaikan index menggunakan blank identifier (_)
	for _, name := range names {
		fmt.Println("Nama:", name)
	}

	// 6. For Range pada Map
	fmt.Println("--- 6. For Range pada Map ---")
	person := map[string]string{
		"nama":   "Adit",
		"kota":   "Jakarta",
		"profesi": "Developer",
	}
	for key, value := range person {
		fmt.Printf("%s: %s\n", key, value)
	}

	// 7. For Range Integer (Fitur modern Go 1.22+)
	// Sintaks ringkas untuk iterasi dari 0 hingga n-1
	fmt.Println("--- 7. For Range Integer (Go 1.22+) ---")
	for i := range 3 { // Menghasilkan 0, 1, 2
		fmt.Println("Range i:", i)
	}

	// 8. Labelled Break (Keluar dari Nested Loop)
	fmt.Println("--- 8. Labelled Break ---")
OuterLoop:
	for baris := 1; baris <= 3; baris++ {
		for kolom := 1; kolom <= 3; kolom++ {
			if baris == 2 && kolom == 2 {
				fmt.Println("Stop di baris 2 kolom 2 menggunakan Labelled Break!")
				break OuterLoop // Langsung keluar dari OuterLoop, bukan hanya loop kolom
			}
			fmt.Printf("[%d, %d] ", baris, kolom)
		}
		fmt.Println()
	}
}

/*
==================================================
PENJELASAN FOR LOOP DI GO
==================================================

1. SATU-SATUNYA KATA KUNCI PERULANGAN
   - Di Go, TIDAK ADA kata kunci `while` atau `do-while`.
   - SEMUA bentuk perulangan di Go dibuat menggunakan satu kata kunci saja: `for`.

2. VARIASI BENTUK FOR LOOP
   a. Standar (3 komponen):
      `for init; kondisi; post { ... }`
      - `init` : Dijalankan sekali di awal (biasanya deklarasi counter: `i := 0`).
      - `kondisi` : Dievaluasi setiap awal putaran; jika true loop lanjut, jika false loop berhenti.
      - `post` : Dijalankan di akhir setiap putaran (biasanya increment: `i++`).

   b. For Kondisi Saja (Pengganti `while`):
      `for kondisi { ... }`
      - Mengeksekusi blok selama kondisi bernilai `true`.

   c. For Tanpa Kondisi (Infinite Loop):
      `for { ... }`
      - Berjalan selamanya hingga dihentikan secara eksplisit oleh `break`, `return`, atau `panic`.

3. KATA KUNCI KONTROL
   - `break`    : Menghentikan perulangan secara langsung dan keluar dari blok loop.
   - `continue` : Mengabaikan sisa instruksi di iterasi saat ini dan langsung melompat ke iterasi berikutnya.

4. FOR RANGE
   Cara paling idiomatik di Go untuk mengiterasi data koleksi:
   - Slice / Array : `for index, value := range collection`
   - Map           : `for key, value := range mapData`
   - String        : `for byteIndex, runeChar := range text`
   - Integer (1.22+): `for i := range N` (iterasi dari 0 sampai N-1)
   - Gunakan blank identifier `_` jika tidak membutuhkan index/key atau value-nya.

5. LABELLED BREAK & CONTINUE
   - Di perulangan bertingkat (nested loop), perintah `break` biasa hanya akan menghentikan perulangan terdalam.
   - Dengan memberi label di perulangan luar (`NamaLabel: for ...`), kita bisa memanggil `break NamaLabel`
     untuk menghentikan seluruh tingkatan perulangan sekaligus.

CARA MENJALANKAN FILE INI:
  $ go test -v -run TestForLoop
==================================================
*/

