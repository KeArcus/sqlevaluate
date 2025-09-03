// Code generated from grammar/SQL.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // SQL

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

type SQLParser struct {
	*antlr.BaseParser
}

var SQLParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func sqlParserInit() {
	staticData := &SQLParserStaticData
	staticData.LiteralNames = []string{
		"", "','", "'('", "')'", "'='", "'!='", "'<>'", "'<'", "'<='", "'>'",
		"'>='", "", "", "", "", "", "", "", "", "", "", "", "", "'*'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "SELECT", "FROM", "WHERE",
		"AND", "OR", "NOT", "IN", "LIKE", "IS", "NULL", "TRUE", "FALSE", "STAR",
		"STRING", "NUMBER", "IDENTIFIER", "WS",
	}
	staticData.RuleNames = []string{
		"sql_statement", "select_statement", "select_list", "column_list", "table_name",
		"column_name", "where_clause", "condition", "comparison_expression",
		"expression", "comparison_operator", "value", "value_list", "string_literal",
		"number_literal", "boolean_literal", "function_call", "function_name",
		"function_args", "identifier",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 27, 170, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 1, 0, 1, 0, 1,
		0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 50, 8, 1, 1, 2, 1, 2, 3, 2,
		54, 8, 2, 1, 3, 1, 3, 1, 3, 5, 3, 59, 8, 3, 10, 3, 12, 3, 62, 9, 3, 1,
		4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 107, 8, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 5, 7, 115, 8, 7, 10, 7, 12, 7, 118, 9, 7, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 3, 9, 127, 8, 9, 1, 10, 1, 10, 1, 11,
		1, 11, 1, 11, 1, 11, 3, 11, 135, 8, 11, 1, 12, 1, 12, 1, 12, 5, 12, 140,
		8, 12, 10, 12, 12, 12, 143, 9, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1,
		15, 1, 16, 1, 16, 1, 16, 3, 16, 154, 8, 16, 1, 16, 1, 16, 1, 17, 1, 17,
		1, 18, 1, 18, 1, 18, 5, 18, 163, 8, 18, 10, 18, 12, 18, 166, 9, 18, 1,
		19, 1, 19, 1, 19, 0, 1, 14, 20, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
		22, 24, 26, 28, 30, 32, 34, 36, 38, 0, 2, 1, 0, 4, 10, 1, 0, 21, 22, 169,
		0, 40, 1, 0, 0, 0, 2, 43, 1, 0, 0, 0, 4, 53, 1, 0, 0, 0, 6, 55, 1, 0, 0,
		0, 8, 63, 1, 0, 0, 0, 10, 65, 1, 0, 0, 0, 12, 67, 1, 0, 0, 0, 14, 106,
		1, 0, 0, 0, 16, 119, 1, 0, 0, 0, 18, 126, 1, 0, 0, 0, 20, 128, 1, 0, 0,
		0, 22, 134, 1, 0, 0, 0, 24, 136, 1, 0, 0, 0, 26, 144, 1, 0, 0, 0, 28, 146,
		1, 0, 0, 0, 30, 148, 1, 0, 0, 0, 32, 150, 1, 0, 0, 0, 34, 157, 1, 0, 0,
		0, 36, 159, 1, 0, 0, 0, 38, 167, 1, 0, 0, 0, 40, 41, 3, 2, 1, 0, 41, 42,
		5, 0, 0, 1, 42, 1, 1, 0, 0, 0, 43, 44, 5, 11, 0, 0, 44, 45, 3, 4, 2, 0,
		45, 46, 5, 12, 0, 0, 46, 49, 3, 8, 4, 0, 47, 48, 5, 13, 0, 0, 48, 50, 3,
		12, 6, 0, 49, 47, 1, 0, 0, 0, 49, 50, 1, 0, 0, 0, 50, 3, 1, 0, 0, 0, 51,
		54, 5, 23, 0, 0, 52, 54, 3, 6, 3, 0, 53, 51, 1, 0, 0, 0, 53, 52, 1, 0,
		0, 0, 54, 5, 1, 0, 0, 0, 55, 60, 3, 10, 5, 0, 56, 57, 5, 1, 0, 0, 57, 59,
		3, 10, 5, 0, 58, 56, 1, 0, 0, 0, 59, 62, 1, 0, 0, 0, 60, 58, 1, 0, 0, 0,
		60, 61, 1, 0, 0, 0, 61, 7, 1, 0, 0, 0, 62, 60, 1, 0, 0, 0, 63, 64, 3, 38,
		19, 0, 64, 9, 1, 0, 0, 0, 65, 66, 3, 38, 19, 0, 66, 11, 1, 0, 0, 0, 67,
		68, 3, 14, 7, 0, 68, 13, 1, 0, 0, 0, 69, 70, 6, 7, -1, 0, 70, 71, 5, 2,
		0, 0, 71, 72, 3, 14, 7, 0, 72, 73, 5, 3, 0, 0, 73, 107, 1, 0, 0, 0, 74,
		107, 3, 16, 8, 0, 75, 76, 3, 18, 9, 0, 76, 77, 5, 17, 0, 0, 77, 78, 5,
		2, 0, 0, 78, 79, 3, 24, 12, 0, 79, 80, 5, 3, 0, 0, 80, 107, 1, 0, 0, 0,
		81, 82, 3, 18, 9, 0, 82, 83, 5, 16, 0, 0, 83, 84, 5, 17, 0, 0, 84, 85,
		5, 2, 0, 0, 85, 86, 3, 24, 12, 0, 86, 87, 5, 3, 0, 0, 87, 107, 1, 0, 0,
		0, 88, 89, 3, 18, 9, 0, 89, 90, 5, 18, 0, 0, 90, 91, 3, 18, 9, 0, 91, 107,
		1, 0, 0, 0, 92, 93, 3, 18, 9, 0, 93, 94, 5, 16, 0, 0, 94, 95, 5, 18, 0,
		0, 95, 96, 3, 18, 9, 0, 96, 107, 1, 0, 0, 0, 97, 98, 3, 18, 9, 0, 98, 99,
		5, 19, 0, 0, 99, 100, 5, 20, 0, 0, 100, 107, 1, 0, 0, 0, 101, 102, 3, 18,
		9, 0, 102, 103, 5, 19, 0, 0, 103, 104, 5, 16, 0, 0, 104, 105, 5, 20, 0,
		0, 105, 107, 1, 0, 0, 0, 106, 69, 1, 0, 0, 0, 106, 74, 1, 0, 0, 0, 106,
		75, 1, 0, 0, 0, 106, 81, 1, 0, 0, 0, 106, 88, 1, 0, 0, 0, 106, 92, 1, 0,
		0, 0, 106, 97, 1, 0, 0, 0, 106, 101, 1, 0, 0, 0, 107, 116, 1, 0, 0, 0,
		108, 109, 10, 10, 0, 0, 109, 110, 5, 14, 0, 0, 110, 115, 3, 14, 7, 11,
		111, 112, 10, 9, 0, 0, 112, 113, 5, 15, 0, 0, 113, 115, 3, 14, 7, 10, 114,
		108, 1, 0, 0, 0, 114, 111, 1, 0, 0, 0, 115, 118, 1, 0, 0, 0, 116, 114,
		1, 0, 0, 0, 116, 117, 1, 0, 0, 0, 117, 15, 1, 0, 0, 0, 118, 116, 1, 0,
		0, 0, 119, 120, 3, 18, 9, 0, 120, 121, 3, 20, 10, 0, 121, 122, 3, 18, 9,
		0, 122, 17, 1, 0, 0, 0, 123, 127, 3, 32, 16, 0, 124, 127, 3, 10, 5, 0,
		125, 127, 3, 22, 11, 0, 126, 123, 1, 0, 0, 0, 126, 124, 1, 0, 0, 0, 126,
		125, 1, 0, 0, 0, 127, 19, 1, 0, 0, 0, 128, 129, 7, 0, 0, 0, 129, 21, 1,
		0, 0, 0, 130, 135, 3, 26, 13, 0, 131, 135, 3, 28, 14, 0, 132, 135, 3, 30,
		15, 0, 133, 135, 5, 20, 0, 0, 134, 130, 1, 0, 0, 0, 134, 131, 1, 0, 0,
		0, 134, 132, 1, 0, 0, 0, 134, 133, 1, 0, 0, 0, 135, 23, 1, 0, 0, 0, 136,
		141, 3, 18, 9, 0, 137, 138, 5, 1, 0, 0, 138, 140, 3, 18, 9, 0, 139, 137,
		1, 0, 0, 0, 140, 143, 1, 0, 0, 0, 141, 139, 1, 0, 0, 0, 141, 142, 1, 0,
		0, 0, 142, 25, 1, 0, 0, 0, 143, 141, 1, 0, 0, 0, 144, 145, 5, 24, 0, 0,
		145, 27, 1, 0, 0, 0, 146, 147, 5, 25, 0, 0, 147, 29, 1, 0, 0, 0, 148, 149,
		7, 1, 0, 0, 149, 31, 1, 0, 0, 0, 150, 151, 3, 34, 17, 0, 151, 153, 5, 2,
		0, 0, 152, 154, 3, 36, 18, 0, 153, 152, 1, 0, 0, 0, 153, 154, 1, 0, 0,
		0, 154, 155, 1, 0, 0, 0, 155, 156, 5, 3, 0, 0, 156, 33, 1, 0, 0, 0, 157,
		158, 3, 38, 19, 0, 158, 35, 1, 0, 0, 0, 159, 164, 3, 18, 9, 0, 160, 161,
		5, 1, 0, 0, 161, 163, 3, 18, 9, 0, 162, 160, 1, 0, 0, 0, 163, 166, 1, 0,
		0, 0, 164, 162, 1, 0, 0, 0, 164, 165, 1, 0, 0, 0, 165, 37, 1, 0, 0, 0,
		166, 164, 1, 0, 0, 0, 167, 168, 5, 26, 0, 0, 168, 39, 1, 0, 0, 0, 11, 49,
		53, 60, 106, 114, 116, 126, 134, 141, 153, 164,
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

// SQLParserInit initializes any static state used to implement SQLParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSQLParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SQLParserInit() {
	staticData := &SQLParserStaticData
	staticData.once.Do(sqlParserInit)
}

// NewSQLParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSQLParser(input antlr.TokenStream) *SQLParser {
	SQLParserInit()
	this := new(SQLParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &SQLParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "SQL.g4"

	return this
}

// SQLParser tokens.
const (
	SQLParserEOF        = antlr.TokenEOF
	SQLParserT__0       = 1
	SQLParserT__1       = 2
	SQLParserT__2       = 3
	SQLParserT__3       = 4
	SQLParserT__4       = 5
	SQLParserT__5       = 6
	SQLParserT__6       = 7
	SQLParserT__7       = 8
	SQLParserT__8       = 9
	SQLParserT__9       = 10
	SQLParserSELECT     = 11
	SQLParserFROM       = 12
	SQLParserWHERE      = 13
	SQLParserAND        = 14
	SQLParserOR         = 15
	SQLParserNOT        = 16
	SQLParserIN         = 17
	SQLParserLIKE       = 18
	SQLParserIS         = 19
	SQLParserNULL       = 20
	SQLParserTRUE       = 21
	SQLParserFALSE      = 22
	SQLParserSTAR       = 23
	SQLParserSTRING     = 24
	SQLParserNUMBER     = 25
	SQLParserIDENTIFIER = 26
	SQLParserWS         = 27
)

// SQLParser rules.
const (
	SQLParserRULE_sql_statement         = 0
	SQLParserRULE_select_statement      = 1
	SQLParserRULE_select_list           = 2
	SQLParserRULE_column_list           = 3
	SQLParserRULE_table_name            = 4
	SQLParserRULE_column_name           = 5
	SQLParserRULE_where_clause          = 6
	SQLParserRULE_condition             = 7
	SQLParserRULE_comparison_expression = 8
	SQLParserRULE_expression            = 9
	SQLParserRULE_comparison_operator   = 10
	SQLParserRULE_value                 = 11
	SQLParserRULE_value_list            = 12
	SQLParserRULE_string_literal        = 13
	SQLParserRULE_number_literal        = 14
	SQLParserRULE_boolean_literal       = 15
	SQLParserRULE_function_call         = 16
	SQLParserRULE_function_name         = 17
	SQLParserRULE_function_args         = 18
	SQLParserRULE_identifier            = 19
)

// ISql_statementContext is an interface to support dynamic dispatch.
type ISql_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Select_statement() ISelect_statementContext
	EOF() antlr.TerminalNode

	// IsSql_statementContext differentiates from other interfaces.
	IsSql_statementContext()
}

