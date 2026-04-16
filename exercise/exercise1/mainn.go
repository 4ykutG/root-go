package main

import "fmt"

var i int
var durum string

func main() {
	for i = 0; i <= 100; i++ {
		kontrolEt()
		fmt.Printf("sayi: %d, durum: %s\n", i, durum)
	}
}

func kontrolEt() {
	if i%2 == 0 {
		durum = "cift"
	} else {
		durum = "tek"
	}
}
