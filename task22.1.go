package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const a = 10

func numSearch(arrayTwo [a]int, n int) (answer string) {
	for i := 0; i < len(arrayTwo); i++ {
		if arrayTwo[i] == n {
			answer = strconv.Itoa(n) + " искомое число, " + strconv.Itoa(i+1) + " элемент в масссиве, " + strconv.Itoa(arrayTwo[len(arrayTwo)-i]) + " элементов до конца массива"
			break
		} else {
			answer = "0, искомое число не найдено."
		}
	}
	return
}

func main() {

	fmt.Println("22.4 Практическая работа")
	fmt.Println("")
	fmt.Println("Задание 1. Подсчёт чётных и нечётных чисел в массиве") // рандомное название =)
	fmt.Println("------------")

	var x int
	fmt.Println("Введите искомое число")
	fmt.Scan(&x)

	rand.Seed(time.Now().UnixNano())
	var array [a]int
	for i, _ := range array {
		array[i] = rand.Intn(len(array)-1) + 1
	}
	fmt.Println(array)
	fmt.Println(numSearch(array, x))
}
