package main

import "fmt"

func validMountArr(arr [7]int) bool {
	strictlyInc := false
	strictlyDec := false

	for i := 0; i < len(arr)-1; i++ {
		if arr[i+1] < arr[i] {
			if !strictlyInc {
				return false
			}
			strictlyDec = true
		} else if arr[i+1] > arr[i] {
			if strictlyDec {
				return false
			}
			strictlyInc = true
		} else {
			return false
		}
	}

	if strictlyDec && strictlyInc {
		return true
	} else {
		return false
	}
	return true
}

func main() {
	arr := [7]int{} //{1, 5, 5, 15, 9, 6, 2}
	for i := 0; i < len(arr); i++ {
		fmt.Scan(&arr[i])
	}
	if validMountArr(arr) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
