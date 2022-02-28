package lexer

type LexemeType int64

const (
	Undefined LexemeType = iota

	// Single-character lexemes
	LeftParenthesis
	RightParenthesis
	LeftBrace
	RightBrace
	Comma
	Dot
	Minus
	Plus
	Semicolon
	Slash
	Star

	// One-or-two character lexemes
	Bang
	BangEqual
	Equal
	EqualEqual
	Greater
	GreaterEqual
	Less
	LessEqual

	// Literals
	Identifier
	LiteralString
	LiteralNumber

	// Keywords
	And
	Class
	Else
	False
	Fun
	For
	If
	Nil
	Or
	Print
	Return
	Super
	This
	True
	Var
	While

	EOF
)

var keywords = map[string]LexemeType{
	And.String():    And,
	Class.String():  Class,
	Else.String():   Else,
	False.String():  False,
	For.String():    For,
	Fun.String():    Fun,
	If.String():     If,
	Nil.String():    Nil,
	Or.String():     Or,
	Print.String():  Print,
	Return.String(): Return,
	Super.String():  Super,
	This.String():   This,
	True.String():   True,
	Var.String():    Var,
	While.String():  While,
}

// Return the string form of the LexemeType.
func (lexeme_type LexemeType) String() string {
	switch lexeme_type {
	case Undefined:
		return "Undefined"
	case LeftParenthesis:
		return "LeftParenthesis"
	case RightParenthesis:
		return "RightParenthesis"
	case LeftBrace:
		return "LeftBrace"
	case RightBrace:
		return "RightBrace"
	case Comma:
		return "Comma"
	case Dot:
		return "Dot"
	case Minus:
		return "Minus"
	case Plus:
		return "Plus"
	case Semicolon:
		return "Semicolon"
	case Slash:
		return "Slash"
	case Star:
		return "Star"
	case Bang:
		return "Bang"
	case BangEqual:
		return "BangEqual"
	case Equal:
		return "Equal"
	case EqualEqual:
		return "EqualEqual"
	case Greater:
		return "Greater"
	case GreaterEqual:
		return "GreaterEqual"
	case Less:
		return "Less"
	case LessEqual:
		return "LessEqual"
	case Identifier:
		return "Identifier"
	case LiteralString:
		return "String"
	case LiteralNumber:
		return "Number"
	case And:
		return "And"
	case Class:
		return "Class"
	case Else:
		return "Else"
	case False:
		return "False"
	case Fun:
		return "Fun"
	case For:
		return "For"
	case If:
		return "If"
	case Nil:
		return "Nil"
	case Or:
		return "Or"
	case Print:
		return "Print"
	case Return:
		return "Return"
	case Super:
		return "Super"
	case This:
		return "This"
	case True:
		return "True"
	case Var:
		return "Var"
	case While:
		return "While"
	case EOF:
		return "EOF"
	}
	return "Unknown"
}
