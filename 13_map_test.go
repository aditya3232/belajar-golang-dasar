package belajar_golang_dasar

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	// 1. Membuat Map Menggunakan Map Literal
	// Format: map[TipeKey]TipeValue{...}
	person := map[string]string{
		"name":    "Adit",
		"address": "Jakarta",
	}

	fmt.Println("Person map       =", person)
	fmt.Println("Nama             =", person["name"])
	fmt.Println("Alamat           =", person["address"])

	// 2. Menambah & Mengubah Data pada Map
	person["title"] = "Software Engineer" // Menambah key baru
	person["address"] = "Bandung"          // Mengubah nilai key yang sudah ada
	fmt.Println("Setelah diubah   =", person)

	// 3. Mengakses Key yang Tidak Ada (Mengembalikan Zero Value)
	// Jika key tidak ditemukan, Go TIDAK melempar error/exception, melainkan mengembalikan zero value
	fmt.Println("Key 'telepon'    =", person["telepon"]) // Output string kosong: ""

	// 4. Mengecek Keberadaan Key (Idiom "Comma Ok")
	// val, exists := map[key]
	// `exists` bertipe boolean (true jika key ada, false jika tidak ada)
	val, exists := person["title"]
	if exists {
		fmt.Println("Key 'title' ditemukan dengan nilai:", val)
	} else {
		fmt.Println("Key 'title' tidak ditemukan")
	}

	_, existsTelepon := person["telepon"]
	fmt.Println("Apakah key 'telepon' ada? =", existsTelepon) // false

	// 5. Menghapus Key Menggunakan fungsi bawaan delete(map, key)
	delete(person, "title")
	fmt.Println("Setelah delete 'title' =", person)

	// 6. Membuat Map Menggunakan fungsi make(map[K]V)
	book := make(map[string]string)
	book["title"] = "Belajar Go-Lang"
	book["author"] = "Eko Kurniawan"
	book["ups"] = "Salah Masuk"

	fmt.Println("Book map   =", book)
	fmt.Println("Jumlah key =", len(book)) // 3

	delete(book, "ups")
	fmt.Println("Jumlah key setelah delete =", len(book)) // 2

	// 7. Iterasi Map Menggunakan for range
	// CATATAN: Urutan iterasi map di Go TIDAK TERJAMIN (acak / non-deterministic)
	fmt.Println("--- Iterasi Map person ---")
	for key, value := range person {
		fmt.Printf("Key: %-8s | Value: %s\n", key, value)
	}

	// 8. PERINGATAN: Nil Map vs Empty Map
	// var nilMap map[string]string  -> Bernilai nil. Menulis nilMap["a"] = "b" akan PANIC!
	// emptyMap := make(map[string]string) -> Siap digunakan untuk menulis data.
}

/*
==================================================
PENJELASAN TIPE DATA MAP DI GO
==================================================

1. DEFINISI
   - Map adalah tipe data koleksi di Go yang menyimpan data dalam bentuk pasangan kunci dan nilai (Key-Value Pair).
   - Format: `map[TipeKey]TipeValue`
   - Setiap elemen diakses melalui kata kuncinya (Key), bukan nomor indeks (seperti Array/Slice).
   - Key dalam map bersifat unik (tidak boleh ada duplikasi key).

2. SYARAT TIPE DATA KEY
   - Tipe data yang digunakan sebagai Key WAJIB mendukung operasi perbandingan kesetaraan (`==` / comparable).
   - Tipe yang VALID sebagai key: `string`, `int`, `float`, `bool`, `pointer`, `channel`, `interface`, dan `struct` (yang seluruh field-nya comparable).
   - Tipe yang INVALID sebagai key: `slice`, `map`, dan `func` (akan menghasilkan compile error).

3. FUNGSI & OPERASI BAWAAN PADA MAP
   - `len(map)`            : Mengambil jumlah pasangan key-value dalam map.
   - `map[key] = value`    : Menambah data baru atau memperbarui data yang sudah ada.
   - `map[key]`            : Mengambil nilai berdasarkan key (jika tidak ada, mengembalikan zero value).
   - `delete(map, key)`    : Menghapus pasangan key-value (aman dijalankan meski key tidak ada).
   - `value, ok := map[key]`: Idiom "comma ok" untuk memastikan apakah key benar-benar ada di map atau sekadar zero value.

4. SIFAT-SIFAT PENTING MAP
   a. Urutan Tidak Terjamin (Unordered):
      - Go sengaja mengacak urutan iterasi map untuk mencegah developer bergantung pada keteraturan susunan data.
   b. Reference Type:
      - Map adalah tipe referensi (pointer ke struktur hmap runtime di balik layar).
      - Mengirimkan map ke fungsi atau menyalin ke variabel lain tidak menduplikasi data, melainkan merujuk ke data map yang sama.
   c. Hati-hati dengan Nil Map:
      - `var m map[string]string` menghasilkan nil map.
      - Membaca nil map aman (mengembalikan zero value), tetapi MENULIS ke nil map akan menyebabkan RUNTIME PANIC: assignment to entry in nil map.
      - Selalu inisialisasi map dengan `make(map[K]V)` atau literal `{}` sebelum menulis ke dalamnya.

CARA MENJALANKAN FILE INI:
  $ go test -v -run TestMap
==================================================
*/

