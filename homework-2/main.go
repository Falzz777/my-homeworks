package main

import (
	"fmt"
	"reflect"
)

func main() {

}

func getUserInput() (float64) {
	var userInput float64
	fmt.Print("Введите ваше значение")
	fmt.Scan(&userInput)
	return userInput
}

func calculatUserInfo(amount float64, fromCurrency string, toCurrency string) (result float64) {

}