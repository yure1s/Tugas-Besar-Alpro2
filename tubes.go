package main

import "fmt"

const NMAX int = 100

type item struct {
	nama     string
	kategori string
	warna    string
	ukuran   string
	acara    string
}

type ootd [NMAX]item

func main() {
	var outfit ootd
	var nData int
	var pilih int

	selesai := false
	for !selesai {
		fmt.Println("\n--- Sistem Manajemen Butik ---")
		fmt.Println("1. Tambah Item Fashion")
		fmt.Println("2. Lihat Semua Item")
		fmt.Println("3. Rekomendasi Outfit")//Sequential Search
		fmt.Println("4. Urutkan data")//Insertion Sort (Kategori) & Selection Sort (Acara)
		fmt.Println("5. Edit Item Fashion")//Sequential Search
		fmt.Println("6. Cari berdasarkan Warna") //Sequential Search
		fmt.Println("7. Cari berdasarkan Ukuran ")//Binary Search
		fmt.Println("8. Keluar")

		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilih)

		if pilih == 1 {
			inputData(&outfit, &nData)
		} else if pilih == 2 {
			cetakData(outfit, nData)
		} else if pilih == 3 {
			rekomendasi(outfit, nData)
		} else if pilih == 4 {
			urutData(&outfit, nData)
			cetakData(outfit, nData)
		} else if pilih == 5 {
			editData(&outfit, &nData)
			cetakData(outfit, nData)
		} else if pilih == 6 {
			cariWarna(outfit, nData)
		} else if pilih == 7 {
			cariUkuran(&outfit, nData)
		} else if pilih == 8 {
			selesai = true
		}
	}
}

func inputData(A *ootd, n *int) {
	var tambah int
	fmt.Println("Berapa banyak item yang akan ditambahkan?")
	fmt.Scan(&tambah)
	fmt.Println("Nama item(gunakan _ jika lebih dari 2 kata), Kategori(Atasan/Aksesoris/Heels), Acara(Formal/Kasual/Pesta), Warna(Biru/Merah/Putih/Hitam), Ukuran(S/M/L/XL/XXL): ")
	for i := 0; i < tambah; i++ {
		fmt.Scan(&A[*n].nama, &A[*n].kategori, &A[*n].acara, &A[*n].warna, &A[*n].ukuran)
		*n++
	}
	fmt.Println("Input berhasil ^w^") 
	fmt.Printf("Total item sekarang : %d\n", *n)
}

func cetakData(A ootd, n int) {
	fmt.Printf("%-20s %-15s %-15s %-15s %-15s\n", "Nama Item", "Kategori", "Warna", "Ukuran", "Acara")
	for i := 0; i < 51; i++ {
		fmt.Print("-")
	}
	fmt.Println()
	for i := 0; i < n; i++ {
		fmt.Printf("%-20s %-15s %-15s %-15s %-15s\n", A[i].nama, A[i].kategori, A[i].warna, A[i].ukuran, A[i].acara)
	}
}

func urutData(A *ootd, n int) {
	var target string
	fmt.Printf("Diurutkan Berdasarkan? (Kategori / Acara)\n")
	fmt.Scan(&target)
	if target == "Kategori" {
		for i := 1; i < n; i++ {
			temp := A[i]
			j := i
			for j > 0 && temp.kategori < A[j-1].kategori {
				A[j] = A[j-1]
				j--
			}
			A[j] = temp
		}
	} else if target == "Acara" {
		for i := 0; i < n-1; i++ {
			idxMin := i
			for j := i + 1; j < n; j++ {
				if A[j].acara < A[idxMin].acara {
					idxMin = j
				}
			}
			temp := A[i]
			A[i] = A[idxMin]
			A[idxMin] = temp
		}
	}
}

func rekomendasi(A ootd, n int) {
	var targetAcara string
	var atasanShown, aksesorisShown, heelsShown bool
	var i int

	fmt.Print("Masukkan jenis acara (Formal/Kasual/Pesta): ")
	fmt.Scan(&targetAcara)

	fmt.Printf("\n--- Rekomendasi untuk Acara '%s' ---\n", targetAcara)
	fmt.Printf("%-20s %-15s %-15s\n", "Nama Item", "Kategori", "Acara")

	for k := 0; k < 51; k++ {
		fmt.Print("-")
	}
	fmt.Println()

	for i < n && !(atasanShown && aksesorisShown && heelsShown) {
		if A[i].acara == targetAcara {

			if A[i].kategori == "Atasan" && !atasanShown {
				fmt.Printf("%-20s %-15s %-15s\n",
					A[i].nama, A[i].kategori, A[i].acara)
				atasanShown = true

			} else if A[i].kategori == "Aksesoris" && !aksesorisShown {
				fmt.Printf("%-20s %-15s %-15s\n",
					A[i].nama, A[i].kategori, A[i].acara)
				aksesorisShown = true

			} else if A[i].kategori == "Heels" && !heelsShown {
				fmt.Printf("%-20s %-15s %-15s\n",
					A[i].nama, A[i].kategori, A[i].acara)
				heelsShown = true
			}
		}
		i++
	}
}

