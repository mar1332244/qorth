package main

import (
    "bufio"
    "fmt"
    "os"
    "path/filepath"
    "strconv"
    "strings"
)

import (
    "github.com/mar1332244/qorth/pkg/queue"
)

type TokenType int

const (
    OP_QUEUE_PUSH  = TokenType(iota) // any int
	OP_QUEUE_POP   = TokenType(iota) // '
	OP_QUEUE_SIZE  = TokenType(iota) // ?
	OP_QUEUE_DUP   = TokenType(iota) // @
	OP_QUEUE_BACK  = TokenType(iota) // #
	OP_QUEUE_CLEAR = TokenType(iota) // _

    OP_INT_ADD = TokenType(iota) // +
	OP_INT_SUB = TokenType(iota) // -
	OP_INT_MUL = TokenType(iota) // *
	OP_INT_DIV = TokenType(iota) // /
	OP_INT_MOD = TokenType(iota) // %
	OP_INT_POW = TokenType(iota) // **

	OP_INT_AND = TokenType(iota) // &
	OP_INT_OR  = TokenType(iota) // |
	OP_INT_XOR = TokenType(iota) // ^
	OP_INT_NOT = TokenType(iota) // ~
	OP_INT_RS  = TokenType(iota) // <<
	OP_INT_LS  = TokenType(iota) // >>

	OP_BOOL_AND = TokenType(iota) // &&
	OP_BOOL_OR  = TokenType(iota) // ||
	OP_BOOL_NOT = TokenType(iota) // !
	OP_BOOL_GT  = TokenType(iota) // <
	OP_BOOL_GE  = TokenType(iota) // <=
	OP_BOOL_LT  = TokenType(iota) // >
	OP_BOOL_LE  = TokenType(iota) // >=
	OP_BOOL_EQ  = TokenType(iota) // ==
	OP_BOOL_NEQ = TokenType(iota) // !=

    OP_INT_DUMP  = TokenType(iota) // .
	OP_CHAR_DUMP = TokenType(iota) // $
	OP_CHAR_READ = TokenType(iota) // ,

	BLOCK_WHILE  = TokenType(iota) // while
	BLOCK_DO     = TokenType(iota) // do
	BLOCK_END    = TokenType(iota) // end
	BLOCK_UNLESS = TokenType(iota) // unless
	BLOCK_ELSE   = TokenType(iota) // else

	COMMENT_BEGIN = TokenType(iota) // (
	COMMENT_END   = TokenType(iota) // )
)

/*

10
[10]
2
[10 2]
dup
[10 2 10]
%
[10 0]
back
[0 10]
drop
[10]
.
[10]
1
[10 1]
-
[9]
2
[9 2]
dup
[9 2 9]
%
[9 1]
back
[1 9]
back
[9 1]
dup
[9 1 9]
+
[9 10]
back
[10 9]
.
[10 9]
drop
[9]
10
[9 10]
back
[10 9]
$
[10 9]
drop
[9]
1
[9 1]
-
[8]

10 do
	2 @ % #
	if 0 == do
		' .
	else
		# @ + # . '
	end
	10 # $ '
	1 -
while 0 > end
*/

var KnownOperations map[string]TokenType = map[string]TokenType {
    "+": OP_INT_ADD,
    ".": OP_INT_DUMP,
	"'": OP_QUEUE_POP,
	"$": OP_CHAR_DUMP,
}

type Token struct {
    File  string
    Line  int
    Pos   int
    Repr  string
    Value int
    Type  TokenType
}

func (t Token) String() string {
    return "{" + t.Repr + "}"
}

func ValidateToken(t *Token) error {
    if opType, ok := KnownOperations[t.Repr]; ok {
        t.Type = opType
        return nil
    }
    if n, err := strconv.Atoi(t.Repr); err == nil {
        t.Value = n
        t.Type = OP_QUEUE_PUSH
        return nil
    }
    return fmt.Errorf(
		"%s:%d:%d: unknown token `%s` encountered", t.File, t.Line, t.Pos, t.Repr,
	)
}

