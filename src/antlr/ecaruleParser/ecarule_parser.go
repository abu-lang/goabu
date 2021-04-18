// Code generated from /home/michelep/Gitlab/steel-lang/src/antlr/ecarule.g4 by ANTLR 4.9.1. DO NOT EDIT.

package ecaruleParser // ecarule
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 38, 122,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 3, 2, 3, 2, 3, 2, 3, 2, 5, 2, 23, 10, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 5, 2, 29, 10, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 7, 3, 38, 10, 3, 12, 3, 14, 3, 41, 11, 3, 3, 4, 3, 4, 5, 4, 45, 10,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 7, 5, 56, 10,
	5, 12, 5, 14, 5, 59, 11, 5, 3, 6, 5, 6, 62, 10, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 73, 10, 6, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 5, 7, 91, 10, 7, 3, 7, 3, 7, 3, 7, 5, 7, 96, 10, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	7, 7, 113, 10, 7, 12, 7, 14, 7, 116, 11, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3,
	9, 2, 3, 12, 10, 2, 4, 6, 8, 10, 12, 14, 16, 2, 9, 3, 2, 7, 8, 3, 2, 28,
	29, 3, 2, 26, 27, 3, 2, 21, 24, 3, 2, 19, 20, 3, 2, 16, 17, 4, 2, 15, 15,
	34, 37, 2, 131, 2, 18, 3, 2, 2, 2, 4, 32, 3, 2, 2, 2, 6, 42, 3, 2, 2, 2,
	8, 50, 3, 2, 2, 2, 10, 72, 3, 2, 2, 2, 12, 95, 3, 2, 2, 2, 14, 117, 3,
	2, 2, 2, 16, 119, 3, 2, 2, 2, 18, 19, 7, 10, 2, 2, 19, 22, 5, 14, 8, 2,
	20, 21, 7, 4, 2, 2, 21, 23, 5, 14, 8, 2, 22, 20, 3, 2, 2, 2, 22, 23, 3,
	2, 2, 2, 23, 24, 3, 2, 2, 2, 24, 25, 7, 5, 2, 2, 25, 28, 5, 4, 3, 2, 26,
	27, 7, 3, 2, 2, 27, 29, 5, 8, 5, 2, 28, 26, 3, 2, 2, 2, 28, 29, 3, 2, 2,
	2, 29, 30, 3, 2, 2, 2, 30, 31, 5, 6, 4, 2, 31, 3, 3, 2, 2, 2, 32, 33, 5,
	14, 8, 2, 33, 39, 7, 11, 2, 2, 34, 35, 5, 14, 8, 2, 35, 36, 7, 11, 2, 2,
	36, 38, 3, 2, 2, 2, 37, 34, 3, 2, 2, 2, 38, 41, 3, 2, 2, 2, 39, 37, 3,
	2, 2, 2, 39, 40, 3, 2, 2, 2, 40, 5, 3, 2, 2, 2, 41, 39, 3, 2, 2, 2, 42,
	44, 7, 9, 2, 2, 43, 45, 9, 2, 2, 2, 44, 43, 3, 2, 2, 2, 44, 45, 3, 2, 2,
	2, 45, 46, 3, 2, 2, 2, 46, 47, 5, 12, 7, 2, 47, 48, 7, 6, 2, 2, 48, 49,
	5, 8, 5, 2, 49, 7, 3, 2, 2, 2, 50, 51, 5, 10, 6, 2, 51, 57, 7, 11, 2, 2,
	52, 53, 5, 10, 6, 2, 53, 54, 7, 11, 2, 2, 54, 56, 3, 2, 2, 2, 55, 52, 3,
	2, 2, 2, 56, 59, 3, 2, 2, 2, 57, 55, 3, 2, 2, 2, 57, 58, 3, 2, 2, 2, 58,
	9, 3, 2, 2, 2, 59, 57, 3, 2, 2, 2, 60, 62, 7, 13, 2, 2, 61, 60, 3, 2, 2,
	2, 61, 62, 3, 2, 2, 2, 62, 63, 3, 2, 2, 2, 63, 64, 5, 14, 8, 2, 64, 65,
	7, 25, 2, 2, 65, 66, 5, 12, 7, 2, 66, 73, 3, 2, 2, 2, 67, 68, 7, 14, 2,
	2, 68, 69, 5, 14, 8, 2, 69, 70, 7, 25, 2, 2, 70, 71, 5, 12, 7, 2, 71, 73,
	3, 2, 2, 2, 72, 61, 3, 2, 2, 2, 72, 67, 3, 2, 2, 2, 73, 11, 3, 2, 2, 2,
	74, 75, 8, 7, 1, 2, 75, 76, 7, 31, 2, 2, 76, 77, 5, 12, 7, 2, 77, 78, 7,
	32, 2, 2, 78, 96, 3, 2, 2, 2, 79, 96, 5, 16, 9, 2, 80, 81, 7, 18, 2, 2,
	81, 96, 5, 12, 7, 11, 82, 83, 7, 30, 2, 2, 83, 84, 7, 31, 2, 2, 84, 85,
	5, 12, 7, 2, 85, 86, 7, 33, 2, 2, 86, 87, 5, 12, 7, 2, 87, 88, 7, 32, 2,
	2, 88, 96, 3, 2, 2, 2, 89, 91, 7, 13, 2, 2, 90, 89, 3, 2, 2, 2, 90, 91,
	3, 2, 2, 2, 91, 92, 3, 2, 2, 2, 92, 96, 5, 14, 8, 2, 93, 94, 7, 14, 2,
	2, 94, 96, 5, 14, 8, 2, 95, 74, 3, 2, 2, 2, 95, 79, 3, 2, 2, 2, 95, 80,
	3, 2, 2, 2, 95, 82, 3, 2, 2, 2, 95, 90, 3, 2, 2, 2, 95, 93, 3, 2, 2, 2,
	96, 114, 3, 2, 2, 2, 97, 98, 12, 10, 2, 2, 98, 99, 9, 3, 2, 2, 99, 113,
	5, 12, 7, 11, 100, 101, 12, 9, 2, 2, 101, 102, 9, 4, 2, 2, 102, 113, 5,
	12, 7, 10, 103, 104, 12, 8, 2, 2, 104, 105, 9, 5, 2, 2, 105, 113, 5, 12,
	7, 9, 106, 107, 12, 7, 2, 2, 107, 108, 9, 6, 2, 2, 108, 113, 5, 12, 7,
	8, 109, 110, 12, 6, 2, 2, 110, 111, 9, 7, 2, 2, 111, 113, 5, 12, 7, 7,
	112, 97, 3, 2, 2, 2, 112, 100, 3, 2, 2, 2, 112, 103, 3, 2, 2, 2, 112, 106,
	3, 2, 2, 2, 112, 109, 3, 2, 2, 2, 113, 116, 3, 2, 2, 2, 114, 112, 3, 2,
	2, 2, 114, 115, 3, 2, 2, 2, 115, 13, 3, 2, 2, 2, 116, 114, 3, 2, 2, 2,
	117, 118, 7, 38, 2, 2, 118, 15, 3, 2, 2, 2, 119, 120, 9, 8, 2, 2, 120,
	17, 3, 2, 2, 2, 13, 22, 28, 39, 44, 57, 61, 72, 90, 95, 112, 114,
}
var literalNames = []string{
	"", "'default'", "'in'", "'on'", "'do'", "'all'", "'some'", "'for'", "'rule'",
	"';'", "", "'this.'", "'ext.'", "'null'", "'and'", "'or'", "'not'", "'=='",
	"'=/='", "'<'", "'<='", "'>'", "'>='", "'='", "'+'", "'-'", "'/'", "'*'",
	"'concat'", "'('", "')'", "','",
}
var symbolicNames = []string{
	"", "DEFAULT", "IN", "ON", "DO", "ALL", "SOME", "FOR", "RULE", "SEMICOLON",
	"WS", "THIS", "EXT", "UNDEF", "AND", "OR", "NOT", "EQ", "NEQ", "LT", "LEQ",
	"GT", "GEQ", "ASSIGN", "PLUS", "MINUS", "DIV", "MUL", "CONCAT", "ROUNDLEFT",
	"ROUNDRIGHT", "COMMA", "BOOL", "INT", "DEC", "STR", "ID",
}

