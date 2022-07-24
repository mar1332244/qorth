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
	openStr = "[ERROR] %s:%d:%d: string %s encountered"
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

func generateTokens(buffer *bytes.Buffer, src string) ([]Token, error) {
	var builder strings.Builder
	var inComment, inString bool
	tokens := make([]Token, 0, initialSliceCap)
	lineNum, charNum := 1, 1
	token := Token{File: filpath.Base(src), Line: 1, Pos: 1}
	for c := byte(0); 0 < buffer.Len(); c, _ = buffer.ReadByte() {
	}
	if inString {
		return nil, fmt.Errorf(
			openStr, token.File, token.Line, token.Pos, builder.String(),
		)
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
	buffer.Grow(initialSliceCap)
	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		buffer.WriteString(scanner.Text())
		buffer.WriteByte('\n')
	}
	if err = scanner.Err(); err != nil {
		return nil, fmt.Errorf(scanErr, err)
	}
	return generateTokens(buffer, src)
}
