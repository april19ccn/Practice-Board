// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package eval

import (
	"fmt"
	"strconv"
	"strings"
	"text/scanner"
)

// ---- lexer ----

// This lexer is similar to the one described in Chapter 13. 这个词法分析器类似于第13章中描述的词法分析器
type lexer struct {
	scan  scanner.Scanner
	token rune // current lookahead token
}

// 这段代码定义了lexer类型的next方法，功能是调用lex.scan.Scan()方法获取下一个词法单元，并将结果存储到lex.token字段中。
// 简单来说，就是向前移动词法分析器的指针并保存当前识别到的token。
func (lex *lexer) next() { lex.token = lex.scan.Scan() }

// 这段代码定义了lexer类型的text()方法，功能是返回当前扫描到的词法单元的文本内容。
// 它通过调用lex.scan.TokenText()方法获取并返回词法分析器当前识别到的token文本。
func (lex *lexer) text() string { return lex.scan.TokenText() }

type lexPanic string

// describe returns a string describing the current token, for use in errors.
// 描述返回一个描述当前令牌的字符串，用于错误。
func (lex *lexer) describe() string {
	switch lex.token {
	case scanner.EOF:
		return "end of file"
	case scanner.Ident:
		return fmt.Sprintf("identifier %s", lex.text())
	case scanner.Int, scanner.Float:
		return fmt.Sprintf("number %s", lex.text())
	}
	return fmt.Sprintf("%q", rune(lex.token)) // any other rune
}

// 这段代码定义了一个precedence函数，用于返回运算符的优先级
func precedence(op rune) int {
	switch op {
	case '*', '/':
		return 2
	case '+', '-':
		return 1
	}
	return 0
}

// ---- parser ----

// Parse parses the input string as an arithmetic expression. Parse 将输入字符串解析为算术表达式。
//
//	expr = num                         a literal number, e.g., 3.14159
//	     | id                          a variable name, e.g., x
//	     | id '(' expr ',' ... ')'     a function call
//	     | '-' expr                    a unary operator (+-)
//	     | expr '+' expr               a binary operator (+-*/)
func Parse(input string) (_ Expr, err error) {
	// https://golang-china.github.io/gopl-zh/ch5/ch5-10.html
	// Recover捕获异常
	defer func() {
		switch x := recover().(type) {
		case nil:
			// no panic
		case lexPanic:
			err = fmt.Errorf("%s", x)
		default:
			// unexpected panic: resume state of panic.
			panic(x)
		}
	}()

	lex := new(lexer)
	lex.scan.Init(strings.NewReader(input))
	lex.scan.Mode = scanner.ScanIdents | scanner.ScanInts | scanner.ScanFloats
	lex.next() // initial lookahead
	e := parseExpr(lex)
	if lex.token != scanner.EOF {
		return nil, fmt.Errorf("unexpected %s", lex.describe())
	}
	return e, nil
}

func parseExpr(lex *lexer) Expr { return parseBinary(lex, 1) }

// binary = unary ('+' binary)*
// parseBinary stops when it encounters an
// operator of lower precedence than prec1.
func parseBinary(lex *lexer, prec1 int) Expr {
	lhs := parseUnary(lex)
	for prec := precedence(lex.token); prec >= prec1; prec-- {
		for precedence(lex.token) == prec {
			op := lex.token
			lex.next() // consume operator
			rhs := parseBinary(lex, prec+1)
			lhs = binary{op, lhs, rhs}
		}
	}
	return lhs
}

// unary = '+' expr | primary
func parseUnary(lex *lexer) Expr {
	if lex.token == '+' || lex.token == '-' {
		op := lex.token
		lex.next() // consume '+' or '-'
		return unary{op, parseUnary(lex)}
	}
	return parsePrimary(lex)
}

// primary = id
//
//	| id '(' expr ',' ... ',' expr ')'
//	| num
//	| '(' expr ')'
func parsePrimary(lex *lexer) Expr {
	switch lex.token {
	case scanner.Ident:
		id := lex.text()
		lex.next() // consume Ident
		if lex.token != '(' {
			return Var(id)
		}
		lex.next() // consume '('
		var args []Expr
		if lex.token != ')' {
			for {
				args = append(args, parseExpr(lex))
				if lex.token != ',' {
					break
				}
				lex.next() // consume ','
			}
			if lex.token != ')' {
				msg := fmt.Sprintf("got %s, want ')'", lex.describe())
				panic(lexPanic(msg))
			}
		}
		lex.next() // consume ')'
		return call{id, args}

	case scanner.Int, scanner.Float:
		f, err := strconv.ParseFloat(lex.text(), 64)
		if err != nil {
			panic(lexPanic(err.Error()))
		}
		lex.next() // consume number
		return literal(f)

	case '(':
		lex.next() // consume '('
		e := parseExpr(lex)
		if lex.token != ')' {
			msg := fmt.Sprintf("got %s, want ')'", lex.describe())
			panic(lexPanic(msg))
		}
		lex.next() // consume ')'
		return e
	}
	msg := fmt.Sprintf("unexpected %s", lex.describe())
	panic(lexPanic(msg))
}
