package eval

// An Expr is an arithmetic expression.
type Expr interface {
	// Eval returns the value of this Expr in the environment env.
	Eval(env Env) float64
	// Check reports errors in this Expr and adds its Vars to the set.
	// Check(vars map[Var]bool) error
}

//!+ast

// A Var identifies a variable, e.g., x. Var标识一个变量，例如x。
type Var string

// A literal is a numeric constant, e.g., 3.141. 文字是一个数字常量，例如 3.141。
type literal float64

// A unary represents a unary operator expression, e.g., -x. 一元表示一元运算符表达式，例如-x。
type unary struct {
	op rune // one of '+', '-'
	x  Expr
}

// A binary represents a binary operator expression, e.g., x+y. 二进制表示二进制运算符表达式，例如 x+y。
type binary struct {
	op   rune // one of '+', '-', '*', '/'
	x, y Expr
}

// A call represents a function call expression, e.g., sin(x). 调用表示函数调用表达式，例如 sin（x）
type call struct {
	fn   string // one of "pow", "sin", "sqrt"
	args []Expr
}

//!-ast
