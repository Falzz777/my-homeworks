package main

import (
	"errors"
	"fmt"
)

func main() {
	for {
		amount, fromCurrency, toCurrency := getUserInput()
		err := checkError(amount, fromCurrency, toCurrency)
		if err != nil {
			fmt.Println("Введите правильные значения!!!")
			continue
		}
		result := calculatUserInfo(amount, fromCurrency, toCurrency)
		fmt.Printf("%.1f %s = %.1f %s\n", amount, fromCurrency, result, toCurrency)
		isRepeateCalculation := checkRepeatCalculation()
		if !isRepeateCalculation{
			break
		}
	}
}

func getUserInput() (float64, string,  string) {
	var amount float64
	var fromCurrency string
	var toCurrency string
	fmt.Print("Введите вашу исходную валюту(USD, EUR, RUB): ")
	fmt.Scan(&fromCurrency)
	fmt.Print("Введите ваше значение: ")
	fmt.Scan(&amount)
	fmt.Print("Введите вашу целевую валюту(USD, EUR, RUB): ")
	fmt.Scan(&toCurrency)
	return amount, fromCurrency, toCurrency
}

func calculatUserInfo(amount float64, fromCurrency string, toCurrency string) (float64) {
	switch {
		case fromCurrency == "USD" && toCurrency == "EUR":
			return amount * 0.85
		case fromCurrency == "EUR" && toCurrency == "USD":
			return amount * 1.18
		case fromCurrency == "RUB" && toCurrency == "USD":
			return amount * 0.012
		case fromCurrency == "RUB" && toCurrency == "EUR":
			return amount * 0.011
		case fromCurrency == "EUR" && toCurrency == "RUB":
			return amount * 93.69
		default:
			return amount * 81.01
 	}
}

func checkRepeatCalculation() (bool) {
	var userChoise string
	fmt.Print("Вы хотите сделать еще расчёт (y/n)")
	fmt.Scan(&userChoise)
	if userChoise == "y" || userChoise == "Y" {
		return true
	}
	return false
}


func checkError (amount float64, fromCurrency string, toCurrency string) (error) {
	if amount <= 0 || fromCurrency == "" || toCurrency == ""{
		return errors.New("")
	}
	return nil
}