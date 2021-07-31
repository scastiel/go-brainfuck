package brainfuck

import (
	"io"

	"github.com/scastiel/go-brainfuck/lib/memory"
)

func Execute(code string, position int, memory *memory.Memory, out io.ByteWriter, in io.ByteReader) error {
	i := 0
	for i < len(code) {
		switch code[i] {
		case '>':
			position++
		case '<':
			position--
		case '+':
			memory.Inc(position)
		case '-':
			memory.Dec(position)
		case '.':
			b := memory.Read(position)
			err := out.WriteByte(b)
			if err != nil {
				return err
			}
		case ',':
			byte, err := in.ReadByte()
			if err != nil {
				return err
			}
			memory.Write(position, byte)
		case '[':
			if memory.Read(position) == 0 {
				brakets := 1
				for {
					i++
					if code[i] == '[' {
						brakets++
					}
					if code[i] == ']' {
						brakets--
						if brakets == 0 {
							break
						}
					}
				}
			}
		case ']':
			if memory.Read(position) != 0 {
				brakets := 1
				for {
					i--
					if code[i] == ']' {
						brakets++
					}
					if code[i] == '[' {
						brakets--
						if brakets == 0 {
							break
						}
					}
				}
			}
		}
		i++
	}
	return nil
}
