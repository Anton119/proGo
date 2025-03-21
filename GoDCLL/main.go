package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 1 2 4
// 2 3 4
func main() {
	scnr := bufio.NewScanner(os.Stdin)
	scnr.Scan()
	strInput := strings.Split(scnr.Text(), " ")

	firstList := list.New()

	for _, v := range strInput {
		if v == "" {
			continue
		}
		iVal, _ := strconv.Atoi(v)
		firstList.PushBack(iVal)
	}

	scnr.Scan()
	strInput2 := strings.Split(scnr.Text(), " ")

	secondList := list.New()

	for _, v := range strInput2 {
		if v == "" {
			continue
		}
		iVal2, _ := strconv.Atoi(v)
		secondList.PushBack(iVal2)
	}

	finalList := firstList

	for el := secondList.Front(); el != nil; el = el.Next() {
		finalList.PushBack(el.Value)
	}

	for end := finalList.Back(); end != finalList.Front(); end = end.Prev() {
		for el := finalList.Front(); el != nil && el.Next() != nil; el = el.Next() {
			v1 := el.Value.(int)
			v2 := el.Next().Value.(int)

			if v1 > v2 {
				el.Value, el.Next().Value = v2, v1
			}
		}
	}

	for el := finalList.Front(); el != nil; el = el.Next() {
		fmt.Print(el.Value, " ")
	}

}
