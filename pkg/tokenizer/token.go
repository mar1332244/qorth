package tokenizer

import (
	"fmt"
	"strings"
)

type Token struct {
	File string
	Line int
	Pos  int
	Repr string
}

func (t Token) String() string {
	return fmt.Sprintf("%s:%d:%d", t.File, t.Line, t.Pos)
}

func Add(t1, t2 Token) Token {
	var builder strings.Builder
	t := t1
	builder.WriteString(t1.Repr)
	builder.WriteString(t2.Repr)
	t.Repr = builder.String()
    return t
}
