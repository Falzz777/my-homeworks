package main

import (
	"errors"
	"fmt"
)

func main() {
	for {
		amount := getAmount()
		err := checkError(amount)
		if err != nil {
			fmt.Println("Ошибка:", err)
			continue
		}
		fromCurrency := getFromCurrency()
		err = checkErrorValidFromCurrency(fromCurrency)
		if err != nil {
			fmt.Println("Ошибка:", err)
			continue
		}
		toCurrency := getToCurrency()
		err = checkErrorValidToCurrency(toCurrency)
		if err != nil {
			fmt.Println("Ошибка:", err)
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

func getToCurrency() string {
	var toCurrency string
	fmt.Print("Введите вашу целевую валюту(USD, EUR, RUB): ")
	fmt.Scan(&toCurrency)
	return toCurrency
}

func getAmount () float64 {
	var amount float64
	fmt.Print("Введите ваше значение: ")
	fmt.Scan(&amount)
	return amount
}	

func getFromCurrency() string{
	var fromCurrency string
	fmt.Print("Введите вашу исходную валюту(USD, EUR, RUB): ")
	fmt.Scan(&fromCurrency)
	return fromCurrency
}

func calculatUserInfo(amount float64, fromCurrency string, toCurrency string) float64 {
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

func checkRepeatCalculation() bool {
	var userChoise string
	fmt.Print("Вы хотите сделать еще расчёт (y/n)")
	fmt.Scan(&userChoise)
	if userChoise == "y" || userChoise == "Y" {
		return true
	}
	return false
}


func checkError(amount float64) error {
    // Проверка суммы
    if amount <= 0 {
        return errors.New("сумма должна быть больше 0")
	}
    return nil
}

func checkErrorValidFromCurrency(fromCurrency string) error {
    if !isValidCurrency(fromCurrency) {
        return errors.New("неверная исходная валюта (доступны: USD, EUR, RUB)")
    }
	return nil
}

func checkErrorValidToCurrency(toCurrency string) error {
    if !isValidCurrency(toCurrency) {
        return errors.New("неверная целевая валюта (доступны: USD, EUR, RUB)")
    }
	return nil
}

func isValidCurrency(currency string) bool {
    validCurrencies := []string{"USD", "EUR", "RUB"}
    for _, valid := range validCurrencies {
        if currency == valid {
            return true
        }
    }
    return false
}