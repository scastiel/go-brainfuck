package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"

	brainfuck "github.com/scastiel/go-brainfuck/lib/brainfuck"
	memory "github.com/scastiel/go-brainfuck/lib/memory"
)

func main() {
	code, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal("Can't read input file")
	}

	memory := memory.NewMemory()
	out := bufio.NewWriterSize(os.Stdout, 1)
	in := bufio.NewReaderSize(os.Stdin, 1)

	brainfuck.Execute(string(code), 0, &memory, out, in)
	out.Flush()
}