type Sql_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySql_statementContext() *Sql_statementContext {
	var p = new(Sql_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_sql_statement
	return p
}

func InitEmptySql_statementContext(p *Sql_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_sql_statement
}

func (*Sql_statementContext) IsSql_statementContext() {}

func NewSql_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Sql_statementContext {
	var p = new(Sql_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_sql_statement

	return p
}

func (s *Sql_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Sql_statementContext) Select_statement() ISelect_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelect_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelect_statementContext)
}

func (s *Sql_statementContext) EOF() antlr.TerminalNode {
	return s.GetToken(SQLParserEOF, 0)
}

func (s *Sql_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Sql_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Sql_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterSql_statement(s)
	}
}

func (s *Sql_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitSql_statement(s)
	}
}

func (p *SQLParser) Sql_statement() (localctx ISql_statementContext) {
	localctx = NewSql_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SQLParserRULE_sql_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(40)
		p.Select_statement()
	}
	{
		p.SetState(41)
		p.Match(SQLParserEOF)
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

// ISelect_statementContext is an interface to support dynamic dispatch.
type ISelect_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SELECT() antlr.TerminalNode
	Select_list() ISelect_listContext
	FROM() antlr.TerminalNode
	Table_name() ITable_nameContext
	WHERE() antlr.TerminalNode
	Where_clause() IWhere_clauseContext

	// IsSelect_statementContext differentiates from other interfaces.
	IsSelect_statementContext()
}

