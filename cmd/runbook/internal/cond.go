package internal

import (
	"fmt"
	"regexp"
	"strconv"
)

// Bool is a tri-state used by condition evaluation so that "we genuinely
// don't know" (undefined threshold param, empty result, comparison error) is
// not silently collapsed into true or false.
type Bool int

const (
	False Bool = iota
	True
	Unknown
)

func (b Bool) String() string {
	switch b {
	case True:
		return "true"
	case False:
		return "false"
	default:
		return "unknown"
	}
}

// CondEnv is the evaluation context for a condition expression.
type CondEnv struct {
	// Series is the (per-check or per-step) result set. Empty lists yield
	// vacuous quantifier results, which is exactly the empty-result rule.
	Series []Series
	// Params holds numeric parameters such as fRoleDON.
	Params map[string]float64
}

// MatchResult evaluates cond against env, returning a tri-state Bool.
func MatchResult(cond string, env CondEnv) (Bool, error) {
	p, err := newCondParser(cond)
	if err != nil {
		return Unknown, err
	}
	e, err := p.parse()
	if err != nil {
		return Unknown, fmt.Errorf("condition %q: %w", cond, err)
	}
	return e.eval(env), nil
}

// --- AST ---

type condNode interface{ eval(env CondEnv) Bool }

type andNode struct{ l, r condNode }
type orNode struct{ l, r condNode }
type notNode struct{ c condNode }
type knownNode struct{ param string }

type comparison struct {
	op   cmpOp
	l, r operand
}

type quantNode struct {
	kind quantKind // any | all | no
	ls   *labelSel
	op   cmpOp
	rhs  numExpr
}

type labelSel struct {
	name string
	op   labelOp
	re   *regexp.Regexp
	str  string
}

type labelOp int

const (
	labelEq labelOp = iota
	labelNeq
	labelRe
	labelNre
)

// operand is either the "result" value (max over series) or a numeric expr.
type operand struct {
	isResult bool
	expr     numExpr
}

type quantKind int

const (
	quantAny quantKind = iota
	quantAll
	quantNo
)

func (n andNode) eval(env CondEnv) Bool { return bAnd(n.l.eval(env), n.r.eval(env)) }
func (n orNode) eval(env CondEnv) Bool  { return bOr(n.l.eval(env), n.r.eval(env)) }
func (n notNode) eval(env CondEnv) Bool {
	switch n.c.eval(env) {
	case True:
		return False
	case False:
		return True
	default:
		return Unknown
	}
}
func (n knownNode) eval(env CondEnv) Bool {
	_, ok := env.Params[n.param]
	return boolOf(ok)
}

func (c comparison) eval(env CondEnv) Bool {
	lv, lu := c.l.value(env)
	rv, ru := c.r.value(env)
	if lu || ru {
		return Unknown
	}
	return boolOf(c.op.apply(lv, rv))
}

func (o operand) value(env CondEnv) (float64, bool) {
	if o.isResult {
		// The engine only evaluates a bare `result` comparison on a set it has
		// already decided to treat as meaningful: an always_emitted:false
		// counter's empty result means "the event never fired", i.e. value 0
		// (annual). always_emitted:true empty (pipeline broken) is short-
		// circuited to UNKNOWN by the checklist before reaching here, and the
		// decision graph refuses to evaluate on empty at all. So max-of-empty
		// is 0, not unknown.
		max := 0.0
		for _, s := range env.Series {
			if s.Value > max {
				max = s.Value
			}
		}
		return max, false
	}
	v, err := o.expr.eval(env)
	if err != nil {
		return 0, true
	}
	return v, false
}

func (q quantNode) eval(env CondEnv) Bool {
	rhs, err := q.rhs.eval(env)
	verboten := err != nil // undefined param / division by zero
	for _, s := range env.Series {
		if q.ls != nil && !q.ls.matches(s) {
			continue
		}
		if verboten {
			continue // can't judge with a bad rhs; treat series as not qualifying
		}
		res := q.op.apply(s.Value, rhs)
		switch q.kind {
		case quantAny:
			if res {
				return True
			}
		case quantAll:
			if !res {
				return False
			}
		case quantNo:
			if res {
				return False
			}
		}
	}
	// Vacuous over no matching series.
	switch q.kind {
	case quantAny:
		if verboten {
			return Unknown
		}
		return False
	case quantAll, quantNo:
		if verboten {
			return Unknown
		}
		return True
	}
	return Unknown
}

