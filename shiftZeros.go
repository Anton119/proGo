// https://stepik.org/lesson/1134769/step/9?unit=1146392
package main

import "fmt"

func main() {
	arr := [5]int{}
	for i := 0; i < len(arr); i++ {
		fmt.Scan(&arr[i])
	}
	zero := 0
	tmp := 0
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] == zero {
				tmp = arr[j+1]
				arr[j+1] = arr[j]
				arr[j] = tmp
			}
		}
	}
	for _, el := range arr {
		fmt.Print(el, " ")
	}
}