type Select_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelect_statementContext() *Select_statementContext {
	var p = new(Select_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_select_statement
	return p
}

func InitEmptySelect_statementContext(p *Select_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_select_statement
}

func (*Select_statementContext) IsSelect_statementContext() {}

func NewSelect_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Select_statementContext {
	var p = new(Select_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_select_statement

	return p
}

func (s *Select_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Select_statementContext) SELECT() antlr.TerminalNode {
	return s.GetToken(SQLParserSELECT, 0)
}

func (s *Select_statementContext) Select_list() ISelect_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelect_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelect_listContext)
}

func (s *Select_statementContext) FROM() antlr.TerminalNode {
	return s.GetToken(SQLParserFROM, 0)
}

func (s *Select_statementContext) Table_name() ITable_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITable_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITable_nameContext)
}

func (s *Select_statementContext) WHERE() antlr.TerminalNode {
	return s.GetToken(SQLParserWHERE, 0)
}

func (s *Select_statementContext) Where_clause() IWhere_clauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhere_clauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhere_clauseContext)
}

func (s *Select_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Select_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Select_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterSelect_statement(s)
	}
}

func (s *Select_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitSelect_statement(s)
	}
}

func (p *SQLParser) Select_statement() (localctx ISelect_statementContext) {
	localctx = NewSelect_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SQLParserRULE_select_statement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(43)
		p.Match(SQLParserSELECT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(44)
		p.Select_list()
	}
	{
		p.SetState(45)
		p.Match(SQLParserFROM)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(46)
		p.Table_name()
	}
	p.SetState(49)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SQLParserWHERE {
		{
			p.SetState(47)
			p.Match(SQLParserWHERE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(48)
			p.Where_clause()
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

// ISelect_listContext is an interface to support dynamic dispatch.
type ISelect_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STAR() antlr.TerminalNode
	Column_list() IColumn_listContext

	// IsSelect_listContext differentiates from other interfaces.
	IsSelect_listContext()
}

type Select_listContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelect_listContext() *Select_listContext {
	var p = new(Select_listContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_select_list
	return p
}

func InitEmptySelect_listContext(p *Select_listContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_select_list
}

func (*Select_listContext) IsSelect_listContext() {}

func NewSelect_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Select_listContext {
	var p = new(Select_listContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_select_list

	return p
}

func (s *Select_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Select_listContext) STAR() antlr.TerminalNode {
	return s.GetToken(SQLParserSTAR, 0)
}

func (s *Select_listContext) Column_list() IColumn_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumn_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumn_listContext)
}

func (s *Select_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Select_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Select_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterSelect_list(s)
	}
}

func (s *Select_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitSelect_list(s)
	}
}

func (p *SQLParser) Select_list() (localctx ISelect_listContext) {
	localctx = NewSelect_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SQLParserRULE_select_list)
	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SQLParserSTAR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(51)
			p.Match(SQLParserSTAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SQLParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(52)
			p.Column_list()
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

// IColumn_listContext is an interface to support dynamic dispatch.
type IColumn_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllColumn_name() []IColumn_nameContext
	Column_name(i int) IColumn_nameContext

	// IsColumn_listContext differentiates from other interfaces.
	IsColumn_listContext()
}

type Column_listContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumn_listContext() *Column_listContext {
	var p = new(Column_listContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_column_list
	return p
}

func InitEmptyColumn_listContext(p *Column_listContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_column_list
}

func (*Column_listContext) IsColumn_listContext() {}

func NewColumn_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Column_listContext {
	var p = new(Column_listContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_column_list

	return p
}

func (s *Column_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Column_listContext) AllColumn_name() []IColumn_nameContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IColumn_nameContext); ok {
			len++
		}
	}

	tst := make([]IColumn_nameContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IColumn_nameContext); ok {
			tst[i] = t.(IColumn_nameContext)
			i++
		}
	}

	return tst
}

