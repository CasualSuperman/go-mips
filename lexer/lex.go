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

const eof = -1

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
	l := lexer{name, input}
	go l.run()
	return &l
}

type stateFn func(*lexer) stateFn

func (l *lexer) run() {
	for state := lexCode; state != nil; {
		state = state(l)
	}
	close(l.items)
}

// Emit the string currently selected between start and pos with type typ.
func (l *lexer) emit(typ itemType) {
	l.items <- item{typ, l.input[l.start:l.pos]}
	l.start = l.pos
}

func (l *lexer) ignore() {
	l.start = l.pos
}

func (l *lexer) backup() {
	l.pos -= l.width
}

func (l *lexer) next() rune {
	if l.pos >= len(l.input) {
		l.width = 0
		return eof
	}

	r, l.width = utf8.DecodeRuneInString(l.input[pos:])
	l.pos += l.width
	return r
}

func (l *lexer) peek() rune {
	w := l.width
	r := l.next()
	l.backup()
	l.width = w
	return r
}

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
