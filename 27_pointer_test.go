package belajar_golang_dasar

import (
	"errors"
	"fmt"
	"testing"
)

type Location struct {
	City, Province, Country string
}

type Man struct {
	Name string
}

// 1. Fungsi dengan Pointer Parameter
// Parameter bertipe `*Location` memungkinkan fungsi mengubah data variabel aslinya
func changeCountryToIndonesia(loc *Location) {
	loc.Country = "Indonesia"
}

// 2. Struct Method: Value Receiver (Tidak mengubah data asli)
func (man Man) NonPointerMarried() {
	man.Name = "Mr. " + man.Name
}

// 3. Struct Method: Pointer Receiver (Mengubah data struct aslinya)
func (man *Man) PointerMarried() {
	man.Name = "Mr. " + man.Name
}

// 4. Standar Method Proyek Go (Request Value & Response Pointer)
type LoginRequest struct {
	Username string
	Password string
}

type UserResponse struct {
	Username string
	Token    string
}

type AuthService struct {
	AppName string
}

// Pola standar method di proyek Go:
// - Receiver bertipe Pointer (*AuthService) agar efisien dan konsisten
// - Parameter bertipe Value (LoginRequest) sebagai input read-only (aman dari mutasi & nil)
// - Return bertipe Pointer (*UserResponse) agar bisa mengembalikan nil saat error
// - Return error diletakkan di posisi terakhir
func (s *AuthService) SignIn(req LoginRequest) (*UserResponse, error) {
	// Validasi input
	if req.Username == "" || req.Password == "" {
		// Mengembalikan nil untuk response karena terjadi error
		return nil, errors.New("username dan password wajib diisi")
	}

	if req.Username != "admin" || req.Password != "secret123" {
		return nil, errors.New("kredensial login salah")
	}

	// Berhasil: kembalikan pointer ke UserResponse dan error bernilai nil
	return &UserResponse{
		Username: req.Username,
		Token:    "jwt-token-rahasia-" + s.AppName,
	}, nil
}

func TestPointer(t *testing.T) {
	// ==========================================
	// 1. Konsep Dasar Pass by Value vs Pointer
	// ==========================================
	fmt.Println("=== 1. Operator & (Reference) dan * (Dereference) ===")
	angkaA := 10

	// Operator `&` mengambil alamat memori (Memory Address) dari variabel angkaA
	// Tipe data `ptrAngka` adalah `*int` (Pointer ke integer)
	var ptrAngka *int = &angkaA

	fmt.Println("Nilai angkaA            =", angkaA)
	fmt.Println("Alamat memori angkaA (&)=", ptrAngka)

	// Operator `*` (Dereference) digunakan untuk mengakses/mengubah nilai di alamat tersebut
	*ptrAngka = 25 // Mengubah nilai di alamat yang ditunjuk

	fmt.Println("Nilai angkaA setelah *ptr diubah =", angkaA) // angkaA ikut berubah menjadi 25!

	// ==========================================
	// 2. Pointer pada Struct
	// ==========================================
	fmt.Println("\n=== 2. Pointer pada Struct ===")
	loc1 := Location{"Jakarta Selatan", "DKI Jakarta", "Indonesia"}

	// loc2 menunjuk ke alamat memori yang sama dengan loc1
	loc2 := &loc1

	loc2.City = "Bandung" // Otomatis dereference di Go (tidak perlu (*loc2).City)

	fmt.Println("loc1 (ikut berubah) =", loc1) // {Bandung DKI Jakarta Indonesia}
	fmt.Println("loc2                =", loc2)

	// Mengarahkan ulang seluruh isi memori menggunakan operator `*`
	// Semua variabel yang menunjuk ke alamat ini akan ikut berubah datanya
	*loc2 = Location{"Malang", "Jawa Timur", "Indonesia"}
	fmt.Println("loc1 setelah *loc2 di-assign ulang =", loc1) // {Malang Jawa Timur Indonesia}

	// ==========================================
	// 3. Fungsi bawaan: new(TipeData)
	// ==========================================
	fmt.Println("\n=== 3. Fungsi new ===")
	// `new` mengalokasikan memori kosong dan mengembalikan pointer ke zero value
	locBaru := new(Location)
	locBaru.City = "Surabaya"
	fmt.Println("Hasil new(Location) =", locBaru) // &[Surabaya  ]

	// ==========================================
	// 4. Pointer sebagai Parameter Fungsi
	// ==========================================
	fmt.Println("\n=== 4. Pointer di Fungsi ===")
	locKantor := Location{"Tokyo", "Kanto", "Jepang"}

	// Berikan operator `&` saat memanggil fungsi yang menerima parameter pointer
	changeCountryToIndonesia(&locKantor)
	fmt.Println("Negara locKantor sekarang =", locKantor.Country) // Berubah menjadi "Indonesia"

	// ==========================================
	// 5. Pointer Receiver pada Struct Method
	// ==========================================
	fmt.Println("\n=== 5. Pointer Receiver pada Method ===")
	adit := Man{Name: "Adit"}

	// a. Memanggil method dengan Value Receiver (Gagal mengubah nama)
	adit.NonPointerMarried()
	fmt.Println("Setelah NonPointerMarried =", adit.Name) // Tetap "Adit"

	// b. Memanggil method dengan Pointer Receiver (Berhasil mengubah nama asli)
	adit.PointerMarried()
	fmt.Println("Setelah PointerMarried    =", adit.Name) // Berubah menjadi "Mr. Adit"

	// ==========================================
	// 6. Pola Standar Method di Proyek Nyata (Request Value & Response Pointer)
	// ==========================================
	fmt.Println("\n=== 6. Standar Method Request & Response ===")
	authService := &AuthService{AppName: "SistemAkademik"}

	// Kasus A: Login Gagal (Response bernilai nil dan error terisi)
	reqGagal := LoginRequest{
		Username: "admin",
		Password: "password_salah",
	}
	resGagal, errGagal := authService.SignIn(reqGagal)
	if errGagal != nil {
		fmt.Println("Status Login : GAGAL ->", errGagal)
		fmt.Println("Hasil res    : nil? ->", resGagal == nil) // true (aman, response adalah pointer nil)
	}

	// Kasus B: Login Sukses (Response terisi data pointer dan error bernilai nil)
	reqSukses := LoginRequest{
		Username: "admin",
		Password: "secret123",
	}
	resSukses, errSukses := authService.SignIn(reqSukses)
	if errSukses == nil {
		fmt.Println("Status Login : SUKSES")
		fmt.Printf("Data User    : Username=%s, Token=%s\n", resSukses.Username, resSukses.Token)
	}
}

