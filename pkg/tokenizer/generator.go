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

func skipComment(buffer *bytes.Buffer) (int, bool) {
	charsRead := 1
	for 0 < buffer.Len() {
		charsRead++
		if c, _ := buffer.ReadByte(); c == commentEnd {
			break
		}
	}
	return charsRead, 0 < buffer.Len()
}

func generateTokens(buffer *bytes.Buffer) ([]Token, error) {
	var builder strings.Builder
	var lineNum, linePos, startLine, startPos int
	filename, _ := buffer.ReadString('\n')
	filename = filename[:len(filename)-1]
	tokens := make([]Token, 0, initialSliceCap)
	for 0 < buffer.Len() {
		c, _ := buffer.ReadByte()
		switch c {
		}
	}
	return tokens[:len(tokens):len(tokens)], nil
}

func GetTokensFromFile(filename string) ([]Token, error) {
	inFile, err := os.Open(filename)
	if err, ok := err.(*os.PathError); ok {
		return nil, fmt.Errorf(openErr, err.Path, err.Err)
	}
	defer inFile.Close()
	buffer := bytes.NewBufferString(filepath.Base(filename))
	buffer.WriteByte('\n')
	buffer.Grow(initialSliceCap)
	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		buffer.WriteString(scanner.Text())
		buffer.WriteByte('\n')
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf(scanErr, err)
	}
	return generateTokens(buffer)
}