func (s *Column_listContext) Column_name(i int) IColumn_nameContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumn_nameContext); ok {
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

	return t.(IColumn_nameContext)
}

func (s *Column_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Column_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Column_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterColumn_list(s)
	}
}

func (s *Column_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitColumn_list(s)
	}
}

func (p *SQLParser) Column_list() (localctx IColumn_listContext) {
	localctx = NewColumn_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SQLParserRULE_column_list)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(55)
		p.Column_name()
	}
	p.SetState(60)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SQLParserT__0 {
		{
			p.SetState(56)
			p.Match(SQLParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(57)
			p.Column_name()
		}

		p.SetState(62)
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

// ITable_nameContext is an interface to support dynamic dispatch.
type ITable_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() IIdentifierContext

	// IsTable_nameContext differentiates from other interfaces.
	IsTable_nameContext()
}

type Table_nameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTable_nameContext() *Table_nameContext {
	var p = new(Table_nameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_table_name
	return p
}

func InitEmptyTable_nameContext(p *Table_nameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_table_name
}

func (*Table_nameContext) IsTable_nameContext() {}

func NewTable_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Table_nameContext {
	var p = new(Table_nameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_table_name

	return p
}

func (s *Table_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Table_nameContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Table_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Table_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Table_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterTable_name(s)
	}
}

func (s *Table_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitTable_name(s)
	}
}

func (p *SQLParser) Table_name() (localctx ITable_nameContext) {
	localctx = NewTable_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SQLParserRULE_table_name)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(63)
		p.Identifier()
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

// IColumn_nameContext is an interface to support dynamic dispatch.
type IColumn_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() IIdentifierContext

	// IsColumn_nameContext differentiates from other interfaces.
	IsColumn_nameContext()
}

type Column_nameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumn_nameContext() *Column_nameContext {
	var p = new(Column_nameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_column_name
	return p
}

func InitEmptyColumn_nameContext(p *Column_nameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_column_name
}

func (*Column_nameContext) IsColumn_nameContext() {}

func NewColumn_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Column_nameContext {
	var p = new(Column_nameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_column_name

	return p
}

func (s *Column_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Column_nameContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Column_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Column_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Column_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterColumn_name(s)
	}
}

func (s *Column_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitColumn_name(s)
	}
}

func (p *SQLParser) Column_name() (localctx IColumn_nameContext) {
	localctx = NewColumn_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SQLParserRULE_column_name)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(65)
		p.Identifier()
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

// IWhere_clauseContext is an interface to support dynamic dispatch.
type IWhere_clauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Condition() IConditionContext

	// IsWhere_clauseContext differentiates from other interfaces.
	IsWhere_clauseContext()
}

type Where_clauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhere_clauseContext() *Where_clauseContext {
	var p = new(Where_clauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_where_clause
	return p
}

func InitEmptyWhere_clauseContext(p *Where_clauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_where_clause
}

func (*Where_clauseContext) IsWhere_clauseContext() {}

func NewWhere_clauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Where_clauseContext {
	var p = new(Where_clauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_where_clause

	return p
}

func (s *Where_clauseContext) GetParser() antlr.Parser { return s.parser }

func (s *Where_clauseContext) Condition() IConditionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *Where_clauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Where_clauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Where_clauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterWhere_clause(s)
	}
}

func (s *Where_clauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitWhere_clause(s)
	}
}

func (p *SQLParser) Where_clause() (localctx IWhere_clauseContext) {
	localctx = NewWhere_clauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SQLParserRULE_where_clause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(67)
		p.condition(0)
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

// IConditionContext is an interface to support dynamic dispatch.
type IConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsConditionContext differentiates from other interfaces.
	IsConditionContext()
}

type ConditionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionContext() *ConditionContext {
	var p = new(ConditionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_condition
	return p
}

func InitEmptyConditionContext(p *ConditionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_condition
}

func (*ConditionContext) IsConditionContext() {}

func NewConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionContext {
	var p = new(ConditionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_condition

	return p
}

func (s *ConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionContext) CopyAll(ctx *ConditionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type InConditionContext struct {
	ConditionContext
}

func NewInConditionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InConditionContext {
	var p = new(InConditionContext)

	InitEmptyConditionContext(&p.ConditionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConditionContext))

	return p
}

func (s *InConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InConditionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *InConditionContext) IN() antlr.TerminalNode {
	return s.GetToken(SQLParserIN, 0)
}

func (s *InConditionContext) Value_list() IValue_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValue_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValue_listContext)
}

