package parser

import (
	"strconv"
)

import (
	"github.com/mar1332244/qorth/tokenizer"
)

func CreateCrossReferences(program Program) error {
	for ip := program.First; ip != nil; ip = ip.Next {
		switch ip.Type {
		case BLOCK_IF:
		default:
			break
		}
	}
}

func ParseTokens(tokens []tokenizer.Token) (Program, error) {
	var program Program
	for _, t := range tokens {
		node := new(Instruction)
		tType, value := GetTokenType(t)
		if tType == TOKEN_UNKNOWN {
			return program, fmt.Errorf("[ERROR] qorth: unknown token `%s` encountered", t.Repr)
		}
		node.Type = tType
		node.Value = value
		program.Append(node)
	}
	return program, nil
}

func GetTokenType(t tokenizer.Token) (TokenType, int) {
	if tType, ok := KnownInstructions[t.Repr]; ok {
		return tType, 0
	}
	if n, err := strconv.Atoi(t.Repr); err == nil {
		return OP_QUEUE_PUSH, n
	}
	return TOKEN_UNKNOWN, -1
}
