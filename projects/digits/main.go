package main

import (
	"fmt"
)

func main() {
	var input string
	fmt.Scanln(&input)

	maxDigit := '0' // Инициализируем максимальную цифру нулем

	// Проходим по всем символам строки
	for _, char := range input {
		if char > maxDigit {
			maxDigit = char // Обновляем максимальную цифру
		}
	}

	fmt.Println(string(maxDigit)) // Выводим максимальную цифру
}
