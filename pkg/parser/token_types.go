package parser

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

	BLOCK_WHILE = TokenType(iota) // while
	BLOCK_DO    = TokenType(iota) // do
	BLOCK_END   = TokenType(iota) // end
	BLOCK_IF    = TokenType(iota) // if
	BLOCK_ELSE  = TokenType(iota) // else

	TOKEN_UNKNOWN = TokenType(iota)
)
