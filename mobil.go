package main

import "fmt"

// Struktur untuk mobil
type Mobil struct {
	RodaDepanKiri     Roda
	RodaDepanKanan    Roda
	RodaBelakangKiri  Roda
	RodaBelakangKanan Roda
	PintuKiri         Pintu
	PintuKanan        Pintu
}

// Interface untuk roda
type Roda interface {
	Bunyi() string
}

// Struktur untuk ban karet
type BanKaret struct{}

func (bk BanKaret) Bunyi() string {
	return "Tok tok"
}

// Struktur untuk ban kayu
type BanKayu struct{}

func (bk BanKayu) Bunyi() string {
	return "Tak tak"
}

// Struktur untuk ban besi
type BanBesi struct{}

func (bb BanBesi) Bunyi() string {
	return "Dok dok"
}

// Interface untuk pintu
type Pintu interface {
	Ketuk() string
	Buka() string
}

// Struktur untuk pintu kiri
type PintuKiri struct{}

func (pk PintuKiri) Ketuk() string {
	return "Beep"
}

func (pk PintuKiri) Buka() string {
	return "Knock"
}

// Struktur untuk pintu kanan
type PintuKanan struct{}

func (pk PintuKanan) Ketuk() string {
	return "Knock"
}

func (pk PintuKanan) Buka() string {
	return "Beep"
}

// Fungsi untuk membuat mobil baru dengan roda yang dipilih
func NewMobil(roda string) Mobil {
	var mobil Mobil

	switch roda {
	case "karet\n":
		mobil.RodaDepanKiri = BanKaret{}
		mobil.RodaDepanKanan = BanKaret{}
		mobil.RodaBelakangKiri = BanKaret{}
		mobil.RodaBelakangKanan = BanKaret{}
	case "kayu\n":
		mobil.RodaDepanKiri = BanKayu{}
		mobil.RodaDepanKanan = BanKayu{}
		mobil.RodaBelakangKiri = BanKayu{}
		mobil.RodaBelakangKanan = BanKayu{}
	case "besi\n":
		mobil.RodaDepanKiri = BanBesi{}
		mobil.RodaDepanKanan = BanBesi{}
		mobil.RodaBelakangKiri = BanBesi{}
		mobil.RodaBelakangKanan = BanBesi{}
	default:
		fmt.Println("Jenis roda tidak valid.")
	}

	mobil.PintuKiri = PintuKiri{}
	mobil.PintuKanan = PintuKanan{}

	return mobil
}

// Fungsi untuk menampilkan informasi bunyi roda dan pintu saat diketuk dan dibuka
func (mobil Mobil) TampilkanInfo() {
	fmt.Println("Suara Roda Depan Kiri: " + mobil.RodaDepanKiri.Bunyi())
	fmt.Println("Suara Roda Depan Kanan: " + mobil.RodaDepanKanan.Bunyi())
	fmt.Println("Suara Roda Belakang Kiri: " + mobil.RodaBelakangKiri.Bunyi())
	fmt.Println("Suara Roda Belakang Kanan: " + mobil.RodaBelakangKanan.Bunyi())

	fmt.Println("Ketukan Pintu Kiri: " + mobil.PintuKiri.Ketuk())
	fmt.Println("Buka Pintu Kiri: " + mobil.PintuKiri.Buka())
	fmt.Println("Ketukan Pintu Kanan: " + mobil.PintuKanan.Ketuk())
	fmt.Println("Buka Pintu Kanan: " + mobil.PintuKanan.Buka())
}
