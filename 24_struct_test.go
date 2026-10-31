package belajar_golang_dasar

import (
	"fmt"
	"testing"
)

// 1. Definisi Struct Dasar
// Template data gabungan beberapa tipe field (mirip blueprint/class)
type Customer struct {
	Name, Address string
	Age           int
}

// 2. Struct Method (Menempelkan fungsi ke struct)
// `(c Customer)` disebut sebagai "Receiver"
func (c Customer) sayHello(guestName string) {
	fmt.Printf("Halo %s! Nama saya adalah %s, asal %s.\n", guestName, c.Name, c.Address)
}

// 3. Embedded Struct (Komposisi / Composition - Pengganti Inheritance di Go)
type Address struct {
	City, Country string
}

type Employee struct {
	Name     string
	Position string
	Address  // Embedded struct (field-field dari Address otomatis terangkat ke Employee)
}

func TestStruct(t *testing.T) {
	// ==========================================
	// 1. Inisialisasi Struct Manual (Satu per satu)
	// ==========================================
	fmt.Println("=== 1. Inisialisasi Manual ===")
	var adit Customer
	// Sebelum diisi, semua field berisi zero value (string: "", int: 0)
	fmt.Println("Zero value Customer =", adit) // {  0}

	adit.Name = "Adit Kurniawan"
	adit.Address = "Bandung"
	adit.Age = 25

	fmt.Println("Nama    =", adit.Name)
	fmt.Println("Alamat  =", adit.Address)
	fmt.Println("Usia    =", adit.Age)

	// ==========================================
	// 2. Inisialisasi Menggunakan Struct Literal
	// ==========================================
	fmt.Println("\n=== 2. Struct Literal ===")
	// a. Menggunakan Nama Field (Sangat Direkomendasikan karena aman terhadap perubahan urutan field)
	joko := Customer{
		Name:    "Joko Widodo",
		Address: "Solo",
		Age:     60,
	}
	fmt.Println("Customer Joko =", joko)

	// b. Inisialisasi Posisi Langsung (Harus urut sesuai definisi struct)
	budi := Customer{"Budi Santoso", "Surabaya", 30}
	fmt.Println("Customer Budi =", budi)

	// ==========================================
	// 3. Memanggil Struct Method
	// ==========================================
	fmt.Println("\n=== 3. Struct Method ===")
	adit.sayHello("Budi") // Halo Budi! Nama saya adalah Adit Kurniawan, asal Bandung.
	joko.sayHello("Adit")

	// ==========================================
	// 4. Anonymous Struct (Struct Sekali Pakai)
	// ==========================================
	fmt.Println("\n=== 4. Anonymous Struct ===")
	// Dibuat tanpa deklarasi `type ... struct`, berguna untuk payload DTO/HTTP response lokal
	response := struct {
		StatusCode int
		Message    string
	}{
		StatusCode: 200,
		Message:    "Data berhasil diambil",
	}
	fmt.Printf("Status: %d | Pesan: %s\n", response.StatusCode, response.Message)

	// ==========================================
	// 5. Embedded Struct (Composition)
	// ==========================================
	fmt.Println("\n=== 5. Embedded Struct ===")
	pegawai := Employee{
		Name:     "Siti Aminah",
		Position: "Backend Developer",
		Address: Address{
			City:    "Jakarta Selatan",
			Country: "Indonesia",
		},
	}

	// Field dari Address langsung dapat diakses dari struct Employee
	fmt.Println("Nama Karyawan =", pegawai.Name)
	fmt.Println("Kota          =", pegawai.City)          // Langsung via embedded
	fmt.Println("Negara        =", pegawai.Address.Country) // Atau lewat nama struct aslinya
}

/*
==================================================
PENJELASAN STRUCT DI GO
==================================================

1. DEFINISI & PERAN STRUCT
   - Go BUKAN bahasa berorientasi objek murni (tidak ada kata kunci `class` maupun pewarisan `inheritance`).
   - `struct` adalah tipe data buatan (koleksi beberapa field data) yang menjadi fondasi utama
     pemodelan data dan entitas di Go.

2. METODE INISIALISASI
   a. Deklarasi Kosong:
      var c Customer -> Seluruh field otomatis bernilai zero value masing-masing tipe data.
   b. Named Fields Literal (Paling Direkomendasikan):
      c := Customer{Name: "Adit", Age: 25} -> Jelas, eksplisit, dan tidak rusak jika urutan field di struct berubah.
   c. Positional Literal:
      c := Customer{"Adit", "Bandung", 25} -> Rentan error jika ada field baru yang ditambahkan di kemudian hari.

3. STRUCT METHOD
   - Kita bisa menempelkan fungsi ke struct dengan menambahkan "Receiver" sebelum nama fungsi:
     `func (c Customer) sayHello(...)`
   - Receiver ini memungkinkan method mengakses field-field milik struct tersebut (mirip `this` atau `self`).
   - (Catatan: Pointer Receiver untuk memodifikasi data struct asli akan dipelajari pada materi Pointer).

4. ANONYMOUS STRUCT
   - Struct yang dibuat langsung saat inisialisasi tanpa membuat deklarasi `type` di level package.
   - Sangat sering digunakan untuk data sementara, seperti request/response JSON atau unit test table data.

5. EMBEDDED STRUCT (COMPOSITION)
   - Go menerapkan prinsip "Composition over Inheritance".
   - Untuk mewariskan/menggabungkan properti dari struct lain, cukup cantumkan tipe struct tersebut tanpa nama field.
   - Field dari struct yang di-embed otomatis dapat diakses seolah-olah milik struct pembungkusnya secara langsung.

CARA MENJALANKAN FILE INI:
  $ go test -v -run TestStruct
==================================================
*/
