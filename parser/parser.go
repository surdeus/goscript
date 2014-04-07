//line parser.go.y:2
package parser

import __yyfmt__ "fmt"

//line parser.go.y:2
import (
	"github.com/mattn/anko/ast"
	"unsafe"
)

//line parser.go.y:23
type yySymType struct {
	yys                    int
	stmt_var               ast.Stmt
	stmt_if                ast.Stmt
	stmt_for               ast.Stmt
	stmt_try_catch_finally ast.Stmt
	stmts                  []ast.Stmt
	stmt                   ast.Stmt
	teim                   ast.Expr
	expr                   ast.Expr
	tok                    Token
	idents                 []string
	exprs                  []ast.Expr
	pair                   ast.Expr
	pairs                  []ast.Expr
}

const IDENT = 57346
const NUMBER = 57347
const STRING = 57348
const ARRAY = 57349
const VARARG = 57350
const FUNC = 57351
const RETURN = 57352
const VAR = 57353
const THROW = 57354
const IF = 57355
const ELSE = 57356
const FOR = 57357
const IN = 57358
const EQEQ = 57359
const NEQ = 57360
const GE = 57361
const LE = 57362
const OROR = 57363
const ANDAND = 57364
const NEW = 57365
const TRUE = 57366
const FALSE = 57367
const NIL = 57368
const MODULE = 57369
const TRY = 57370
const CATCH = 57371
const FINALLY = 57372
const PLUSEQ = 57373
const MINUSEQ = 57374
const MULEQ = 57375
const DIVEQ = 57376
const ANDEQ = 57377
const OREQ = 57378
const BREAK = 57379
const CONTINUE = 57380
const PLUSPLUS = 57381
const MINUSMINUS = 57382
const POW = 57383
const SHIFTLEFT = 57384
const SHIFTRIGHT = 57385
const UNARY = 57386

var yyToknames = []string{
	"IDENT",
	"NUMBER",
	"STRING",
	"ARRAY",
	"VARARG",
	"FUNC",
	"RETURN",
	"VAR",
	"THROW",
	"IF",
	"ELSE",
	"FOR",
	"IN",
	"EQEQ",
	"NEQ",
	"GE",
	"LE",
	"OROR",
	"ANDAND",
	"NEW",
	"TRUE",
	"FALSE",
	"NIL",
	"MODULE",
	"TRY",
	"CATCH",
	"FINALLY",
	"PLUSEQ",
	"MINUSEQ",
	"MULEQ",
	"DIVEQ",
	"ANDEQ",
	"OREQ",
	"BREAK",
	"CONTINUE",
	"PLUSPLUS",
	"MINUSMINUS",
	"POW",
	"SHIFTLEFT",
	"SHIFTRIGHT",
	" =",
	" ?",
	" :",
	" >",
	" <",
	" +",
	" -",
	" *",
	" /",
	" %",
	"UNARY",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line parser.go.y:492

//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 109,
	17, 0,
	18, 0,
	-2, 64,
	-1, 110,
	17, 0,
	18, 0,
	-2, 65,
}

const yyNprod = 85
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 951

