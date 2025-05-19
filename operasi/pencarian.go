package operasi

import (
    "tubeskredit/pinjaman"
    "strings"
)

// Pencarian Sequential
func CariPinjamanSequential(daftar *pinjaman.DaftarPinjaman, namaCari string) []pinjaman.Pinjaman {
    hasilPencarian := []pinjaman.Pinjaman{}
    
    for _, p := range daftar.Data {
        // Pencarian tidak Case Sensitive
        if strings.Contains(strings.ToLower(p.NamaPeminjam), strings.ToLower(namaCari)) {
            hasilPencarian = append(hasilPencarian, p)
        }
    }
    
    return hasilPencarian
}

// Pencarian Binary Search (diasumsikan data sudah terurut)
func CariPinjamanBinary(daftar []pinjaman.Pinjaman, namaCari string) *pinjaman.Pinjaman {
    kiri := 0
    kanan := len(daftar) - 1

    for kiri <= kanan {
        tengah := (kiri + kanan) / 2
        
        // Bandingkan nama
        if daftar[tengah].NamaPeminjam == namaCari {
            return &daftar[tengah]
        }
        
        // Tentukan arah pencarian
        if daftar[tengah].NamaPeminjam < namaCari {
            kiri = tengah + 1
        } else {
            kanan = tengah - 1
        }
    }
    
    return nil
}