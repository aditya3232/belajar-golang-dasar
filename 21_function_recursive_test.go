package belajar_golang_dasar

import (
	"fmt"
	"strings"
	"testing"
)

// 1. Pendekatan Iteratif Biasa (Menggunakan For Loop)
func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

// 2. Pendekatan Rekursif (Memanggil fungsi itu sendiri)
// Contoh: 5! = 5 * 4 * 3 * 2 * 1
func factorialRecursive(value int) int {
	// Base Case: Kondisi berhenti agar tidak terjadi perulangan tanpa batas
	if value <= 1 {
		return 1
	}

	// Recursive Step: Memanggil diri sendiri dengan argumen yang semakin mengecil
	return value * factorialRecursive(value-1)
}

// 3. Rekursif: Deret Fibonacci (0, 1, 1, 2, 3, 5, 8, 13, ...)
func fibonacci(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// 4. Contoh Nyata: Menampilkan Struktur Kategori Bertingkat (Tree Traversal)
type Category struct {
	Name          string
	SubCategories []Category
}

func printCategoryTree(category Category, level int) {
	indent := strings.Repeat("  ", level)
	fmt.Printf("%s- %s\n", indent, category.Name)

	// Base case terjadi secara alami ketika SubCategories kosong
	for _, sub := range category.SubCategories {
		printCategoryTree(sub, level+1)
	}
}

func TestFunctionRecursive(t *testing.T) {
	// 1. Perbandingan Faktorial Loop vs Rekursif
	angka := 5
	hasilLoop := factorialLoop(angka)
	hasilRecursive := factorialRecursive(angka)

	fmt.Println("Faktorial Loop (5!)      =", hasilLoop)      // 120
	fmt.Println("Faktorial Rekursif (5!)  =", hasilRecursive) // 120

	// 2. Menguji Fibonacci
	fmt.Println("Fibonacci ke-7           =", fibonacci(7))   // 13

	// 3. Contoh Nyata Traversal Struktur Data Bersarang
	fmt.Println("--- Tree Traversal Kategori Produk ---")
	menuKategori := Category{
		Name: "Elektronik",
		SubCategories: []Category{
			{
				Name: "Komputer & Laptop",
				SubCategories: []Category{
					{Name: "Laptop Gaming"},
					{Name: "Ultrabook"},
				},
			},
			{
				Name: "Handphone & Tablet",
				SubCategories: []Category{
					{Name: "Smartphone Android"},
					{Name: "iPhone / iOS"},
				},
			},
		},
	}

	printCategoryTree(menuKategori, 0)
}

/*
==================================================
PENJELASAN FUNCTION RECURSIVE DI GO
==================================================

1. DEFINISI
   - Function Recursive (fungsi rekursif) adalah fungsi yang memanggil dirinya sendiri di dalam badannya.
   - Sangat elegan untuk menyelesaikan masalah yang dapat dipecah menjadi sub-masalah identik yang lebih kecil.

2. DUA KOMPONEN WAJIB FUNGSI REKURSIF
   a. Base Case (Kondisi Berhenti):
      - Syarat di mana rekursi harus berhenti dan mengembalikan nilai akhir tanpa memanggil dirinya lagi.
      - TANPA Base Case, program akan terus memanggil dirinya tanpa henti hingga memicu error "Stack Overflow".
   b. Recursive Step:
      - Logika di mana fungsi memanggil dirinya sendiri, dengan nilai parameter yang terus bergerak mendekati Base Case.

3. RESIKO: STACK OVERFLOW
   - Setiap kali sebuah fungsi dipanggil, Go mengalokasikan bingkai memori baru di "Call Stack".
   - Jika fungsi rekursif terlalu dalam (misal memanggil hingga ratusan ribu / jutaan kali) atau lupa memberi Base Case,
     memori stack akan habis dan aplikasi mengalami CRASH karena runtime error:
     "fatal error: stack overflow"

4. KAPAN SEBAIKNYA MENGGUNAKAN REKURSIF?
   - Kapan TIDAK disarankan:
     Operasi matematika sederhana seperti faktorial atau iterasi biasa sebenarnya lebih cepat dan hemat memori
     jika menggunakan `for loop` biasa (iteratif).
   - Kapan SANGAT disarankan:
     Untuk memproses data hierarki berulang atau struktur pohon (*tree/graph*), seperti:
     - Struktur folder dan file sistem direktori.
     - Kategori & sub-kategori produk e-commerce bertingkat (seperti contoh di atas).
     - Parser dokumen (HTML DOM, XML, atau JSON tree).

CARA MENJALANKAN FILE INI:
  $ go test -v -run TestFunctionRecursive
==================================================
*/