var yyAct = []int{

	1, 91, 138, 31, 33, 34, 35, 36, 38, 39,
	67, 123, 48, 180, 124, 204, 62, 163, 124, 70,
	43, 44, 45, 46, 47, 143, 221, 124, 142, 61,
	188, 192, 187, 95, 41, 9, 42, 57, 59, 182,
	179, 162, 48, 87, 151, 147, 63, 65, 146, 141,
	89, 69, 45, 46, 47, 84, 85, 86, 219, 61,
	218, 215, 63, 212, 41, 93, 42, 57, 59, 209,
	207, 71, 127, 72, 129, 206, 98, 99, 119, 101,
	102, 103, 104, 105, 106, 107, 108, 109, 110, 111,
	112, 113, 114, 115, 116, 117, 118, 63, 148, 139,
	137, 205, 54, 56, 88, 125, 199, 196, 128, 195,
	130, 131, 132, 133, 134, 135, 136, 194, 173, 63,
	171, 161, 159, 154, 48, 49, 50, 96, 214, 210,
	53, 55, 43, 44, 45, 46, 47, 152, 153, 144,
	155, 61, 202, 48, 167, 193, 41, 191, 42, 57,
	59, 190, 165, 178, 176, 174, 63, 63, 217, 63,
	61, 122, 158, 169, 73, 41, 211, 42, 57, 59,
	97, 177, 126, 181, 164, 184, 185, 186, 92, 189,
	168, 197, 63, 170, 90, 172, 166, 156, 140, 198,
	100, 200, 201, 94, 203, 68, 66, 51, 52, 54,
	56, 58, 60, 208, 8, 7, 6, 5, 2, 0,
	0, 213, 0, 0, 0, 216, 0, 0, 0, 0,
	220, 48, 49, 50, 0, 40, 0, 53, 55, 43,
	44, 45, 46, 47, 0, 0, 0, 0, 61, 183,
	0, 0, 0, 41, 0, 42, 57, 59, 51, 52,
	54, 56, 58, 60, 0, 0, 0, 0, 75, 76,
	77, 78, 79, 80, 0, 0, 81, 82, 0, 0,
	0, 74, 48, 49, 50, 0, 40, 0, 53, 55,
	43, 44, 45, 46, 47, 83, 0, 175, 0, 61,
	0, 0, 0, 0, 41, 0, 42, 57, 59, 51,
	52, 54, 56, 58, 60, 0, 0, 0, 0, 75,
	76, 77, 78, 79, 80, 0, 0, 81, 82, 0,
	0, 0, 74, 48, 49, 50, 0, 40, 0, 53,
	55, 43, 44, 45, 46, 47, 83, 0, 121, 0,
	61, 160, 0, 0, 0, 41, 0, 42, 57, 59,
	51, 52, 54, 56, 58, 60, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 48, 49, 50, 0, 40, 0,
	53, 55, 43, 44, 45, 46, 47, 0, 0, 0,
	0, 61, 157, 0, 0, 0, 41, 0, 42, 57,
	59, 51, 52, 54, 56, 58, 60, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 48, 49, 50, 0, 40,
	0, 53, 55, 43, 44, 45, 46, 47, 0, 0,
	0, 0, 61, 0, 0, 0, 0, 41, 150, 42,
	57, 59, 51, 52, 54, 56, 58, 60, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 48, 49, 50, 0,
	40, 149, 53, 55, 43, 44, 45, 46, 47, 0,
	0, 0, 0, 61, 0, 0, 0, 0, 41, 0,
	42, 57, 59, 51, 52, 54, 56, 58, 60, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 48, 49, 50,
	0, 40, 0, 53, 55, 43, 44, 45, 46, 47,
	0, 0, 0, 0, 61, 145, 0, 0, 0, 41,
	0, 42, 57, 59, 51, 52, 54, 56, 58, 60,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 48, 49,
	50, 0, 40, 0, 53, 55, 43, 44, 45, 46,
	47, 0, 0, 0, 0, 61, 0, 120, 0, 0,
	41, 0, 42, 57, 59, 51, 52, 54, 56, 58,
	60, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 48,
	49, 50, 0, 40, 0, 53, 55, 43, 44, 45,
	46, 47, 51, 52, 54, 56, 61, 60, 0, 0,
	0, 41, 0, 42, 57, 59, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 48, 49, 50, 0,
	0, 0, 53, 55, 43, 44, 45, 46, 47, 0,
	0, 0, 0, 61, 0, 0, 0, 0, 41, 0,
	42, 57, 59, 18, 17, 22, 0, 0, 27, 10,
	13, 11, 14, 37, 15, 0, 0, 0, 0, 0,
	0, 0, 30, 23, 24, 25, 12, 16, 0, 0,
	0, 0, 0, 0, 0, 0, 3, 4, 0, 0,
	0, 0, 0, 0, 0, 18, 17, 22, 0, 19,
	27, 10, 13, 11, 14, 28, 15, 29, 0, 0,
	20, 21, 26, 0, 30, 23, 24, 25, 12, 16,
	0, 0, 0, 0, 0, 0, 0, 0, 3, 4,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 19, 0, 0, 0, 0, 32, 28, 0, 29,
	0, 0, 20, 21, 26, 18, 17, 22, 0, 0,
	27, 10, 13, 11, 14, 0, 15, 0, 0, 0,
	0, 0, 0, 0, 30, 23, 24, 25, 12, 16,
	0, 0, 0, 0, 0, 0, 0, 0, 3, 4,
	0, 51, 52, 54, 56, 0, 0, 0, 0, 0,
	0, 19, 0, 0, 0, 0, 0, 28, 0, 29,
	0, 0, 20, 21, 26, 48, 49, 50, 0, 0,
	0, 53, 55, 43, 44, 45, 46, 47, 18, 17,
	22, 0, 61, 27, 0, 0, 0, 41, 0, 42,
	57, 59, 0, 0, 0, 0, 0, 30, 23, 24,
	25, 64, 17, 22, 0, 0, 27, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	30, 23, 24, 25, 19, 0, 0, 0, 0, 0,
	28, 0, 29, 0, 0, 20, 21, 26, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 19, 0, 0,
	0, 0, 0, 28, 0, 29, 0, 0, 20, 21,
	26,
}
var yyPact = []int{

	791, -1000, 731, 791, 791, 791, 689, 791, 791, 588,
	887, 864, 192, 191, -7, 15, 108, -1000, 227, 864,
	864, 864, -1000, -1000, -1000, -1000, 887, 46, 172, 864,
	189, -1000, 791, -1000, -1000, -1000, -1000, 114, -1000, -1000,
	864, 864, 186, 864, 864, 864, 864, 864, 864, 864,
	864, 864, 864, 864, 864, 864, 864, 864, 864, 864,
	864, 887, -1000, 537, 278, 588, 105, -33, -1000, 864,
	156, 791, 864, 791, 864, 864, 864, 864, 864, 864,
	864, -1000, -1000, 887, 102, 102, 102, -62, 184, -9,
	-32, -1000, 93, 486, -10, -1000, -13, 791, 435, 384,
	-1000, 1, 1, 102, 102, 102, 588, -29, -29, 83,
	83, -29, -29, -29, -29, 588, 625, 588, 814, -15,
	887, 887, 791, 887, 183, 333, 864, 65, 282, 64,
	588, 588, 588, 588, 588, 588, 588, -18, -1000, -42,
	166, 182, 172, -1000, 864, -1000, 887, 864, 63, 864,
	-1000, -1000, -1000, -1000, 61, -1000, -1000, 99, 231, -1000,
	98, 142, -1000, 97, -19, -46, 165, -1000, 588, -20,
	180, -1000, 588, -1000, 791, 791, 791, -26, 791, 95,
	91, -28, -1000, 89, 60, 52, 50, 177, 791, 49,
	791, 791, 86, 791, -1000, -1000, -1000, -44, 44, -1000,
	18, 13, 791, 12, 73, 136, -1000, -1000, 6, -1000,
	791, 72, -1000, 4, 791, 128, 3, 2, -1000, 791,
	-31, -1000,
}
var yyPgo = []int{

	0, 0, 208, 207, 206, 205, 204, 35, 16, 1,
	184, 10,
}
var yyR1 = []int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	2, 2, 2, 2, 3, 4, 4, 4, 5, 5,
	5, 6, 6, 6, 6, 9, 10, 10, 10, 11,
	11, 11, 8, 8, 8, 8, 7, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
	7, 7, 7, 7, 7,
}
var yyR2 = []int{

	0, 0, 2, 3, 2, 2, 2, 2, 2, 2,
	1, 2, 2, 5, 4, 9, 5, 7, 7, 4,
	7, 15, 12, 11, 8, 3, 0, 1, 3, 0,
	1, 3, 0, 1, 3, 3, 1, 1, 2, 2,
	2, 1, 1, 1, 1, 5, 4, 3, 3, 7,
	8, 8, 9, 3, 3, 5, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 2, 2, 3,
	3, 3, 3, 4, 4,
}
var yyChk = []int{

	-1000, -1, -2, 37, 38, -3, -4, -5, -6, -7,
	10, 12, 27, 11, 13, 15, 28, 5, 4, 50,
	61, 62, 6, 24, 25, 26, 63, 9, 56, 58,
	23, -1, 55, -1, -1, -1, -1, 14, -1, -1,
	45, 63, 65, 49, 50, 51, 52, 53, 41, 42,
	43, 17, 18, 47, 19, 48, 20, 66, 21, 67,
	22, 58, -8, -7, 4, -7, 4, -11, 4, 58,
	4, 56, 58, 56, 44, 31, 32, 33, 34, 35,
	36, 39, 40, 58, -7, -7, -7, -8, 58, 4,
	-10, -9, 6, -7, 4, -1, 13, 56, -7, -7,
	4, -7, -7, -7, -7, -7, -7, -7, -7, -7,
	-7, -7, -7, -7, -7, -7, -7, -7, -7, -8,
	60, 60, 56, 44, 60, -7, 16, -1, -7, -1,
	-7, -7, -7, -7, -7, -7, -7, -8, 64, -11,
	4, 58, 60, 57, 46, 59, 58, 58, -1, 46,
	64, 59, -8, -8, -1, -8, 4, 59, -7, 57,
	59, 57, 59, 59, 8, -11, 4, -9, -7, -8,
	-7, 57, -7, 57, 56, 56, 56, 29, 56, 59,
	59, 8, 59, 59, -1, -1, -1, 58, 56, -1,
	56, 56, 59, 56, 57, 57, 57, 4, -1, 57,
	-1, -1, 56, -1, 59, 57, 57, 57, -1, 57,
	56, 30, 57, -1, 56, 57, -1, 30, 57, 56,
	-1, 57,
}
var yyDef = []int{

	1, -2, 1, 1, 1, 1, 1, 1, 1, 10,
	32, 0, 0, 29, 0, 0, 0, 36, 37, 0,
	0, 0, 41, 42, 43, 44, 32, 0, 26, 0,
	0, 2, 1, 4, 5, 6, 7, 0, 8, 9,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 32, 11, 33, 37, 12, 0, 0, 30, 0,
	0, 1, 0, 1, 0, 0, 0, 0, 0, 0,
	0, 77, 78, 32, 38, 39, 40, 0, 29, 0,
	0, 27, 0, 0, 0, 3, 0, 1, 0, 0,
	47, 56, 57, 58, 59, 60, 61, 62, 63, -2,
	-2, 66, 67, 68, 69, 79, 80, 81, 82, 0,
	32, 32, 1, 32, 0, 0, 0, 0, 0, 0,
	70, 71, 72, 73, 74, 75, 76, 0, 48, 0,
	30, 29, 0, 53, 0, 54, 32, 0, 0, 0,
	46, 84, 34, 35, 0, 14, 31, 0, 0, 19,
	0, 0, 83, 0, 0, 0, 30, 28, 25, 0,
	0, 16, 45, 13, 1, 1, 1, 0, 1, 0,
	0, 0, 55, 0, 0, 0, 0, 0, 1, 0,
	1, 1, 0, 1, 17, 18, 20, 0, 0, 49,
	0, 0, 1, 0, 0, 24, 50, 51, 0, 15,
	1, 0, 52, 0, 1, 23, 0, 0, 22, 1,
	0, 21,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 61, 3, 3, 3, 53, 67, 3,
	58, 59, 51, 49, 60, 50, 65, 52, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 46, 55,
	48, 44, 47, 45, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 63, 3, 64, 62, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 56, 66, 57,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 54,
}
var yyTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

