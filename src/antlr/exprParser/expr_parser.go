// Code generated from /home/michelep/go/src/mSteelProto/antlr/expr.g4 by ANTLR 4.9.1. DO NOT EDIT.

package exprParser // expr
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 29, 56, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 5, 2, 25, 10,
	2, 3, 2, 3, 2, 3, 2, 5, 2, 30, 10, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 7, 2, 47, 10,
	2, 12, 2, 14, 2, 50, 11, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 2, 3, 2, 5, 2,
	4, 6, 2, 8, 3, 2, 18, 19, 3, 2, 16, 17, 3, 2, 11, 14, 3, 2, 9, 10, 3, 2,
	6, 7, 4, 2, 5, 5, 24, 27, 2, 63, 2, 29, 3, 2, 2, 2, 4, 51, 3, 2, 2, 2,
	6, 53, 3, 2, 2, 2, 8, 9, 8, 2, 1, 2, 9, 10, 7, 21, 2, 2, 10, 11, 5, 2,
	2, 2, 11, 12, 7, 22, 2, 2, 12, 30, 3, 2, 2, 2, 13, 30, 5, 6, 4, 2, 14,
	15, 7, 8, 2, 2, 15, 30, 5, 2, 2, 11, 16, 17, 7, 20, 2, 2, 17, 18, 7, 21,
	2, 2, 18, 19, 5, 2, 2, 2, 19, 20, 7, 23, 2, 2, 20, 21, 5, 2, 2, 2, 21,
	22, 7, 22, 2, 2, 22, 30, 3, 2, 2, 2, 23, 25, 7, 3, 2, 2, 24, 23, 3, 2,
	2, 2, 24, 25, 3, 2, 2, 2, 25, 26, 3, 2, 2, 2, 26, 30, 5, 4, 3, 2, 27, 28,
	7, 4, 2, 2, 28, 30, 5, 4, 3, 2, 29, 8, 3, 2, 2, 2, 29, 13, 3, 2, 2, 2,
	29, 14, 3, 2, 2, 2, 29, 16, 3, 2, 2, 2, 29, 24, 3, 2, 2, 2, 29, 27, 3,
	2, 2, 2, 30, 48, 3, 2, 2, 2, 31, 32, 12, 10, 2, 2, 32, 33, 9, 2, 2, 2,
	33, 47, 5, 2, 2, 11, 34, 35, 12, 9, 2, 2, 35, 36, 9, 3, 2, 2, 36, 47, 5,
	2, 2, 10, 37, 38, 12, 8, 2, 2, 38, 39, 9, 4, 2, 2, 39, 47, 5, 2, 2, 9,
	40, 41, 12, 7, 2, 2, 41, 42, 9, 5, 2, 2, 42, 47, 5, 2, 2, 8, 43, 44, 12,
	6, 2, 2, 44, 45, 9, 6, 2, 2, 45, 47, 5, 2, 2, 7, 46, 31, 3, 2, 2, 2, 46,
	34, 3, 2, 2, 2, 46, 37, 3, 2, 2, 2, 46, 40, 3, 2, 2, 2, 46, 43, 3, 2, 2,
	2, 47, 50, 3, 2, 2, 2, 48, 46, 3, 2, 2, 2, 48, 49, 3, 2, 2, 2, 49, 3, 3,
	2, 2, 2, 50, 48, 3, 2, 2, 2, 51, 52, 7, 28, 2, 2, 52, 5, 3, 2, 2, 2, 53,
	54, 9, 7, 2, 2, 54, 7, 3, 2, 2, 2, 6, 24, 29, 46, 48,
}
var literalNames = []string{
	"", "'this.'", "'ext.'", "'null'", "'and'", "'or'", "'not'", "'=='", "'=/='",
	"'<'", "'<='", "'>'", "'>='", "'='", "'+'", "'-'", "'/'", "'*'", "'concat'",
	"'('", "')'", "','",
}
var symbolicNames = []string{
	"", "THIS", "EXT", "UNDEF", "AND", "OR", "NOT", "EQ", "NEQ", "LT", "LEQ",
	"GT", "GEQ", "ASSIGN", "PLUS", "MINUS", "DIV", "MUL", "CONCAT", "ROUNDLEFT",
	"ROUNDRIGHT", "COMMA", "BOOL", "INT", "DEC", "STR", "ID", "WS",
}

var ruleNames = []string{
	"exp", "id", "val",
}

type exprParser struct {
	*antlr.BaseParser
}

// NewexprParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *exprParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewexprParser(input antlr.TokenStream) *exprParser {
	this := new(exprParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "expr.g4"

	return this
}

