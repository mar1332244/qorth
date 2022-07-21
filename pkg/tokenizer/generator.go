package tokenizer

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

const initialTokensCap = 1024

const (
	commentBegin = '('
	commentEnd   = ')'
)

func GetTokensFromFile(fname string) ([]Token, error) {
	inFile, err := os.Open(fname)
	if err, ok := err.(*os.PathError); ok {
		return nil, fmt.Errorf("[ERROR]: failed to open `%s` (%s)", err.Path, err.Err)
	}
	defer inFile.Close()
	fname = filepath.Base(fname)
	inComment := new(bool)
	inString := new(bool)
	tokens := make([]Token, 0, initialTokensCap)
	scanner := bufio.NewScanner(inFile)
	for lineNo := 1; scanner.Scan(); lineNo++ {
		for _, token := range splitLine(scanner.Text(), inComment, inString) {
			token.Line = lineNo
			token.File = fname
			tokens = append(tokens, token)
		}
	}
	if err = scanner.Err(); err != nil {
		return tokens, fmt.Errorf("[FATAL]:%v", err)
	}
	return tokens, nil
}

func splitLine(line string, inComment, inString *bool) []Token {
	tokens := make([]Token, 0, len(line))
	for linePos := 0; linePos < len(line); linePos++ {
		if line[linePos] == commentBegin {
			*inComment = true
		}
		if line[linePos] == ' ' || *inComment {
			if line[linePos] == commentEnd {
				*inComment = false
			}
			continue
		}
		start := linePos
		if line[linePos] == '"' {
			*inString = true
		}
		endToken := ' '
		if *inString {
			endToken = '"'
		}
		for linePos < len(line) && line[linePos] != endToken {
			linePos++
		}
		tokens = append(tokens, Token {
			Pos:  start + 1,
			Repr: line[start:linePos],
		})
	}
	return tokens
}
