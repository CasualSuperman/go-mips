package lexer

type itemType int

const (
	itemError itemType = iota
	itemEOF
	itemNumber
	itemLabel
	itemDirective
	itemInstruction
	itemRegister
	itemString
	itemComment
)

type item struct {
	typ itemType
	val string
}

type lexer struct {
	name  string    // used only for error reports.
	input string    // the string being scanned.
	start int       // start position of this item.
	pos   int       // current position in the input.
	width int       // width of last rune read from input.
	items chan item // channel of scanned items.
}

func Lex(name, input string) *lexer {

}

func (l *lexer) run() {
	for state := lexText; state != nil; {
		state = state(l)
	}
	close(l.items)
}

type stateFn func(*lexer) stateFn

/*
* Labels
* Directives
* register
* number
* hex literal
* octal literal
* comment
* decimal literal
* binary literal
* instruction
 */
