package operasi

import (
	"strings"
	"tubeskredit/pinjaman"
)

const MaksHasil = 100

// Pencarian Sequential (case-insensitive, menggunakan strings.Contains)
func CariPinjamanSequential(daftar *pinjaman.DaftarPinjaman, namaCari string, hasil *[MaksHasil]pinjaman.Pinjaman, n *int) {
	*n = 0
	for i := 0; i < daftar.N; i++ {
		p := daftar.Data[i]
		// Pencarian tidak case sensitive, gunakan strings.Contains
		if strings.Contains(strings.ToLower(p.NamaPeminjam), strings.ToLower(namaCari)) {
			hasil[*n] = p
			*n++
		}
	}
}

// Pencarian Binary Search (diasumsikan data sudah terurut)
func CariPinjamanBinary(daftar *pinjaman.DaftarPinjaman, namaCari string) int {
	kiri := 0
	kanan := daftar.N - 1
	ketemu := -1

	for kiri <= kanan && ketemu == -1 {
		tengah := (kiri + kanan) / 2
		if daftar.Data[tengah].NamaPeminjam == namaCari {
			ketemu = tengah
		} else if daftar.Data[tengah].NamaPeminjam < namaCari {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}
	return ketemu
}
