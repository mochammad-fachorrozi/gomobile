package main

import "fmt"

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

// Struktur untuk mobil
type Mobil struct {
	RodaDepanKiri     Roda
	RodaDepanKanan    Roda
	RodaBelakangKiri  Roda
	RodaBelakangKanan Roda
	PintuKiri         Pintu
	PintuKanan        Pintu
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

func main() {
	// Inisialisasi ban
	banKaret := BanKaret{}

	// Inisialisasi pintu
	pintuKiri := PintuKiri{}
	pintuKanan := PintuKanan{}

	// Inisialisasi mobil dengan roda karet
	mobil := Mobil{
		RodaDepanKiri:     banKaret,
		RodaDepanKanan:    banKaret,
		RodaBelakangKiri:  banKaret,
		RodaBelakangKanan: banKaret,
		PintuKiri:         pintuKiri,
		PintuKanan:        pintuKanan,
	}

	// Output suara roda saat diketuk
	fmt.Println("Suara Roda Depan Kiri: " + mobil.RodaDepanKiri.Bunyi())
	fmt.Println("Suara Roda Depan Kanan: " + mobil.RodaDepanKanan.Bunyi())
	fmt.Println("Suara Roda Belakang Kiri: " + mobil.RodaBelakangKiri.Bunyi())
	fmt.Println("Suara Roda Belakang Kanan: " + mobil.RodaBelakangKanan.Bunyi())

	// Output bunyi pintu saat diketuk dan dibuka
	fmt.Println("Ketukan Pintu Kiri: " + mobil.PintuKiri.Ketuk())
	fmt.Println("Buka Pintu Kiri: " + mobil.PintuKiri.Buka())
	fmt.Println("Ketukan Pintu Kanan: " + mobil.PintuKanan.Ketuk())
	fmt.Println("Buka Pintu Kanan: " + mobil.PintuKanan.Buka())
}
