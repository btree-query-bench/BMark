// Code generated from QueryLang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package antlr

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type QueryLangLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var QueryLangLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func querylanglexerLexerInit() {
	staticData := &QueryLangLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'CREATE'", "'BENCHMARK'", "'SEARCH'", "','", "'INSERT'", "'DELETE'",
		"'START'", "'('", "')'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "BENCHMARK_NAME", "INDEX_TYPE",
		"INT_LIT", "WS",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"BENCHMARK_NAME", "INDEX_TYPE", "INT_LIT", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 13, 113, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 5, 9, 80,
		8, 9, 10, 9, 12, 9, 83, 9, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 100,
		8, 10, 1, 11, 4, 11, 103, 8, 11, 11, 11, 12, 11, 104, 1, 12, 4, 12, 108,
		8, 12, 11, 12, 12, 12, 109, 1, 12, 1, 12, 0, 0, 13, 1, 1, 3, 2, 5, 3, 7,
		4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 1,
		0, 4, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97,
		122, 1, 0, 48, 57, 3, 0, 9, 10, 13, 13, 32, 32, 117, 0, 1, 1, 0, 0, 0,
		0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0,
		0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0,
		0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0,
		0, 0, 1, 27, 1, 0, 0, 0, 3, 34, 1, 0, 0, 0, 5, 44, 1, 0, 0, 0, 7, 51, 1,
		0, 0, 0, 9, 53, 1, 0, 0, 0, 11, 60, 1, 0, 0, 0, 13, 67, 1, 0, 0, 0, 15,
		73, 1, 0, 0, 0, 17, 75, 1, 0, 0, 0, 19, 77, 1, 0, 0, 0, 21, 99, 1, 0, 0,
		0, 23, 102, 1, 0, 0, 0, 25, 107, 1, 0, 0, 0, 27, 28, 5, 67, 0, 0, 28, 29,
		5, 82, 0, 0, 29, 30, 5, 69, 0, 0, 30, 31, 5, 65, 0, 0, 31, 32, 5, 84, 0,
		0, 32, 33, 5, 69, 0, 0, 33, 2, 1, 0, 0, 0, 34, 35, 5, 66, 0, 0, 35, 36,
		5, 69, 0, 0, 36, 37, 5, 78, 0, 0, 37, 38, 5, 67, 0, 0, 38, 39, 5, 72, 0,
		0, 39, 40, 5, 77, 0, 0, 40, 41, 5, 65, 0, 0, 41, 42, 5, 82, 0, 0, 42, 43,
		5, 75, 0, 0, 43, 4, 1, 0, 0, 0, 44, 45, 5, 83, 0, 0, 45, 46, 5, 69, 0,
		0, 46, 47, 5, 65, 0, 0, 47, 48, 5, 82, 0, 0, 48, 49, 5, 67, 0, 0, 49, 50,
		5, 72, 0, 0, 50, 6, 1, 0, 0, 0, 51, 52, 5, 44, 0, 0, 52, 8, 1, 0, 0, 0,
		53, 54, 5, 73, 0, 0, 54, 55, 5, 78, 0, 0, 55, 56, 5, 83, 0, 0, 56, 57,
		5, 69, 0, 0, 57, 58, 5, 82, 0, 0, 58, 59, 5, 84, 0, 0, 59, 10, 1, 0, 0,
		0, 60, 61, 5, 68, 0, 0, 61, 62, 5, 69, 0, 0, 62, 63, 5, 76, 0, 0, 63, 64,
		5, 69, 0, 0, 64, 65, 5, 84, 0, 0, 65, 66, 5, 69, 0, 0, 66, 12, 1, 0, 0,
		0, 67, 68, 5, 83, 0, 0, 68, 69, 5, 84, 0, 0, 69, 70, 5, 65, 0, 0, 70, 71,
		5, 82, 0, 0, 71, 72, 5, 84, 0, 0, 72, 14, 1, 0, 0, 0, 73, 74, 5, 40, 0,
		0, 74, 16, 1, 0, 0, 0, 75, 76, 5, 41, 0, 0, 76, 18, 1, 0, 0, 0, 77, 81,
		7, 0, 0, 0, 78, 80, 7, 1, 0, 0, 79, 78, 1, 0, 0, 0, 80, 83, 1, 0, 0, 0,
		81, 79, 1, 0, 0, 0, 81, 82, 1, 0, 0, 0, 82, 20, 1, 0, 0, 0, 83, 81, 1,
		0, 0, 0, 84, 85, 5, 66, 0, 0, 85, 86, 5, 84, 0, 0, 86, 87, 5, 82, 0, 0,
		87, 88, 5, 69, 0, 0, 88, 100, 5, 69, 0, 0, 89, 90, 5, 66, 0, 0, 90, 91,
		5, 43, 0, 0, 91, 92, 5, 84, 0, 0, 92, 93, 5, 82, 0, 0, 93, 94, 5, 69, 0,
		0, 94, 100, 5, 69, 0, 0, 95, 96, 5, 76, 0, 0, 96, 97, 5, 73, 0, 0, 97,
		98, 5, 83, 0, 0, 98, 100, 5, 84, 0, 0, 99, 84, 1, 0, 0, 0, 99, 89, 1, 0,
		0, 0, 99, 95, 1, 0, 0, 0, 100, 22, 1, 0, 0, 0, 101, 103, 7, 2, 0, 0, 102,
		101, 1, 0, 0, 0, 103, 104, 1, 0, 0, 0, 104, 102, 1, 0, 0, 0, 104, 105,
		1, 0, 0, 0, 105, 24, 1, 0, 0, 0, 106, 108, 7, 3, 0, 0, 107, 106, 1, 0,
		0, 0, 108, 109, 1, 0, 0, 0, 109, 107, 1, 0, 0, 0, 109, 110, 1, 0, 0, 0,
		110, 111, 1, 0, 0, 0, 111, 112, 6, 12, 0, 0, 112, 26, 1, 0, 0, 0, 5, 0,
		81, 99, 104, 109, 1, 6, 0, 0,
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

// QueryLangLexerInit initializes any static state used to implement QueryLangLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewQueryLangLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func QueryLangLexerInit() {
	staticData := &QueryLangLexerLexerStaticData
	staticData.once.Do(querylanglexerLexerInit)
}

// NewQueryLangLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewQueryLangLexer(input antlr.CharStream) *QueryLangLexer {
	QueryLangLexerInit()
	l := new(QueryLangLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &QueryLangLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "QueryLang.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// QueryLangLexer tokens.
const (
	QueryLangLexerT__0           = 1
	QueryLangLexerT__1           = 2
	QueryLangLexerT__2           = 3
	QueryLangLexerT__3           = 4
	QueryLangLexerT__4           = 5
	QueryLangLexerT__5           = 6
	QueryLangLexerT__6           = 7
	QueryLangLexerT__7           = 8
	QueryLangLexerT__8           = 9
	QueryLangLexerBENCHMARK_NAME = 10
	QueryLangLexerINDEX_TYPE     = 11
	QueryLangLexerINT_LIT        = 12
	QueryLangLexerWS             = 13
)
