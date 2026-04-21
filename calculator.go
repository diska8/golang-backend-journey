package main

import "fmt"

func main() {
	var AngkaPertama int
	var AngkaKedua int

	AngkaPertama = 20
	AngkaKedua = 4

	fmt.Println("Angka 1 = 20")
	fmt.Println("Angka 2 = 4")

	Penjumlahan := AngkaPertama + AngkaKedua
	Pengurangan := AngkaPertama - AngkaKedua
	Perkalian := AngkaPertama * AngkaKedua
	Pembagian := float64(AngkaPertama) / float64(AngkaKedua)

	fmt.Printf("Hasil dari Penjumlahan : %d\n", Penjumlahan)
	fmt.Printf("Hasil dari Pengurangan : %d\n", Pengurangan)
	fmt.Printf("Hasil dari Perkalian   : %d\n", Perkalian)
	fmt.Printf("Hasil dari Pembagian   : %.2f\n", Pembagian)
}