// exprParser tokens.
const (
	exprParserEOF        = antlr.TokenEOF
	exprParserTHIS       = 1
	exprParserEXT        = 2
	exprParserUNDEF      = 3
	exprParserAND        = 4
	exprParserOR         = 5
	exprParserNOT        = 6
	exprParserEQ         = 7
	exprParserNEQ        = 8
	exprParserLT         = 9
	exprParserLEQ        = 10
	exprParserGT         = 11
	exprParserGEQ        = 12
	exprParserASSIGN     = 13
	exprParserPLUS       = 14
	exprParserMINUS      = 15
	exprParserDIV        = 16
	exprParserMUL        = 17
	exprParserCONCAT     = 18
	exprParserROUNDLEFT  = 19
	exprParserROUNDRIGHT = 20
	exprParserCOMMA      = 21
	exprParserBOOL       = 22
	exprParserINT        = 23
	exprParserDEC        = 24
	exprParserSTR        = 25
	exprParserID         = 26
	exprParserWS         = 27
)

// exprParser rules.
const (
	exprParserRULE_exp = 0
	exprParserRULE_id  = 1
	exprParserRULE_val = 2
)

// IExpContext is an interface to support dynamic dispatch.
type IExpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// IsExpContext differentiates from other interfaces.
	IsExpContext()
}

type ExpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyExpContext() *ExpContext {
	var p = new(ExpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = exprParserRULE_exp
	return p
}

func (*ExpContext) IsExpContext() {}

func NewExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpContext {
	var p = new(ExpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = exprParserRULE_exp

	return p
}

func (s *ExpContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpContext) GetOp() antlr.Token { return s.op }

func (s *ExpContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExpContext) ROUNDLEFT() antlr.TerminalNode {
	return s.GetToken(exprParserROUNDLEFT, 0)
}

func (s *ExpContext) AllExp() []IExpContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpContext)(nil)).Elem())
	var tst = make([]IExpContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpContext)
		}
	}

	return tst
}

func (s *ExpContext) Exp(i int) IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *ExpContext) ROUNDRIGHT() antlr.TerminalNode {
	return s.GetToken(exprParserROUNDRIGHT, 0)
}

func (s *ExpContext) Val() IValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValContext)
}

func (s *ExpContext) NOT() antlr.TerminalNode {
	return s.GetToken(exprParserNOT, 0)
}

func (s *ExpContext) CONCAT() antlr.TerminalNode {
	return s.GetToken(exprParserCONCAT, 0)
}

func (s *ExpContext) COMMA() antlr.TerminalNode {
	return s.GetToken(exprParserCOMMA, 0)
}

func (s *ExpContext) Id() IIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *ExpContext) THIS() antlr.TerminalNode {
	return s.GetToken(exprParserTHIS, 0)
}

func (s *ExpContext) EXT() antlr.TerminalNode {
	return s.GetToken(exprParserEXT, 0)
}

func (s *ExpContext) MUL() antlr.TerminalNode {
	return s.GetToken(exprParserMUL, 0)
}

func (s *ExpContext) DIV() antlr.TerminalNode {
	return s.GetToken(exprParserDIV, 0)
}

func (s *ExpContext) PLUS() antlr.TerminalNode {
	return s.GetToken(exprParserPLUS, 0)
}

func (s *ExpContext) MINUS() antlr.TerminalNode {
	return s.GetToken(exprParserMINUS, 0)
}

func (s *ExpContext) LT() antlr.TerminalNode {
	return s.GetToken(exprParserLT, 0)
}

func (s *ExpContext) LEQ() antlr.TerminalNode {
	return s.GetToken(exprParserLEQ, 0)
}

func (s *ExpContext) GT() antlr.TerminalNode {
	return s.GetToken(exprParserGT, 0)
}

func (s *ExpContext) GEQ() antlr.TerminalNode {
	return s.GetToken(exprParserGEQ, 0)
}

func (s *ExpContext) EQ() antlr.TerminalNode {
	return s.GetToken(exprParserEQ, 0)
}

func (s *ExpContext) NEQ() antlr.TerminalNode {
	return s.GetToken(exprParserNEQ, 0)
}

func (s *ExpContext) AND() antlr.TerminalNode {
	return s.GetToken(exprParserAND, 0)
}

func (s *ExpContext) OR() antlr.TerminalNode {
	return s.GetToken(exprParserOR, 0)
}

func (s *ExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case exprVisitor:
		return t.VisitExp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *exprParser) Exp() (localctx IExpContext) {
	return p.exp(0)
}

