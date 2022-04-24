package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKTPku NoKTP = "3644783463924324"
	var status Married = true

	fmt.Println(noKTPku)
	fmt.Println(status)

}
