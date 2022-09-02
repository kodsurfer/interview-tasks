package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := []int{3, 4, 5, 6, 7}

	fmt.Printf("%v and %v intersection: %v \n", a, b, intersection(a, b))

	a = []int{1, 2, 1, 2, 1}
	b = []int{3, 2, 1}

	fmt.Printf("%v and %v intersection: %v \n", a, b, intersection(a, b))
}

func intersection(a, b []int) []int {
	cntr := make(map[int]int)
	res := make([]int, 0)

	for _, v := range a {
		if _, ok := cntr[v]; !ok {
			cntr[v] = 1
		} else {
			cntr[v]++
		}
	}

	for _, v := range b {
		if cnt, ok := cntr[v]; ok && cnt > 0 {
			cntr[v]--
			res = append(res, v)
		}
	}

	return res
}
