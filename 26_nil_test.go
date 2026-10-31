package belajar_golang_dasar

import (
	"fmt"
	"testing"
)

// 1. Fungsi yang mengembalikan `nil` jika parameter tidak valid
func newPersonMap(name string) map[string]string {
	if name == "" {
		return nil // Mengembalikan nil sebagai representasi "data kosong"
	}
	return map[string]string{
		"name": name,
	}
}

// 2. Fungsi yang mengembalikan pointer bernilai nil jika data tidak ditemukan
func findCustomerName(id int) *string {
	if id == 1 {
		name := "Adit"
		return &name
	}
	return nil // ID selain 1 dianggap tidak ditemukan
}

func TestNil(t *testing.T) {
	// ==========================================
	// 1. Nil pada Map
	// ==========================================
	fmt.Println("=== 1. Nil pada Map ===")
	dataUser1 := newPersonMap("")
	if dataUser1 == nil {
		fmt.Println("Data user 1 KOSONG (bernilai nil)")
	} else {
		fmt.Println("Data user 1 ditemukan:", dataUser1["name"])
	}

	dataUser2 := newPersonMap("Adit Kurniawan")
	if dataUser2 != nil {
		fmt.Println("Data user 2 ditemukan:", dataUser2["name"])
	}

	// ==========================================
	// 2. Nil Slice vs Empty Slice
	// ==========================================
	fmt.Println("\n=== 2. Nil Slice vs Empty Slice ===")
	var nilSlice []int      // Tidak diinisialisasi -> bernilai nil
	emptySlice := []int{}   // Diinisialisasi dengan literal kosong -> BUKAN nil

	fmt.Printf("nilSlice   -> len: %d, cap: %d, isNil: %t\n", len(nilSlice), cap(nilSlice), nilSlice == nil)     // true
	fmt.Printf("emptySlice -> len: %d, cap: %d, isNil: %t\n", len(emptySlice), cap(emptySlice), emptySlice == nil) // false
	// Keduanya sama-sama aman untuk fungsi append():
	nilSlice = append(nilSlice, 10)
	fmt.Println("nilSlice setelah append:", nilSlice)

	// ==========================================
	// 3. Nil pada Pointer & Interface
	// ==========================================
	fmt.Println("\n=== 3. Nil pada Pointer ===")
	customer := findCustomerName(99)
	if customer == nil {
		fmt.Println("Customer dengan ID 99 TIDAK DITEMUKAN (pointer nil)")
	}

	var dataKosong any = nil
	fmt.Println("Interface any bernilai nil =", dataKosong == nil) // true

	// ==========================================
	// 4. JEBAKAN KLASIK: "Typed Nil" pada Interface
	// ==========================================
	fmt.Println("\n=== 4. Jebakan Typed Nil pada Interface ===")
	// Di balik layar, interface memiliki 2 komponen: (type, value)
	// Interface bernilai nil HANYA JIKA kedua komponen bernilai nil (nil, nil)
	var ptr *int = nil      // Pointer bertipe *int bernilai nil
	var iface any = ptr     // Dimasukkan ke interface -> (type: *int, value: nil)

	// iface TIDAK SAMA DENGAN nil karena tipe datanya (*int) terisi!
	fmt.Println("Apakah ptr == nil?   =", ptr == nil)   // true
	fmt.Println("Apakah iface == nil? =", iface == nil) // FALSE! (Hati-hati dengan jebakan ini)
}

/*
==================================================
PENJELASAN NIL DI GO
==================================================

1. DEFINISI NIL
   - Di banyak bahasa lain, data kosong dikenal dengan nama `null` atau `None`.
   - Di Go, data kosong direpresentasikan dengan kata kunci: `nil`.
   - `nil` menandakan bahwa sebuah variabel belum mengarah ke memori atau data manapun.

2. TIPE DATA YANG BISA MEMILIKI NILAI NIL
   HANYA tipe data berikut yang dapat diisi `nil`:
   - `pointer`
   - `interface` (termasuk `any`)
   - `slice`
   - `map`
   - `channel`
   - `func` (fungsi)

3. TIPE DATA YANG TIDAK BISA BERNILAI NIL!
   Tipe data primitif dan struct TIDAK BISA bernilai `nil`:
   - `int`, `float`  -> Zero value-nya adalah `0`
   - `bool`          -> Zero value-nya adalah `false`
   - `string`        -> Zero value-nya adalah string kosong `""`
   - `struct`        -> Zero value-nya adalah struct kosong dengan semua field default
   Jika Anda mencoba menulis `var umur int = nil`, Go compiler akan mengeluarkan COMPILE ERROR!

4. NIL SLICE VS EMPTY SLICE
   - `var s []int`   : nil slice (tidak memakan alokasi heap memori, `s == nil`).
   - `s := []int{}`  : empty slice (pointer teralokasi, `s != nil`).
   - Keduanya sama-sama memiliki `len == 0` dan aman digunakan untuk perulangan `for range` atau `append()`.

5. JEBAKAN "TYPED NIL" PADA INTERFACE (VERY IMPORTANT)
   - Interface menyimpan dua informasi: `(tipe_data, nilai_asli)`.
   - Sebuah interface dianggap `nil` HANYA JIKA tipenya nil DAN nilainya nil `(nil, nil)`.
   - Jika Anda memasukkan pointer nil bertipe konkret (misal `*MyStruct` yang bernilai nil) ke dalam interface,
     maka interface tersebut bernilai `(*MyStruct, nil)`, yang artinya `interface != nil`.

CARA MENJALANKAN FILE INI:
  $ go test -v -run TestNil
==================================================
*/
