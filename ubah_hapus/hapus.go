package ubah_hapus

import (
	"fmt"
	"tubeskredit/pinjaman"
)

func HapusPinjaman(daftar *pinjaman.DaftarPinjaman, nomor int) {
	// Tampilkan daftar peminjam
	daftar.TampilkanPinjaman()

	// Minta input nomor pinjaman yang ingin dihapus
	fmt.Print("\nMasukkan nomor pinjaman yang ingin dihapus: ")
	fmt.Scan(&nomor)

	// Cari index pinjaman dengan nomor yang sesuai
	index := -1
	i := 0
	for i < daftar.N {
		if daftar.Data[i].Nomor == nomor {
			index = i
			i = daftar.N // keluar loop
		} else {
			i++
		}
	}

	if index == -1 {
		fmt.Println("Pinjaman dengan nomor tersebut tidak ditemukan.")
		return
	}

	// Hapus data pinjaman
	for i := index; i < daftar.N-1; i++ {
		daftar.Data[i] = daftar.Data[i+1]
	}
	daftar.N--

	fmt.Println("Data pinjaman berhasil dihapus!")
}
