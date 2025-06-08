package main

import (
	"fmt"
	"log"
	"math"
)

func main() {
	var weight float64
	fmt.Printf("Введите ваш вес (кг): ")
	if _, err := fmt.Scan(&weight); err != nil {
		log.Fatal("Ошибка ввода веса %s\n", err)
	}
	fmt.Println(weight)

	var heightCm float64
	fmt.Printf("Введите ваш рост (см): ")
	if _, err := fmt.Scan(&heightCm); err != nil {
		log.Fatal("Ошибка ввода роста %s\n", err)
	}

	bmi := weight / math.Pow(heightCm/100, 2)

	var category string
	switch {
	case bmi < 18.5:
		category = "Недостаточный вес"
	case bmi < 24.9:
		category = "Нормальный вес"
	case bmi < 29.9:
		category = "Избыточный вес"
	default:
		category = "Ожирение"
	}
	fmt.Printf("Ваш ИМТ: %.2f\n", bmi)
	fmt.Printf("Категория: %s\n", category)
}
