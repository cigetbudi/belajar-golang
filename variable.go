package main

import "fmt"

func main() {
	//cara klasik
	var name string

	name = "Sigit Budi"
	fmt.Println(name)

	name = "Syifa Amelia"
	fmt.Println(name)

	//cara singkat
	var nama = "Jancuk"
	fmt.Println(nama)

	var umur = 25
	fmt.Println(umur)

	//cara lebih singkat :=
	negara := "indonesia"
	fmt.Println(negara)

	negara = "malaysia"
	fmt.Println(negara)

	//multiple variable
	var (
		namaDepan    = "Temon"
		namaBelakang = "Anjeng"
	)

	fmt.Println(namaDepan)
	fmt.Println(namaBelakang)

}
