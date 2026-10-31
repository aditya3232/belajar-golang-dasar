package belajar_golang_dasar

import (
	"fmt"
	"testing"
)

// Fungsi pembantu 1: Menerima string dan mengembalikan string
func getGoodBye(name string) string {
	return "Good Bye " + name
}

// Fungsi pembantu 2 & 3: Operasi matematika dengan tipe yang sama: func(int, int) int
func tambah(a, b int) int {
	return a + b
}

func kali(a, b int) int {
	return a * b
}

func TestFunctionAsValue(t *testing.T) {
	// 1. Menyimpan Fungsi ke dalam Variabel
	// Perhatikan: JANGAN sertakan tanda kurung `()`, cukup tuliskan nama fungsinya
	// Tipe data dari variabel `sayGoodbye` adalah: func(string) string
	sayGoodbye := getGoodBye

	// Memanggil fungsi melalui variabel
	pesan := sayGoodbye("Adit")
	fmt.Println("Hasil sayGoodbye =", pesan)

	// Deklarasi eksplisit dengan tipe data fungsi
	var farewell func(string) string = getGoodBye
	fmt.Println("Hasil farewell   =", farewell("Budi"))

	// 2. Mengubah Nilai Variabel Berisi Fungsi
	// Variabel bertipe fungsi dapat diarahkan ke fungsi lain asalkan signature-nya identik
	var operasiMatematika func(int, int) int

	operasiMatematika = tambah
	fmt.Println("Hasil tambah (10, 5) =", operasiMatematika(10, 5)) // 15

	operasiMatematika = kali
	fmt.Println("Hasil kali   (10, 5) =", operasiMatematika(10, 5)) // 50

	// 3. Contoh Praktis: Menyimpan Fungsi ke dalam Map (Dispatch Table / Router)
	// Berguna untuk memilih operasi berdasarkan input teks tanpa if-else bertingkat
	kalkulator := map[string]func(int, int) int{
		"penjumlahan": tambah,
		"perkalian":   kali,
	}

	operasi := "perkalian"
	fn, exists := kalkulator[operasi]
	if exists {
		hasil := fn(6, 7)
		fmt.Printf("Kalkulator '%s' dari 6 dan 7 = %d\n", operasi, hasil) // 42
	}
}

/*
==================================================
PENJELASAN FUNCTION AS VALUE DI GO
==================================================

1. FIRST-CLASS CITIZEN (FIRST-CLASS FUNCTIONS)
   - Di Go, fungsi diperlakukan sebagai "first-class citizen".
   - Artinya, fungsi dianggap setara dengan tipe data lainnya (seperti `int`, `string`, `bool`).
   - Fungsi dapat:
     a. Disimpan ke dalam variabel.
     b. Disimpan di dalam slice, array, atau map.
     c. Dijadikan argumen/parameter ke fungsi lain (*Function as Parameter*).
     d. Dikembalikan dari fungsi lain (*Returning Functions*).

2. TIPE DATA FUNGSI (FUNCTION SIGNATURE)
   - Tipe data dari fungsi ditentukan oleh kombinasi parameter dan nilai baliknya (signature).
   - Contoh:
     - `func(string) string`   : Fungsi yang menerima 1 string dan mengembalikan 1 string.
     - `func(int, int) int`    : Fungsi yang menerima 2 int dan mengembalikan 1 int.
     - `func()`                : Fungsi tanpa parameter dan tanpa return value.

3. HATI-HATI: NAMA FUNGSI VS PEMANGGILAN FUNGSI
   - `fn := getGoodBye`   -> Menyimpan referensi fungsinya ke variabel `fn` (Function as Value).
   - `fn := getGoodBye()` -> MENGEKSEKUSI fungsi dan menyimpan nilai return-nya ke variabel `fn`.

4. MANFAAT PRAKTIS
   - Membuat *dispatch table* (seperti router/kalkulator map di atas).
   - Menjadi dasar untuk materi pemrograman fungsional berikutnya di Go:
     *Function as Parameter* (Higher-Order Functions), *Anonymous Functions*, dan *Closures*.

CARA MENJALANKAN FILE INI:
  $ go test -v -run TestFunctionAsValue
==================================================
*/

