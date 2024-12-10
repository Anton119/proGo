// https://stepik.org/lesson/1134769/step/8?unit=1146392
package main

import "fmt"

func main() {
	arr := [5]int{} //{1, 0, 1, 1, 1}
	repeat := []int{}
	var count int
	for i := 0; i < len(arr); i++ {
		fmt.Scan(&arr[i])
	}

	for _, num := range arr {
		if num == 1 {
			count += 1
		} else {
			count = 0
		}
		repeat = append(repeat, count)
	}
	maxInd := 0
	for i := 1; i < len(repeat); i++ {
		if repeat[i] > maxInd {
			maxInd = i
		}
	}
	fmt.Println(repeat[maxInd])

}
