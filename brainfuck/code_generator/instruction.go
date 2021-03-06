package generator

import (
	"fmt"
	"io"
)

type Instruction interface {
	Eval(t Tape, in io.ByteReader, out io.ByteWriter)

	String() string
}

type InstMoveHead struct {
	V int
}

func (i InstMoveHead) Eval(t Tape, in io.ByteReader, out io.ByteWriter) {
	t.MoveHead(i.V)
}

func (i InstMoveHead) String() string {
	return fmt.Sprintf("InstMoveHead{%d}", i.V)
}

type InstAddToByte struct {
	V int
}

func (i InstAddToByte) Eval(t Tape, in io.ByteReader, out io.ByteWriter) {
	t.AddToByte(i.V)
}

func (i InstAddToByte) String() string {
	return fmt.Sprintf("InstAddToByte{%d}", i.V)
}

type InstWriteToOutput struct{}

func (i InstWriteToOutput) Eval(t Tape, in io.ByteReader, out io.ByteWriter) {
	b := t.GetByte()
	out.WriteByte(b)
}

func (i InstWriteToOutput) String() string {
	return "InstWriteToOutput"
}

type InstReadFromInput struct{}

func (i InstReadFromInput) Eval(t Tape, in io.ByteReader, out io.ByteWriter) {
	b, _ := in.ReadByte()
	if b == byte(0) {
		return
	}
	t.SetByte(b)
}

func (i InstReadFromInput) String() string {
	return "InstReadFromInput"
}

type InstLoop struct {
	Insts []Instruction
}

func (i InstLoop) Eval(t Tape, in io.ByteReader, out io.ByteWriter) {
	for {
		if t.GetByte() == byte(0) {
			break
		}

		for _, ii := range i.Insts {
			ii.Eval(t, in, out)
		}

		if t.GetByte() == byte(0) {
			break
		}

	}
}

func (i InstLoop) String() string {
	s := "InstLoop\n"
	for _, ii := range i.Insts {
		s += fmt.Sprintf("%s\n", ii)
	}
	return s
}
