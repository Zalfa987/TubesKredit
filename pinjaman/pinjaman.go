package pinjaman

import "fmt"

const MaksPinjaman = 100 // Tambahkan ini

type Pinjaman struct {
	Nomor          int
	NamaPeminjam   string
	JumlahPinjaman float64
	LamaPinjaman   int
	SukuBunga      float64
}

// Struktur untuk menyimpan daftar pinjaman
type DaftarPinjaman struct {
	Data [MaksPinjaman]Pinjaman
	N    int // Tambahkan ini
}

// Metode untuk menambah pinjaman
func (dp *DaftarPinjaman) TambahPinjaman(p Pinjaman) {
	if dp.N < MaksPinjaman {
		p.Nomor = dp.N + 1
		dp.Data[dp.N] = p
		dp.N++
		fmt.Println("Pinjaman berhasil ditambahkan!")
	} else {
		fmt.Println("Kapasitas pinjaman penuh!")
	}
}

// Metode untuk menampilkan semua pinjaman
func (dp *DaftarPinjaman) TampilkanPinjaman() {
	fmt.Println("Daftar Pinjaman:")
	for i := 0; i < dp.N; i++ {
		pinjaman := dp.Data[i]
		fmt.Printf("Nomor: %d | Nama: %s | Pinjaman: Rp %.2f\n",
			pinjaman.Nomor,
			pinjaman.NamaPeminjam,
			pinjaman.JumlahPinjaman)
	}
}
