package tokenizer

type Token struct {
	File string
	Line int
	Pos  int
	Repr string
}

func (t Token) String() string {
	return "{" + t.Repr + "}"
}
