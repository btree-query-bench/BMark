// Code generated from QueryLang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package antlr // QueryLang
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by QueryLangParser.
type QueryLangVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by QueryLangParser#command.
	VisitCommand(ctx *CommandContext) interface{}

	// Visit a parse tree produced by QueryLangParser#createBenchmark.
	VisitCreateBenchmark(ctx *CreateBenchmarkContext) interface{}

	// Visit a parse tree produced by QueryLangParser#insertAll.
	VisitInsertAll(ctx *InsertAllContext) interface{}

	// Visit a parse tree produced by QueryLangParser#insert.
	VisitInsert(ctx *InsertContext) interface{}

	// Visit a parse tree produced by QueryLangParser#lookup.
	VisitLookup(ctx *LookupContext) interface{}

	// Visit a parse tree produced by QueryLangParser#delete.
	VisitDelete(ctx *DeleteContext) interface{}

	// Visit a parse tree produced by QueryLangParser#compareLookup.
	VisitCompareLookup(ctx *CompareLookupContext) interface{}

	// Visit a parse tree produced by QueryLangParser#dataList.
	VisitDataList(ctx *DataListContext) interface{}

	// Visit a parse tree produced by QueryLangParser#dataEntry.
	VisitDataEntry(ctx *DataEntryContext) interface{}

	// Visit a parse tree produced by QueryLangParser#keyList.
	VisitKeyList(ctx *KeyListContext) interface{}
}
