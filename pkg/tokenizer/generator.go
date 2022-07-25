package tokenizer

import (
	"bytes"
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const initialSliceCap = 1024

const (
	commentBegin = '('
	commentEnd   = ')'
)

const (
	openErr = "[ERROR] qorth: failed to open `%s` (%s)"
	scanErr = "[FATAL] internal error: %v"
	openStr = "[ERROR] %v: string %v\" is never closed"
)

func getString(buffer *bytes.Buffer) (string, bool) {
	var builder strings.Builder
	builder.WriteByte('"')
	for 0 < buffer.Len() {
		c, _ := buffer.ReadByte()
		builder.WriteByte(c)
		if c == '"' {
			return builder.String(), true
		}
	}
	return builder.String(), false
}

func readFilename(buffer *bytes.Buffer) string {
	filename, _ := buffer.ReadString('\n')
	return filename[:len(filename)-1]
}

func generateTokens(buffer *bytes.Buffer) ([]Token, error) {
	var builder strings.Builder
	token := &Token{File: readFilename(buffer), Line: 1, Pos: 1}
	inString := false
	tokens := make([]Token, 0, initialSliceCap)
	for lineNum, charNum := 1, 1; 0 < buffer.Len(); charNum++ {
		c, _ := buffer.ReadByte()
		if c == '\n' {
			if 0 < builder.Len() {
				token.Repr = builder.String()
				tokens = append(tokens, *token)
				builder.Reset()
			}
			charNum = 0
			lineNum++
		}
		if c != ' ' && builder.Len() == 0 {
			token.Line = lineNum
			token.Pos = charNum
		}
		if c == ' ' && 0 < builder.Len() {
			token.Repr = builder.String()
			tokens = append(tokens, *token)
			builder.Reset()
		}
		builder.WriteByte(c)
	}
	if inString {
		return nil, fmt.Errorf(openStr, token, builder)
	}
	return tokens[:len(tokens):len(tokens)], nil
}

func GetTokensFromFile(src string) ([]Token, error) {
	inFile, err := os.Open(src)
	if err, ok := err.(*os.PathError); ok {
		return nil, fmt.Errorf(openErr, err.Path, err.Err)
	}
	defer inFile.Close()
	buffer := new(bytes.Buffer)
	fmt.Fprintln(buffer, filepath.Base(src))
	buffer.Grow(initialSliceCap)
	scanner := bufio.NewScanner(inFile)
	for inComment := false; scanner.Scan(); {
		line := scanner.Text()
		for _, char := range line {
			if char == commentBegin && !inComment {
				inComment = true
			}
			if inComment {
				if char == commentEnd {
					inComment = false
				}
				buffer.WriteByte(' ')
				continue
			}
			buffer.WriteRune(char)
		}
		buffer.WriteByte('\n')
	}
	if err = scanner.Err(); err != nil {
		return nil, fmt.Errorf(scanErr, err)
	}
	return generateTokens(buffer)
}
