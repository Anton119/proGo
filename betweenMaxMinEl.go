// https://stepik.org/lesson/1134769/step/7?unit=1146392
package main

import "fmt"

func main() {
	arr := [5]int{}
	var sum int
	for i := 0; i < len(arr); i++ {
		fmt.Scan(&arr[i])
	}
	maxInd := 0
	minInd := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[minInd] {
			minInd = i
		}
		if arr[i] > arr[maxInd] {
			maxInd = i
		}
	}
	if maxInd < minInd {
		maxInd, minInd = minInd, maxInd
	}

	for i := minInd + 1; i < maxInd; i++ {
		sum += arr[i]
	}
	fmt.Println(sum)
}
