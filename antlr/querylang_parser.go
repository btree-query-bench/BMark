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
		"", "'CREATE'", "'BENCHMARK'", "'SEARCH'", "','", "'INSERT'", "'DELETE'",
		"'START'", "'('", "')'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "BENCHMARK_NAME", "INDEX_TYPE",
		"INT_LIT", "WS",
	}
	staticData.RuleNames = []string{
		"query", "createBenchmark", "search", "modify", "insert", "delete",
		"start", "dataList", "dataEntry", "keyList",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 13, 87, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 1, 0, 1,
		0, 1, 0, 1, 0, 3, 0, 25, 8, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 3, 2, 36, 8, 2, 1, 2, 1, 2, 1, 3, 1, 3, 3, 3, 42, 8, 3, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 49, 8, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 3, 5, 58, 8, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1,
		7, 5, 7, 68, 8, 7, 10, 7, 12, 7, 71, 9, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 9, 1, 9, 1, 9, 5, 9, 82, 8, 9, 10, 9, 12, 9, 85, 9, 9, 1, 9, 0,
		0, 10, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 0, 0, 85, 0, 24, 1, 0, 0, 0,
		2, 26, 1, 0, 0, 0, 4, 30, 1, 0, 0, 0, 6, 41, 1, 0, 0, 0, 8, 43, 1, 0, 0,
		0, 10, 52, 1, 0, 0, 0, 12, 61, 1, 0, 0, 0, 14, 64, 1, 0, 0, 0, 16, 72,
		1, 0, 0, 0, 18, 78, 1, 0, 0, 0, 20, 25, 3, 2, 1, 0, 21, 25, 3, 4, 2, 0,
		22, 25, 3, 6, 3, 0, 23, 25, 3, 12, 6, 0, 24, 20, 1, 0, 0, 0, 24, 21, 1,
		0, 0, 0, 24, 22, 1, 0, 0, 0, 24, 23, 1, 0, 0, 0, 25, 1, 1, 0, 0, 0, 26,
		27, 5, 1, 0, 0, 27, 28, 5, 2, 0, 0, 28, 29, 5, 10, 0, 0, 29, 3, 1, 0, 0,
		0, 30, 31, 5, 3, 0, 0, 31, 32, 5, 10, 0, 0, 32, 35, 5, 4, 0, 0, 33, 34,
		5, 11, 0, 0, 34, 36, 5, 4, 0, 0, 35, 33, 1, 0, 0, 0, 35, 36, 1, 0, 0, 0,
		36, 37, 1, 0, 0, 0, 37, 38, 3, 18, 9, 0, 38, 5, 1, 0, 0, 0, 39, 42, 3,
		8, 4, 0, 40, 42, 3, 10, 5, 0, 41, 39, 1, 0, 0, 0, 41, 40, 1, 0, 0, 0, 42,
		7, 1, 0, 0, 0, 43, 44, 5, 5, 0, 0, 44, 45, 5, 10, 0, 0, 45, 48, 5, 4, 0,
		0, 46, 47, 5, 11, 0, 0, 47, 49, 5, 4, 0, 0, 48, 46, 1, 0, 0, 0, 48, 49,
		1, 0, 0, 0, 49, 50, 1, 0, 0, 0, 50, 51, 3, 14, 7, 0, 51, 9, 1, 0, 0, 0,
		52, 53, 5, 6, 0, 0, 53, 54, 5, 10, 0, 0, 54, 57, 5, 4, 0, 0, 55, 56, 5,
		11, 0, 0, 56, 58, 5, 4, 0, 0, 57, 55, 1, 0, 0, 0, 57, 58, 1, 0, 0, 0, 58,
		59, 1, 0, 0, 0, 59, 60, 3, 18, 9, 0, 60, 11, 1, 0, 0, 0, 61, 62, 5, 7,
		0, 0, 62, 63, 5, 10, 0, 0, 63, 13, 1, 0, 0, 0, 64, 69, 3, 16, 8, 0, 65,
		66, 5, 4, 0, 0, 66, 68, 3, 16, 8, 0, 67, 65, 1, 0, 0, 0, 68, 71, 1, 0,
		0, 0, 69, 67, 1, 0, 0, 0, 69, 70, 1, 0, 0, 0, 70, 15, 1, 0, 0, 0, 71, 69,
		1, 0, 0, 0, 72, 73, 5, 8, 0, 0, 73, 74, 5, 12, 0, 0, 74, 75, 5, 4, 0, 0,
		75, 76, 5, 12, 0, 0, 76, 77, 5, 9, 0, 0, 77, 17, 1, 0, 0, 0, 78, 83, 5,
		12, 0, 0, 79, 80, 5, 4, 0, 0, 80, 82, 5, 12, 0, 0, 81, 79, 1, 0, 0, 0,
		82, 85, 1, 0, 0, 0, 83, 81, 1, 0, 0, 0, 83, 84, 1, 0, 0, 0, 84, 19, 1,
		0, 0, 0, 85, 83, 1, 0, 0, 0, 7, 24, 35, 41, 48, 57, 69, 83,
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
	QueryLangParserBENCHMARK_NAME = 10
	QueryLangParserINDEX_TYPE     = 11
	QueryLangParserINT_LIT        = 12
	QueryLangParserWS             = 13
)