var ruleNames = []string{
	"prule", "evt", "task", "actslist", "act", "exp", "id", "val",
}

type ecaruleParser struct {
	*antlr.BaseParser
}

// NewecaruleParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *ecaruleParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewecaruleParser(input antlr.TokenStream) *ecaruleParser {
	this := new(ecaruleParser)
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
	this.GrammarFileName = "ecarule.g4"

	return this
}

// ecaruleParser tokens.
const (
	ecaruleParserEOF        = antlr.TokenEOF
	ecaruleParserDEFAULT    = 1
	ecaruleParserIN         = 2
	ecaruleParserON         = 3
	ecaruleParserDO         = 4
	ecaruleParserALL        = 5
	ecaruleParserSOME       = 6
	ecaruleParserFOR        = 7
	ecaruleParserRULE       = 8
	ecaruleParserSEMICOLON  = 9
	ecaruleParserWS         = 10
	ecaruleParserTHIS       = 11
	ecaruleParserEXT        = 12
	ecaruleParserUNDEF      = 13
	ecaruleParserAND        = 14
	ecaruleParserOR         = 15
	ecaruleParserNOT        = 16
	ecaruleParserEQ         = 17
	ecaruleParserNEQ        = 18
	ecaruleParserLT         = 19
	ecaruleParserLEQ        = 20
	ecaruleParserGT         = 21
	ecaruleParserGEQ        = 22
	ecaruleParserASSIGN     = 23
	ecaruleParserPLUS       = 24
	ecaruleParserMINUS      = 25
	ecaruleParserDIV        = 26
	ecaruleParserMUL        = 27
	ecaruleParserCONCAT     = 28
	ecaruleParserROUNDLEFT  = 29
	ecaruleParserROUNDRIGHT = 30
	ecaruleParserCOMMA      = 31
	ecaruleParserBOOL       = 32
	ecaruleParserINT        = 33
	ecaruleParserDEC        = 34
	ecaruleParserSTR        = 35
	ecaruleParserID         = 36
)

