// Code generated from QueryLang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package antlr // QueryLang
import "github.com/antlr4-go/antlr/v4"

type BaseQueryLangVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseQueryLangVisitor) VisitQuery(ctx *QueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQueryLangVisitor) VisitCreateBenchmark(ctx *CreateBenchmarkContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQueryLangVisitor) VisitSearch(ctx *SearchContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQueryLangVisitor) VisitModify(ctx *ModifyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQueryLangVisitor) VisitInsert(ctx *InsertContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQueryLangVisitor) VisitDelete(ctx *DeleteContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQueryLangVisitor) VisitStart(ctx *StartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQueryLangVisitor) VisitDataList(ctx *DataListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQueryLangVisitor) VisitDataEntry(ctx *DataEntryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQueryLangVisitor) VisitKeyList(ctx *KeyListContext) interface{} {
	return v.VisitChildren(ctx)
}
