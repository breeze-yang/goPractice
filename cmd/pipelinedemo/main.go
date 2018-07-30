package main

import (
	"goPractice/pipeline"
	"fmt"
	"os"
	"bufio"
)

func main(){
	const filename = "large.in"
	const n = 100000000
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	p := pipeline.RandomSource(n)
	writer := bufio.NewWriter(file)
	pipeline.WriterSink(writer, p)
	writer.Flush()

	file, err = os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p = pipeline.ReaderSource(bufio.NewReader(file), -1)
	count := 0
	for v := range p {
		fmt.Println(v)
		count++
		if count >= 100 {
			break
		}
	}
}

func mergeDemo() {
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