// ecaruleParser rules.
const (
	ecaruleParserRULE_prule    = 0
	ecaruleParserRULE_evt      = 1
	ecaruleParserRULE_task     = 2
	ecaruleParserRULE_actslist = 3
	ecaruleParserRULE_act      = 4
	ecaruleParserRULE_exp      = 5
	ecaruleParserRULE_id       = 6
	ecaruleParserRULE_val      = 7
)

// IPruleContext is an interface to support dynamic dispatch.
type IPruleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPruleContext differentiates from other interfaces.
	IsPruleContext()
}

type PruleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPruleContext() *PruleContext {
	var p = new(PruleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ecaruleParserRULE_prule
	return p
}

func (*PruleContext) IsPruleContext() {}

func NewPruleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PruleContext {
	var p = new(PruleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ecaruleParserRULE_prule

	return p
}

func (s *PruleContext) GetParser() antlr.Parser { return s.parser }

func (s *PruleContext) RULE() antlr.TerminalNode {
	return s.GetToken(ecaruleParserRULE, 0)
}

func (s *PruleContext) AllId() []IIdContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIdContext)(nil)).Elem())
	var tst = make([]IIdContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIdContext)
		}
	}

	return tst
}

func (s *PruleContext) Id(i int) IIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *PruleContext) ON() antlr.TerminalNode {
	return s.GetToken(ecaruleParserON, 0)
}

func (s *PruleContext) Evt() IEvtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEvtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEvtContext)
}

func (s *PruleContext) Task() ITaskContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITaskContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITaskContext)
}

func (s *PruleContext) IN() antlr.TerminalNode {
	return s.GetToken(ecaruleParserIN, 0)
}

func (s *PruleContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(ecaruleParserDEFAULT, 0)
}

func (s *PruleContext) Actslist() IActslistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IActslistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IActslistContext)
}

