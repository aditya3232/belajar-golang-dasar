package belajar_golang_dasar

import (
	"errors"
	"fmt"
	"testing"
)

// 1. Sentinel Error (Error Konstan yang dapat dibandingkan)
var (
	ErrUserNotFound = errors.New("user tidak ditemukan")
	ErrUnauthorized = errors.New("akses ditolak")
)

// 2. Fungsi dengan Return Error Standar (errors.New)
func pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("pembagian dengan angka nol tidak diizinkan")
	}
	return nilai / pembagi, nil
}

// 3. Custom Error Struct
// Siapa pun yang memiliki method `Error() string` otomatis menjadi `error`
type ValidationError struct {
	Field   string
	Message string
}

func (v *ValidationError) Error() string {
	return fmt.Sprintf("validasi gagal pada field '%s': %s", v.Field, v.Message)
}

func validateUser(username string) error {
	if username == "" {
		return &ValidationError{
			Field:   "username",
			Message: "tidak boleh kosong",
		}
	}
	return nil
}

// 4. Fungsi dengan Error Wrapping (fmt.Errorf dan verb %w)
func queryDatabase(id int) error {
	if id == 404 {
		// Membungkus ErrUserNotFound dengan konteks tambahan
		return fmt.Errorf("repositori gagal: %w", ErrUserNotFound)
	}
	return nil
}

func TestError(t *testing.T) {
	// ==========================================
	// 1. Penanganan Error Standar (errors.New)
	// ==========================================
	fmt.Println("=== 1. Error Standar ===")
	// Kasus Berhasil
	hasilSukses, err1 := pembagian(10, 2)
	if err1 == nil {
		fmt.Println("Hasil 10 / 2 =", hasilSukses)
	}

	// Kasus Gagal
	hasilGagal, err2 := pembagian(10, 0)
	if err2 != nil {
		fmt.Println("Terjadi error pembagian :", err2.Error())
		fmt.Println("Hasil zero value         :", hasilGagal) // 0
	}

	// ==========================================
	// 2. Custom Error & Pengecekan dengan errors.As
	// ==========================================
	fmt.Println("\n=== 2. Custom Error (errors.As) ===")
	errValidation := validateUser("")
	if errValidation != nil {
		fmt.Println("Pesan error:", errValidation.Error())

		// errors.As digunakan untuk meng-cast error ke tipe struct aslinya
		var vErr *ValidationError
		if errors.As(errValidation, &vErr) {
			fmt.Printf("Detail -> Field yang salah: '%s' | Alasan: '%s'\n", vErr.Field, vErr.Message)
		}
	}

	// ==========================================
	// 3. Error Wrapping & Pengecekan dengan errors.Is
	// ==========================================
	fmt.Println("\n=== 3. Error Wrapping & errors.Is ===")
	errRepo := queryDatabase(404)
	if errRepo != nil {
		fmt.Println("Error asli tertutup context :", errRepo)

		// errors.Is dapat menembus pembungkus (%w) untuk mencocokkan error sentinel aslinya
		if errors.Is(errRepo, ErrUserNotFound) {
			fmt.Println("Mendeteksi: Error ini berakar dari ErrUserNotFound!")
		}
	}
}

/*
==================================================
PENJELASAN ERROR HANDLING DI GO
==================================================

1. FILOSOFI: "ERRORS ARE VALUES"
   - Go TIDAK MENGGUNAKAN mekanisme `try-catch` / `exception` seperti Java, Python, atau C#.
   - Di Go, error diperlakukan sebagai NILAI BIASA (first-class value) yang dikembalikan sebagai nilai balik terakhir dari fungsi.
   - Pengecekan dilakukan secara eksplisit:
       result, err := doSomething()
       if err != nil {
           // tangani error di sini
       }

2. TIPE DATA ERROR
   - `error` di Go sebenarnya hanyalah sebuah interface bawaan sederhana:
       type error interface {
           Error() string
       }
   - Tipe data atau struct apapun yang memiliki method `Error() string` otomatis dianggap sebagai `error`.

3. MEMBUAT ERROR
   a. `errors.New("pesan")`: Membuat error teks sederhana.
   b. `fmt.Errorf("format: %v", data)`: Membuat error dengan format string dinamis.
   c. Custom Struct: Membuat struct khusus untuk menyimpan metadata error (seperti `ValidationError`).

4. ERROR WRAPPING, errors.Is & errors.As (Go 1.13+)
   a. Error Wrapping (`%w`):
      - Membungkus error asal dengan pesan tambahan tanpa kehilangan jejak error aslinya:
        `fmt.Errorf("gagal simpan order: %w", err)`
   b. `errors.Is(err, target)`:
      - Memeriksa apakah error cocok dengan error acuan (sentinel error) tertentu,
        bahkan jika error tersebut sudah dibungkus bertingkat.
   c. `errors.As(err, &targetStruct)`:
      - Mencari dan mengonversi error ke tipe custom error struct tertentu dari rantai pembungkus.

CARA MENJALANKAN FILE INI:
  $ go test -v -run TestError
==================================================
*/
