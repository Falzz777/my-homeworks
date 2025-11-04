package main

import "fmt"

func main() {
	const usdInEur float64 = 0.87
	const usdInRub float64 = 80.9
	const eurInRub float64 = usdInEur / usdInRub
	fmt.Println(eurInRub)
}