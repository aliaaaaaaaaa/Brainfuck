package main

import (
	"Brainfuck/code_generator"
	"Brainfuck/parser"
	"Brainfuck/reader"
	"bufio"
	"io"
	"os"
)

var file io.Reader
var err error

type byteWriterFlusher struct {
	w *bufio.Writer
}

func (bw byteWriterFlusher) WriteByte(b byte) error {
	err := bw.w.WriteByte(b)
	bw.w.Flush()
	return err
}
func main() {
	file, err = os.Open("brainfuck-test.bf")
	lexer := Reader.New(file)
	parser := parser.New(lexer)
	instructions := parser.Parse()
	tape := branfuck.New()
	in := bufio.NewReader(os.Stdin)
	out := byteWriterFlusher{bufio.NewWriter(os.Stdin)}
	for _, i := range instructions {
		i.Eval(tape, in, out)
	}
}
