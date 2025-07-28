// Code generated from QueryLang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package antlr // QueryLang
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type QueryLangParser struct {
	*antlr.BaseParser
}

var QueryLangParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func querylangParserInit() {
	staticData := &QueryLangParserStaticData
	staticData.LiteralNames = []string{
		"", "'CREATE'", "'BENCHMARK'", "'INSERT_ALL'", "','", "'INSERT'", "'LOOKUP'",
		"'DELETE'", "'COMPARE'", "'('", "')'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "BENCHMARK_NAME", "BTREE_TYPE",
		"INT_LIT", "WS",
	}
	staticData.RuleNames = []string{
		"command", "createBenchmark", "insertAll", "insert", "lookup", "delete",
		"compareLookup", "dataList", "dataEntry", "keyList",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 14, 87, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 3, 0, 27, 8, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 5, 7, 68, 8,
		7, 10, 7, 12, 7, 71, 9, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1,
		9, 1, 9, 5, 9, 82, 8, 9, 10, 9, 12, 9, 85, 9, 9, 1, 9, 0, 0, 10, 0, 2,
		4, 6, 8, 10, 12, 14, 16, 18, 0, 0, 83, 0, 26, 1, 0, 0, 0, 2, 28, 1, 0,
		0, 0, 4, 32, 1, 0, 0, 0, 6, 37, 1, 0, 0, 0, 8, 44, 1, 0, 0, 0, 10, 51,
		1, 0, 0, 0, 12, 58, 1, 0, 0, 0, 14, 64, 1, 0, 0, 0, 16, 72, 1, 0, 0, 0,
		18, 78, 1, 0, 0, 0, 20, 27, 3, 2, 1, 0, 21, 27, 3, 4, 2, 0, 22, 27, 3,
		6, 3, 0, 23, 27, 3, 8, 4, 0, 24, 27, 3, 10, 5, 0, 25, 27, 3, 12, 6, 0,
		26, 20, 1, 0, 0, 0, 26, 21, 1, 0, 0, 0, 26, 22, 1, 0, 0, 0, 26, 23, 1,
		0, 0, 0, 26, 24, 1, 0, 0, 0, 26, 25, 1, 0, 0, 0, 27, 1, 1, 0, 0, 0, 28,
		29, 5, 1, 0, 0, 29, 30, 5, 2, 0, 0, 30, 31, 5, 11, 0, 0, 31, 3, 1, 0, 0,
		0, 32, 33, 5, 3, 0, 0, 33, 34, 5, 11, 0, 0, 34, 35, 5, 4, 0, 0, 35, 36,
		3, 14, 7, 0, 36, 5, 1, 0, 0, 0, 37, 38, 5, 5, 0, 0, 38, 39, 5, 11, 0, 0,
		39, 40, 5, 4, 0, 0, 40, 41, 5, 12, 0, 0, 41, 42, 5, 4, 0, 0, 42, 43, 3,
		14, 7, 0, 43, 7, 1, 0, 0, 0, 44, 45, 5, 6, 0, 0, 45, 46, 5, 11, 0, 0, 46,
		47, 5, 4, 0, 0, 47, 48, 5, 12, 0, 0, 48, 49, 5, 4, 0, 0, 49, 50, 3, 18,
		9, 0, 50, 9, 1, 0, 0, 0, 51, 52, 5, 7, 0, 0, 52, 53, 5, 11, 0, 0, 53, 54,
		5, 4, 0, 0, 54, 55, 5, 12, 0, 0, 55, 56, 5, 4, 0, 0, 56, 57, 3, 18, 9,
		0, 57, 11, 1, 0, 0, 0, 58, 59, 5, 8, 0, 0, 59, 60, 5, 6, 0, 0, 60, 61,
		5, 11, 0, 0, 61, 62, 5, 4, 0, 0, 62, 63, 3, 18, 9, 0, 63, 13, 1, 0, 0,
		0, 64, 69, 3, 16, 8, 0, 65, 66, 5, 4, 0, 0, 66, 68, 3, 16, 8, 0, 67, 65,
		1, 0, 0, 0, 68, 71, 1, 0, 0, 0, 69, 67, 1, 0, 0, 0, 69, 70, 1, 0, 0, 0,
		70, 15, 1, 0, 0, 0, 71, 69, 1, 0, 0, 0, 72, 73, 5, 9, 0, 0, 73, 74, 5,
		13, 0, 0, 74, 75, 5, 4, 0, 0, 75, 76, 5, 13, 0, 0, 76, 77, 5, 10, 0, 0,
		77, 17, 1, 0, 0, 0, 78, 83, 5, 13, 0, 0, 79, 80, 5, 4, 0, 0, 80, 82, 5,
		13, 0, 0, 81, 79, 1, 0, 0, 0, 82, 85, 1, 0, 0, 0, 83, 81, 1, 0, 0, 0, 83,
		84, 1, 0, 0, 0, 84, 19, 1, 0, 0, 0, 85, 83, 1, 0, 0, 0, 3, 26, 69, 83,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// QueryLangParserInit initializes any static state used to implement QueryLangParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewQueryLangParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func QueryLangParserInit() {
	staticData := &QueryLangParserStaticData
	staticData.once.Do(querylangParserInit)
}

// NewQueryLangParser produces a new parser instance for the optional input antlr.TokenStream.
func NewQueryLangParser(input antlr.TokenStream) *QueryLangParser {
	QueryLangParserInit()
	this := new(QueryLangParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &QueryLangParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "QueryLang.g4"

	return this
}

// QueryLangParser tokens.
const (
	QueryLangParserEOF            = antlr.TokenEOF
	QueryLangParserT__0           = 1
	QueryLangParserT__1           = 2
	QueryLangParserT__2           = 3
	QueryLangParserT__3           = 4
	QueryLangParserT__4           = 5
	QueryLangParserT__5           = 6
	QueryLangParserT__6           = 7
	QueryLangParserT__7           = 8
	QueryLangParserT__8           = 9
	QueryLangParserT__9           = 10
	QueryLangParserBENCHMARK_NAME = 11
	QueryLangParserBTREE_TYPE     = 12
	QueryLangParserINT_LIT        = 13
	QueryLangParserWS             = 14
)

// QueryLangParser rules.
const (
	QueryLangParserRULE_command         = 0
	QueryLangParserRULE_createBenchmark = 1
	QueryLangParserRULE_insertAll       = 2
	QueryLangParserRULE_insert          = 3
	QueryLangParserRULE_lookup          = 4
	QueryLangParserRULE_delete          = 5
	QueryLangParserRULE_compareLookup   = 6
	QueryLangParserRULE_dataList        = 7
	QueryLangParserRULE_dataEntry       = 8
	QueryLangParserRULE_keyList         = 9
)

// ICommandContext is an interface to support dynamic dispatch.
type ICommandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CreateBenchmark() ICreateBenchmarkContext
	InsertAll() IInsertAllContext
	Insert() IInsertContext
	Lookup() ILookupContext
	Delete_() IDeleteContext
	CompareLookup() ICompareLookupContext

	// IsCommandContext differentiates from other interfaces.
	IsCommandContext()
}

type CommandContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommandContext() *CommandContext {
	var p = new(CommandContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLangParserRULE_command
	return p
}

func InitEmptyCommandContext(p *CommandContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLangParserRULE_command
}

func (*CommandContext) IsCommandContext() {}

func NewCommandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommandContext {
	var p = new(CommandContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryLangParserRULE_command

	return p
}

func (s *CommandContext) GetParser() antlr.Parser { return s.parser }

func (s *CommandContext) CreateBenchmark() ICreateBenchmarkContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICreateBenchmarkContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICreateBenchmarkContext)
}

func (s *CommandContext) InsertAll() IInsertAllContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInsertAllContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInsertAllContext)
}