func (s *PruleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PruleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PruleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ecaruleVisitor:
		return t.VisitPrule(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ecaruleParser) Prule() (localctx IPruleContext) {
	localctx = NewPruleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ecaruleParserRULE_prule)
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
		p.SetState(16)
		p.Match(ecaruleParserRULE)
	}
	{
		p.SetState(17)
		p.Id()
	}
	p.SetState(20)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ecaruleParserIN {
		{
			p.SetState(18)
			p.Match(ecaruleParserIN)
		}
		{
			p.SetState(19)
			p.Id()
		}

	}
	{
		p.SetState(22)
		p.Match(ecaruleParserON)
	}
	{
		p.SetState(23)
		p.Evt()
	}
	p.SetState(26)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ecaruleParserDEFAULT {
		{
			p.SetState(24)
			p.Match(ecaruleParserDEFAULT)
		}
		{
			p.SetState(25)
			p.Actslist()
		}

	}
	{
		p.SetState(28)
		p.Task()
	}

	return localctx
}

// IEvtContext is an interface to support dynamic dispatch.
type IEvtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEvtContext differentiates from other interfaces.
	IsEvtContext()
}

type EvtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEvtContext() *EvtContext {
	var p = new(EvtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ecaruleParserRULE_evt
	return p
}

func (*EvtContext) IsEvtContext() {}

func NewEvtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EvtContext {
	var p = new(EvtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ecaruleParserRULE_evt

	return p
}

func (s *EvtContext) GetParser() antlr.Parser { return s.parser }

func (s *EvtContext) AllId() []IIdContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIdContext)(nil)).Elem())
	var tst = make([]IIdContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIdContext)
		}
	}

	return tst
}

func (s *EvtContext) Id(i int) IIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *EvtContext) AllSEMICOLON() []antlr.TerminalNode {
	return s.GetTokens(ecaruleParserSEMICOLON)
}

func (s *EvtContext) SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(ecaruleParserSEMICOLON, i)
}

func (s *EvtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EvtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EvtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ecaruleVisitor:
		return t.VisitEvt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ecaruleParser) Evt() (localctx IEvtContext) {
	localctx = NewEvtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ecaruleParserRULE_evt)
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
		p.SetState(30)
		p.Id()
	}
	{
		p.SetState(31)
		p.Match(ecaruleParserSEMICOLON)
	}
	p.SetState(37)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ecaruleParserID {
		{
			p.SetState(32)
			p.Id()
		}
		{
			p.SetState(33)
			p.Match(ecaruleParserSEMICOLON)
		}

		p.SetState(39)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITaskContext is an interface to support dynamic dispatch.
type ITaskContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTaskContext differentiates from other interfaces.
	IsTaskContext()
}

type TaskContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTaskContext() *TaskContext {
	var p = new(TaskContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ecaruleParserRULE_task
	return p
}

func (*TaskContext) IsTaskContext() {}

func NewTaskContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TaskContext {
	var p = new(TaskContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ecaruleParserRULE_task

	return p
}

func (s *TaskContext) GetParser() antlr.Parser { return s.parser }

func (s *TaskContext) FOR() antlr.TerminalNode {
	return s.GetToken(ecaruleParserFOR, 0)
}

func (s *TaskContext) Exp() IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *TaskContext) DO() antlr.TerminalNode {
	return s.GetToken(ecaruleParserDO, 0)
}

func (s *TaskContext) Actslist() IActslistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IActslistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IActslistContext)
}

func (s *TaskContext) SOME() antlr.TerminalNode {
	return s.GetToken(ecaruleParserSOME, 0)
}

func (s *TaskContext) ALL() antlr.TerminalNode {
	return s.GetToken(ecaruleParserALL, 0)
}

func (s *TaskContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TaskContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TaskContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ecaruleVisitor:
		return t.VisitTask(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ecaruleParser) Task() (localctx ITaskContext) {
	localctx = NewTaskContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ecaruleParserRULE_task)
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
		p.SetState(40)
		p.Match(ecaruleParserFOR)
	}
	p.SetState(42)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ecaruleParserALL || _la == ecaruleParserSOME {
		{
			p.SetState(41)
			_la = p.GetTokenStream().LA(1)

			if !(_la == ecaruleParserALL || _la == ecaruleParserSOME) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	{
		p.SetState(44)
		p.exp(0)
	}
	{
		p.SetState(45)
		p.Match(ecaruleParserDO)
	}
	{
		p.SetState(46)
		p.Actslist()
	}

	return localctx
}

// IActslistContext is an interface to support dynamic dispatch.
type IActslistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsActslistContext differentiates from other interfaces.
	IsActslistContext()
}

type ActslistContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyActslistContext() *ActslistContext {
	var p = new(ActslistContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ecaruleParserRULE_actslist
	return p
}

func (*ActslistContext) IsActslistContext() {}

func NewActslistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ActslistContext {
	var p = new(ActslistContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ecaruleParserRULE_actslist

	return p
}

func (s *ActslistContext) GetParser() antlr.Parser { return s.parser }

func (s *ActslistContext) AllAct() []IActContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IActContext)(nil)).Elem())
	var tst = make([]IActContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IActContext)
		}
	}

	return tst
}

func (s *ActslistContext) Act(i int) IActContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IActContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IActContext)
}

func (s *ActslistContext) AllSEMICOLON() []antlr.TerminalNode {
	return s.GetTokens(ecaruleParserSEMICOLON)
}

func (s *ActslistContext) SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(ecaruleParserSEMICOLON, i)
}

func (s *ActslistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ActslistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ActslistContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ecaruleVisitor:
		return t.VisitActslist(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ecaruleParser) Actslist() (localctx IActslistContext) {
	localctx = NewActslistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ecaruleParserRULE_actslist)
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
		p.SetState(48)
		p.Act()
	}
	{
		p.SetState(49)
		p.Match(ecaruleParserSEMICOLON)
	}
	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-11)&-(0x1f+1)) == 0 && ((1<<uint((_la-11)))&((1<<(ecaruleParserTHIS-11))|(1<<(ecaruleParserEXT-11))|(1<<(ecaruleParserID-11)))) != 0 {
		{
			p.SetState(50)
			p.Act()
		}
		{
			p.SetState(51)
			p.Match(ecaruleParserSEMICOLON)
		}

		p.SetState(57)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IActContext is an interface to support dynamic dispatch.
type IActContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsActContext differentiates from other interfaces.
	IsActContext()
}

type ActContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyActContext() *ActContext {
	var p = new(ActContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ecaruleParserRULE_act
	return p
}

func (*ActContext) IsActContext() {}

func NewActContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ActContext {
	var p = new(ActContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ecaruleParserRULE_act

	return p
}

func (s *ActContext) GetParser() antlr.Parser { return s.parser }

func (s *ActContext) Id() IIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *ActContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(ecaruleParserASSIGN, 0)
}

func (s *ActContext) Exp() IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *ActContext) THIS() antlr.TerminalNode {
	return s.GetToken(ecaruleParserTHIS, 0)
}

func (s *ActContext) EXT() antlr.TerminalNode {
	return s.GetToken(ecaruleParserEXT, 0)
}

func (s *ActContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ActContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ActContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ecaruleVisitor:
		return t.VisitAct(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ecaruleParser) Act() (localctx IActContext) {
	localctx = NewActContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ecaruleParserRULE_act)
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

	p.SetState(70)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ecaruleParserTHIS, ecaruleParserID:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(59)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ecaruleParserTHIS {
			{
				p.SetState(58)
				p.Match(ecaruleParserTHIS)
			}

		}
		{
			p.SetState(61)
			p.Id()
		}
		{
			p.SetState(62)
			p.Match(ecaruleParserASSIGN)
		}
		{
			p.SetState(63)
			p.exp(0)
		}

	case ecaruleParserEXT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(65)
			p.Match(ecaruleParserEXT)
		}
		{
			p.SetState(66)
			p.Id()
		}
		{
			p.SetState(67)
			p.Match(ecaruleParserASSIGN)
		}
		{
			p.SetState(68)
			p.exp(0)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

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
	p.RuleIndex = ecaruleParserRULE_exp
	return p
}

func (*ExpContext) IsExpContext() {}

func NewExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpContext {
	var p = new(ExpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ecaruleParserRULE_exp

	return p
}

func (s *ExpContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpContext) GetOp() antlr.Token { return s.op }

func (s *ExpContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExpContext) ROUNDLEFT() antlr.TerminalNode {
	return s.GetToken(ecaruleParserROUNDLEFT, 0)
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
	return s.GetToken(ecaruleParserROUNDRIGHT, 0)
}

func (s *ExpContext) Val() IValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValContext)
}

func (s *ExpContext) NOT() antlr.TerminalNode {
	return s.GetToken(ecaruleParserNOT, 0)
}

func (s *ExpContext) CONCAT() antlr.TerminalNode {
	return s.GetToken(ecaruleParserCONCAT, 0)
}

func (s *ExpContext) COMMA() antlr.TerminalNode {
	return s.GetToken(ecaruleParserCOMMA, 0)
}

func (s *ExpContext) Id() IIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *ExpContext) THIS() antlr.TerminalNode {
	return s.GetToken(ecaruleParserTHIS, 0)
}

func (s *ExpContext) EXT() antlr.TerminalNode {
	return s.GetToken(ecaruleParserEXT, 0)
}

func (s *ExpContext) MUL() antlr.TerminalNode {
	return s.GetToken(ecaruleParserMUL, 0)
}

func (s *ExpContext) DIV() antlr.TerminalNode {
	return s.GetToken(ecaruleParserDIV, 0)
}

func (s *ExpContext) PLUS() antlr.TerminalNode {
	return s.GetToken(ecaruleParserPLUS, 0)
}

func (s *ExpContext) MINUS() antlr.TerminalNode {
	return s.GetToken(ecaruleParserMINUS, 0)
}

func (s *ExpContext) LT() antlr.TerminalNode {
	return s.GetToken(ecaruleParserLT, 0)
}

func (s *ExpContext) LEQ() antlr.TerminalNode {
	return s.GetToken(ecaruleParserLEQ, 0)
}

func (s *ExpContext) GT() antlr.TerminalNode {
	return s.GetToken(ecaruleParserGT, 0)
}

func (s *ExpContext) GEQ() antlr.TerminalNode {
	return s.GetToken(ecaruleParserGEQ, 0)
}

func (s *ExpContext) EQ() antlr.TerminalNode {
	return s.GetToken(ecaruleParserEQ, 0)
}

func (s *ExpContext) NEQ() antlr.TerminalNode {
	return s.GetToken(ecaruleParserNEQ, 0)
}

func (s *ExpContext) AND() antlr.TerminalNode {
	return s.GetToken(ecaruleParserAND, 0)
}

func (s *ExpContext) OR() antlr.TerminalNode {
	return s.GetToken(ecaruleParserOR, 0)
}

func (s *ExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ecaruleVisitor:
		return t.VisitExp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ecaruleParser) Exp() (localctx IExpContext) {
	return p.exp(0)
}