func (ls *labelSel) matches(s Series) bool {
	v, ok := s.Labels[ls.name]
	switch ls.op {
	case labelEq:
		return ok && v == ls.str
	case labelNeq:
		return !ok || v != ls.str
	case labelRe:
		return ok && ls.re.MatchString(v)
	case labelNre:
		return !ok || !ls.re.MatchString(v)
	}
	return false
}

// --- numeric expressions ---

type numExpr interface {
	eval(env CondEnv) (float64, error)
}

type numLit struct{ v float64 }
type numParam struct{ name string }
type numBin struct {
	op   byte // + - * /
	l, r numExpr
}

func (n numLit) eval(env CondEnv) (float64, error) { return n.v, nil }
func (n numParam) eval(env CondEnv) (float64, error) {
	if v, ok := env.Params[n.name]; ok {
		return v, nil
	}
	return 0, fmt.Errorf("undefined parameter %q", n.name)
}
func (n numBin) eval(env CondEnv) (float64, error) {
	l, err := n.l.eval(env)
	if err != nil {
		return 0, err
	}
	r, err := n.r.eval(env)
	if err != nil {
		return 0, err
	}
	switch n.op {
	case '+':
		return l + r, nil
	case '-':
		return l - r, nil
	case '*':
		return l * r, nil
	case '/':
		if r == 0 {
			return 0, fmt.Errorf("division by zero")
		}
		return l / r, nil
	}
	return 0, fmt.Errorf("bad numeric operator %q", n.op)
}

// --- lexical tokens ---

type tokKind int

const (
	tokIdent tokKind = iota
	tokNum
	tokStr
	tokOp
	tokEOF
)

type token struct {
	kind tokKind
	text string
}

var twoCharOps = []string{">=", "<=", "==", "!=", "=~", "!~"}

func lexCondition(s string) ([]token, error) {
	var out []token
	rs := []rune(s)
	i := 0
	for i < len(rs) {
		c := rs[i]
		switch {
		case c == ' ' || c == '\t' || c == '\n':
			i++
		case c == '(' || c == ')' || c == '+' || c == '-' || c == '*' || c == '/':
			out = append(out, token{tokOp, string(c)})
			i++
		case c == '>' || c == '<' || c == '=' || c == '!':
			if i+1 < len(rs) {
				two := string(rs[i]) + string(rs[i+1])
				match := false
				for _, op := range twoCharOps {
					if op == two {
						match = true
						break
					}
				}
				if match {
					out = append(out, token{tokOp, two})
					i += 2
					continue
				}
			}
			out = append(out, token{tokOp, string(c)})
			i++
		case c >= '0' && c <= '9' || c == '.':
			j := i
			for j < len(rs) && (rs[j] >= '0' && rs[j] <= '9' || rs[j] == '.') {
				j++
			}
			out = append(out, token{tokNum, string(rs[i:j])})
			i = j
		case isIdentStart(c):
			j := i
			for j < len(rs) && isIdentChar(rs[j]) {
				j++
			}
			out = append(out, token{tokIdent, string(rs[i:j])})
			i = j
		case c == '"' || c == '\'':
			quote := c
			j := i + 1
			for j < len(rs) && rs[j] != quote {
				j++
			}
			if j >= len(rs) {
				return nil, fmt.Errorf("unterminated string at pos %d", i)
			}
			out = append(out, token{tokStr, string(rs[i+1 : j])})
			i = j + 1
		default:
			return nil, fmt.Errorf("unexpected character %q at pos %d", string(c), i)
		}
	}
	out = append(out, token{tokEOF, ""})
	return out, nil
}

func isIdentStart(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '_'
}
func isIdentChar(c rune) bool { return isIdentStart(c) || (c >= '0' && c <= '9') }

// --- comparison operators ---

type cmpOp int

const (
	opGT cmpOp = iota
	opGE
	opLT
	opLE
	opEQ
	opNE
)

func (o cmpOp) apply(a, b float64) bool {
	switch o {
	case opGT:
		return a > b
	case opGE:
		return a >= b
	case opLT:
		return a < b
	case opLE:
		return a <= b
	case opEQ:
		return a == b
	case opNE:
		return a != b
	}
	return false
}

// --- parser ---

type condParser struct {
	toks []token
	i    int
}

func newCondParser(s string) (*condParser, error) {
	toks, err := lexCondition(s)
	if err != nil {
		return nil, err
	}
	return &condParser{toks: toks}, nil
}

