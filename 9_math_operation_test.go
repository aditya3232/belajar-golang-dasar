package belajar_golang_dasar

import (
	"fmt"
	"math"
	"testing"
)

func TestMathOperation(t *testing.T) {
	// 1. Operasi Aritmatika Dasar (+, -, *, /, %)
	a := 10
	b := 3

	fmt.Println("Penjumlahan (10 + 3) =", a+b) // 13
	fmt.Println("Pengurangan (10 - 3) =", a-b) // 7
	fmt.Println("Perkalian   (10 * 3) =", a*b) // 30
	fmt.Println("Pembagian   (10 / 3) =", a/b) // 3 (karena integer, pecahan dibuang)
	fmt.Println("Modulo/Sisa (10 % 3) =", a%b) // 1

	// Pembagian bilangan desimal (float)
	var x float64 = 10.0
	var y float64 = 3.0
	fmt.Println("Pembagian float (10.0 / 3.0) =", x/y) // 3.3333333333333335

	// 2. Augmented Assignments (Operasi Penugasan Gabungan)
	// a = a + 10  ->  a += 10
	c := 10
	c += 5 // c = 10 + 5
	fmt.Println("c += 5  ->", c) // 15

	c -= 3 // c = 15 - 3
	fmt.Println("c -= 3  ->", c) // 12

	c *= 2 // c = 12 * 2
	fmt.Println("c *= 2  ->", c) // 24

	c /= 4 // c = 24 / 4
	fmt.Println("c /= 4  ->", c) // 6

	c %= 4 // c = 6 % 4
	fmt.Println("c %= 4  ->", c) // 2

	// 3. Unary Operator
	// Di Go, ++ dan -- adalah STATEMENT, bukan expression
	d := 1
	d++ // sama dengan d = d + 1
	fmt.Println("d++ =", d) // 2

	d-- // sama dengan d = d - 1
	fmt.Println("d-- =", d) // 1

	// Unary minus (-) dan plus (+)
	positif := +100
	negatif := -100
	fmt.Println("Positif =", positif)
	fmt.Println("Negatif =", negatif)

	// 4. Operasi Matematika Lanjutan dengan Package bawaan `math`
	// Fungsi math bawaan umumnya bekerja dengan tipe data `float64`
	fmt.Println("math.Sqrt(16)   =", math.Sqrt(16))     // Akar kuadrat = 4
	fmt.Println("math.Pow(2, 3)  =", math.Pow(2, 3))    // 2 pangkat 3 = 8
	fmt.Println("math.Round(4.7) =", math.Round(4.7))   // Pembulatan ke terdekat = 5
	fmt.Println("math.Floor(4.9) =", math.Floor(4.9))   // Pembulatan ke bawah = 4
	fmt.Println("math.Ceil(4.1)  =", math.Ceil(4.1))    // Pembulatan ke atas = 5
}

/*
==================================================
PENJELASAN OPERASI MATEMATIKA DI GO
==================================================

1. OPERATOR ARITMATIKA DASAR
   - `+` : Penjumlahan
   - `-` : Pengurangan
   - `*` : Perkalian
   - `/` : Pembagian
   - `%` : Modulo (sisa bagi) -> Hanya bisa digunakan untuk tipe integer.

2. ATURAN PEMBAGIAN INTEGER
   - Di Go, jika dua bilangan bertipe integer dibagi (misal 10 / 3), hasilnya tetap berupa integer (3).
   - Nilai pecahan di belakang koma langsung dibuang (dibulatkan ke nol/truncated), bukan dibulatkan ke atas.
   - Jika ingin mendapatkan hasil pecahan, setidaknya salah satu atau kedua operan harus bertipe floating-point (misal 10.0 / 3.0).

3. AUGMENTED ASSIGNMENTS
   Singkatan untuk memodifikasi nilai variabel itu sendiri:
   - `a += b`  ->  `a = a + b`
   - `a -= b`  ->  `a = a - b`
   - `a *= b`  ->  `a = a * b`
   - `a /= b`  ->  `a = a / b`
   - `a %= b`  ->  `a = a % b`

4. ATURAN KHUSUS UNARY OPERATOR (`++` dan `--`) DI GO
   - Berbeda dengan bahasa C/Java/JavaScript:
     a. Di Go, `++` dan `--` adalah STATEMENT, BUKAN EXPRESSION.
        Artinya, tidak bisa dimasukkan ke dalam assignment atau argumen fungsi:
        Contoh ERROR:
          var x = y++        // Compile error!
          fmt.Println(d++)   // Compile error!
     b. Hanya ada format postfix (`d++`, `d--`), TIDAK ADA prefix (`++d` atau `--d` akan error).

5. PACKAGE BAWAAN `math`
   - Untuk operasi matematika yang lebih kompleks (seperti perpangkatan, akar, trigonometri, nilai mutlak),
     Go menyediakan package standar `math`.
   - Hampir seluruh fungsi dalam package `math` menerima dan mengembalikan tipe data `float64`.

CARA MENJALANKAN FILE INI:
  $ go test -v -run TestMathOperation
==================================================
*/

