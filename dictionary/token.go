package dictionary

type Token struct {
	text      string
	frequency float64
	pos       string
}

func (t Token) Text() string {
	return t.text
}

func (t Token) Frequency() float64 {
	return t.frequency
}

func (t Token) Pos() string {
	return t.pos
}

func NewToken(text string, frequency float64, pos string) Token {
	return Token{text: text, frequency: frequency, pos: pos}
}
