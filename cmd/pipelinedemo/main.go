package main

import (
	"goPractice/pipeline"
	"fmt"
)

func main() {
	arr1 := pipeline.ArraySource(3, 2, 6, 7, 4)
	arr2 := pipeline.ArraySource(3, 2, 7, 4, 0, 18, 3, 13)
	p := pipeline.Merge(
		pipeline.InMemSort(arr1),
		pipeline.InMemSort(arr2),
		)
	for v := range p {
		fmt.Println(v)
	}
}