func (s *InConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterInCondition(s)
	}
}

func (s *InConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitInCondition(s)
	}
}

type OrConditionContext struct {
	ConditionContext
}

func NewOrConditionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrConditionContext {
	var p = new(OrConditionContext)

	InitEmptyConditionContext(&p.ConditionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConditionContext))

	return p
}

func (s *OrConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrConditionContext) AllCondition() []IConditionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConditionContext); ok {
			len++
		}
	}

	tst := make([]IConditionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConditionContext); ok {
			tst[i] = t.(IConditionContext)
			i++
		}
	}

	return tst
}

func (s *OrConditionContext) Condition(i int) IConditionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionContext); ok {
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

	return t.(IConditionContext)
}

func (s *OrConditionContext) OR() antlr.TerminalNode {
	return s.GetToken(SQLParserOR, 0)
}

func (s *OrConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterOrCondition(s)
	}
}

func (s *OrConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitOrCondition(s)
	}
}

type ParenthesisConditionContext struct {
	ConditionContext
}

func NewParenthesisConditionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenthesisConditionContext {
	var p = new(ParenthesisConditionContext)

	InitEmptyConditionContext(&p.ConditionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConditionContext))

	return p
}

func (s *ParenthesisConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenthesisConditionContext) Condition() IConditionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *ParenthesisConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterParenthesisCondition(s)
	}
}

func (s *ParenthesisConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitParenthesisCondition(s)
	}
}

type AndConditionContext struct {
	ConditionContext
}

func NewAndConditionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndConditionContext {
	var p = new(AndConditionContext)

	InitEmptyConditionContext(&p.ConditionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConditionContext))

	return p
}

func (s *AndConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndConditionContext) AllCondition() []IConditionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConditionContext); ok {
			len++
		}
	}

	tst := make([]IConditionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConditionContext); ok {
			tst[i] = t.(IConditionContext)
			i++
		}
	}

	return tst
}

func (s *AndConditionContext) Condition(i int) IConditionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionContext); ok {
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

	return t.(IConditionContext)
}

func (s *AndConditionContext) AND() antlr.TerminalNode {
	return s.GetToken(SQLParserAND, 0)
}

func (s *AndConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterAndCondition(s)
	}
}

func (s *AndConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitAndCondition(s)
	}
}

type IsNullConditionContext struct {
	ConditionContext
}

func NewIsNullConditionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IsNullConditionContext {
	var p = new(IsNullConditionContext)

	InitEmptyConditionContext(&p.ConditionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConditionContext))

	return p
}

func (s *IsNullConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IsNullConditionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IsNullConditionContext) IS() antlr.TerminalNode {
	return s.GetToken(SQLParserIS, 0)
}

func (s *IsNullConditionContext) NULL() antlr.TerminalNode {
	return s.GetToken(SQLParserNULL, 0)
}

func (s *IsNullConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterIsNullCondition(s)
	}
}

func (s *IsNullConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitIsNullCondition(s)
	}
}

type LikeConditionContext struct {
	ConditionContext
}

func NewLikeConditionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LikeConditionContext {
	var p = new(LikeConditionContext)

	InitEmptyConditionContext(&p.ConditionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConditionContext))

	return p
}

func (s *LikeConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LikeConditionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *LikeConditionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *LikeConditionContext) LIKE() antlr.TerminalNode {
	return s.GetToken(SQLParserLIKE, 0)
}

func (s *LikeConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterLikeCondition(s)
	}
}

func (s *LikeConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitLikeCondition(s)
	}
}

type IsNotNullConditionContext struct {
	ConditionContext
}

func NewIsNotNullConditionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IsNotNullConditionContext {
	var p = new(IsNotNullConditionContext)

	InitEmptyConditionContext(&p.ConditionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConditionContext))

	return p
}

func (s *IsNotNullConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IsNotNullConditionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IsNotNullConditionContext) IS() antlr.TerminalNode {
	return s.GetToken(SQLParserIS, 0)
}

func (s *IsNotNullConditionContext) NOT() antlr.TerminalNode {
	return s.GetToken(SQLParserNOT, 0)
}

func (s *IsNotNullConditionContext) NULL() antlr.TerminalNode {
	return s.GetToken(SQLParserNULL, 0)
}

func (s *IsNotNullConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterIsNotNullCondition(s)
	}
}

func (s *IsNotNullConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitIsNotNullCondition(s)
	}
}

type ComparisonConditionContext struct {
	ConditionContext
}

func NewComparisonConditionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ComparisonConditionContext {
	var p = new(ComparisonConditionContext)

	InitEmptyConditionContext(&p.ConditionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConditionContext))

	return p
}

func (s *ComparisonConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonConditionContext) Comparison_expression() IComparison_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparison_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComparison_expressionContext)
}

func (s *ComparisonConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterComparisonCondition(s)
	}
}

func (s *ComparisonConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitComparisonCondition(s)
	}
}

type NotInConditionContext struct {
	ConditionContext
}

func NewNotInConditionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotInConditionContext {
	var p = new(NotInConditionContext)

	InitEmptyConditionContext(&p.ConditionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConditionContext))

	return p
}

