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
		case unicode.IsSpace(r):
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
			return lexLabel
		case r == '$':
			l.ignore()
			return lexRegister
		}
	}
}

func lexDirective(l *lexer) stateFn {

	for {
		switch r := l.next(); {
		case unicode.IsSpace(r) || r == eof || r == '\n':
			l.emit(itemDirective)
			return lexCode
		case unicode.IsLetter(r) || '0' <= r && r <= '9':
			// Continue
		}
	}
}

func lexLabel(l *lexer) stateFn {
	for {
		switch r := l.next(); {
		case r == eof || r == '\n':
			return l.errorf("Unfinished label.")
		}
	}
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