func cariWarna(A ootd, n int) {
	var warnaCari string
	var found bool

	fmt.Print("Masukkan warna yang dicari: ")
	fmt.Scan(&warnaCari)

	fmt.Println("\nHasil pencarian:")

	for i := 0; i < n; i++ {
		if A[i].warna == warnaCari {
			fmt.Printf("%s | %s | %s | %s | %s\n",
				A[i].nama,
				A[i].kategori,
				A[i].warna,
				A[i].ukuran,
				A[i].acara)
			found = true
		}
	}

	if !found {
		fmt.Println("Produk tidak ditemukan")
	}
}

func nilaiUkuran(u string) int {
	if u == "S" {
		return 1
	} else if u == "M" {
		return 2
	} else if u == "L" {
		return 3
	} else if u == "XL" {
		return 4
	} else if u == "XXL" {
		return 5
	}
	return 0
}

func sortUkuran(A *ootd, n int) {
	for i := 1; i < n; i++ {
		temp := A[i]
		j := i

		for j > 0 && nilaiUkuran(A[j-1].ukuran) > nilaiUkuran(temp.ukuran) {
			A[j] = A[j-1]
			j--
		}

		A[j] = temp
	}
}

func cariUkuran(A *ootd, n int) {
	var ukuranCari string

	sortUkuran(A, n)

	fmt.Print("Masukkan ukuran yang dicari (S/M/L/XL/XXL): ")
	fmt.Scan(&ukuranCari)

	left := 0
	right := n - 1
	found := -1

	for left <= right {
		mid := (left + right) / 2

		if nilaiUkuran(A[mid].ukuran) == nilaiUkuran(ukuranCari) {
			found = mid
			break
		} else if nilaiUkuran(A[mid].ukuran) < nilaiUkuran(ukuranCari) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if found != -1 {
		fmt.Println("Produk ditemukan:")
		fmt.Printf("%s | %s | %s | %s | %s\n",
			A[found].nama,
			A[found].kategori,
			A[found].warna,
			A[found].ukuran,
			A[found].acara)
	} else {
		fmt.Println("Produk tidak ditemukan")
	}
}

func editData(A *ootd, n *int) {
	var target string
	var pilihan, ubah int
	var found bool = false

	fmt.Printf("\nPilih apa yang ingin dilakukan\n")
	fmt.Printf("1. Edit\n")
	fmt.Printf("2. Hapus\n")
	fmt.Print("Opsi nomor (1/2)? ")
	fmt.Scan(&pilihan)

	if pilihan == 1 {
		fmt.Printf("Masukkan nama item yang ingin diubah: \n")
		fmt.Scan(&target)
		for i := 0; i < *n; i++ {
			if A[i].nama == target {
				fmt.Printf("Item ditemukan: %s - %s - %s\n", A[i].nama, A[i].kategori, A[i].acara)
				fmt.Printf("\nApa yang ingin anda ubah:\n 1. Nama\n 2. Kategori\n 3. Acara\n")
				fmt.Print("Pilih menu: ")
				fmt.Scan(&ubah)
				if ubah == 1 {
					fmt.Print("Masukkan nama baru (gunakan _ jika lebih dari 2 kata): ")
					fmt.Scan(&A[i].nama)
					fmt.Println("Nama item berhasil diubah")
				} else if ubah == 2 {
					fmt.Print("Masukkan kategori baru (Pakaian/Aksesoris/Heels): ")
					fmt.Scan(&A[i].kategori)
					fmt.Println("Kategori item berhasil diubah")
				} else if ubah == 3 {
					fmt.Print("Masukkan acara baru (Formal/Kasual/Pesta): ")
					fmt.Scan(&A[i].acara)
					fmt.Println("Acara Item berhasil diubah")
				}
			}
		}
	} else if pilihan == 2 {
		fmt.Printf("Masukkan nama item yang ingin dihapus: \n")
		fmt.Scan(&target)
		for i := 0; i < *n; i++ {
			if A[i].nama == target {
				for j := i; j < *n-1; j++ {
					A[j] = A[j+1]
				}
				*n--
				found = true
				i = *n
			}
		}
		if found {
			fmt.Printf("Item '%s' berhasil dihapus.\n", target)
			fmt.Printf("Total item setelah dihapus : %d\n", *n)
		} else {
			fmt.Println("Item tidak ditemukan")
		}
	}
}
