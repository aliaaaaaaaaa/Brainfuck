package branfuck

// Tape represents a tape a Brainfuck program executes on.
type Tape interface {
	MoveHead(i int)

	AddToByte(i int)

	SetByte(b byte)

	GetByte() byte

	GetHead() int

	SetHead(i int)
}

type TapeImpl struct {
	head int // position of the head in the tape
	tape []byte
}

func New() Tape {
	return &TapeImpl{
		head: 0,
		tape: make([]byte, 1),
	}
}

func (t *TapeImpl) MoveHead(i int) {
	if t.head+i < 0 {
		panic("tape: out of bounds, cannot have a negative tape header")
	}

	if t.head+i+1 > len(t.tape) {
		t.tape = append(t.tape, make([]byte, i+1)...)
	}
	t.head += i
}

func (t *TapeImpl) AddToByte(i int) {
	cb := t.tape[t.head]
	if (int(cb)+i)%256 < 0 {
		t.tape[t.head] = byte(int(cb) + (i % 256) + 256)
		return
	}
	t.tape[t.head] = byte((int(cb) + i) % 256)
}

func (t *TapeImpl) GetByte() byte {
	b := t.tape[t.head]
	return b
}

func (t *TapeImpl) SetByte(b byte) {
	t.tape[t.head] = b
}

func (t *TapeImpl) GetHead() int {
	return t.head
}

func (t *TapeImpl) SetHead(i int) {
	if i < 0 || i > len(t.tape) {
		panic("tape: cannot set head to negative index")
	}
	t.head = i
}