func (s *NotInConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotInConditionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *NotInConditionContext) NOT() antlr.TerminalNode {
	return s.GetToken(SQLParserNOT, 0)
}

func (s *NotInConditionContext) IN() antlr.TerminalNode {
	return s.GetToken(SQLParserIN, 0)
}

func (s *NotInConditionContext) Value_list() IValue_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValue_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValue_listContext)
}

func (s *NotInConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterNotInCondition(s)
	}
}

func (s *NotInConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitNotInCondition(s)
	}
}

type NotLikeConditionContext struct {
	ConditionContext
}

func NewNotLikeConditionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotLikeConditionContext {
	var p = new(NotLikeConditionContext)

	InitEmptyConditionContext(&p.ConditionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConditionContext))

	return p
}

func (s *NotLikeConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotLikeConditionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *NotLikeConditionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *NotLikeConditionContext) NOT() antlr.TerminalNode {
	return s.GetToken(SQLParserNOT, 0)
}

func (s *NotLikeConditionContext) LIKE() antlr.TerminalNode {
	return s.GetToken(SQLParserLIKE, 0)
}

func (s *NotLikeConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterNotLikeCondition(s)
	}
}

func (s *NotLikeConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitNotLikeCondition(s)
	}
}

func (p *SQLParser) Condition() (localctx IConditionContext) {
	return p.condition(0)
}

func (p *SQLParser) condition(_p int) (localctx IConditionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewConditionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IConditionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 14
	p.EnterRecursionRule(localctx, 14, SQLParserRULE_condition, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		localctx = NewParenthesisConditionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(70)
			p.Match(SQLParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(71)
			p.condition(0)
		}
		{
			p.SetState(72)
			p.Match(SQLParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewComparisonConditionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(74)
			p.Comparison_expression()
		}

	case 3:
		localctx = NewInConditionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(75)
			p.Expression()
		}
		{
			p.SetState(76)
			p.Match(SQLParserIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(77)
			p.Match(SQLParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(78)
			p.Value_list()
		}
		{
			p.SetState(79)
			p.Match(SQLParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewNotInConditionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(81)
			p.Expression()
		}
		{
			p.SetState(82)
			p.Match(SQLParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(83)
			p.Match(SQLParserIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(84)
			p.Match(SQLParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(85)
			p.Value_list()
		}
		{
			p.SetState(86)
			p.Match(SQLParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewLikeConditionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(88)
			p.Expression()
		}
		{
			p.SetState(89)
			p.Match(SQLParserLIKE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(90)
			p.Expression()
		}

	case 6:
		localctx = NewNotLikeConditionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(92)
			p.Expression()
		}
		{
			p.SetState(93)
			p.Match(SQLParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(94)
			p.Match(SQLParserLIKE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(95)
			p.Expression()
		}

	case 7:
		localctx = NewIsNullConditionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(97)
			p.Expression()
		}
		{
			p.SetState(98)
			p.Match(SQLParserIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(99)
			p.Match(SQLParserNULL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		localctx = NewIsNotNullConditionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(101)
			p.Expression()
		}
		{
			p.SetState(102)
			p.Match(SQLParserIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(103)
			p.Match(SQLParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(104)
			p.Match(SQLParserNULL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(116)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(114)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
			case 1:
				localctx = NewAndConditionContext(p, NewConditionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, SQLParserRULE_condition)
				p.SetState(108)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(109)
					p.Match(SQLParserAND)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(110)
					p.condition(11)
				}

			case 2:
				localctx = NewOrConditionContext(p, NewConditionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, SQLParserRULE_condition)
				p.SetState(111)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(112)
					p.Match(SQLParserOR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(113)
					p.condition(10)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(118)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
		if p.HasError() {
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
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IComparison_expressionContext is an interface to support dynamic dispatch.
type IComparison_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	Comparison_operator() IComparison_operatorContext

	// IsComparison_expressionContext differentiates from other interfaces.
	IsComparison_expressionContext()
}

type Comparison_expressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparison_expressionContext() *Comparison_expressionContext {
	var p = new(Comparison_expressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_comparison_expression
	return p
}

func InitEmptyComparison_expressionContext(p *Comparison_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_comparison_expression
}

func (*Comparison_expressionContext) IsComparison_expressionContext() {}

func NewComparison_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Comparison_expressionContext {
	var p = new(Comparison_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_comparison_expression

	return p
}

func (s *Comparison_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Comparison_expressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *Comparison_expressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *Comparison_expressionContext) Comparison_operator() IComparison_operatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparison_operatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComparison_operatorContext)
}

func (s *Comparison_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Comparison_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Comparison_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterComparison_expression(s)
	}
}

func (s *Comparison_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitComparison_expression(s)
	}
}

func (p *SQLParser) Comparison_expression() (localctx IComparison_expressionContext) {
	localctx = NewComparison_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, SQLParserRULE_comparison_expression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(119)
		p.Expression()
	}
	{
		p.SetState(120)
		p.Comparison_operator()
	}
	{
		p.SetState(121)
		p.Expression()
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

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Function_call() IFunction_callContext
	Column_name() IColumn_nameContext
	Value() IValueContext

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Function_call() IFunction_callContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunction_callContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunction_callContext)
}

func (s *ExpressionContext) Column_name() IColumn_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumn_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumn_nameContext)
}

func (s *ExpressionContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *SQLParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, SQLParserRULE_expression)
	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(123)
			p.Function_call()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(124)
			p.Column_name()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(125)
			p.Value()
		}

	case antlr.ATNInvalidAltNumber:
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

// IComparison_operatorContext is an interface to support dynamic dispatch.
type IComparison_operatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsComparison_operatorContext differentiates from other interfaces.
	IsComparison_operatorContext()
}

type Comparison_operatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparison_operatorContext() *Comparison_operatorContext {
	var p = new(Comparison_operatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_comparison_operator
	return p
}

func InitEmptyComparison_operatorContext(p *Comparison_operatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_comparison_operator
}

func (*Comparison_operatorContext) IsComparison_operatorContext() {}

func NewComparison_operatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Comparison_operatorContext {
	var p = new(Comparison_operatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_comparison_operator

	return p
}

func (s *Comparison_operatorContext) GetParser() antlr.Parser { return s.parser }
func (s *Comparison_operatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Comparison_operatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Comparison_operatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterComparison_operator(s)
	}
}

func (s *Comparison_operatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitComparison_operator(s)
	}
}

func (p *SQLParser) Comparison_operator() (localctx IComparison_operatorContext) {
	localctx = NewComparison_operatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, SQLParserRULE_comparison_operator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(128)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2032) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
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

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	String_literal() IString_literalContext
	Number_literal() INumber_literalContext
	Boolean_literal() IBoolean_literalContext
	NULL() antlr.TerminalNode

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_value
	return p
}

func InitEmptyValueContext(p *ValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_value
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) String_literal() IString_literalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IString_literalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IString_literalContext)
}

func (s *ValueContext) Number_literal() INumber_literalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumber_literalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumber_literalContext)
}

func (s *ValueContext) Boolean_literal() IBoolean_literalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolean_literalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoolean_literalContext)
}

func (s *ValueContext) NULL() antlr.TerminalNode {
	return s.GetToken(SQLParserNULL, 0)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitValue(s)
	}
}

func (p *SQLParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, SQLParserRULE_value)
	p.SetState(134)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SQLParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(130)
			p.String_literal()
		}

	case SQLParserNUMBER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(131)
			p.Number_literal()
		}

	case SQLParserTRUE, SQLParserFALSE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(132)
			p.Boolean_literal()
		}

	case SQLParserNULL:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(133)
			p.Match(SQLParserNULL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

// IValue_listContext is an interface to support dynamic dispatch.
type IValue_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext

	// IsValue_listContext differentiates from other interfaces.
	IsValue_listContext()
}