func (p *exprParser) exp(_p int) (localctx IExpContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 0
	p.EnterRecursionRule(localctx, 0, exprParserRULE_exp, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(27)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case exprParserROUNDLEFT:
		{
			p.SetState(7)
			p.Match(exprParserROUNDLEFT)
		}
		{
			p.SetState(8)
			p.exp(0)
		}
		{
			p.SetState(9)
			p.Match(exprParserROUNDRIGHT)
		}

	case exprParserUNDEF, exprParserBOOL, exprParserINT, exprParserDEC, exprParserSTR:
		{
			p.SetState(11)
			p.Val()
		}

	case exprParserNOT:
		{
			p.SetState(12)
			p.Match(exprParserNOT)
		}
		{
			p.SetState(13)
			p.exp(9)
		}

	case exprParserCONCAT:
		{
			p.SetState(14)
			p.Match(exprParserCONCAT)
		}
		{
			p.SetState(15)
			p.Match(exprParserROUNDLEFT)
		}
		{
			p.SetState(16)
			p.exp(0)
		}
		{
			p.SetState(17)
			p.Match(exprParserCOMMA)
		}
		{
			p.SetState(18)
			p.exp(0)
		}
		{
			p.SetState(19)
			p.Match(exprParserROUNDRIGHT)
		}

	case exprParserTHIS, exprParserID:
		p.SetState(22)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == exprParserTHIS {
			{
				p.SetState(21)
				p.Match(exprParserTHIS)
			}

		}
		{
			p.SetState(24)
			p.Id()
		}

	case exprParserEXT:
		{
			p.SetState(25)
			p.Match(exprParserEXT)
		}
		{
			p.SetState(26)
			p.Id()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(44)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, exprParserRULE_exp)
				p.SetState(29)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(30)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == exprParserDIV || _la == exprParserMUL) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(31)
					p.exp(9)
				}

			case 2:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, exprParserRULE_exp)
				p.SetState(32)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(33)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == exprParserPLUS || _la == exprParserMINUS) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(34)
					p.exp(8)
				}

			case 3:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, exprParserRULE_exp)
				p.SetState(35)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(36)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<exprParserLT)|(1<<exprParserLEQ)|(1<<exprParserGT)|(1<<exprParserGEQ))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(37)
					p.exp(7)
				}

			case 4:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, exprParserRULE_exp)
				p.SetState(38)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(39)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == exprParserEQ || _la == exprParserNEQ) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(40)
					p.exp(6)
				}

			case 5:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, exprParserRULE_exp)
				p.SetState(41)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(42)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == exprParserAND || _la == exprParserOR) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(43)
					p.exp(5)
				}

			}

		}
		p.SetState(48)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
	}

	return localctx
}

// IIdContext is an interface to support dynamic dispatch.
type IIdContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdContext differentiates from other interfaces.
	IsIdContext()
}

type IdContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdContext() *IdContext {
	var p = new(IdContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = exprParserRULE_id
	return p
}

func (*IdContext) IsIdContext() {}

func NewIdContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdContext {
	var p = new(IdContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = exprParserRULE_id

	return p
}

func (s *IdContext) GetParser() antlr.Parser { return s.parser }

func (s *IdContext) ID() antlr.TerminalNode {
	return s.GetToken(exprParserID, 0)
}

func (s *IdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case exprVisitor:
		return t.VisitId(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *exprParser) Id() (localctx IIdContext) {
	localctx = NewIdContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, exprParserRULE_id)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(49)
		p.Match(exprParserID)
	}

	return localctx
}

// IValContext is an interface to support dynamic dispatch.
type IValContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValContext differentiates from other interfaces.
	IsValContext()
}

type ValContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValContext() *ValContext {
	var p = new(ValContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = exprParserRULE_val
	return p
}

func (*ValContext) IsValContext() {}

func NewValContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValContext {
	var p = new(ValContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = exprParserRULE_val

	return p
}

func (s *ValContext) GetParser() antlr.Parser { return s.parser }

func (s *ValContext) UNDEF() antlr.TerminalNode {
	return s.GetToken(exprParserUNDEF, 0)
}

func (s *ValContext) BOOL() antlr.TerminalNode {
	return s.GetToken(exprParserBOOL, 0)
}

func (s *ValContext) INT() antlr.TerminalNode {
	return s.GetToken(exprParserINT, 0)
}

func (s *ValContext) DEC() antlr.TerminalNode {
	return s.GetToken(exprParserDEC, 0)
}

func (s *ValContext) STR() antlr.TerminalNode {
	return s.GetToken(exprParserSTR, 0)
}

func (s *ValContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case exprVisitor:
		return t.VisitVal(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *exprParser) Val() (localctx IValContext) {
	localctx = NewValContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, exprParserRULE_val)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(51)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<exprParserUNDEF)|(1<<exprParserBOOL)|(1<<exprParserINT)|(1<<exprParserDEC)|(1<<exprParserSTR))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

func (p *exprParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 0:
		var t *ExpContext = nil
		if localctx != nil {
			t = localctx.(*ExpContext)
		}
		return p.Exp_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *exprParser) Exp_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
