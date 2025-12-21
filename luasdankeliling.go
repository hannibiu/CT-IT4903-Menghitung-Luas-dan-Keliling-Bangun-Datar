// Nama : Ni Komang Yuni Handayani (103032530007)//
// Nama : Yemima Denta Arella (103032500080)//

package main

import "fmt"

func main() {
	var pilihan int
	var lanjut string
	var jalan = true

	for jalan {
		fmt.Println("=== Selamat Datang di Program Menghitung Luas dan Keliling Bangun Datar ===")
		fmt.Println("1. Persegi")
		fmt.Println("2. Segitiga")
		fmt.Println("3. Lingkaran")
		fmt.Println("4. Trapesium")
		fmt.Println("5. Exit")
		fmt.Print("Silahkan masukan angka yang ingin anda hitung : ")
		fmt.Scan(&pilihan)

		if pilihan < 1 || pilihan > 5 {
			fmt.Printf("Maaf opsi %d tidak tersedia, mohon masukan opsi yang valid\n\n", pilihan)

		} else if pilihan == 5 {
			fmt.Println("Terimakasih telah menggunakan program ini.")
			jalan = false

			//hitung persegi//
		} else if pilihan == 1 {
			var sisi float64
			for sisi <= 0 {
				fmt.Print("Masukan nilai sisi (sisi > 0): ")
				fmt.Scan(&sisi)
			}
			fmt.Println("Luas persegi : ", sisi*sisi)
			fmt.Println("Keliling persegi : ", 4*sisi)

			//hitung segitiga//
		} else if pilihan == 2 {
			var alas, tinggi, s1, s2, s3 float64

			for alas <= 0 {
				fmt.Print("Masukan nilai alas (alas > 0) : ")
				fmt.Scan(&alas)

			}
			for tinggi <= 0 {
				fmt.Print("Masukan nilai tinggi (tinggi > 0) : ")
				fmt.Scan(&tinggi)
			}
			for s1 <= 0 {
				fmt.Print("Masukan nilai sisi 1 (1 > 0 ) : ")
				fmt.Scan(&s1)
			}
			for s2 <= 0 {
				fmt.Print("Masukan nilai sisi 2 (2 > 0) : ")
				fmt.Scan(&s2)
			}
			for s3 <= 0 {
				fmt.Print("Masukan nilai sisi 3 (3 > 0) : ")
				fmt.Scan(&s3)
			}

			fmt.Println("Luas segitiga : ", 0.5*alas*tinggi)
			fmt.Println("Keliling segitiga : ", s1+s2+s3)

			//hitung lingkaran//
		} else if pilihan == 3 {
			var r float64
			const phi = 3.14

			for r <= 0 {
				fmt.Print("Masukan nilai jari-jari (jari jari > 0) : ")
				fmt.Scan(&r)
			}

			fmt.Println("Luas lingkaran : ", phi*r*r)
			fmt.Println("Keliling lingkaran : ", 2*phi*r)

			//hitung trapesium//
		} else if pilihan == 4 {
			var a, b, tinggi, s1, s2 float64

			for a <= 0 {
				fmt.Print("Masukan nilai sisi sejajar A (A > 0) : ")
				fmt.Scan(&a)
			}
			for b <= 0 {
				fmt.Print("Masukan nilai sisi sejajar B (B > 0) : ")
				fmt.Scan(&b)
			}
			for tinggi <= 0 {
				fmt.Print("Masukan nilai tinggi (tinggi > 0 ) : ")
				fmt.Scan(&tinggi)
			}
			for s1 <= 0 {
				fmt.Print("Masukann nilai sisi miring A (A > 0 ) : ")
				fmt.Scan(&s1)
			}
			for s2 <= 0 {
				fmt.Print("Masukan nilai sisi miring B (B > 0) : ")
				fmt.Scan(&s2)
			}

			fmt.Println("Luas trapesium : ", 0.5*(a+b)*tinggi)
			fmt.Println("Keliling trapesium : ", a+b+s1+s2)

		}
		if jalan {
			lanjut = ""
			for lanjut != "y" && lanjut != "Y" && lanjut != "n" && lanjut != "N" {
				fmt.Print("\n Apakah ingin lanjut menghitung (y/n) : ")
				fmt.Scan(&lanjut)
			}

			if lanjut != "y" && lanjut != "Y" && lanjut != "n" && lanjut != "N" {
				fmt.Println("Input tidak valid, silakan masukan y/n.")
			}
			if lanjut == "n" || lanjut == "N" {
				fmt.Println("Terima kasih telah menggunakan program ini.")
				jalan = false
			}

		}
	}
}
