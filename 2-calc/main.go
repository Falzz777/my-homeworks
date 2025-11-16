package main

import (
	"fmt"
	"sort"
)

func main() {
	transactions := []float64{}
	for {
		transaction := scanTransaction()
		if transaction == 0 {
			break
		}
		transactions = append(transactions, transaction)
	}
	userInput := userInputInfo()
	switch{
		case userInput == "AVG":
			valueAvg := calculateAvg(transactions)
			fmt.Printf("AVG ваших значений равна: %.1f", valueAvg)
		case userInput == "SUM":
			valueSum := calculateSum(transactions)
			fmt.Printf("SUM ваших значений равна: %.1f", valueSum)
		case userInput == "MED":
			valueMed := calculateMed(transactions)
			fmt.Printf("MED ваших значений равна: %.1f", valueMed)
	}
}

func scanTransaction() float64 {
	var transaction float64
	fmt.Println("Введите ваши значения (n для выхода): ")
	fmt.Scan(&transaction)
	return transaction
}


func userInputInfo()string{
	var userInfo string
	fmt.Println("Укажите что вы хотите сделать с этими значениями: AVG, SUM, MED")
	fmt.Scan(&userInfo)
	return userInfo
}

func calculateSum (transaction []float64) float64{
	var valueSum float64
	for _, value := range transaction {
		valueSum += value
	}
	return valueSum
}


func calculateAvg(transaction []float64) float64 {
	var valueAvg float64
	for _, value := range transaction {
		valueAvg += value
	}
	valueAvg = valueAvg / float64(len(transaction))
	return valueAvg
}

func calculateMed(transaction []float64) float64 {
	sort.Float64s(transaction)
	n := len(transaction)
	
	if n%2 == 1 {
		return transaction[n/2]
	} else {
		return (transaction[n/2-1] + transaction[n/2]) / 2
	}
}