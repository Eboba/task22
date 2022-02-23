package main

import "fmt"

const a = 12

func indexSearch(array [a]int, n int) int {

	var (
		max    = a - 1
		min    int
		middle int
		index  = -1
	)

	for max > min { // ищет число в массиве
		middle = (max + min) / 2
		if array[middle] == n {
			index = middle
			break
		} else if array[middle] > n {
			max = middle //пришлось убрать +- 1, чтоб не было слепых зон
		} else {
			min = middle
		}
	}

	for { // ищем крайнее число
		if middle != 0 && array[middle] == array[middle-1] {
			middle--
			continue
		}
		index = middle
		break
	}
	return index
}

func main() {

	fmt.Println("22.4 Практическая работа")
	fmt.Println("")
	fmt.Println("Задание 2. Нахождение первого вхождения числа в упорядоченном массиве.")
	fmt.Println("------------")

	var x int
	fmt.Println("Введите искомое число:")
	fmt.Scan(&x)

	array := [a]int{1, 2, 2, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(array)
	fmt.Println(indexSearch(array, x))
}