/*
==================================================
PENJELASAN POINTER DI GO
==================================================

1. DEFINISI & KARAKTERISTIK
   - Secara default di Go, semua operasi assignment dan passing parameter bersifat "PASS BY VALUE"
     (seluruh isi data diduplikasi/di-copy ke variabel baru).
   - Pointer adalah variabel yang menyimpan ALAMAT MEMORI dari variabel lain, bukan nilainya langsung.
   - Dengan pointer, kita dapat melakukan "PASS BY REFERENCE", yaitu memanipulasi data di sumber aslinya.

2. DUA OPERATOR KUNCI
   a. Operator `&` (Address-of / Reference):
      - Mengambil alamat memori dari sebuah variabel (contoh: `ptr := &data`).
   b. Operator `*` (Dereference):
      - Mengambil atau mengubah nilai asli yang tersimpan di alamat yang ditunjuk pointer (contoh: `*ptr = 100`).
      - Juga dipakai untuk mendefinisikan tipe data pointer (contoh: `var p *int`).

3. FUNGSI BAWAAN `new()`
   - Digunakan untuk membuat pointer baru secara instan.
   - Mengembalikan pointer yang mengarah ke data kosong (*zero value*).
   - Contoh: `p := new(Location)` sama hasilnya dengan `p := &Location{}`.

4. KAPAN HARUS MENGGUNAKAN POINTER?
   a. Ingin Mengubah Data Asli:
      - Ketika sebuah fungsi atau method perlu memodifikasi data variabel/struct yang dikirimkan.
   b. Efisiensi Performa & Memori (Ukuran Data Besar):
      - Jika struct memiliki banyak field atau data berukuran besar, menduplikasinya setiap kali dipanggil
        akan membebani memori dan Garbage Collector. Mengirim pointer hanya menyalin alamat memori (8 byte).

5. VALUE RECEIVER VS POINTER RECEIVER PADA METHOD
   - `func (m Man) Method()`    : Value receiver. `m` adalah salinan/copy. Perubahan di dalam method TIDAK mengubah struct asli.
   - `func (m *Man) Method()`   : Pointer receiver. Perubahan pada `m` LANGSUNG mengubah struct aslinya.
   - BEST PRACTICE GO: Sangat disarankan selalu menggunakan **Pointer Receiver** untuk method struct,
     kecuali struct memang sengaja didesain kecil dan immutable.

6. KEAMANAN POINTER DI GO
   - Pointer di Go JAUH LEBIH AMAN dibanding bahasa C/C++.
   - Go TIDAK MENDUKUNG "Pointer Arithmetic" (seperti `ptr++` atau `ptr + 4`), sehingga tidak ada risiko
     sengaja/tidak sengaja membaca segmen memori terlarang yang menyebabkan memory corruption.

7. STANDAR PENULISAN METHOD DI PROYEK GO (REQUEST & RESPONSE)
   Format: func (s *Service) Method(ctx context.Context, req Request) (*Response, error)
   a. Receiver Pointer (*Service):
      - Menghindari duplikasi struct service/repository di setiap pemanggilan method.
   b. Parameter Request Value (req Request):
      - Menjamin input bersifat read-only (immutable snapshot).
      - Service terhindar dari bug modifikasi data luar dan aman dari crash `nil pointer dereference`.
   c. Return Pointer (*Response):
      - BUKAN semata-mata untuk dimutasi di luar, melainkan AGAR BISA MENGEMBALIKAN `nil` saat terjadi error.
      - Menghindari copy alokasi memori berlebih saat data mengalir antar layer (Repo -> Service -> Handler).
   d. Return error di urutan terakhir:
      - Konvensi standar penanganan kesalahan di Go.

CARA MENJALANKAN FILE INI:
  $ go test -v -run TestPointer
==================================================
*/
