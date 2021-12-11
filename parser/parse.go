package parser

import (
	"Brainfuck/code_generator"
	"Brainfuck/reader"
)

type Parser struct {
	reader *Reader.Reader
}

// New creates a new Parser.
func New(l *Reader.Reader) *Parser {
	return &Parser{
		reader: l,
	}
}

// next returns the next token
func (p *Parser) next() Reader.Token {
	tok := <-p.reader.Tokens

	return tok
}

func (p *Parser) nextInst(tok Reader.Token) branfuck.Instruction {
	switch tok.Type {
	case Reader.IncTape:
		return branfuck.InstMoveHead{1}
	case Reader.DecTape:
		return branfuck.InstMoveHead{-1}
	case Reader.IncByte:
		return branfuck.InstAddToByte{1}
	case Reader.DecByte:
		return branfuck.InstAddToByte{-1}
	case Reader.WriteByte:
		return branfuck.InstWriteToOutput{}
	case Reader.StoreByte:
		return branfuck.InstReadFromInput{}
	case Reader.LoopEnter:
		return p.parseLoop()
	case Reader.LoopExit:
		return nil
	}
	panic("parser: unreachable")
}

func (p *Parser) parseLoop() branfuck.Instruction {
	insts := make([]branfuck.Instruction, 0)
	for tok := p.next(); tok.Type != Reader.EOF; tok = p.next() {
		i := p.nextInst(tok)
		if i == nil { // exit loop
			break
		}
		insts = append(insts, i)
	}
	return branfuck.InstLoop{insts}
}

func (p *Parser) Parse() []branfuck.Instruction {
	prog := make([]branfuck.Instruction, 0)
	for tok := p.next(); tok.Type != Reader.EOF; tok = p.next() {
		i := p.nextInst(tok)

		prog = append(prog, i)
	}
	return prog
}
