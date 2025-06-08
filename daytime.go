package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var input int
	fmt.Println("Введите время в часах (0-23): ")
	_, err := fmt.Scan(&input)

	if err != nil || input < 0 || input > 23 {
		log.Fatal("Неверно задано время")
		os.Exit(1)
	}

	var timeOfDay string
	switch {
	case input >= 6 && input < 12:
		timeOfDay = "утро"
	case input >= 12 && input < 18:
		timeOfDay = "день"
	case input >= 18 && input < 23:
		timeOfDay = "вечер"
	default:
		timeOfDay = "ночь"
	}
	fmt.Printf("Сейчас %02dч. - %s\n", input, timeOfDay)
}
