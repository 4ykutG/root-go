package main

import "fmt"

var i int
var durum string

func main() {
	for i = 1; i <= 10; i++ {
		kontrolEt()
		fmt.Printf("Sayi: %d, durum: %s\n", i, durum)
	}
}

func kontrolEt() {
	if i%2 == 0 {
		durum = "cift"
	} else {
		durum = "tek"
	}

}
