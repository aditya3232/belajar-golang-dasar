package belajar_golang_dasar

import (
	"fmt"
	"testing"
)

func TestSwitchExpression(t *testing.T) {
	// 1. Switch Sederhana
	// Di Go, tidak perlu menulis kata kunci `break` di setiap akhir case (otomatis break)
	name := "Adit"

	switch name {
	case "Adit":
		fmt.Println("Halo Adit")
	case "Budi":
		fmt.Println("Halo Budi")
	case "Joko":
		fmt.Println("Halo Joko")
	default:
		fmt.Println("Halo, salam kenal!")
	}

	// 2. Switch dengan Banyak Nilai dalam Satu Case (Multiple Values)
	day := "Sabtu"

	switch day {
	case "Senin", "Selasa", "Rabu", "Kamis", "Jumat":
		fmt.Println(day, "adalah Hari Kerja (Weekday)")
	case "Sabtu", "Minggu":
		fmt.Println(day, "adalah Hari Libur (Weekend)")
	default:
		fmt.Println("Nama hari tidak valid")
	}

	// 3. Switch dengan Short Statement
	// Deklarasi variabel inisialisasi sementara langsung di header switch
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama cukup panjang")
	case false:
		fmt.Println("Nama sudah pas dan ringkas")
	}

	// 4. Switch Tanpa Kondisi (Tagless Switch)
	// Digunakan sebagai alternatif yang jauh lebih rapi dibanding rangkaian if-else if yang panjang
	nilai := 85

	switch {
	case nilai >= 90:
		fmt.Println("Nilai Anda: A")
	case nilai >= 75:
		fmt.Println("Nilai Anda: B")
	case nilai >= 60:
		fmt.Println("Nilai Anda: C")
	default:
		fmt.Println("Nilai Anda: D (Remedial)")
	}

	// 5. Penggunaan Kata Kunci `fallthrough`
	// Digunakan jika ingin memaksa mengeksekusi case berikutnya TANPA mengecek kondisinya lagi
	angka := 2

	fmt.Println("--- Demonstrasi Fallthrough ---")
	switch angka {
	case 1:
		fmt.Println("Angka adalah 1")
	case 2:
		fmt.Println("Angka adalah 2")
		fallthrough // Lanjut mengeksekusi case 3 di bawahnya
	case 3:
		fmt.Println("Ini tercetak karena fallthrough dari case 2!")
	default:
		fmt.Println("Angka lainnya")
	}
}

/*
==================================================
PENJELASAN SWITCH EXPRESSION DI GO
==================================================

1. DEFINISI & KELEBIHAN DI GO
   - Switch expression adalah percabangan yang lebih sederhana dan rapi untuk memeriksa kesamaan nilai.
   - FITUR KHAS GO: Setiap blok `case` OTOMATIS BERHENTI (tidak butuh kata kunci `break`).
     Hal ini mencegah bug klasik "fallthrough tak sengaja" yang sering terjadi di bahasa C/Java/JS.

2. FITUR-FITUR UTAMA SWITCH DI GO
   a. Multiple Values per Case:
      - Satu case dapat menampung beberapa nilai sekaligus yang dipisahkan dengan tanda koma `,`.
      - Contoh: `case "Sabtu", "Minggu":`

   b. Switch dengan Short Statement:
      - Mirip dengan `if`, variabel dapat dideklarasikan sebelum ekspresi:
        `switch length := len(str); length > 5 { ... }`
      - Variabel hanya hidup di dalam lingkup blok `switch` tersebut.

   c. Switch Tanpa Kondisi (Tagless Switch):
      - Cukup menulis `switch { case kondisi1: ... case kondisi2: ... }`.
      - Menghilangkan kebutuhan menulis `if ... else if ... else if ...` yang panjang dan berulang.

   d. Kata Kunci `fallthrough`:
      - Kebalikan dari perilaku default Go. Jika Anda sengaja ingin eksekusi berlanjut ke case tepat di bawahnya,
        tambahkan kata kunci `fallthrough`.
      - PENTING: `fallthrough` akan langsung mengeksekusi kode di case berikutnya TANPA memeriksa kondisi case tersebut bernilai true atau false.

CARA MENJALANKAN FILE INI:
  $ go test -v -run TestSwitchExpression
==================================================
*/

