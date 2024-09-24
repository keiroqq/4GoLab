package main

import (
	"fmt"
	"strings"
)

// Функция для вычисления среднего возраста
func averageAge(people map[string]int) float64 {
	if len(people) == 0 {
		return 0
	}

	var total int
	for _, age := range people {
		total += age
	}
	return float64(total) / float64(len(people))
}

// Функция для удаления записи по имени
func deletePerson(people map[string]int, name string) {
	delete(people, name)
}

func main() {
	// Инициализация карты с именами и возрастами
	people := make(map[string]int)
	people["Максим"] = 20
	people["Данила"] = 27
	people["Иван"] = 23

	// Добавление нового человека
	people["Леха"] = 26

	// Вывод всех записей
	fmt.Println("Записи:")
	for name, age := range people {
		fmt.Printf("%s: %d\n", name, age)
	}

	// Вычисление и вывод среднего возраста
	avgAge := averageAge(people)
	fmt.Printf("Средний возраст: %.2f\n", avgAge)

	// Удаление записи по имени
	var nameToDelete string
	fmt.Print("Введите имя для удаления: ")
	fmt.Scan(&nameToDelete)
	deletePerson(people, nameToDelete)

	// Вывод всех записей после удаления
	fmt.Println("Записи после удаления:")
	for name, age := range people {
		fmt.Printf("%s: %d\n", name, age)
	}

	// Считывание строки и вывод в верхнем регистре
	var inputString string
	fmt.Print("Введите строку: ")
	fmt.Scan(&inputString)
	fmt.Println("Верхний регистр:", strings.ToUpper(inputString))

	// Считывание нескольких чисел и вывод суммы
	var numbers []int
	var num int
	fmt.Println("Введите числа (введите '0' для завершения ввода):")
	for {
		fmt.Scan(&num)
		if num == 0 {
			break
		}
		numbers = append(numbers, num)
	}

	// Вычисление суммы
	totalSum := 0
	for _, n := range numbers {
		totalSum += n
	}
	fmt.Println("Сумма введенных чисел:", totalSum)

	// Считывание массива целых чисел и вывод в обратном порядке
	var arraySize int
	fmt.Print("Введите размер массива: ")
	fmt.Scan(&arraySize)

	arr := make([]int, arraySize)
	fmt.Println("Введите элементы массива:")
	for i := 0; i < arraySize; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Println("Массив в обратном порядке:")
	for i := arraySize - 1; i >= 0; i-- {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}