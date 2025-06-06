package pinjaman

import "fmt"

const MaksPinjaman = 100

type Pinjaman struct {
	Nomor            int
	NamaPeminjam     string
	JumlahPinjaman   float64
	LamaPinjaman     int
	SukuBunga        float64
	StatusPembayaran string
}

type DaftarPinjaman struct {
	Data [MaksPinjaman]Pinjaman
	N    int
}

func (dp *DaftarPinjaman) TambahPinjaman(p Pinjaman) {
	if dp.N < MaksPinjaman {
		p.StatusPembayaran = "Belum Lunas"
		p.Nomor = dp.N + 1
		dp.Data[dp.N] = p
		dp.N++
		fmt.Println("Pinjaman berhasil ditambahkan!")
	} else {
		fmt.Println("Kapasitas pinjaman penuh!")
	}
}

func (dp *DaftarPinjaman) TampilkanPinjaman() {
	fmt.Println("Daftar Pinjaman:")
	for i := 0; i < dp.N; i++ {
		pinjaman := dp.Data[i]
		fmt.Printf("Nomor: %d | Nama: %s | Pinjaman: Rp %.2f | Status: %s\n",
			pinjaman.Nomor,
			pinjaman.NamaPeminjam,
			pinjaman.JumlahPinjaman,
			pinjaman.StatusPembayaran)
	}
}