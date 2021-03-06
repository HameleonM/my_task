package main

import (
	"fmt"
	"math"
)

func max_multiplication(list []int) int {
	max1, max2 := 0, 0
	min1, min2 := 0, 0
	for _, i := range list {
		if i > max1 {
			max1 = i
			if max1 > max2 {
				max1, max2 = max2, max1
			}
		}
		if i < min1 {
			min1 = i
			if min1 < min2 {
				min1, min2 = min2, min1
			}
		}
		//fmt.Println(i, max1, max2, min1, min2)
	}
	task1 := float64(max1 * max2)
	task2 := float64(min1 * min2)
	//fmt.Println("Максимальное произведение двух целых чисел из списка предложенных:", math.Max(task1, task2))
	max := int(math.Max(task1, task2))
	return max
}

func main() {
	fmt.Println(max_multiplication([]int{1, 3, 8, 5, -8, -10}))
}

