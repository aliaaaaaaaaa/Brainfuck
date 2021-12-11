package parser

import (
	"Brainfuck/brainfuck/code_generator"
	"Brainfuck/brainfuck/reader"
	"io"
)

type Parser struct {
	reader *Reader.Reader
}

// New creates a new Parser.
func New(l io.Reader) *Parser {
	lexer := Reader.New(l)

	return &Parser{
		reader: lexer,
	}
}

// next returns the next token
func (p *Parser) next() Reader.Token {
	tok := <-p.reader.Tokens

	return tok
}

func (p *Parser) nextInst(tok Reader.Token) generator.Instruction {
	switch tok.Type {
	case Reader.IncTape:
		return generator.InstMoveHead{1}
	case Reader.DecTape:
		return generator.InstMoveHead{-1}
	case Reader.IncByte:
		return generator.InstAddToByte{1}
	case Reader.DecByte:
		return generator.InstAddToByte{-1}
	case Reader.WriteByte:
		return generator.InstWriteToOutput{}
	case Reader.StoreByte:
		return generator.InstReadFromInput{}
	case Reader.LoopEnter:
		return p.parseLoop()
	case Reader.LoopExit:
		return nil
	}
	panic("parser: unreachable")
}

func (p *Parser) parseLoop() generator.Instruction {
	insts := make([]generator.Instruction, 0)
	for tok := p.next(); tok.Type != Reader.EOF; tok = p.next() {
		i := p.nextInst(tok)
		if i == nil { // exit loop
			break
		}
		insts = append(insts, i)
	}
	return generator.InstLoop{insts}
}

func (p *Parser) Parse() []generator.Instruction {
	prog := make([]generator.Instruction, 0)
	for tok := p.next(); tok.Type != Reader.EOF; tok = p.next() {
		i := p.nextInst(tok)

		prog = append(prog, i)
	}
	return prog
}