func (p *condParser) peek() token { return p.toks[p.i] }
func (p *condParser) next() token { t := p.toks[p.i]; p.i++; return t }
func (p *condParser) isIdent(kw string) bool {
	t := p.peek()
	return t.kind == tokIdent && t.text == kw
}
func (p *condParser) expectIdent(kw string) error {
	if !p.isIdent(kw) {
		return fmt.Errorf("expected %q, got %q", kw, p.peek().text)
	}
	p.i++
	return nil
}

func (p *condParser) parse() (condNode, error) {
	n, err := p.parseOr()
	if err != nil {
		return nil, err
	}
	if p.peek().kind != tokEOF {
		return nil, fmt.Errorf("unexpected trailing token %q", p.peek().text)
	}
	return n, nil
}

func (p *condParser) parseOr() (condNode, error) {
	l, err := p.parseAnd()
	if err != nil {
		return nil, err
	}
	for p.isIdent("or") {
		p.i++
		r, err := p.parseAnd()
		if err != nil {
			return nil, err
		}
		l = orNode{l, r}
	}
	return l, nil
}

func (p *condParser) parseAnd() (condNode, error) {
	l, err := p.parseNot()
	if err != nil {
		return nil, err
	}
	for p.isIdent("and") {
		p.i++
		r, err := p.parseNot()
		if err != nil {
			return nil, err
		}
		l = andNode{l, r}
	}
	return l, nil
}

func (p *condParser) parseNot() (condNode, error) {
	if p.isIdent("not") {
		p.i++
		c, err := p.parseNot()
		if err != nil {
			return nil, err
		}
		return notNode{c}, nil
	}
	return p.parseAtom()
}

func (p *condParser) parseAtom() (condNode, error) {
	t := p.peek()
	if t.kind == tokOp && t.text == "(" {
		p.i++
		e, err := p.parseOr()
		if err != nil {
			return nil, err
		}
		if err := expectOp(p, ")"); err != nil {
			return nil, err
		}
		return e, nil
	}
	if t.kind == tokNum {
		// A general comparison such as `2*fRoleDON+1 <= result < 3*fRoleDON+1`
		// that starts with a numeric expression rather than `result`.
		return p.parseGeneralComparison()
	}
	if t.kind != tokIdent {
		return nil, fmt.Errorf("unexpected token %q", t.text)
	}
	switch t.text {
	case "known":
		p.i++
		if p.peek().kind != tokIdent {
			return nil, fmt.Errorf("known requires a parameter name")
		}
		return knownNode{param: p.next().text}, nil
	case "result":
		// parseResultComparison reads the "result" token itself.
		return p.parseResultComparison()
	case "any", "all", "no":
		return p.parseQuantifier()
	default:
		// General comparison that starts with a parameter expression.
		return p.parseGeneralComparison()
	}
}

func (p *condParser) parseGeneralComparison() (condNode, error) {
	comps, err := p.parseComparisonChain()
	if err != nil {
		return nil, err
	}
	return foldComparisons(comps), nil
}

func foldComparisons(comps []condNode) condNode {
	if len(comps) == 1 {
		return comps[0]
	}
	n := comps[0]
	for _, c := range comps[1:] {
		n = andNode{n, c}
	}
	return n
}

func (p *condParser) parseResultComparison() (condNode, error) {
	comps, err := p.parseComparisonChain()
	if err != nil {
		return nil, err
	}
	return foldComparisons(comps), nil
}

func (p *condParser) parseComparisonChain() ([]condNode, error) {
	var comps []condNode
	left, err := p.parseOperand()
	if err != nil {
		return nil, err
	}
	for {
		op, ok := p.tryRelop()
		if !ok {
			break
		}
		right, err := p.parseOperand()
		if err != nil {
			return nil, err
		}
		comps = append(comps, comparison{op: op, l: left, r: right})
		left = right
	}
	if len(comps) == 0 {
		return nil, fmt.Errorf("expected a comparison")
	}
	return comps, nil
}

func (p *condParser) parseOperand() (operand, error) {
	if p.isIdent("result") {
		p.i++
		return operand{isResult: true}, nil
	}
	expr, err := p.parseNumExpr()
	if err != nil {
		return operand{}, err
	}
	return operand{expr: expr}, nil
}

func (p *condParser) parseQuantifier() (condNode, error) {
	kindTok := p.next().text
	var kind quantKind
	switch kindTok {
	case "any":
		kind = quantAny
	case "all":
		kind = quantAll
	case "no":
		kind = quantNo
	}
	if err := p.expectIdent("series"); err != nil {
		return nil, err
	}
	var ls *labelSel
	if p.isIdent("with") {
		p.i++
		v, err := p.parseLabelSel()
		if err != nil {
			return nil, err
		}
		ls = &v
	}
	op, ok := p.tryRelop()
	if !ok {
		return nil, fmt.Errorf("quantifier missing comparison operator")
	}
	rhs, err := p.parseNumExpr()
	if err != nil {
		return nil, err
	}
	return quantNode{kind: kind, ls: ls, op: op, rhs: rhs}, nil
}