func (p *ecaruleParser) exp(_p int) (localctx IExpContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 10
	p.EnterRecursionRule(localctx, 10, ecaruleParserRULE_exp, _p)
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
	p.SetState(93)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ecaruleParserROUNDLEFT:
		{
			p.SetState(73)
			p.Match(ecaruleParserROUNDLEFT)
		}
		{
			p.SetState(74)
			p.exp(0)
		}
		{
			p.SetState(75)
			p.Match(ecaruleParserROUNDRIGHT)
		}

	case ecaruleParserUNDEF, ecaruleParserBOOL, ecaruleParserINT, ecaruleParserDEC, ecaruleParserSTR:
		{
			p.SetState(77)
			p.Val()
		}

	case ecaruleParserNOT:
		{
			p.SetState(78)
			p.Match(ecaruleParserNOT)
		}
		{
			p.SetState(79)
			p.exp(9)
		}

	case ecaruleParserCONCAT:
		{
			p.SetState(80)
			p.Match(ecaruleParserCONCAT)
		}
		{
			p.SetState(81)
			p.Match(ecaruleParserROUNDLEFT)
		}
		{
			p.SetState(82)
			p.exp(0)
		}
		{
			p.SetState(83)
			p.Match(ecaruleParserCOMMA)
		}
		{
			p.SetState(84)
			p.exp(0)
		}
		{
			p.SetState(85)
			p.Match(ecaruleParserROUNDRIGHT)
		}

	case ecaruleParserTHIS, ecaruleParserID:
		p.SetState(88)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ecaruleParserTHIS {
			{
				p.SetState(87)
				p.Match(ecaruleParserTHIS)
			}

		}
		{
			p.SetState(90)
			p.Id()
		}

	case ecaruleParserEXT:
		{
			p.SetState(91)
			p.Match(ecaruleParserEXT)
		}
		{
			p.SetState(92)
			p.Id()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(112)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(110)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ecaruleParserRULE_exp)
				p.SetState(95)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(96)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == ecaruleParserDIV || _la == ecaruleParserMUL) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(97)
					p.exp(9)
				}

			case 2:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ecaruleParserRULE_exp)
				p.SetState(98)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(99)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == ecaruleParserPLUS || _la == ecaruleParserMINUS) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(100)
					p.exp(8)
				}

			case 3:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ecaruleParserRULE_exp)
				p.SetState(101)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(102)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ecaruleParserLT)|(1<<ecaruleParserLEQ)|(1<<ecaruleParserGT)|(1<<ecaruleParserGEQ))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(103)
					p.exp(7)
				}

			case 4:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ecaruleParserRULE_exp)
				p.SetState(104)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(105)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == ecaruleParserEQ || _la == ecaruleParserNEQ) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(106)
					p.exp(6)
				}

			case 5:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ecaruleParserRULE_exp)
				p.SetState(107)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(108)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == ecaruleParserAND || _la == ecaruleParserOR) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(109)
					p.exp(5)
				}

			}

		}
		p.SetState(114)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())
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
	p.RuleIndex = ecaruleParserRULE_id
	return p
}

func (*IdContext) IsIdContext() {}

func NewIdContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdContext {
	var p = new(IdContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ecaruleParserRULE_id

	return p
}

func (s *IdContext) GetParser() antlr.Parser { return s.parser }

func (s *IdContext) ID() antlr.TerminalNode {
	return s.GetToken(ecaruleParserID, 0)
}

func (s *IdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ecaruleVisitor:
		return t.VisitId(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ecaruleParser) Id() (localctx IIdContext) {
	localctx = NewIdContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ecaruleParserRULE_id)

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
		p.SetState(115)
		p.Match(ecaruleParserID)
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
	p.RuleIndex = ecaruleParserRULE_val
	return p
}

func (*ValContext) IsValContext() {}

func NewValContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValContext {
	var p = new(ValContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ecaruleParserRULE_val

	return p
}

func (s *ValContext) GetParser() antlr.Parser { return s.parser }

func (s *ValContext) UNDEF() antlr.TerminalNode {
	return s.GetToken(ecaruleParserUNDEF, 0)
}

func (s *ValContext) BOOL() antlr.TerminalNode {
	return s.GetToken(ecaruleParserBOOL, 0)
}

func (s *ValContext) INT() antlr.TerminalNode {
	return s.GetToken(ecaruleParserINT, 0)
}

func (s *ValContext) DEC() antlr.TerminalNode {
	return s.GetToken(ecaruleParserDEC, 0)
}

func (s *ValContext) STR() antlr.TerminalNode {
	return s.GetToken(ecaruleParserSTR, 0)
}

func (s *ValContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ecaruleVisitor:
		return t.VisitVal(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ecaruleParser) Val() (localctx IValContext) {
	localctx = NewValContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ecaruleParserRULE_val)
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
		p.SetState(117)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-13)&-(0x1f+1)) == 0 && ((1<<uint((_la-13)))&((1<<(ecaruleParserUNDEF-13))|(1<<(ecaruleParserBOOL-13))|(1<<(ecaruleParserINT-13))|(1<<(ecaruleParserDEC-13))|(1<<(ecaruleParserSTR-13)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

func (p *ecaruleParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 5:
		var t *ExpContext = nil
		if localctx != nil {
			t = localctx.(*ExpContext)
		}
		return p.Exp_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *ecaruleParser) Exp_Sempred(localctx antlr.RuleContext, predIndex int) bool {
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