// QueryLangParser rules.
const (
	QueryLangParserRULE_query           = 0
	QueryLangParserRULE_createBenchmark = 1
	QueryLangParserRULE_search          = 2
	QueryLangParserRULE_modify          = 3
	QueryLangParserRULE_insert          = 4
	QueryLangParserRULE_delete          = 5
	QueryLangParserRULE_start           = 6
	QueryLangParserRULE_dataList        = 7
	QueryLangParserRULE_dataEntry       = 8
	QueryLangParserRULE_keyList         = 9
)

// IQueryContext is an interface to support dynamic dispatch.
type IQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CreateBenchmark() ICreateBenchmarkContext
	Search() ISearchContext
	Modify() IModifyContext
	Start_() IStartContext

	// IsQueryContext differentiates from other interfaces.
	IsQueryContext()
}

type QueryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQueryContext() *QueryContext {
	var p = new(QueryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLangParserRULE_query
	return p
}

func InitEmptyQueryContext(p *QueryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLangParserRULE_query
}

func (*QueryContext) IsQueryContext() {}

func NewQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QueryContext {
	var p = new(QueryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryLangParserRULE_query

	return p
}

func (s *QueryContext) GetParser() antlr.Parser { return s.parser }

func (s *QueryContext) CreateBenchmark() ICreateBenchmarkContext {
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

func (s *QueryContext) Search() ISearchContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISearchContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISearchContext)
}

func (s *QueryContext) Modify() IModifyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IModifyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IModifyContext)
}

func (s *QueryContext) Start_() IStartContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStartContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStartContext)
}

