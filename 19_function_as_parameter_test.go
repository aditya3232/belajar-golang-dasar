package belajar_golang_dasar

import (
	"fmt"
	"strings"
	"testing"
)

// 1. Deklarasi Type untuk Fungsi (Function Type)
// Membantu mempersingkat penulisan parameter jika signature fungsinya panjang
type Filter func(string) string
type NumberPredicate func(int) bool

// 2. Fungsi yang Menerima Fungsi Lain sebagai Parameter
// Parameter `filter` bertipe Filter (yaitu func(string) string)
func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

// 3. Fungsi-fungsi yang akan dijadikan argumen filter
func spamFilter(name string) string {
	if strings.ToLower(name) == "kasar" || strings.ToLower(name) == "anjing" {
		return "..."
	}
	return name
}

// 4. Contoh Lain: Fungsi Filter Angka (Functional Programming Style)
func filterNumbers(numbers []int, predicate NumberPredicate) []int {
	var result []int
	for _, num := range numbers {
		if predicate(num) {
			result = append(result, num)
		}
	}
	return result
}

func isEven(n int) bool {
	return n%2 == 0
}

func isGreaterThanTen(n int) bool {
	return n > 10
}

func TestFunctionAsParameter(t *testing.T) {
	// 1. Memanggil fungsi dengan memasukkan fungsi `spamFilter` sebagai parameter
	sayHelloWithFilter("Adit", spamFilter)     // Hello Adit
	sayHelloWithFilter("Kasar", spamFilter)    // Hello ...
	sayHelloWithFilter("anjing", spamFilter)   // Hello ...

	// 2. Memasukkan fungsi bawaan standar (seperti strings.ToUpper)
	// strings.ToUpper memiliki signature yang cocok: func(string) string
	sayHelloWithFilter("budi", strings.ToUpper) // Hello BUDI

	// 3. Menggunakan fungsi filter koleksi angka
	angkaList := []int{2, 5, 8, 11, 14, 17, 20}

	// Filter angka genap
	genap := filterNumbers(angkaList, isEven)
	fmt.Println("Angka genap           =", genap) // [2 8 14 20]

	// Filter angka lebih besar dari 10
	lebihDari10 := filterNumbers(angkaList, isGreaterThanTen)
	fmt.Println("Angka lebih dari 10   =", lebihDari10) // [11 14 17 20]
}

/*
==================================================
PENJELASAN FUNCTION AS PARAMETER DI GO
==================================================

1. HIGHER-ORDER FUNCTIONS
   - Di Go, sebuah fungsi dapat menerima fungsi lain sebagai parameter atau mengembalikan fungsi.
   - Fungsi yang menerima fungsi lain disebut sebagai "Higher-Order Function".
   - Pola ini sangat populer untuk:
     a. Filter dan sanitasi data.
     b. Callback & event handler.
     c. Pola desain Strategy Pattern.
     d. Pembuatan Middleware (pada web framework seperti Gin, Fiber, atau net/http).

2. PENULISAN PARAMETER BERUPA FUNGSI
   a. Penulisan Langsung di Parameter:
      func sayHello(name string, filter func(string) string)
      - Terkadang terlihat panjang dan sulit dibaca jika signature-nya kompleks.

   b. Menggunakan Type Declaration (Function Type):
      type Filter func(string) string
      func sayHello(name string, filter Filter)
      - Kode menjadi jauh lebih rapi, terstruktur, dan mudah digunakan kembali (reusable).

3. KESESUAIAN SIGNATURE (SIGNATURE MATCHING)
   - Fungsi apapun dapat dijadikan parameter selama:
     - Jumlah parameter dan tipe datanya sama persis.
     - Jumlah nilai balik dan tipe datanya sama persis.
   - Itulah mengapa fungsi bawaan seperti `strings.ToUpper` bisa langsung dimasukkan sebagai argumen
     karena memiliki signature yang identik: `func(string) string`.

CARA MENJALANKAN FILE INI:
  $ go test -v -run TestFunctionAsParameter
==================================================
*/

