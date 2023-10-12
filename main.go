package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Menampilkan pesan instruksi kepada pengguna
	fmt.Println("Selamat datang di aplikasi simulasi mobil!")
	fmt.Println("Silakan pilih jenis roda untuk mobil Anda: [karet/kayu/besi]")

	// Membaca input dari pengguna
	reader := bufio.NewReader(os.Stdin)
	roda, _ := reader.ReadString('\n')

	// Membuat mobil dengan roda yang dipilih
	mobil := NewMobil(roda)

	// Menampilkan bunyi roda dan pintu saat diketuk dan dibuka
	mobil.TampilkanInfo()
}
