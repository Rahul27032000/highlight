// Code generated from ./antlr/SearchGrammar.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser

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

type SearchGrammarLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var SearchGrammarLexerLexerStaticData struct {
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

func searchgrammarlexerLexerInit() {
	staticData := &SearchGrammarLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'AND'", "'OR'", "'NOT'", "'EXISTS'", "'!'", "'='", "'!='", "'<'",
		"'<='", "'>'", "'>='", "'('", "')'", "':'",
	}
	staticData.SymbolicNames = []string{
		"", "AND", "OR", "NOT", "EXISTS", "BANG", "EQ", "NEQ", "LT", "LTE",
		"GT", "GTE", "LPAREN", "RPAREN", "COLON", "ID", "STRING", "VALUE", "WS",
		"ERROR_CHARACTERS",
	}
	staticData.RuleNames = []string{
		"AND", "OR", "NOT", "EXISTS", "BANG", "EQ", "NEQ", "LT", "LTE", "GT",
		"GTE", "LPAREN", "RPAREN", "COLON", "ID", "STRING", "VALUE", "WS", "ERROR_CHARACTERS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 19, 108, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 1, 0, 1, 0, 1, 0, 1, 0,
		1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8,
		1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12,
		1, 13, 1, 13, 1, 14, 4, 14, 82, 8, 14, 11, 14, 12, 14, 83, 1, 15, 1, 15,
		5, 15, 88, 8, 15, 10, 15, 12, 15, 91, 9, 15, 1, 15, 1, 15, 1, 16, 4, 16,
		96, 8, 16, 11, 16, 12, 16, 97, 1, 17, 4, 17, 101, 8, 17, 11, 17, 12, 17,
		102, 1, 17, 1, 17, 1, 18, 1, 18, 1, 89, 0, 19, 1, 1, 3, 2, 5, 3, 7, 4,
		9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14,
		29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 1, 0, 13, 2, 0, 65, 65, 97, 97,
		2, 0, 78, 78, 110, 110, 2, 0, 68, 68, 100, 100, 2, 0, 79, 79, 111, 111,
		2, 0, 82, 82, 114, 114, 2, 0, 84, 84, 116, 116, 2, 0, 69, 69, 101, 101,
		2, 0, 88, 88, 120, 120, 2, 0, 73, 73, 105, 105, 2, 0, 83, 83, 115, 115,
		6, 0, 42, 42, 45, 46, 48, 57, 65, 90, 95, 95, 97, 122, 6, 0, 9, 10, 12,
		13, 32, 33, 40, 41, 58, 58, 60, 62, 3, 0, 9, 10, 12, 13, 32, 32, 111, 0,
		1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0,
		9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0,
		0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0,
		0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0,
		0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 1, 39, 1,
		0, 0, 0, 3, 43, 1, 0, 0, 0, 5, 46, 1, 0, 0, 0, 7, 50, 1, 0, 0, 0, 9, 57,
		1, 0, 0, 0, 11, 59, 1, 0, 0, 0, 13, 61, 1, 0, 0, 0, 15, 64, 1, 0, 0, 0,
		17, 66, 1, 0, 0, 0, 19, 69, 1, 0, 0, 0, 21, 71, 1, 0, 0, 0, 23, 74, 1,
		0, 0, 0, 25, 76, 1, 0, 0, 0, 27, 78, 1, 0, 0, 0, 29, 81, 1, 0, 0, 0, 31,
		85, 1, 0, 0, 0, 33, 95, 1, 0, 0, 0, 35, 100, 1, 0, 0, 0, 37, 106, 1, 0,
		0, 0, 39, 40, 7, 0, 0, 0, 40, 41, 7, 1, 0, 0, 41, 42, 7, 2, 0, 0, 42, 2,
		1, 0, 0, 0, 43, 44, 7, 3, 0, 0, 44, 45, 7, 4, 0, 0, 45, 4, 1, 0, 0, 0,
		46, 47, 7, 1, 0, 0, 47, 48, 7, 3, 0, 0, 48, 49, 7, 5, 0, 0, 49, 6, 1, 0,
		0, 0, 50, 51, 7, 6, 0, 0, 51, 52, 7, 7, 0, 0, 52, 53, 7, 8, 0, 0, 53, 54,
		7, 9, 0, 0, 54, 55, 7, 5, 0, 0, 55, 56, 7, 9, 0, 0, 56, 8, 1, 0, 0, 0,
		57, 58, 5, 33, 0, 0, 58, 10, 1, 0, 0, 0, 59, 60, 5, 61, 0, 0, 60, 12, 1,
		0, 0, 0, 61, 62, 5, 33, 0, 0, 62, 63, 5, 61, 0, 0, 63, 14, 1, 0, 0, 0,
		64, 65, 5, 60, 0, 0, 65, 16, 1, 0, 0, 0, 66, 67, 5, 60, 0, 0, 67, 68, 5,
		61, 0, 0, 68, 18, 1, 0, 0, 0, 69, 70, 5, 62, 0, 0, 70, 20, 1, 0, 0, 0,
		71, 72, 5, 62, 0, 0, 72, 73, 5, 61, 0, 0, 73, 22, 1, 0, 0, 0, 74, 75, 5,
		40, 0, 0, 75, 24, 1, 0, 0, 0, 76, 77, 5, 41, 0, 0, 77, 26, 1, 0, 0, 0,
		78, 79, 5, 58, 0, 0, 79, 28, 1, 0, 0, 0, 80, 82, 7, 10, 0, 0, 81, 80, 1,
		0, 0, 0, 82, 83, 1, 0, 0, 0, 83, 81, 1, 0, 0, 0, 83, 84, 1, 0, 0, 0, 84,
		30, 1, 0, 0, 0, 85, 89, 5, 34, 0, 0, 86, 88, 9, 0, 0, 0, 87, 86, 1, 0,
		0, 0, 88, 91, 1, 0, 0, 0, 89, 90, 1, 0, 0, 0, 89, 87, 1, 0, 0, 0, 90, 92,
		1, 0, 0, 0, 91, 89, 1, 0, 0, 0, 92, 93, 5, 34, 0, 0, 93, 32, 1, 0, 0, 0,
		94, 96, 8, 11, 0, 0, 95, 94, 1, 0, 0, 0, 96, 97, 1, 0, 0, 0, 97, 95, 1,
		0, 0, 0, 97, 98, 1, 0, 0, 0, 98, 34, 1, 0, 0, 0, 99, 101, 7, 12, 0, 0,
		100, 99, 1, 0, 0, 0, 101, 102, 1, 0, 0, 0, 102, 100, 1, 0, 0, 0, 102, 103,
		1, 0, 0, 0, 103, 104, 1, 0, 0, 0, 104, 105, 6, 17, 0, 0, 105, 36, 1, 0,
		0, 0, 106, 107, 9, 0, 0, 0, 107, 38, 1, 0, 0, 0, 5, 0, 83, 89, 97, 102,
		1, 6, 0, 0,
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

// SearchGrammarLexerInit initializes any static state used to implement SearchGrammarLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewSearchGrammarLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func SearchGrammarLexerInit() {
	staticData := &SearchGrammarLexerLexerStaticData
	staticData.once.Do(searchgrammarlexerLexerInit)
}

// NewSearchGrammarLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewSearchGrammarLexer(input antlr.CharStream) *SearchGrammarLexer {
	SearchGrammarLexerInit()
	l := new(SearchGrammarLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &SearchGrammarLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "SearchGrammar.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// SearchGrammarLexer tokens.
const (
	SearchGrammarLexerAND              = 1
	SearchGrammarLexerOR               = 2
	SearchGrammarLexerNOT              = 3
	SearchGrammarLexerEXISTS           = 4
	SearchGrammarLexerBANG             = 5
	SearchGrammarLexerEQ               = 6
	SearchGrammarLexerNEQ              = 7
	SearchGrammarLexerLT               = 8
	SearchGrammarLexerLTE              = 9
	SearchGrammarLexerGT               = 10
	SearchGrammarLexerGTE              = 11
	SearchGrammarLexerLPAREN           = 12
	SearchGrammarLexerRPAREN           = 13
	SearchGrammarLexerCOLON            = 14
	SearchGrammarLexerID               = 15
	SearchGrammarLexerSTRING           = 16
	SearchGrammarLexerVALUE            = 17
	SearchGrammarLexerWS               = 18
	SearchGrammarLexerERROR_CHARACTERS = 19
)
