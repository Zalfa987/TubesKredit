package ubah_hapus

import (
	"fmt"
	"tubeskredit/pinjaman"
)

func UbahPinjaman(daftar *pinjaman.DaftarPinjaman) {
	daftar.TampilkanPinjaman()

	var nomor int
	fmt.Print("\nMasukkan nomor peminjam yang ingin diubah: ")
	fmt.Scan(&nomor)

	index := -1
	i := 0
	for i < daftar.N {
		if daftar.Data[i].Nomor == nomor {
			index = i
			i = daftar.N
		} else {
			i++
		}
	}

	if index == -1 {
		fmt.Println("Pinjaman dengan nomor tersebut tidak ditemukan.")
		return
	}

	fmt.Println("\nBagian mana yang ingin diubah?")
	fmt.Println("1. Nama Peminjam")
	fmt.Println("2. Jumlah Pinjaman")
	fmt.Println("3. Lama Pinjaman")
	fmt.Println("4. Suku Bunga")
	fmt.Println("5. Status Pembayaran")
	fmt.Print("Pilih menu: ")
	var menuEdit int
	fmt.Scan(&menuEdit)

	switch menuEdit {
	case 1:
		fmt.Print("Nama Peminjam Baru: ")
		fmt.Scan(&daftar.Data[index].NamaPeminjam)
	case 2:
		fmt.Print("Jumlah Pinjaman Baru: ")
		fmt.Scan(&daftar.Data[index].JumlahPinjaman)
	case 3:
		fmt.Print("Lama Pinjaman Baru (bulan): ")
		fmt.Scan(&daftar.Data[index].LamaPinjaman)
	case 4:
		fmt.Print("Suku Bunga Baru (%): ")
		fmt.Scan(&daftar.Data[index].SukuBunga)
	case 5:
		fmt.Println("Pilih Status Pembayaran:")
		fmt.Println("1. Lunas")
		fmt.Println("2. Belum Lunas")
		fmt.Print("Masukkan pilihan: ")
		var statusPembayaran int
		fmt.Scan(&statusPembayaran)
		if statusPembayaran == 1 {
			daftar.Data[index].StatusPembayaran = "Lunas"
		} else {
			daftar.Data[index].StatusPembayaran = "Belum Lunas"
		}
	default:
		fmt.Println("Pilihan tidak valid!")
		return
	}

	fmt.Println("Data pinjaman berhasil diubah!")
}
