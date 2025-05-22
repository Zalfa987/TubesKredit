package main

import (
	"fmt"
	"tubeskredit/operasi"
	"tubeskredit/pinjaman"
)

func main() {
	// Inisialisasi daftar pinjaman
	daftarPinjaman := &pinjaman.DaftarPinjaman{}

	// Variabel untuk input
	var pilihan int
	var namaPeminjam string
	var jumlahPinjaman float64
	var lamaPinjaman int
	var sukuBunga float64

	for {
		// Menu Utama
		fmt.Println("\n--- SIMULASI PINJAMAN ---")
		fmt.Println("1. Tambah Pinjaman")
		fmt.Println("2. Tampilkan Daftar Pinjaman")
		fmt.Println("3. Cari Pinjaman (Sequential)")
		fmt.Println("4. Urutkan Pinjaman berdasarkan Jumlah")
		fmt.Println("5. Urutkan Pinjaman berdasarkan Nama")
		fmt.Println("6. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			// Input data pinjaman
			fmt.Print("Nama Peminjam: ")
			fmt.Scan(&namaPeminjam)
			fmt.Print("Jumlah Pinjaman: ")
			fmt.Scan(&jumlahPinjaman)
			fmt.Print("Lama Pinjaman (bulan): ")
			fmt.Scan(&lamaPinjaman)
			fmt.Print("Suku Bunga (%): ")
			fmt.Scan(&sukuBunga)

			// Buat objek pinjaman
			pinjamBaru := pinjaman.Pinjaman{
				NamaPeminjam:   namaPeminjam,
				JumlahPinjaman: jumlahPinjaman,
				LamaPinjaman:   lamaPinjaman,
				SukuBunga:      sukuBunga,
			}

			// Tambah pinjaman
			daftarPinjaman.TambahPinjaman(pinjamBaru)

		case 2:
			// Tampilkan daftar pinjaman
			daftarPinjaman.TampilkanPinjaman()

		case 3:
			// Cari pinjaman
			fmt.Print("Masukkan nama yang dicari: ")
			fmt.Scan(&namaPeminjam)
			var hasil [operasi.MaksHasil]pinjaman.Pinjaman
			var nHasil int
			operasi.CariPinjamanSequential(daftarPinjaman, namaPeminjam, &hasil, &nHasil)
			if nHasil > 0 {
				fmt.Println("Pinjaman ditemukan:")
				for i := 0; i < nHasil; i++ {
					fmt.Printf("Nama: %s | Pinjaman: Rp %.2f\n", hasil[i].NamaPeminjam, hasil[i].JumlahPinjaman)
				}
			} else {
				fmt.Println("Pinjaman tidak ditemukan")
			}

		case 4:
			// Urutkan berdasarkan jumlah pinjaman
			operasi.UrutPinjamanSelectionSort(&daftarPinjaman.Data, daftarPinjaman.N)
			fmt.Println("Daftar Pinjaman Terurut (Jumlah Pinjaman):")
			for i := 0; i < daftarPinjaman.N; i++ {
				p := daftarPinjaman.Data[i]
				fmt.Printf("Nama: %s | Pinjaman: Rp %.2f\n", p.NamaPeminjam, p.JumlahPinjaman)
			}

		case 5:
			// Urutkan berdasarkan nama
			operasi.UrutPinjamanInsertionSort(&daftarPinjaman.Data, daftarPinjaman.N)
			fmt.Println("Daftar Pinjaman Terurut (Nama):")
			for i := 0; i < daftarPinjaman.N; i++ {
				p := daftarPinjaman.Data[i]
				fmt.Printf("Nama: %s | Pinjaman: Rp %.2f\n", p.NamaPeminjam, p.JumlahPinjaman)
			}

		case 6:
			// Keluar dari program
			fmt.Println("Terima kasih!")
			return

		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}
