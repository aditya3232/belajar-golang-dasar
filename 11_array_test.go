package belajar_golang_dasar

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	// 1. Deklarasi Array Manual dengan Menentukan Ukuran (Fixed Size)
	var names [3]string
	names[0] = "Adit"
	names[1] = "Kurniawan"
	names[2] = "Nugroho"

	fmt.Println("Elemen ke-0 =", names[0])
	fmt.Println("Elemen ke-1 =", names[1])
	fmt.Println("Elemen ke-2 =", names[2])
	fmt.Println("Seluruh isi names =", names)

	// 2. Deklarasi Array Langsung (Array Literal)
	values := [3]int{90, 80, 95}
	fmt.Println("Values =", values)

	// 3. Deklarasi Array Tanpa Menuliskan Jumlah Elemen ([...])
	// Ukuran array akan dihitung secara otomatis oleh compiler berdasarkan jumlah nilai yang diisi
	fruits := [...]string{"Apel", "Jeruk", "Mangga", "Pisang"}
	fmt.Println("Fruits =", fruits)
	fmt.Println("Jumlah buah (panjang array) =", len(fruits)) // 4

	// 4. Inisialisasi Elemen pada Indeks Tertentu (Keyed Elements)
	// Indeks yang tidak disebutkan otomatis bernilai zero value (0)
	numbers := [5]int{0: 10, 4: 50}
	fmt.Println("Numbers (indeks 0 dan 4 terisi) =", numbers) // [10 0 0 0 50]

	// 5. Mengubah Nilai Elemen Array
	values[0] = 100
	fmt.Println("Values setelah index 0 diubah =", values)

	// 6. Iterasi Array Menggunakan for range
	fmt.Println("--- Iterasi Array fruits ---")
	for index, fruit := range fruits {
		fmt.Printf("Index %d: %s\n", index, fruit)
	}

	// 7. Array Multidimensi (Contoh Matriks 2x3)
	matrix := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("Matrix row 0, col 2 =", matrix[0][2]) // 3

	// 8. Sifat Array di Go: VALUE TYPE (Bukan Reference)
	// Saat array di-copy ke variabel baru, seluruh datanya diduplikasi.
	// Mengubah salinan tidak akan mempengaruhi array asli.
	original := [3]int{1, 2, 3}
	copied := original
	copied[0] = 999

	fmt.Println("Original tetap =", original) // [1 2 3]
	fmt.Println("Copied berubah =", copied)   // [999 2 3]
}

/*
==================================================
PENJELASAN TIPE DATA ARRAY DI GO
==================================================

1. DEFINISI & KARAKTERISTIK ARRAY
   - Array adalah kumpulan data yang bertipe sama dan memiliki UKURAN TETAP (fixed length).
   - Kapasitas array ditentukan saat pembuatan dan TIDAK BISA bertambah atau berkurang setelah dibuat.
   - Ukuran array merupakan bagian dari tipe datanya:
     Tipe `[3]int` dan `[4]int` adalah DUA TIPE YANG BERBEDA sehingga tidak bisa saling ditugaskan.

2. METODE DEKLARASI ARRAY
   a. Menentukan ukuran di awal:
      var names [3]string
      names[0] = "Adit"

   b. Menggunakan array literal:
      values := [3]int{90, 80, 95}

   c. Menggunakan ellipsis `[...]`:
      fruits := [...]string{"Apel", "Jeruk"} // otomatis berukuran 2

   d. Menentukan indeks tertentu:
      numbers := [5]int{0: 10, 4: 50}

3. ZERO VALUE PADA ARRAY
   - Saat array dideklarasikan tanpa nilai awal, setiap elemennya otomatis diisi zero value dari tipenya:
     - `[3]int`    -> `[0, 0, 0]`
     - `[3]string` -> `["", "", ""]`
     - `[3]bool`   -> `[false, false, false]`

4. OPERASI BAWAAN PADA ARRAY
   - `len(array)` : Mengembalikan panjang/kapasitas array.
   - `array[i]`   : Mengambil elemen pada indeks ke-i (dimulai dari 0).
   - `array[i] = v`: Mengubah nilai elemen pada indeks ke-i.
   - Akses indeks di luar batas (misal index 3 pada array ukuran 3) akan menghasilkan compile error atau panic: index out of range.

5. ARRAY VS SLICE (PENTING!)
   - Array berukuran tetap dan bersifat VALUE TYPE (diduplikasi saat di-copy atau dijadikan argumen fungsi).
   - Dalam praktiknya di Go, Array jarang digunakan secara langsung; programmer Go lebih sering menggunakan SLICE
     karena Slice bersifat dinamis dan fleksibel (akan dipelajari pada materi berikutnya).

CARA MENJALANKAN FILE INI:
  $ go test -v -run TestArray
==================================================
*/

