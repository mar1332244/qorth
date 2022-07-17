package main

import (
    "fmt"
    "bufio"
    "strconv"
    "path/filepath"
    "os"
    "strings"
)

type TokenType int

const (
    OP_INT_PUSH = TokenType(iota)
    OP_INT_ADD  = TokenType(iota)
    OP_INT_DUMP = TokenType(iota)
)

var KnownOperations map[string]TokenType = map[string]TokenType {
    "+": OP_INT_ADD,
    ".": OP_INT_DUMP,
}

type Token struct {
    File  string
    Line  int
    Pos   int
    Repr  string
    Value interface{}
    Type  TokenType
}

func (t Token) String() string {
    return "{" + t.Repr + "}"
}

func ValidateToken(token *Token) error {
    if opType, ok := KnownOperations[token.Repr]; ok {
        token.Value = token.Repr
        token.Type = opType
        return nil
    }
    if n, err := strconv.Atoi(token.Repr); err == nil {
        token.Value = n
        token.Type = OP_INT_PUSH
        return nil
    }
    return fmt.Errorf(
        "%s:%d:%d: unknown token `%s` encountered",
        token.File, token.Line, token.Pos, token.Repr,
    )
}

// TODO: benchmark strings.Builder vs string slicing
func SplitLineIntoTokens(fname, line string, lineNo int) ([]Token, error) {
    tokens := make([]Token, len(line))
    tokenSliceSize := 0
    var builder strings.Builder
    for linePos := 0; linePos < len(line); linePos++ {
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
        err := ValidateToken(&token)
        if err != nil {
            return tokens[:tokenSliceSize], err
        }
        tokens[tokenSliceSize] = token
        tokenSliceSize++
        builder.Reset()
    }
    return tokens[:tokenSliceSize], nil
}

func GetTokensFromFile(fname string) ([]Token, error) {
    fp, err := os.Open(fname)
    if err != nil {
        err := err.(*os.PathError)
        return nil, fmt.Errorf(
            "qorth: failed to open `%s` (%w)",
            err.Path, err.Err,
        )
    }
    defer fp.Close()
    scanner := bufio.NewScanner(fp)
    program := make([]Token, 0, 1024)
    fname = filepath.Base(fname)
    for lineNo := 1; scanner.Scan(); lineNo++ {
        tokens, err := SplitLineIntoTokens(
            fname, scanner.Text(), lineNo,
        )
        if err != nil {
            return program, err
        }
        program = append(program, tokens...)
    }
    if err = scanner.Err(); err != nil {
        return program, err
    }
    fmt.Println(program)
    return program, nil
}

func CreateCrossReferences(program []Token) (map[Token]Token, error) {
    return nil, nil
}

func InterpretProgram(program []Token) error {
    return nil
}

func main() {
    if len(os.Args) < 2 {
        fmt.Fprintln(os.Stderr, "qorth: no file provided")
        return
    }
    fstat, err := os.Stat(os.Args[1])
    if err != nil {
        err := err.(*os.PathError)
        fmt.Fprintf(
            os.Stderr, "qorth: failed to open `%s` (%s)\n",
            err.Path, err.Err,
        )
        return
    }
    if fstat.IsDir() {
        fmt.Fprintf(
            os.Stderr, "qorth: `%s` is a directory\n", os.Args[1],
        )
        return
    }
    if !strings.HasSuffix(os.Args[1], ".qorth") {
        fmt.Fprintf(
            os.Stderr, "qorth: `%s` is not a qorth file", os.Args[1],
        )
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
