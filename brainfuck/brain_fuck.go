package brainfuck

import (
	generator "Brainfuck/brainfuck/code_generator"
	"Brainfuck/brainfuck/parser"
	"bufio"
	"io"
	"os"
)

type byteWriterFlusher struct {
	w *bufio.Writer
}

func (bw byteWriterFlusher) WriteByte(b byte) error {
	err := bw.w.WriteByte(b)
	bw.w.Flush()
	return err
}
func Start(reader io.Reader) {
	parser := parser.New(reader)
	instructions := parser.Parse()
	tape := generator.New()
	in := bufio.NewReader(os.Stdin)
	out := byteWriterFlusher{bufio.NewWriter(os.Stdin)}
	for _, i := range instructions {
		i.Eval(tape, in, out)
	}
}