const yyFlag = -1000

func yyTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 4 && c-4 < len(yyToknames) {
		if yyToknames[c-4] != "" {
			return yyToknames[c-4]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yylex1(lex yyLexer, lval *yySymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		c = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			c = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		c = yyTok3[i+0]
		if c == char {
			c = yyTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(c), uint(char))
	}
	return c
}

func yyParse(yylex yyLexer) int {
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, yyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yychar), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yychar < 0 {
		yychar = yylex1(yylex, &yylval)
	}
	yyn += yychar
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yychar { /* valid shift */
		yychar = -1
		yyVAL = yylval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yychar < 0 {
			yychar = yylex1(yylex, &yylval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yychar {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error("syntax error")
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yychar))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}
			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		//line parser.go.y:55
		{
			yyVAL.stmts = nil
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 2:
		//line parser.go.y:62
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
				yyS[yypt-1].stmt.SetPos(l.pos)
			}
		}
	case 3:
		//line parser.go.y:70
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-2].stmt}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
				yyS[yypt-2].stmt.SetPos(l.pos)
			}
		}
	case 4:
		//line parser.go.y:78
		{
			yyVAL.stmts = append([]ast.Stmt{&ast.BreakStmt{}}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 5:
		//line parser.go.y:85
		{
			yyVAL.stmts = append([]ast.Stmt{&ast.ContinueStmt{}}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 6:
		//line parser.go.y:92
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_var}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 7:
		//line parser.go.y:99
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_if}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 8:
		//line parser.go.y:106
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_for}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 9:
		//line parser.go.y:113
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_try_catch_finally}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 10:
		//line parser.go.y:121
		{
			yyVAL.stmt = &ast.ExprStmt{Expr: yyS[yypt-0].expr}
		}
	case 11:
		//line parser.go.y:125
		{
			yyVAL.stmt = &ast.ReturnStmt{Exprs: yyS[yypt-0].exprs}
		}
	case 12:
		//line parser.go.y:129
		{
			yyVAL.stmt = &ast.ThrowStmt{Expr: yyS[yypt-0].expr}
		}
	case 13:
		//line parser.go.y:133
		{
			yyVAL.stmt = &ast.ModuleStmt{Name: yyS[yypt-3].tok.lit, Stmts: yyS[yypt-1].stmts}
		}
	case 14:
		//line parser.go.y:138
		{
			yyVAL.stmt_var = &ast.VarStmt{Names: yyS[yypt-2].idents, Exprs: yyS[yypt-0].exprs}
		}
	case 15:
		//line parser.go.y:143
		{
			yyS[yypt-8].stmt_if.(*ast.IfStmt).ElseIf = append(yyS[yypt-8].stmt_if.(*ast.IfStmt).ElseIf, &ast.IfStmt{If: yyS[yypt-4].expr, Then: yyS[yypt-1].stmts})
		}
	case 16:
		//line parser.go.y:147
		{
			if yyS[yypt-4].stmt_if.(*ast.IfStmt).Else != nil {
				yylex.Error("multiple else statement")
			} else {
				yyS[yypt-4].stmt_if.(*ast.IfStmt).Else = yyS[yypt-1].stmts
			}
		}
	case 17:
		//line parser.go.y:155
		{
			yyVAL.stmt_if = &ast.IfStmt{If: yyS[yypt-4].expr, Then: yyS[yypt-1].stmts}
		}
	case 18:
		//line parser.go.y:160
		{
			yyVAL.stmt_for = &ast.ForStmt{Var: yyS[yypt-5].tok.lit, Value: yyS[yypt-3].expr, Stmts: yyS[yypt-1].stmts}
		}
	case 19:
		//line parser.go.y:164
		{
			yyVAL.stmt_for = &ast.LoopStmt{Stmts: yyS[yypt-1].stmts}
		}
	case 20:
		//line parser.go.y:168
		{
			yyVAL.stmt_for = &ast.LoopStmt{Expr: yyS[yypt-4].expr, Stmts: yyS[yypt-1].stmts}
		}
	case 21:
		//line parser.go.y:173
		{
			yyVAL.stmt_try_catch_finally = &ast.TryStmt{Try: yyS[yypt-12].stmts, Var: yyS[yypt-8].tok.lit, Catch: yyS[yypt-5].stmts, Finally: yyS[yypt-1].stmts}
		}
	case 22:
		//line parser.go.y:177
		{
			yyVAL.stmt_try_catch_finally = &ast.TryStmt{Try: yyS[yypt-9].stmts, Catch: yyS[yypt-5].stmts, Finally: yyS[yypt-1].stmts}
		}
	case 23:
		//line parser.go.y:181
		{
			yyVAL.stmt_try_catch_finally = &ast.TryStmt{Try: yyS[yypt-8].stmts, Var: yyS[yypt-4].tok.lit, Catch: yyS[yypt-1].stmts}
		}
	case 24:
		//line parser.go.y:185
		{
			yyVAL.stmt_try_catch_finally = &ast.TryStmt{Try: yyS[yypt-5].stmts, Catch: yyS[yypt-1].stmts}
		}
	case 25:
		//line parser.go.y:190
		{
			yyVAL.pair = &ast.PairExpr{Key: yyS[yypt-2].tok.lit, Value: yyS[yypt-0].expr}
		}
	case 26:
		//line parser.go.y:195
		{
			yyVAL.pairs = []ast.Expr{}
		}
	case 27:
		//line parser.go.y:199
		{
			yyVAL.pairs = []ast.Expr{yyS[yypt-0].pair}
		}
	case 28:
		//line parser.go.y:203
		{
			yyVAL.pairs = append(yyS[yypt-2].pairs, yyS[yypt-0].pair)
		}
	case 29:
		//line parser.go.y:208
		{
			yyVAL.idents = []string{}
		}
	case 30:
		//line parser.go.y:212
		{
			yyVAL.idents = []string{yyS[yypt-0].tok.lit}
		}
	case 31:
		//line parser.go.y:216
		{
			yyVAL.idents = append(yyS[yypt-2].idents, yyS[yypt-0].tok.lit)
		}
	case 32:
		//line parser.go.y:221
		{
			yyVAL.exprs = []ast.Expr{}
		}
	case 33:
		//line parser.go.y:225
		{
			yyVAL.exprs = []ast.Expr{yyS[yypt-0].expr}
		}
	case 34:
		//line parser.go.y:229
		{
			yyVAL.exprs = append([]ast.Expr{yyS[yypt-2].expr}, yyS[yypt-0].exprs...)
		}
	case 35:
		//line parser.go.y:233
		{
			yyVAL.exprs = append([]ast.Expr{&ast.IdentExpr{Lit: yyS[yypt-2].tok.lit}}, yyS[yypt-0].exprs...)
		}
	case 36:
		//line parser.go.y:238
		{
			yyVAL.expr = &ast.NumberExpr{Lit: yyS[yypt-0].tok.lit}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 37:
		//line parser.go.y:243
		{
			yyVAL.expr = &ast.IdentExpr{Lit: yyS[yypt-0].tok.lit}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 38:
		//line parser.go.y:248
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "-", Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 39:
		//line parser.go.y:253
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "!", Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 40:
		//line parser.go.y:258
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "^", Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 41:
		//line parser.go.y:263
		{
			yyVAL.expr = &ast.StringExpr{Lit: yyS[yypt-0].tok.lit}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 42:
		//line parser.go.y:268
		{
			yyVAL.expr = &ast.ConstExpr{Value: true}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 43:
		//line parser.go.y:273
		{
			yyVAL.expr = &ast.ConstExpr{Value: false}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 44:
		//line parser.go.y:278
		{
			yyVAL.expr = &ast.ConstExpr{Value: unsafe.Pointer(nil)}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 45:
		//line parser.go.y:283
		{
			yyVAL.expr = &ast.TernaryOpExpr{Expr: yyS[yypt-4].expr, Lhs: yyS[yypt-2].expr, Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 46:
		//line parser.go.y:288
		{
			yyVAL.expr = &ast.ItemExpr{Value: yyS[yypt-3].expr, Index: yyS[yypt-1].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 47:
		//line parser.go.y:293
		{
			yyVAL.expr = &ast.MemberExpr{Expr: yyS[yypt-2].expr, Method: yyS[yypt-0].tok.lit}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 48:
		//line parser.go.y:298
		{
			yyVAL.expr = &ast.ArrayExpr{Exprs: yyS[yypt-1].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 49:
		//line parser.go.y:303
		{
			yyVAL.expr = &ast.FuncExpr{Args: yyS[yypt-4].idents, Stmts: yyS[yypt-1].stmts}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 50:
		//line parser.go.y:308
		{
			yyVAL.expr = &ast.FuncExpr{Args: []string{yyS[yypt-5].tok.lit}, Stmts: yyS[yypt-1].stmts, VarArg: true}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 51:
		//line parser.go.y:313
		{
			yyVAL.expr = &ast.FuncExpr{Name: yyS[yypt-6].tok.lit, Args: yyS[yypt-4].idents, Stmts: yyS[yypt-1].stmts}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 52:
		//line parser.go.y:318
		{
			yyVAL.expr = &ast.FuncExpr{Name: yyS[yypt-7].tok.lit, Args: []string{yyS[yypt-5].tok.lit}, Stmts: yyS[yypt-1].stmts, VarArg: true}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 53:
		//line parser.go.y:323
		{
			mapExpr := make(map[string]ast.Expr)
			for _, v := range yyS[yypt-1].pairs {
				mapExpr[v.(*ast.PairExpr).Key] = v.(*ast.PairExpr).Value
			}
			yyVAL.expr = &ast.MapExpr{MapExpr: mapExpr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 54:
		//line parser.go.y:332
		{
			yyVAL.expr = &ast.ParenExpr{SubExpr: yyS[yypt-1].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 55:
		//line parser.go.y:337
		{
			yyVAL.expr = &ast.NewExpr{Name: yyS[yypt-3].tok.lit, SubExprs: yyS[yypt-1].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 56:
		//line parser.go.y:342
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "+", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 57:
		//line parser.go.y:347
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "-", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 58:
		//line parser.go.y:352
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "*", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 59:
		//line parser.go.y:357
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "/", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 60:
		//line parser.go.y:362
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "%", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 61:
		//line parser.go.y:367
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "**", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 62:
		//line parser.go.y:372
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "<<", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 63:
		//line parser.go.y:377
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: ">>", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 64:
		//line parser.go.y:382
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "==", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 65:
		//line parser.go.y:387
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "!=", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 66:
		//line parser.go.y:392
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: ">", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 67:
		//line parser.go.y:397
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: ">=", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 68:
		//line parser.go.y:402
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "<", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 69:
		//line parser.go.y:407
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "<=", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 70:
		//line parser.go.y:417
		{
			yyVAL.expr = &ast.LetExpr{Name: yyS[yypt-2].tok.lit, Operator: "=", Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 71:
		//line parser.go.y:422
		{
			yyVAL.expr = &ast.LetExpr{Name: yyS[yypt-2].tok.lit, Operator: "+", Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 72:
		//line parser.go.y:427
		{
			yyVAL.expr = &ast.LetExpr{Name: yyS[yypt-2].tok.lit, Operator: "-", Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 73:
		//line parser.go.y:432
		{
			yyVAL.expr = &ast.LetExpr{Name: yyS[yypt-2].tok.lit, Operator: "*", Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 74:
		//line parser.go.y:437
		{
			yyVAL.expr = &ast.LetExpr{Name: yyS[yypt-2].tok.lit, Operator: "/", Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 75:
		//line parser.go.y:442
		{
			yyVAL.expr = &ast.LetExpr{Name: yyS[yypt-2].tok.lit, Operator: "&", Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 76:
		//line parser.go.y:447
		{
			yyVAL.expr = &ast.LetExpr{Name: yyS[yypt-2].tok.lit, Operator: "|", Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 77:
		//line parser.go.y:452
		{
			yyVAL.expr = &ast.AssocExpr{Name: yyS[yypt-1].tok.lit, Operator: "++"}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 78:
		//line parser.go.y:457
		{
			yyVAL.expr = &ast.AssocExpr{Name: yyS[yypt-1].tok.lit, Operator: "--"}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 79:
		//line parser.go.y:462
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "|", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 80:
		//line parser.go.y:467
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "||", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 81:
		//line parser.go.y:472
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "&", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 82:
		//line parser.go.y:477
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "&&", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 83:
		//line parser.go.y:482
		{
			yyVAL.expr = &ast.CallExpr{Name: yyS[yypt-3].tok.lit, SubExprs: yyS[yypt-1].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 84:
		//line parser.go.y:487
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyS[yypt-3].expr, SubExprs: yyS[yypt-1].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	}
	goto yystack /* stack new state and value */
}
