package main

import "fmt"

func main() {
	var ujian = 80
	var absensi = 70

	var lulusUjian = ujian >= 80
	var lulusAbsensi = absensi >= 80
	fmt.Println("lulus ujian = ", lulusUjian)
	fmt.Println("lulus absensi = ", lulusAbsensi)

	var lulus = lulusAbsensi && lulusUjian
	fmt.Println("hasil akhir = ", lulus)

	//singkat
	fmt.Println(ujian >= 80 && absensi >= 80)

}
