package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"sync"

	brainfuck "github.com/scastiel/go-brainfuck/lib/brainfuck"
	memory "github.com/scastiel/go-brainfuck/lib/memory"
)

func main() {
	memory := memory.NewMemory()
	out := bufio.NewWriterSize(os.Stdout, 1)
	in := bufio.NewReaderSize(os.Stdin, 1)

	var wg sync.WaitGroup
	for i, file := range os.Args[1:] {
		wg.Add(1)
		go func(i int) {
			code, err := ioutil.ReadFile(file)
			if err != nil {
				log.Fatal("Can't read input file")
			}
			brainfuck.Execute(string(code), i*1000, &memory, out, in)
			wg.Done()
		}(i)
	}
	wg.Wait()

	out.Flush()
}
