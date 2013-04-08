package lexer

import "unicode"

// Every line in a MIPS programs has the format [label]*[instruction][comment]
func lexCode(l *lexer) stateFn {
	for {
		switch r := l.next(); {
		case r == eof:
			l.emit(itemEOF)
			return nil
		case r == '.':
			return lexDirective
		case unicode.IsSpace(r) || r == ',':
			l.ignore()
		case r == '"':
			return lexQuote
		case r == '+' || r == '-' || '0' < r && r <= '9':
			l.backup()
			return lexNumber
		case r == '#':
			l.ignore()
			return lexComment
		case unicode.IsLetter(r):
			l.backup()
			return lexWord
		case r == '$':
			l.ignore()
			return lexRegister
		}
	}
}

func lexRegister (l *lexer) stateFn {

}

func lexDirective(l *lexer) stateFn {
	for {
		switch r := l.next(); {
		case unicode.IsSpace(r) || r == eof || r == '\n':
			l.emit(itemDirective)
			return lexCode
		case '0' <= r && r <= '9':
			return l.errorf("Directives cannot contain numbers.")
		}
	}
}

func lexWord(l *lexer) stateFn {
	for {
		switch r := l.next(); {
		case r == eof || r == '\n':
			return l.errorf("Unfinished label.")
		case r == ':':
			l.emit(itemLabel)
			return lexCode
		case unicode.IsSpace(r):
			l.backup()
			l.emit(itemInstruction)
			return lexCode
		}
	}
}

func lexQuote(l *lexer) stateFn {
Loop:
	for {
		switch l.next() {
		case '\\':
			if r := l.next(); r != eof && r != '\n' {
				break
			}
			fallthrough
		case eof, '\n':
			return l.errorf("unterminated quoted string")
		case '"':
			break Loop
		}
	}
	l.emit(itemString)
	return lexCode
}

func lexComment(l *lexer) stateFn {
	for {
		switch r := l.next(); {
		case r == eof || r == '\n':
			l.emit(itemComment)
			return lexCode
		}
	}
}