func (s *QueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QueryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryLangVisitor:
		return t.VisitQuery(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QueryLangParser) Query() (localctx IQueryContext) {
	localctx = NewQueryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, QueryLangParserRULE_query)
	p.SetState(24)
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
			p.Search()
		}

	case QueryLangParserT__4, QueryLangParserT__5:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(22)
			p.Modify()
		}

	case QueryLangParserT__6:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(23)
			p.Start_()
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
		p.SetState(26)
		p.Match(QueryLangParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(27)
		p.Match(QueryLangParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(28)
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

// ISearchContext is an interface to support dynamic dispatch.
type ISearchContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BENCHMARK_NAME() antlr.TerminalNode
	KeyList() IKeyListContext
	INDEX_TYPE() antlr.TerminalNode

	// IsSearchContext differentiates from other interfaces.
	IsSearchContext()
}

type SearchContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySearchContext() *SearchContext {
	var p = new(SearchContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLangParserRULE_search
	return p
}

func InitEmptySearchContext(p *SearchContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLangParserRULE_search
}

func (*SearchContext) IsSearchContext() {}

func NewSearchContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SearchContext {
	var p = new(SearchContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryLangParserRULE_search

	return p
}

func (s *SearchContext) GetParser() antlr.Parser { return s.parser }

func (s *SearchContext) BENCHMARK_NAME() antlr.TerminalNode {
	return s.GetToken(QueryLangParserBENCHMARK_NAME, 0)
}

func (s *SearchContext) KeyList() IKeyListContext {
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

func (s *SearchContext) INDEX_TYPE() antlr.TerminalNode {
	return s.GetToken(QueryLangParserINDEX_TYPE, 0)
}

func (s *SearchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SearchContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SearchContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryLangVisitor:
		return t.VisitSearch(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QueryLangParser) Search() (localctx ISearchContext) {
	localctx = NewSearchContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, QueryLangParserRULE_search)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(30)
		p.Match(QueryLangParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(31)
		p.Match(QueryLangParserBENCHMARK_NAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(32)
		p.Match(QueryLangParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(35)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == QueryLangParserINDEX_TYPE {
		{
			p.SetState(33)
			p.Match(QueryLangParserINDEX_TYPE)
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

	}
	{
		p.SetState(37)
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

// IModifyContext is an interface to support dynamic dispatch.
type IModifyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Insert() IInsertContext
	Delete_() IDeleteContext

	// IsModifyContext differentiates from other interfaces.
	IsModifyContext()
}

type ModifyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModifyContext() *ModifyContext {
	var p = new(ModifyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLangParserRULE_modify
	return p
}

func InitEmptyModifyContext(p *ModifyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLangParserRULE_modify
}

func (*ModifyContext) IsModifyContext() {}

func NewModifyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModifyContext {
	var p = new(ModifyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryLangParserRULE_modify

	return p
}

func (s *ModifyContext) GetParser() antlr.Parser { return s.parser }

func (s *ModifyContext) Insert() IInsertContext {
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

func (s *ModifyContext) Delete_() IDeleteContext {
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

func (s *ModifyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModifyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModifyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryLangVisitor:
		return t.VisitModify(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QueryLangParser) Modify() (localctx IModifyContext) {
	localctx = NewModifyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, QueryLangParserRULE_modify)
	p.SetState(41)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case QueryLangParserT__4:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(39)
			p.Insert()
		}

	case QueryLangParserT__5:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(40)
			p.Delete_()
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

// IInsertContext is an interface to support dynamic dispatch.
type IInsertContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BENCHMARK_NAME() antlr.TerminalNode
	DataList() IDataListContext
	INDEX_TYPE() antlr.TerminalNode

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

func (s *InsertContext) INDEX_TYPE() antlr.TerminalNode {
	return s.GetToken(QueryLangParserINDEX_TYPE, 0)
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
	p.EnterRule(localctx, 8, QueryLangParserRULE_insert)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(43)
		p.Match(QueryLangParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(44)
		p.Match(QueryLangParserBENCHMARK_NAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(45)
		p.Match(QueryLangParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(48)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == QueryLangParserINDEX_TYPE {
		{
			p.SetState(46)
			p.Match(QueryLangParserINDEX_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(47)
			p.Match(QueryLangParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(50)
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

// IDeleteContext is an interface to support dynamic dispatch.
type IDeleteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BENCHMARK_NAME() antlr.TerminalNode
	KeyList() IKeyListContext
	INDEX_TYPE() antlr.TerminalNode

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

func (s *DeleteContext) INDEX_TYPE() antlr.TerminalNode {
	return s.GetToken(QueryLangParserINDEX_TYPE, 0)
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
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(52)
		p.Match(QueryLangParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(53)
		p.Match(QueryLangParserBENCHMARK_NAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(54)
		p.Match(QueryLangParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(57)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == QueryLangParserINDEX_TYPE {
		{
			p.SetState(55)
			p.Match(QueryLangParserINDEX_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(56)
			p.Match(QueryLangParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(59)
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

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BENCHMARK_NAME() antlr.TerminalNode

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLangParserRULE_start
	return p
}

func InitEmptyStartContext(p *StartContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLangParserRULE_start
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryLangParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) BENCHMARK_NAME() antlr.TerminalNode {
	return s.GetToken(QueryLangParserBENCHMARK_NAME, 0)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryLangVisitor:
		return t.VisitStart(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QueryLangParser) Start_() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, QueryLangParserRULE_start)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(61)
		p.Match(QueryLangParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(62)
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
		p.Match(QueryLangParserT__7)
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
		p.Match(QueryLangParserT__8)
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