type Value_listContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValue_listContext() *Value_listContext {
	var p = new(Value_listContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_value_list
	return p
}

func InitEmptyValue_listContext(p *Value_listContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_value_list
}

func (*Value_listContext) IsValue_listContext() {}

func NewValue_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Value_listContext {
	var p = new(Value_listContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_value_list

	return p
}

func (s *Value_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Value_listContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *Value_listContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *Value_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Value_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Value_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterValue_list(s)
	}
}

func (s *Value_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitValue_list(s)
	}
}

func (p *SQLParser) Value_list() (localctx IValue_listContext) {
	localctx = NewValue_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, SQLParserRULE_value_list)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(136)
		p.Expression()
	}
	p.SetState(141)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SQLParserT__0 {
		{
			p.SetState(137)
			p.Match(SQLParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(138)
			p.Expression()
		}

		p.SetState(143)
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

// IString_literalContext is an interface to support dynamic dispatch.
type IString_literalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode

	// IsString_literalContext differentiates from other interfaces.
	IsString_literalContext()
}

type String_literalContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyString_literalContext() *String_literalContext {
	var p = new(String_literalContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_string_literal
	return p
}

func InitEmptyString_literalContext(p *String_literalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_string_literal
}

func (*String_literalContext) IsString_literalContext() {}

func NewString_literalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *String_literalContext {
	var p = new(String_literalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_string_literal

	return p
}

func (s *String_literalContext) GetParser() antlr.Parser { return s.parser }

func (s *String_literalContext) STRING() antlr.TerminalNode {
	return s.GetToken(SQLParserSTRING, 0)
}

func (s *String_literalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *String_literalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *String_literalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterString_literal(s)
	}
}

func (s *String_literalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitString_literal(s)
	}
}

func (p *SQLParser) String_literal() (localctx IString_literalContext) {
	localctx = NewString_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, SQLParserRULE_string_literal)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(144)
		p.Match(SQLParserSTRING)
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

// INumber_literalContext is an interface to support dynamic dispatch.
type INumber_literalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMBER() antlr.TerminalNode

	// IsNumber_literalContext differentiates from other interfaces.
	IsNumber_literalContext()
}

type Number_literalContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumber_literalContext() *Number_literalContext {
	var p = new(Number_literalContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_number_literal
	return p
}

func InitEmptyNumber_literalContext(p *Number_literalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_number_literal
}

func (*Number_literalContext) IsNumber_literalContext() {}

func NewNumber_literalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Number_literalContext {
	var p = new(Number_literalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_number_literal

	return p
}

func (s *Number_literalContext) GetParser() antlr.Parser { return s.parser }

func (s *Number_literalContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(SQLParserNUMBER, 0)
}

func (s *Number_literalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Number_literalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Number_literalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterNumber_literal(s)
	}
}

func (s *Number_literalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitNumber_literal(s)
	}
}

func (p *SQLParser) Number_literal() (localctx INumber_literalContext) {
	localctx = NewNumber_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, SQLParserRULE_number_literal)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(146)
		p.Match(SQLParserNUMBER)
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

// IBoolean_literalContext is an interface to support dynamic dispatch.
type IBoolean_literalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TRUE() antlr.TerminalNode
	FALSE() antlr.TerminalNode

	// IsBoolean_literalContext differentiates from other interfaces.
	IsBoolean_literalContext()
}

