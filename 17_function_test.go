package belajar_golang_dasar

import (
	"fmt"
	"testing"
)

// 1. Fungsi Sederhana (Tanpa parameter & tanpa return value)
func sayHello() {
	fmt.Println("Halo Dunia!")
}

// 2. Fungsi dengan Parameter
// Jika tipe parameter berurutan sama, cukup tulis tipenya di akhir: (firstName, lastName string)
func sayHelloTo(firstName string, lastName string) {
	fmt.Printf("Halo %s %s!\n", firstName, lastName)
}

// 3. Fungsi dengan Return Value (Nilai Balik Tunggal)
func getHello(name string) string {
	return "Hello " + name
}

// 4. Fungsi Mengembalikan Banyak Nilai (Returning Multiple Values)
// Fitur khas Go: sebuah fungsi bisa mengembalikan lebih dari 1 nilai
func getFullName() (string, string) {
	return "Adit", "Kurniawan"
}

// 5. Named Return Values (Nilai Balik Bernama)
// Variabel return dideklarasikan di header fungsi dan otomatis bernilai zero value
func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Adit"
	middleName = "K"
	lastName = "Nugroho"
	return // "Naked return": otomatis mengembalikan variabel yang didefinisikan di header
}

// 6. Variadic Function (Menerima parameter dengan jumlah dinamis / Varargs)
// Parameter terakhir dengan prefiks `...Tipe` akan diperlakukan sebagai slice di dalam fungsi
func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func TestFunction(t *testing.T) {
	// 1. Memanggil fungsi biasa
	sayHello()

	// 2. Memanggil fungsi berparameter
	sayHelloTo("Adit", "Kurniawan")

	// 3. Memanggil fungsi dengan return value
	pesan := getHello("Golang")
	fmt.Println(pesan)

	// 4. Menangkap Multiple Return Values
	depan, belakang := getFullName()
	fmt.Println("Nama Lengkap:", depan, belakang)

	// Menghiraukan return value yang tidak diinginkan menggunakan blank identifier (_)
	hanyaDepan, _ := getFullName()
	fmt.Println("Hanya Nama Depan:", hanyaDepan)

	// 5. Memanggil fungsi dengan Named Return Values
	fName, mName, lName := getCompleteName()
	fmt.Printf("Named Return: %s %s %s\n", fName, mName, lName)

	// 6. Memanggil Variadic Function
	// a. Memasukkan angka satu per satu secara langsung
	total1 := sumAll(10, 20, 30, 40, 50)
	fmt.Println("Total sumAll (langsung):", total1) // 150

	// b. Memasukkan Slice yang sudah ada dengan operator slice expansion (...)
	angkaSlice := []int{10, 10, 10, 10}
	total2 := sumAll(angkaSlice...) // slice di-unpack menjadi argumen variadic
	fmt.Println("Total sumAll (dari slice):", total2) // 40
}

/*
==================================================
PENJELASAN FUNGSI (FUNCTION) DI GO
==================================================

1. DEFINISI & DEKLARASI
   - Fungsi dideklarasikan menggunakan kata kunci `func`.
   - Format umum:
       func namaFungsi(param1 Tipe, param2 Tipe) TipeKembalian {
           // body
           return nilai
       }
   - Penggabungan tipe parameter: `(a string, b string, c int)` dapat disingkat menjadi `(a, b string, c int)`.

2. MULTIPLE RETURN VALUES
   - Go tidak memiliki konsep tuple, namun fungsi di Go dapat mengembalikan lebih dari satu nilai.
   - Sangat umum digunakan di ekosistem Go untuk mengembalikan `(hasil, error)`.
   - Wajib ditampung ke semua variabel yang sesuai, atau abaikan nilai tertentu dengan blank identifier `_`.

3. NAMED RETURN VALUES
   - Variabel yang akan dikembalikan dapat diberi nama langsung di header fungsi:
     `func hitung() (luas int, keliling int)`
   - Di dalam fungsi, variabel `luas` dan `keliling` sudah otomatis tercipta (zero value).
   - Kata kunci `return` bisa ditulis tanpa argumen (*naked return*), dan Go otomatis mengembalikan nilai variabel-variabel tersebut.
   - CATATAN: Hindari naked return pada fungsi yang panjang karena dapat menurunkan keterbacaan kode (*readability*).

4. VARIADIC FUNCTION (`...Tipe`)
   - Memungkinkan fungsi menerima jumlah argumen variabel (0, 1, atau banyak).
   - Parameter variadic HANYA boleh diletakkan di posisi parameter TERAKHIR.
   - Jika sudah memiliki slice, gunakan sintaks `sliceName...` untuk meneruskannya ke fungsi variadic (*slice expansion*).

CARA MENJALANKAN FILE INI:
  $ go test -v -run TestFunction
==================================================
*/

