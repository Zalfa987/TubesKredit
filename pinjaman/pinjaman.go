package pinjaman

import "fmt"

type Pinjaman struct {
    Nomor           int
    NamaPeminjam    string
    JumlahPinjaman  float64
    LamaPinjaman    int
    SukuBunga       float64
}

// Struktur untuk menyimpan daftar pinjaman
type DaftarPinjaman struct {
    Data []Pinjaman
}

// Metode untuk menambah pinjaman
func (dp *DaftarPinjaman) TambahPinjaman(p Pinjaman) {
    p.Nomor = len(dp.Data) + 1
    dp.Data = append(dp.Data, p)
    println("Pinjaman berhasil ditambahkan!")
}

// Metode untuk menampilkan semua pinjaman
func (dp *DaftarPinjaman) TampilkanPinjaman() {
    println("Daftar Pinjaman:")
    for _, pinjaman := range dp.Data {
        fmt.Printf("Nomor: %d | Nama: %s | Pinjaman: Rp %.2f\n", 
            pinjaman.Nomor, 
            pinjaman.NamaPeminjam, 
            pinjaman.JumlahPinjaman)
    }
}