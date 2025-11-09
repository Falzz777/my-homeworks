package main

import (
	"errors"
	"fmt"
)

func main() {
	for {
		amount, currency := getUserInput()
		err := checkError(amount, currency)
		if err != nil {
			fmt.Println("Введите правильные значения!!!")
			continue
		}
		isRepeateCalculation := checkRepeatCalculation()
		if !isRepeateCalculation{
			break
		}
	}
}

func getUserInput() (float64, string) {
	var amount float64
	var currency string
	fmt.Print("Введите ваше значение: ")
	fmt.Scan(&amount)
	fmt.Print("Введите вашу валюту: ")
	fmt.Scan(&currency)
	return amount, currency
}

// func calculatUserInfo(amount float64, fromCurrency string, toCurrency string) (result float64) {

// }

func checkRepeatCalculation() (bool) {
	var userChoise string
	fmt.Print("Вы хотите сделать еще расчёт (y/n)")
	fmt.Scan(&userChoise)
	if userChoise == "y" || userChoise == "Y" {
		return true
	}
	return false
}


func checkError (amount float64, currency string) (error) {
	if amount <= 0 || currency == "" {
		return errors.New("")
	}
	return nil
}