func (p *condParser) parseLabelSel() (labelSel, error) {
	if p.peek().kind != tokIdent {
		return labelSel{}, fmt.Errorf("expected label name")
	}
	name := p.next().text
	opTok := p.next()
	if opTok.kind != tokOp {
		return labelSel{}, fmt.Errorf("expected label operator after %q", name)
	}
	valTok := p.next()
	if valTok.kind != tokStr && valTok.kind != tokIdent {
		return labelSel{}, fmt.Errorf("expected string value for label %q", name)
	}
	var lo labelOp
	switch opTok.text {
	case "=", "==":
		lo = labelEq
	case "!=":
		lo = labelNeq
	case "=~":
		lo = labelRe
	case "!~":
		lo = labelNre
	default:
		return labelSel{}, fmt.Errorf("unsupported label operator %q", opTok.text)
	}
	ls := labelSel{name: name, op: lo, str: valTok.text}
	if lo == labelRe || lo == labelNre {
		re, err := regexp.Compile(valTok.text)
		if err != nil {
			return labelSel{}, fmt.Errorf("bad label regexp: %w", err)
		}
		ls.re = re
	}
	return ls, nil
}

func (p *condParser) tryRelop() (cmpOp, bool) {
	t := p.peek()
	if t.kind != tokOp {
		return 0, false
	}
	switch t.text {
	case ">":
		p.i++
		return opGT, true
	case ">=":
		p.i++
		return opGE, true
	case "<":
		p.i++
		return opLT, true
	case "<=":
		p.i++
		return opLE, true
	case "==":
		p.i++
		return opEQ, true
	case "!=":
		p.i++
		return opNE, true
	}
	return 0, false
}

// --- numeric expression parsing ---

func (p *condParser) parseNumExpr() (numExpr, error) { return p.parseAddSub() }

func (p *condParser) parseAddSub() (numExpr, error) {
	l, err := p.parseMulDiv()
	if err != nil {
		return nil, err
	}
	for {
		t := p.peek()
		if t.kind == tokOp && (t.text == "+" || t.text == "-") {
			p.i++
			r, err := p.parseMulDiv()
			if err != nil {
				return nil, err
			}
			l = numBin{op: t.text[0], l: l, r: r}
		} else {
			return l, nil
		}
	}
}

func (p *condParser) parseMulDiv() (numExpr, error) {
	l, err := p.parseFactor()
	if err != nil {
		return nil, err
	}
	for {
		t := p.peek()
		if t.kind == tokOp && (t.text == "*" || t.text == "/") {
			p.i++
			r, err := p.parseFactor()
			if err != nil {
				return nil, err
			}
			l = numBin{op: t.text[0], l: l, r: r}
		} else {
			return l, nil
		}
	}
}

func (p *condParser) parseFactor() (numExpr, error) {
	t := p.peek()
	if t.kind == tokOp && t.text == "(" {
		p.i++
		e, err := p.parseNumExpr()
		if err != nil {
			return nil, err
		}
		if err := expectOp(p, ")"); err != nil {
			return nil, err
		}
		return e, nil
	}
	if t.kind == tokNum {
		p.i++
		f, err := strconv.ParseFloat(t.text, 64)
		if err != nil {
			return nil, err
		}
		return numLit{f}, nil
	}
	if t.kind == tokIdent {
		p.i++
		return numParam{t.text}, nil
	}
	return nil, fmt.Errorf("expected number, got %q", t.text)
}

func expectOp(p *condParser, op string) error {
	t := p.peek()
	if t.kind == tokOp && t.text == op {
		p.i++
		return nil
	}
	return fmt.Errorf("expected %q, got %q", op, t.text)
}

// --- boolean helpers ---

func bAnd(l, r Bool) Bool {
	if l == False || r == False {
		return False
	}
	if l == Unknown || r == Unknown {
		return Unknown
	}
	return True
}
func bOr(l, r Bool) Bool {
	if l == True || r == True {
		return True
	}
	if l == Unknown || r == Unknown {
		return Unknown
	}
	return False
}
func boolOf(b bool) Bool {
	if b {
		return True
	}
	return False
}