func (s *CommandContext) Insert() IInsertContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInsertContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInsertContext)
}

func (s *CommandContext) Lookup() ILookupContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILookupContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILookupContext)
}

func (s *CommandContext) Delete_() IDeleteContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeleteContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeleteContext)
}

func (s *CommandContext) CompareLookup() ICompareLookupContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICompareLookupContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICompareLookupContext)
}

func (s *CommandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommandContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryLangVisitor:
		return t.VisitCommand(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QueryLangParser) Command() (localctx ICommandContext) {
	localctx = NewCommandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, QueryLangParserRULE_command)
	p.SetState(26)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case QueryLangParserT__0:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(20)
			p.CreateBenchmark()
		}

	case QueryLangParserT__2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(21)
			p.InsertAll()
		}

	case QueryLangParserT__4:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(22)
			p.Insert()
		}

	case QueryLangParserT__5:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(23)
			p.Lookup()
		}

	case QueryLangParserT__6:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(24)
			p.Delete_()
		}

	case QueryLangParserT__7:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(25)
			p.CompareLookup()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICreateBenchmarkContext is an interface to support dynamic dispatch.
type ICreateBenchmarkContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BENCHMARK_NAME() antlr.TerminalNode

	// IsCreateBenchmarkContext differentiates from other interfaces.
	IsCreateBenchmarkContext()
}

type CreateBenchmarkContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCreateBenchmarkContext() *CreateBenchmarkContext {
	var p = new(CreateBenchmarkContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLangParserRULE_createBenchmark
	return p
}

func InitEmptyCreateBenchmarkContext(p *CreateBenchmarkContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLangParserRULE_createBenchmark
}

func (*CreateBenchmarkContext) IsCreateBenchmarkContext() {}

func NewCreateBenchmarkContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CreateBenchmarkContext {
	var p = new(CreateBenchmarkContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryLangParserRULE_createBenchmark

	return p
}

func (s *CreateBenchmarkContext) GetParser() antlr.Parser { return s.parser }

func (s *CreateBenchmarkContext) BENCHMARK_NAME() antlr.TerminalNode {
	return s.GetToken(QueryLangParserBENCHMARK_NAME, 0)
}

func (s *CreateBenchmarkContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CreateBenchmarkContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CreateBenchmarkContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryLangVisitor:
		return t.VisitCreateBenchmark(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QueryLangParser) CreateBenchmark() (localctx ICreateBenchmarkContext) {
	localctx = NewCreateBenchmarkContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, QueryLangParserRULE_createBenchmark)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(28)
		p.Match(QueryLangParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(29)
		p.Match(QueryLangParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(30)
		p.Match(QueryLangParserBENCHMARK_NAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInsertAllContext is an interface to support dynamic dispatch.
type IInsertAllContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BENCHMARK_NAME() antlr.TerminalNode
	DataList() IDataListContext

	// IsInsertAllContext differentiates from other interfaces.
	IsInsertAllContext()
}

type InsertAllContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInsertAllContext() *InsertAllContext {
	var p = new(InsertAllContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLangParserRULE_insertAll
	return p
}

func InitEmptyInsertAllContext(p *InsertAllContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLangParserRULE_insertAll
}

func (*InsertAllContext) IsInsertAllContext() {}

func NewInsertAllContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InsertAllContext {
	var p = new(InsertAllContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryLangParserRULE_insertAll

	return p
}

func (s *InsertAllContext) GetParser() antlr.Parser { return s.parser }

func (s *InsertAllContext) BENCHMARK_NAME() antlr.TerminalNode {
	return s.GetToken(QueryLangParserBENCHMARK_NAME, 0)
}

func (s *InsertAllContext) DataList() IDataListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDataListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDataListContext)
}

func (s *InsertAllContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InsertAllContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InsertAllContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryLangVisitor:
		return t.VisitInsertAll(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QueryLangParser) InsertAll() (localctx IInsertAllContext) {
	localctx = NewInsertAllContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, QueryLangParserRULE_insertAll)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(32)
		p.Match(QueryLangParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(33)
		p.Match(QueryLangParserBENCHMARK_NAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(34)
		p.Match(QueryLangParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(35)
		p.DataList()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInsertContext is an interface to support dynamic dispatch.
type IInsertContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BENCHMARK_NAME() antlr.TerminalNode
	BTREE_TYPE() antlr.TerminalNode
	DataList() IDataListContext

	// IsInsertContext differentiates from other interfaces.
	IsInsertContext()
}

type InsertContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInsertContext() *InsertContext {
	var p = new(InsertContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLangParserRULE_insert
	return p
}

func InitEmptyInsertContext(p *InsertContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLangParserRULE_insert
}

func (*InsertContext) IsInsertContext() {}

func NewInsertContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InsertContext {
	var p = new(InsertContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryLangParserRULE_insert

	return p
}

func (s *InsertContext) GetParser() antlr.Parser { return s.parser }

func (s *InsertContext) BENCHMARK_NAME() antlr.TerminalNode {
	return s.GetToken(QueryLangParserBENCHMARK_NAME, 0)
}

func (s *InsertContext) BTREE_TYPE() antlr.TerminalNode {
	return s.GetToken(QueryLangParserBTREE_TYPE, 0)
}

func (s *InsertContext) DataList() IDataListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDataListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDataListContext)
}

func (s *InsertContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InsertContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InsertContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryLangVisitor:
		return t.VisitInsert(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QueryLangParser) Insert() (localctx IInsertContext) {
	localctx = NewInsertContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, QueryLangParserRULE_insert)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(37)
		p.Match(QueryLangParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(38)
		p.Match(QueryLangParserBENCHMARK_NAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(39)
		p.Match(QueryLangParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(40)
		p.Match(QueryLangParserBTREE_TYPE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(41)
		p.Match(QueryLangParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(42)
		p.DataList()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILookupContext is an interface to support dynamic dispatch.
type ILookupContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BENCHMARK_NAME() antlr.TerminalNode
	BTREE_TYPE() antlr.TerminalNode
	KeyList() IKeyListContext

	// IsLookupContext differentiates from other interfaces.
	IsLookupContext()
}

type LookupContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLookupContext() *LookupContext {
	var p = new(LookupContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLangParserRULE_lookup
	return p
}

func InitEmptyLookupContext(p *LookupContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLangParserRULE_lookup
}

func (*LookupContext) IsLookupContext() {}

func NewLookupContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LookupContext {
	var p = new(LookupContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryLangParserRULE_lookup

	return p
}

func (s *LookupContext) GetParser() antlr.Parser { return s.parser }

func (s *LookupContext) BENCHMARK_NAME() antlr.TerminalNode {
	return s.GetToken(QueryLangParserBENCHMARK_NAME, 0)
}

func (s *LookupContext) BTREE_TYPE() antlr.TerminalNode {
	return s.GetToken(QueryLangParserBTREE_TYPE, 0)
}

func (s *LookupContext) KeyList() IKeyListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeyListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKeyListContext)
}

func (s *LookupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LookupContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LookupContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryLangVisitor:
		return t.VisitLookup(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QueryLangParser) Lookup() (localctx ILookupContext) {
	localctx = NewLookupContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, QueryLangParserRULE_lookup)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(44)
		p.Match(QueryLangParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(45)
		p.Match(QueryLangParserBENCHMARK_NAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(46)
		p.Match(QueryLangParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(47)
		p.Match(QueryLangParserBTREE_TYPE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(48)
		p.Match(QueryLangParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(49)
		p.KeyList()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeleteContext is an interface to support dynamic dispatch.
type IDeleteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BENCHMARK_NAME() antlr.TerminalNode
	BTREE_TYPE() antlr.TerminalNode
	KeyList() IKeyListContext

	// IsDeleteContext differentiates from other interfaces.
	IsDeleteContext()
}

type DeleteContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeleteContext() *DeleteContext {
	var p = new(DeleteContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLangParserRULE_delete
	return p
}

func InitEmptyDeleteContext(p *DeleteContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLangParserRULE_delete
}

func (*DeleteContext) IsDeleteContext() {}

func NewDeleteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeleteContext {
	var p = new(DeleteContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryLangParserRULE_delete

	return p
}

func (s *DeleteContext) GetParser() antlr.Parser { return s.parser }

func (s *DeleteContext) BENCHMARK_NAME() antlr.TerminalNode {
	return s.GetToken(QueryLangParserBENCHMARK_NAME, 0)
}

func (s *DeleteContext) BTREE_TYPE() antlr.TerminalNode {
	return s.GetToken(QueryLangParserBTREE_TYPE, 0)
}

func (s *DeleteContext) KeyList() IKeyListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeyListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKeyListContext)
}

func (s *DeleteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeleteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeleteContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryLangVisitor:
		return t.VisitDelete(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QueryLangParser) Delete_() (localctx IDeleteContext) {
	localctx = NewDeleteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, QueryLangParserRULE_delete)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(51)
		p.Match(QueryLangParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(52)
		p.Match(QueryLangParserBENCHMARK_NAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(53)
		p.Match(QueryLangParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(54)
		p.Match(QueryLangParserBTREE_TYPE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(55)
		p.Match(QueryLangParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(56)
		p.KeyList()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICompareLookupContext is an interface to support dynamic dispatch.
type ICompareLookupContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BENCHMARK_NAME() antlr.TerminalNode
	KeyList() IKeyListContext

	// IsCompareLookupContext differentiates from other interfaces.
	IsCompareLookupContext()
}

type CompareLookupContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompareLookupContext() *CompareLookupContext {
	var p = new(CompareLookupContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLangParserRULE_compareLookup
	return p
}

func InitEmptyCompareLookupContext(p *CompareLookupContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLangParserRULE_compareLookup
}

func (*CompareLookupContext) IsCompareLookupContext() {}

func NewCompareLookupContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompareLookupContext {
	var p = new(CompareLookupContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryLangParserRULE_compareLookup

	return p
}

func (s *CompareLookupContext) GetParser() antlr.Parser { return s.parser }

func (s *CompareLookupContext) BENCHMARK_NAME() antlr.TerminalNode {
	return s.GetToken(QueryLangParserBENCHMARK_NAME, 0)
}

func (s *CompareLookupContext) KeyList() IKeyListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeyListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKeyListContext)
}

func (s *CompareLookupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompareLookupContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CompareLookupContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryLangVisitor:
		return t.VisitCompareLookup(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QueryLangParser) CompareLookup() (localctx ICompareLookupContext) {
	localctx = NewCompareLookupContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, QueryLangParserRULE_compareLookup)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(58)
		p.Match(QueryLangParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(59)
		p.Match(QueryLangParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(60)
		p.Match(QueryLangParserBENCHMARK_NAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(61)
		p.Match(QueryLangParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(62)
		p.KeyList()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDataListContext is an interface to support dynamic dispatch.
type IDataListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllDataEntry() []IDataEntryContext
	DataEntry(i int) IDataEntryContext

	// IsDataListContext differentiates from other interfaces.
	IsDataListContext()
}

type DataListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDataListContext() *DataListContext {
	var p = new(DataListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLangParserRULE_dataList
	return p
}

func InitEmptyDataListContext(p *DataListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLangParserRULE_dataList
}

func (*DataListContext) IsDataListContext() {}

func NewDataListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DataListContext {
	var p = new(DataListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryLangParserRULE_dataList

	return p
}

func (s *DataListContext) GetParser() antlr.Parser { return s.parser }

func (s *DataListContext) AllDataEntry() []IDataEntryContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDataEntryContext); ok {
			len++
		}
	}

	tst := make([]IDataEntryContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDataEntryContext); ok {
			tst[i] = t.(IDataEntryContext)
			i++
		}
	}

	return tst
}

func (s *DataListContext) DataEntry(i int) IDataEntryContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDataEntryContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDataEntryContext)
}

func (s *DataListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DataListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DataListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryLangVisitor:
		return t.VisitDataList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QueryLangParser) DataList() (localctx IDataListContext) {
	localctx = NewDataListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, QueryLangParserRULE_dataList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(64)
		p.DataEntry()
	}
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == QueryLangParserT__3 {
		{
			p.SetState(65)
			p.Match(QueryLangParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(66)
			p.DataEntry()
		}

		p.SetState(71)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDataEntryContext is an interface to support dynamic dispatch.
type IDataEntryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllINT_LIT() []antlr.TerminalNode
	INT_LIT(i int) antlr.TerminalNode

	// IsDataEntryContext differentiates from other interfaces.
	IsDataEntryContext()
}

type DataEntryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDataEntryContext() *DataEntryContext {
	var p = new(DataEntryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLangParserRULE_dataEntry
	return p
}

func InitEmptyDataEntryContext(p *DataEntryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLangParserRULE_dataEntry
}

func (*DataEntryContext) IsDataEntryContext() {}

func NewDataEntryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DataEntryContext {
	var p = new(DataEntryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryLangParserRULE_dataEntry

	return p
}

func (s *DataEntryContext) GetParser() antlr.Parser { return s.parser }

func (s *DataEntryContext) AllINT_LIT() []antlr.TerminalNode {
	return s.GetTokens(QueryLangParserINT_LIT)
}

func (s *DataEntryContext) INT_LIT(i int) antlr.TerminalNode {
	return s.GetToken(QueryLangParserINT_LIT, i)
}

func (s *DataEntryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DataEntryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DataEntryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryLangVisitor:
		return t.VisitDataEntry(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QueryLangParser) DataEntry() (localctx IDataEntryContext) {
	localctx = NewDataEntryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, QueryLangParserRULE_dataEntry)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(72)
		p.Match(QueryLangParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(73)
		p.Match(QueryLangParserINT_LIT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(74)
		p.Match(QueryLangParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(75)
		p.Match(QueryLangParserINT_LIT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(76)
		p.Match(QueryLangParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IKeyListContext is an interface to support dynamic dispatch.
type IKeyListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllINT_LIT() []antlr.TerminalNode
	INT_LIT(i int) antlr.TerminalNode

	// IsKeyListContext differentiates from other interfaces.
	IsKeyListContext()
}

type KeyListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeyListContext() *KeyListContext {
	var p = new(KeyListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLangParserRULE_keyList
	return p
}

func InitEmptyKeyListContext(p *KeyListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLangParserRULE_keyList
}

func (*KeyListContext) IsKeyListContext() {}

func NewKeyListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyListContext {
	var p = new(KeyListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryLangParserRULE_keyList

	return p
}

func (s *KeyListContext) GetParser() antlr.Parser { return s.parser }

func (s *KeyListContext) AllINT_LIT() []antlr.TerminalNode {
	return s.GetTokens(QueryLangParserINT_LIT)
}

func (s *KeyListContext) INT_LIT(i int) antlr.TerminalNode {
	return s.GetToken(QueryLangParserINT_LIT, i)
}

func (s *KeyListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeyListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryLangVisitor:
		return t.VisitKeyList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QueryLangParser) KeyList() (localctx IKeyListContext) {
	localctx = NewKeyListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, QueryLangParserRULE_keyList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(78)
		p.Match(QueryLangParserINT_LIT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == QueryLangParserT__3 {
		{
			p.SetState(79)
			p.Match(QueryLangParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(80)
			p.Match(QueryLangParserINT_LIT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(85)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
