package main

import (
	"bufio"
	"fmt"
	"os"

	"xiongtaigo@live.com/gointro/pipeline"
)

func main() {
	const filename = "large.in"
	const n = 100000000
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	p := pipeline.RandomSource(n)
	writer := bufio.NewWriter(file)
	pipeline.WriteSink(writer, p)
	writer.Flush()

	file, err = os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	p = pipeline.ReaderSource(reader,-1)
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
	p := pipeline.InMemSort(pipeline.ArraySource(3, 2, 6, 7, 4))
	for {
		if num, ok := <-p; ok {
			fmt.Println(num)
		} else {
			break
		}

	}
}
