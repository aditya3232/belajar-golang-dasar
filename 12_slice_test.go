package belajar_golang_dasar

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	// 1. Membuat Slice dari Array yang Sudah Ada (Slicing: array[low:high])
	// low: indeks awal (inklusif), high: indeks akhir (eksklusif)
	months := [...]string{
		"Januari", "Februari", "Maret", "April", "Mei", "Juni",
		"Juli", "Agustus", "September", "Oktober", "November", "Desember",
	}

	slice1 := months[4:7] // Mengambil index 4, 5, 6 ("Mei", "Juni", "Juli")
	fmt.Println("slice1 =", slice1)
	fmt.Println("len(slice1) =", len(slice1)) // 3 (Mei, Juni, Juli)
	fmt.Println("cap(slice1) =", cap(slice1)) // 8 (dari Mei index 4 sampai Desember index 11)

	// Variasi slicing lainnya:
	// months[2:] -> dari index 2 sampai akhir
	// months[:5] -> dari awal sampai index sebelum 5
	// months[:]  -> seluruh isi array dijadikan slice

	// 2. Sifat Reference pada Slice
	// Mengubah elemen pada slice AKAN mengubah underlying array aslinya
	slice1[0] = "Mei Diubah"
	fmt.Println("Underlying array months[4] ikut berubah =", months[4])

	// 3. Fungsi bawaan: append(slice, data...)
	// Menambahkan elemen ke slice.
	// Jika kapasitas (cap) masih cukup, akan mengubah underlying array.
	// Jika kapasitas (cap) PENUH, Go otomatis membuat underlying array BARU.
	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	daySlice := days[5:] // ["Sabtu", "Minggu"], len=2, cap=2 (karena berada di ujung array)

	// Kapasitas penuh (cap=2, len=2), append akan membuat array baru di balik layar
	daySliceBaru := append(daySlice, "Hari Libur Baru")
	daySliceBaru[0] = "Sabtu Diubah"

	fmt.Println("daySliceBaru =", daySliceBaru)
	fmt.Println("days asli tidak terpengaruh =", days) // Array days tetap "Sabtu", tidak berubah

	// 4. Membuat Slice dengan fungsi make(tipe, len, cap)
	// Sangat aman dan efisien jika kita sudah tahu perkiraan ukuran data
	newSlice := make([]int, 2, 5)
	newSlice[0] = 10
	newSlice[1] = 20

	fmt.Println("newSlice     =", newSlice)
	fmt.Println("len(newSlice)=", len(newSlice)) // 2
	fmt.Println("cap(newSlice)=", cap(newSlice)) // 5

	newSlice2 := append(newSlice, 30)
	fmt.Println("newSlice2 setelah append =", newSlice2) // [10 20 30]

	// 5. Fungsi bawaan: copy(destination, source)
	// Menyalin elemen dari source ke destination.
	// Ukuran slice destination harus sudah disiapkan terlebih dahulu.
	fromSlice := []string{"Adit", "Kurniawan"}
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)
	fmt.Println("Hasil copy toSlice =", toSlice)

	// 6. Perbedaan Penulisan Array vs Slice (Slice Literal)
	contohArray := [...]int{1, 2, 3} // Ada titik tiga [...] atau angka [3] -> ARRAY
	contohSlice := []int{1, 2, 3}    // Kotak kurung kosong [] -> SLICE

	fmt.Println("Tipe Array =", contohArray)
	fmt.Println("Tipe Slice =", contohSlice)
}

/*
==================================================
PENJELASAN TIPE DATA SLICE DI GO
==================================================

1. DEFINISI & STRUKTUR SLICE
   - Slice adalah potongan dari sebuah array (underlying array).
   - Slice berukuran dinamis dan merupakan tipe data koleksi yang PALING SERING digunakan di Go.
   - Di balik layar, sebuah slice terdiri dari 3 komponen:
     a. Pointer  : Penunjuk ke elemen pertama dari underlying array yang menjadi awal slice.
     b. Length   : Panjang elemen yang ada di dalam slice saat ini (`len()`).
     c. Capacity : Kapasitas maksimum dari elemen awal slice hingga ujung underlying array (`cap()`).

2. PEMBUATAN SLICE
   a. Memotong Array / Slice:
      - `data[low:high]` -> dari indeks low s.d. sebelum high.
      - `data[low:]`     -> dari indeks low s.d. akhir.
      - `data[:high]`    -> dari indeks awal (0) s.d. sebelum high.
      - `data[:]`        -> seluruh data.

   b. Menggunakan fungsi `make`:
      - `make([]TipeData, panjang, kapasitas)`
      - Membantu alokasi memori yang efisien tanpa harus membuat array manual terlebih dahulu.

   c. Slice Literal:
      - `slice := []int{1, 2, 3}` (tanpa angka di dalam kurung siku `[]`).

3. OPERASI PENTING
   a. `append(slice, value1, value2, ...)`:
      - Menambahkan elemen ke slice dan mengembalikan slice baru.
      - PENTING:
        - Jika `len < cap`: append akan menimpa data pada underlying array yang sama.
        - Jika `len == cap`: append otomatis membuat underlying array baru dengan kapasitas lebih besar,
          sehingga modifikasi pada slice baru tidak akan merusak array lama.

   b. `copy(destination, source)`:
      - Menyalin data dari `source` ke `destination`.
      - Perhatikan: `len` dari destination harus mencukupi, jika destination memiliki len=0 maka tidak ada data yang tersalin.

4. PERBEDAAN KRUSIAL ARRAY VS SLICE
   - Array : `[5]int{...}` atau `[...]int{...}` (ukuran tetap, value type).
   - Slice : `[]int{...}` (ukuran dinamis, reference type ke underlying array).

CARA MENJALANKAN FILE INI:
  $ go test -v -run TestSlice
==================================================
*/

