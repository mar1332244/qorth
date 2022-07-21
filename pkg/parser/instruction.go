package parser

type Instruction struct {
	Prev  *Instruction
	Next  *Instruction
	Type  TokenType
	Value int
	Refs  map[string]*Instruction
}

type Program struct {
	First *Instruction
	Last  *Instruction
}

func (p *Program) Append(node *Instruction) {
	if p.First == nil {
		p.First = node
		p.Last = p.First
	} else {
		p.Last.Next = node
		node.Prev = p.Last
	}
}

var KnownInstructions = map[string]TokenType {

	"'": OP_QUEUE_POP,
	"?": OP_QUEUE_SIZE, 
	"@": OP_QUEUE_DUP,
	"#": OP_QUEUE_BACK, 
	"_": OP_QUEUE_CLEAR,

	"+":  OP_INT_ADD,
	"-":  OP_INT_SUB,
	"*":  OP_INT_MUL,
	"/":  OP_INT_DIV,
	"%":  OP_INT_MOD,
	"**": OP_INT_POW,

	"&":  OP_INT_AND,
	"|":  OP_INT_OR,
	"^":  OP_INT_XOR,
	"~":  OP_INT_NOT,
	"<<": OP_INT_RS,
	">>": OP_INT_LS,

	"&&": OP_BOOL_AND,
	"||": OP_BOOL_OR,
	"!":  OP_BOOL_NOT,
	">":  OP_BOOL_GT,
	">=": OP_BOOL_GE,
	"<":  OP_BOOL_LT,
	"<=": OP_BOOL_LE,
	"==": OP_BOOL_EQ,
	"!=": OP_BOOL_NEQ,

	".": OP_INT_DUMP, 
	"$": OP_CHAR_DUMP,
	",": OP_CHAR_READ,

	"while": BLOCK_WHILE,
	"do":    BLOCK_DO,
	"end":   BLOCK_END,  
	"if":    BLOCK_IF, 
	"else":  BLOCK_ELSE,
}