type Boolean_literalContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolean_literalContext() *Boolean_literalContext {
	var p = new(Boolean_literalContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_boolean_literal
	return p
}

func InitEmptyBoolean_literalContext(p *Boolean_literalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_boolean_literal
}

func (*Boolean_literalContext) IsBoolean_literalContext() {}

func NewBoolean_literalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Boolean_literalContext {
	var p = new(Boolean_literalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_boolean_literal

	return p
}

func (s *Boolean_literalContext) GetParser() antlr.Parser { return s.parser }

func (s *Boolean_literalContext) TRUE() antlr.TerminalNode {
	return s.GetToken(SQLParserTRUE, 0)
}

func (s *Boolean_literalContext) FALSE() antlr.TerminalNode {
	return s.GetToken(SQLParserFALSE, 0)
}

func (s *Boolean_literalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Boolean_literalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Boolean_literalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterBoolean_literal(s)
	}
}

func (s *Boolean_literalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitBoolean_literal(s)
	}
}

func (p *SQLParser) Boolean_literal() (localctx IBoolean_literalContext) {
	localctx = NewBoolean_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, SQLParserRULE_boolean_literal)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(148)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SQLParserTRUE || _la == SQLParserFALSE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
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

// IFunction_callContext is an interface to support dynamic dispatch.
type IFunction_callContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Function_name() IFunction_nameContext
	Function_args() IFunction_argsContext

	// IsFunction_callContext differentiates from other interfaces.
	IsFunction_callContext()
}

type Function_callContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_callContext() *Function_callContext {
	var p = new(Function_callContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_function_call
	return p
}

func InitEmptyFunction_callContext(p *Function_callContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_function_call
}

func (*Function_callContext) IsFunction_callContext() {}

func NewFunction_callContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_callContext {
	var p = new(Function_callContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_function_call

	return p
}

func (s *Function_callContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_callContext) Function_name() IFunction_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunction_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunction_nameContext)
}

func (s *Function_callContext) Function_args() IFunction_argsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunction_argsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunction_argsContext)
}

func (s *Function_callContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_callContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_callContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterFunction_call(s)
	}
}

func (s *Function_callContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitFunction_call(s)
	}
}

func (p *SQLParser) Function_call() (localctx IFunction_callContext) {
	localctx = NewFunction_callContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, SQLParserRULE_function_call)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(150)
		p.Function_name()
	}
	{
		p.SetState(151)
		p.Match(SQLParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(153)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&124780544) != 0 {
		{
			p.SetState(152)
			p.Function_args()
		}

	}
	{
		p.SetState(155)
		p.Match(SQLParserT__2)
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

// IFunction_nameContext is an interface to support dynamic dispatch.
type IFunction_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() IIdentifierContext

	// IsFunction_nameContext differentiates from other interfaces.
	IsFunction_nameContext()
}

type Function_nameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_nameContext() *Function_nameContext {
	var p = new(Function_nameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_function_name
	return p
}

func InitEmptyFunction_nameContext(p *Function_nameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_function_name
}

func (*Function_nameContext) IsFunction_nameContext() {}

func NewFunction_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_nameContext {
	var p = new(Function_nameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_function_name

	return p
}

func (s *Function_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_nameContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Function_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterFunction_name(s)
	}
}

func (s *Function_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitFunction_name(s)
	}
}

func (p *SQLParser) Function_name() (localctx IFunction_nameContext) {
	localctx = NewFunction_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, SQLParserRULE_function_name)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(157)
		p.Identifier()
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

// IFunction_argsContext is an interface to support dynamic dispatch.
type IFunction_argsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext

	// IsFunction_argsContext differentiates from other interfaces.
	IsFunction_argsContext()
}

type Function_argsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_argsContext() *Function_argsContext {
	var p = new(Function_argsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_function_args
	return p
}

func InitEmptyFunction_argsContext(p *Function_argsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_function_args
}

func (*Function_argsContext) IsFunction_argsContext() {}

func NewFunction_argsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_argsContext {
	var p = new(Function_argsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_function_args

	return p
}

func (s *Function_argsContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_argsContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *Function_argsContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *Function_argsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_argsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_argsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterFunction_args(s)
	}
}

func (s *Function_argsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitFunction_args(s)
	}
}

func (p *SQLParser) Function_args() (localctx IFunction_argsContext) {
	localctx = NewFunction_argsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, SQLParserRULE_function_args)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(159)
		p.Expression()
	}
	p.SetState(164)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SQLParserT__0 {
		{
			p.SetState(160)
			p.Match(SQLParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(161)
			p.Expression()
		}

		p.SetState(166)
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

// IIdentifierContext is an interface to support dynamic dispatch.
type IIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsIdentifierContext differentiates from other interfaces.
	IsIdentifierContext()
}

type IdentifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierContext() *IdentifierContext {
	var p = new(IdentifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_identifier
	return p
}

func InitEmptyIdentifierContext(p *IdentifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_identifier
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SQLParserIDENTIFIER, 0)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterIdentifier(s)
	}
}

func (s *IdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitIdentifier(s)
	}
}

func (p *SQLParser) Identifier() (localctx IIdentifierContext) {
	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, SQLParserRULE_identifier)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(167)
		p.Match(SQLParserIDENTIFIER)
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

func (p *SQLParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 7:
		var t *ConditionContext = nil
		if localctx != nil {
			t = localctx.(*ConditionContext)
		}
		return p.Condition_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *SQLParser) Condition_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 9)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
