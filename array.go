package main

import "fmt"

func main() {
	//manual
	var names [3]string

	names[0] = "Syifa"
	names[1] = "Amelia"
	names[2] = "Latansa"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	//langsung
	var nilai = [3]int{
		34,
		45,
		56,
	}
	fmt.Println(nilai)

	//len
	fmt.Println(len(names))
	fmt.Println(len(nilai))

	var newArray [10]string
	fmt.Println(len(newArray))

}
