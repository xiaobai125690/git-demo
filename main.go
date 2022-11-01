package main

import (
	"fmt"
)

func main() {
	fmt.Println("123")
	//printslice[int]([]int{1, 2, 4, 3, 5})

	v := vector[int]{1, 2, 3}
	v1 := vector[string]{"01", "02", "03", "04"}
	printslice(v)
	printslice(v1)

	m := M[string, int]{"01": 1}
	m["01"] = 2

	c := make(C[int], 10)
	c <- 1

}

func printslice[T any](s []T) {
	for _, v := range s {
		fmt.Println("%v\n", v)
	}
}

type vector[T any] []T
type M[key string, T any] map[key]T
type C[T any]chan T
