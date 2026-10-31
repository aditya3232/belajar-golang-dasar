package belajar_golang_dasar

import (
	"fmt"
	"math"
	"testing"
)

// 1. Definisi Interface Kontrak Dasar
// Siapa pun yang memiliki method `GetName() string` otomatis dianggap sebagai `HasName`
type HasName interface {
	GetName() string
}

// Fungsi Polymorphic: Menerima tipe apa pun yang memenuhi kontrak HasName
func sayHelloToEntity(entity HasName) {
	fmt.Println("Halo,", entity.GetName())
}

// Struct 1: Person
type Person struct {
	Name string
}

// Person mengimplementasikan HasName secara implisit (tanpa keyword `implements`)
func (p Person) GetName() string {
	return p.Name
}

// Struct 2: Animal
type Animal struct {
	Name string
}

// Animal juga mengimplementasikan HasName secara implisit
func (a Animal) GetName() string {
	return "Hewan: " + a.Name
}

// 2. Contoh Nyata: Interface Hitung Bangun Datar
type Hitung interface {
	Luas() float64
	Keliling() float64
}

type PersegiPanjang struct {
	Panjang, Lebar float64
}

func (p PersegiPanjang) Luas() float64 {
	return p.Panjang * p.Lebar
}

func (p PersegiPanjang) Keliling() float64 {
	return 2 * (p.Panjang + p.Lebar)
}

type Lingkaran struct {
	Radius float64
}

func (l Lingkaran) Luas() float64 {
	return math.Pi * math.Pow(l.Radius, 2)
}

func (l Lingkaran) Keliling() float64 {
	return 2 * math.Pi * l.Radius
}

// Fungsi pembantu untuk mencetak hasil kalkulasi bangun datar
func cetakInfoBangunDatar(nama string, h Hitung) {
	fmt.Printf("[%s] Luas: %.2f | Keliling: %.2f\n", nama, h.Luas(), h.Keliling())
}

// 3. Interface Kosong (`interface{}` atau alias modern: `any`)
// Dapat menerima tipe data apapun (karena semua tipe mengimplementasikan 0 method)
func periksaTipeData(data any) {
	// Type Switch untuk memeriksa tipe konkret di balik interface{}
	switch v := data.(type) {
	case string:
		fmt.Printf("Data bertipe STRING: %q\n", v)
	case int:
		fmt.Printf("Data bertipe INTEGER: %d\n", v)
	case bool:
		fmt.Printf("Data bertipe BOOLEAN: %t\n", v)
	default:
		fmt.Printf("Tipe data lain: %v\n", v)
	}
}

func TestInterface(t *testing.T) {
	// ==========================================
	// 1. Polimorfisme Interface Dasar
	// ==========================================
	fmt.Println("=== 1. Interface Dasar & Polimorfisme ===")
	adit := Person{Name: "Adit Kurniawan"}
	kucing := Animal{Name: "Kucing Anggora"}

	sayHelloToEntity(adit)   // Halo, Adit Kurniawan
	sayHelloToEntity(kucing) // Halo, Hewan: Kucing Anggora

	// ==========================================
	// 2. Contoh Nyata: Bangun Datar
	// ==========================================
	fmt.Println("\n=== 2. Interface Bangun Datar ===")
	persegi := PersegiPanjang{Panjang: 10, Lebar: 5}
	lingkaran := Lingkaran{Radius: 7}

	cetakInfoBangunDatar("Persegi Panjang", persegi)
	cetakInfoBangunDatar("Lingkaran", lingkaran)

	// ==========================================
	// 3. Interface Kosong (any / interface{})
	// ==========================================
	fmt.Println("\n=== 3. Interface Kosong (any) ===")
	var kosong any

	kosong = 100
	fmt.Println("Nilai kosong (int)    =", kosong)

	kosong = "Sekarang jadi string"
	fmt.Println("Nilai kosong (string) =", kosong)

	// Memeriksa tipe menggunakan Type Switch
	periksaTipeData("Belajar Golang")
	periksaTipeData(2026)
	periksaTipeData(true)

	// ==========================================
	// 4. Type Assertion (Ekstraksi Tipe Konkret)
	// ==========================================
	fmt.Println("\n=== 4. Type Assertion ===")
	var dataSample any = "Ini adalah teks rahasia"

	// Sintaks: variable.(TipeTujuan)
	// Gunakan pola "comma ok" agar aman dan tidak memicu panic jika tipe salah
	str, ok := dataSample.(string)
	if ok {
		fmt.Println("Type assertion sukses, isi:", str)
	} else {
		fmt.Println("Gagal mengonversi ke string!")
	}

	// Contoh jika tipe salah:
	_, okInt := dataSample.(int)
	fmt.Println("Apakah dataSample bertipe int? =", okInt) // false
}

