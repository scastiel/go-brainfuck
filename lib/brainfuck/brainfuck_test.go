package brainfuck_test

import (
	"bytes"
	"io/ioutil"
	"log"
	"sync"
	"testing"

	"github.com/scastiel/go-brainfuck/lib/brainfuck"
	"github.com/scastiel/go-brainfuck/lib/memory"
)

func TestProgramWithOutput(t *testing.T) {
	code := readFromFile("../../examples/hello.bf")
	memory := memory.NewMemory()

	out := bytes.NewBufferString("")
	in := bytes.NewBufferString("")
	brainfuck.Execute(code, 0, &memory, out, in)

	written := out.String()
	expected := "Hello World!\n"
	if written != expected {
		t.Fatalf("Wrong output, received %q, expected %q", written, expected)
	}
}

func TestProgramWithInputAndOutput(t *testing.T) {
	code := readFromFile("../../examples/echo.bf")
	memory := memory.NewMemory()

	out := bytes.NewBufferString("")
	in := bytes.NewBufferString("Good morning sir!")
	brainfuck.Execute(code, 0, &memory, out, in)

	written := out.String()
	expected := "Good morning sir!"
	if written != expected {
		t.Fatalf("Wrong output, received %q, expected %q", written, expected)
	}

	remaining := in.String()
	expected = ""
	if remaining != expected {
		t.Fatalf("Wrong remaining input, received %q, expected %q", remaining, expected)
	}
}

func TestSeveralProgramsWithOutputInParallel(t *testing.T) {
	code := readFromFile("../../examples/hello.bf")
	memory := memory.NewMemory()
	const n = 1000

	var outs, ins [n]*bytes.Buffer
	for i := 0; i < n; i++ {
		outs[i] = bytes.NewBufferString("")
		ins[i] = bytes.NewBufferString("")
	}

	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			brainfuck.Execute(code, i*1000, &memory, outs[i], ins[i])
			wg.Done()
		}(i)
	}
	wg.Wait()

	expected := "Hello World!\n"
	for i := 0; i < n; i++ {
		written := outs[i].String()
		if written != expected {
			t.Fatalf("Wrong output for program %v, received %q, expected %q", i, written, expected)
		}
	}
}

func TestSeveralProgramsWithInputAndOutputInParallel(t *testing.T) {
	code := readFromFile("../../examples/echo.bf")
	memory := memory.NewMemory()
	const n = 1000

	var outs, ins [n]*bytes.Buffer
	for i := 0; i < n; i++ {
		outs[i] = bytes.NewBufferString("")
		ins[i] = bytes.NewBufferString("Good morning sir!")
	}

	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			brainfuck.Execute(code, i*1000, &memory, outs[i], ins[i])
			wg.Done()
		}(i)
	}
	wg.Wait()

	expected := "Good morning sir!"

	for i := 0; i < n; i++ {
		written := outs[i].String()
		if written != expected {
			t.Fatalf("Wrong output for program %v, received %q, expected %q", i, written, expected)
		}
	}

	expected = ""
	for i := 0; i < n; i++ {
		remaining := ins[i].String()
		if remaining != expected {
			t.Fatalf("Wrong remaining input, received %q, expected %q", remaining, expected)
		}
	}
}

func readFromFile(path string) string {
	code, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Can't read input file")
	}
	return string(code)
}
