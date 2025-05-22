package main

import (
	"fmt"
	"tubeskredit/operasi"
	"tubeskredit/pinjaman"
	"tubeskredit/ubah_hapus"
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
		fmt.Println("\n--- SIMULASI PINJAMAN DAN KREDIT ---")
		fmt.Println("1. Tambah Peminjam")
		fmt.Println("2. Tampilkan Daftar Peminjam")
		fmt.Println("3. Cari Peminjam (Sequential)")
		fmt.Println("4. Urutkan Peminjam berdasarkan Jumlah")
		fmt.Println("5. Urutkan Peminjam berdasarkan Nama")
		fmt.Println("6. Ubah Data Peminjam")
		fmt.Println("7. Hapus Data Peminjam")
		fmt.Println("8. Keluar")
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

			bungaTotal := jumlahPinjaman * (sukuBunga / 100)
			totalBayar := jumlahPinjaman + bungaTotal
			cicilanPerBulan := totalBayar / float64(lamaPinjaman)

			fmt.Printf("Cicilan per bulan: Rp %.2f\n", cicilanPerBulan)

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
			var namaCari string
			fmt.Print("Masukkan nama yang dicari: ")
			fmt.Scan(&namaCari)
			var hasil [operasi.MaksHasil]pinjaman.Pinjaman
			var nHasil int
			operasi.CariPinjamanSequential(daftarPinjaman, namaCari, &hasil, &nHasil)
			if nHasil > 0 {
				fmt.Println("Pinjaman ditemukan:")
				for i := 0; i < nHasil; i++ {
					fmt.Printf("Nomor: %d | Nama: %s | Pinjaman: Rp %.2f | Status: %s\n",
						hasil[i].Nomor,
						hasil[i].NamaPeminjam,
						hasil[i].JumlahPinjaman,
						hasil[i].StatusPembayaran)
				}
			} else {
				fmt.Println("Pinjaman tidak ditemukan")
			}

		case 4:
			// Urutkan berdasarkan jumlah pinjaman
			operasi.UrutPinjamanSelectionSort(&daftarPinjaman.Data, daftarPinjaman.N)
			fmt.Println("Daftar Pinjaman Terurut (Jumlah Pinjaman):")
			daftarPinjaman.TampilkanPinjaman()

		case 5:
			// Urutkan berdasarkan nama
			operasi.UrutPinjamanInsertionSort(&daftarPinjaman.Data, daftarPinjaman.N)
			fmt.Println("Daftar Pinjaman Terurut (Nama):")
			daftarPinjaman.TampilkanPinjaman()

		case 6:
			// Ubah data pinjaman
			ubah_hapus.UbahPinjaman(daftarPinjaman)

		case 7:
			// Hapus data pinjaman
			ubah_hapus.HapusPinjaman(daftarPinjaman,0)

		case 8:
			// Keluar dari program
			fmt.Println("Terima kasih!")
			return

		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}
