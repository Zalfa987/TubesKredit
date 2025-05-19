package operasi

import (
    "tubeskredit/pinjaman"
)

// Pengurutan Selection Sort berdasarkan jumlah pinjaman
func UrutPinjamanSelectionSort(daftar []pinjaman.Pinjaman) []pinjaman.Pinjaman {
    n := len(daftar)
    
    for i := 0; i < n-1; i++ {
        indeksMaks := i
        
        for j := i + 1; j < n; j++ {
            if daftar[j].JumlahPinjaman > daftar[indeksMaks].JumlahPinjaman {
                indeksMaks = j
            }
        }
        
        // Tukar
        daftar[i], daftar[indeksMaks] = daftar[indeksMaks], daftar[i]
    }
    
    return daftar
}

// Pengurutan Insertion Sort berdasarkan nama
func UrutPinjamanInsertionSort(daftar []pinjaman.Pinjaman) []pinjaman.Pinjaman {
    n := len(daftar)
    
    for i := 1; i < n; i++ {
        kunci := daftar[i]
        j := i - 1
        
        for j >= 0 && daftar[j].NamaPeminjam > kunci.NamaPeminjam {
            daftar[j+1] = daftar[j]
            j = j - 1
        }
        
        daftar[j+1] = kunci
    }
    
    return daftar
}