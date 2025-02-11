/*
3
1 2 3
3
4 5 6

*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	// put your code here
	/*var n int
	fmt.Scan(&n)

	a := []int{}
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	var m int
	fmt.Scan(&m)

	b := []int{}
	for i := 0; i < n; i++ {
		fmt.Scan(&b[i])
	}*/

	//n := 3
	//m := 3
	a := []int{1, 23, 100}
	b := []int{1, 3, 37}

	c := []int{}

	for i := 0; i < len(a); i++ {
		c = append(c, a[i])
	}
	for i := 0; i < len(b); i++ {
		c = append(c, b[i])
	}
	sort.Ints(c)
	fmt.Println(c)
}
