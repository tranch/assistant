package interpreter

import (
	"errors"
)

var ErrInterpreter = errors.New("interpreter error")

type Interpreter struct {
	Lexer        *Lexer
	CurrentToken *Token
}

func NewInterpreter(lexer *Lexer) (*Interpreter, error) {
	token, err := lexer.GetNextToken()
	if err != nil {
		return nil, err
	}
	return &Interpreter{Lexer: lexer, CurrentToken: token}, nil
}

func (i *Interpreter) Eat(tokenType string) (err error) {
	if i.CurrentToken.Type == tokenType {
		i.CurrentToken, err = i.Lexer.GetNextToken()
		return
	}

	return ErrInterpreter
}

func (i *Interpreter) Factor() (int, error) {
	token := i.CurrentToken
	if token.Type == INTEGER {
		err := i.Eat(INTEGER)
		if err != nil {
			return 0, nil
		}
		return token.Value.(int), nil
	} else if token.Type == LPAREN {
		err := i.Eat(LPAREN)
		if err != nil {
			return 0, nil
		}
		value, err := i.Expr()
		if err != nil {
			return 0, nil
		}
		err = i.Eat(RPAREN)
		if err != nil {
			return 0, nil
		}
		return value, nil
	}

	return 0, ErrInterpreter
}

func (i *Interpreter) Term() (int, error) {
	value, err := i.Factor()
	if err != nil {
		return 0, err
	}

	for i.CurrentToken.Type == MULTIPLY || i.CurrentToken.Type == DIVIDE {
		tokenType := i.CurrentToken.Type
		if tokenType == MULTIPLY {
			err = i.Eat(MULTIPLY)
			if err != nil {
				return 0, err
			}
			num, err := i.Factor()
			if err != nil {
				return 0, err
			}
			value *= num
		}
		if tokenType == DIVIDE {
			err = i.Eat(DIVIDE)
			if err != nil {
				return 0, err
			}
			num, err := i.Factor()
			if err != nil {
				return 0, err
			}
			value /= num
		}
	}

	return value, nil
}

func (i *Interpreter) Expr() (int, error) {
	value, err := i.Term()
	if err != nil {
		return 0, err
	}

	for i.CurrentToken.Type == PLUS || i.CurrentToken.Type == MINUS {
		tokenType := i.CurrentToken.Type
		if tokenType == PLUS {
			err = i.Eat(PLUS)
			if err != nil {
				return 0, err
			}
			num, err := i.Term()
			if err != nil {
				return 0, err
			}
			value += num
		}
		if tokenType == MINUS {
			err = i.Eat(MINUS)
			if err != nil {
				return 0, err
			}
			num, err := i.Term()
			if err != nil {
				return 0, err
			}
			value -= num
		}
	}

	return value, nil
}

// expr   : term   ((PLUS | MINUS) term)*
// term   : factor ((MUL | DIV) factor)*
// factor : INTEGER | LPAREN expr RPAREN
func (i *Interpreter) Parse() (int, error) {
	return i.Expr()
}