// TODO: benchmark strings.Builder vs string slicing
// TODO: benchmark builder.Reset() vs redeclaration in loop
func SplitLineIntoTokens(fname, line string, lineNo int) ([]Token, error) {
    lineTokens := make([]Token, 0, len(line))
    for linePos := 0; linePos < len(line); linePos++ {
		var builder strings.Builder
        if line[linePos] == ' ' {
            continue
        }
        for linePos < len(line) && line[linePos] != ' ' {
            builder.WriteByte(line[linePos])
            linePos++
        }
        token := Token {
            Repr: builder.String(),
            File: fname,
            Line: lineNo,
            Pos:  linePos + 1,
        }
        if err := ValidateToken(&token); err != nil {
            return lineTokens, err
        }
        lineTokens = append(lineTokens, token)
    }
    return lineTokens, nil
}

func GetTokensFromFile(fname string) ([]Token, error) {
    inFile, err := os.Open(fname)
    if err, ok := err.(*os.PathError); ok {
        return nil, fmt.Errorf("qorth: failed to open `%s` (%w)", err.Path, err.Err)
    }
    defer inFile.Close()
    scanner := bufio.NewScanner(inFile)
    program := make([]Token, 0, 1024)
    fname = filepath.Base(fname)
    for lineNo := 1; scanner.Scan(); lineNo++ {
        tokens, err := SplitLineIntoTokens(fname, scanner.Text(), lineNo)
        if err != nil {
            return program, err
        }
        program = append(program, tokens...)
    }
    if err = scanner.Err(); err != nil {
		return program, fmt.Errorf("qorth: %v", err)
    }
	fmt.Println(program)
    return program, nil
}

func CreateCrossReferences(program []Token) (map[Token]Token, error) {
    return nil, nil
}

func InterpretProgram(program []Token) error {
	var q queue.Queue
	for ip := 0; ip < len(program); ip++ {
		t := program[ip]
		switch t.Type {
		case OP_QUEUE_PUSH:
			q.Push(t.Value)
		case OP_INT_ADD:
			a, err := q.Peek()
			if err != nil {
				return fmt.Errorf("%s:%d:%d: failed to get value for `+`", t.File, t.Line, t.Pos)
			}
			q.Pop()
			b, err := q.Peek()
			if err != nil {
				return fmt.Errorf("%s:%d:%d: failed to get value for `+`", t.File, t.Line, t.Pos)
			}
			q.Pop()
			q.Push(a + b)
		case OP_INT_DUMP:
			a, err := q.Peek()
			if err != nil {
				return fmt.Errorf("%s:%d:%d: failed to get value for `.`", t.File, t.Line, t.Pos)
			}
			fmt.Print(a)
		case OP_QUEUE_POP:
			if err := q.Pop(); err != nil {
				return fmt.Errorf("%s:%d:%d: failed to get value for `'`", t.File, t.Line, t.Pos)
			}
		case OP_CHAR_DUMP:
			a, err := q.Peek()
			if err != nil {
				return fmt.Errorf("%s:%d:%d: failed to get value for `$`", t.File, t.Line, t.Pos)
			}
			fmt.Printf("%c", a)
		default:
			return fmt.Errorf("qorth: unreachable token encountered")
		}
	}
    return nil
}

func main() {
    if len(os.Args) < 2 {
        fmt.Fprintln(os.Stderr, "qorth: no file provided")
        return
    }
    fstat, err := os.Stat(os.Args[1])
    if err, ok := err.(*os.PathError); ok {
        fmt.Fprintf(os.Stderr, "qorth: failed to open `%s` (%s)\n", err.Path, err.Err)
        return
    }
    if fstat.IsDir() {
        fmt.Fprintf(os.Stderr, "qorth: `%s` is a directory\n", os.Args[1])
        return
    }
    if !strings.HasSuffix(os.Args[1], ".qorth") {
        fmt.Fprintf(os.Stderr, "qorth: `%s` is not a qorth file", os.Args[1])
        return
    }
    program, err := GetTokensFromFile(os.Args[1])
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
    if _, err = CreateCrossReferences(program); err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
    if err = InterpretProgram(program); err != nil {
        fmt.Fprintln(os.Stderr, err)
    }
}
