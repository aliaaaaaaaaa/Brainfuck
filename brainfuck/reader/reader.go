package Reader

import (
	"io"
)

type Type int

const (
	EOF       Type = iota // zero value
	IncTape               // '>' increment tape head
	DecTape               // '<' decrement tape head
	IncByte               // '+' increment byte value at tape head
	DecByte               // '-' decrement byte value at tape head
	WriteByte             // '.' write byte to output
	StoreByte             // ',' store byte from input to tape header
	LoopEnter             // '['
	LoopExit              // ']'
)

type Token struct {
	Type    Type
	ByteVal byte
}

const eof = -1

type stateFn func(*Reader) stateFn

type Reader struct {
	Tokens chan Token // lexed tokens
	r      io.Reader  // input stream
	input  byte       // current character
	state  stateFn
}

func New(r io.Reader) *Reader {
	l := &Reader{
		Tokens: make(chan Token),
		r:      r,
	}
	go l.run()
	return l
}

func (l *Reader) run() {
	for l.state = lexMain; l.state != nil; {
		l.state = l.state(l)
	}
	close(l.Tokens)
}

// byte ranges from 0..255
func (l *Reader) next() rune {
	buf := make([]byte, 1)
	_, err := io.ReadFull(l.r, buf)

	if err != nil { // EOF
		return eof
	}

	l.input = buf[0]

	return rune(buf[0])
}

func (l *Reader) send(t Type) {
	tok := Token{
		Type:    t,
		ByteVal: l.input,
	}

	l.Tokens <- tok
}

func lexMain(l *Reader) stateFn {
	r := l.next()

	switch r {
	case eof: // This will terminate the loop in l.run()
		return nil

	case '>':
		l.send(IncTape)
	case '<':
		l.send(DecTape)
	case '+':
		l.send(IncByte)
	case '-':
		l.send(DecByte)
	case '.':
		l.send(WriteByte)
	case ',':
		l.send(StoreByte)
	case '[':
		l.send(LoopEnter)
	case ']':
		l.send(LoopExit)
	}

	return lexMain
}