/*
==================================================
PENJELASAN INTERFACE DI GO
==================================================

1. DEFINISI INTERFACE
   - Interface adalah tipe data abstrak yang mendefinisikan kumpulan tanda tangan method (method signatures).
   - Interface hanya berisi kontrak ("apa yang harus bisa dilakukan"), BUKAN implementasinya ("bagaimana cara melakukannya").

2. IMPLICIT IMPLEMENTATION (CIRI KHAS GO)
   - Di Go, TIDAK ADA kata kunci `implements` seperti di bahasa Java, C#, atau TypeScript.
   - Sebuah struct OTOMATIS mengimplementasikan sebuah interface jika struct tersebut memiliki seluruh method
     yang tertera di dalam interface tersebut dengan nama, parameter, dan return value yang cocok persis.
   - Konsep ini sering disebut "Duck Typing": *If it walks like a duck and quacks like a duck, it's a duck.*

3. INTERFACE KOSONG (`interface{}` / `any`)
   - Interface tanpa satupun method: `interface{}`.
   - Sejak Go 1.18, diperkenalkan kata kunci `any` sebagai alias resmi untuk `interface{}`.
   - Karena tidak memiliki syarat method, MAKA SEMUA TIPE DATA di Go (int, string, struct, slice, dll.)
     dianggap mengimplementasikan `any`.
   - Digunakan ketika kita ingin menerima data dinamis yang tipenya belum diketahui (seperti fungsi `fmt.Println`).

4. TYPE ASSERTION & TYPE SWITCH
   a. Mengapa Butuh Type Assertion?
      - Saat data disimpan di variabel `any` (atau interface), compiler Go "menyembunyikan" tipe aslinya.
      - Anda TIDAK BISA langsung melakukan operasi spesifik tipe (seperti penjumlahan `+`, fungsi `len()`, atau akses field/method struct)
        meskipun Anda tahu nilai di dalamnya adalah angka atau string.
      - Type assertion digunakan untuk memberi tahu compiler: "Bongkar bungkus any ini kembali ke tipe aslinya".

   b. Hati-hati: 1 Variabel (Bisa Panic) vs 2 Variabel (Comma-Ok):
      - CARA BERBAHAYA (1 Variabel):
          `val := data.(string)`
        Jika tebakan tipe salah (misal aslinya int), program AKAN LANGSUNG CRASH / RUNTIME PANIC!
      - CARA AMAN (2 Variabel / Comma-Ok Idiom):
          `val, ok := data.(string)`
        Jika tipe cocok : `ok = true`, `val` berisi nilai aslinya.
        Jika tipe salah : `ok = false`, `val` berisi zero value, dan PROGRAM TETAP AMAN TANPA PANIC!

   c. Kapan Wajib vs Kapan Tidak Perlu?
      - WAJIB       : Saat ingin melakukan operasi spesifik tipe tersebut (kalkulasi, manipulasi string, akses field struct).
      - TIDAK PERLU : Jika hanya meneruskan data ke fungsi yang juga menerima `any` (misal: `fmt.Println` atau `json.Marshal`).

   d. Type Switch:
      - Sintaks: `switch v := data.(type) { case string: ... case int: ... }`
      - Alternatif yang sangat bersih dan rapi jika variabel `any` memiliki banyak kemungkinan tipe data yang harus ditangani.

CARA MENJALANKAN FILE INI:
  $ go test -v -run TestInterface
==================================================
*/
