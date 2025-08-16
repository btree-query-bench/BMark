package query

import (
	"fmt"
	parser "github.com/btree-query-bench/bmark/antlr"
)

type Executor struct {
	parser.BaseQueryLangVisitor
}

func (e *Executor) VisitQuery(ctx *parser.QueryContext) interface{} {
	switch {
	case ctx.CreateBenchmark() != nil:
		return e.VisitCreateBenchmark(ctx.CreateBenchmark())
	case ctx.Search() != nil:
		return e.VisitSearch(ctx.Search())
	case ctx.Modify() != nil:
		return e.VisitModify(ctx.Modify())
	case ctx.Start_() != nil:
		return e.VisitStart(ctx.Start_())
	}
	return nil
}

func (e *Executor) VisitCreateBenchmark(ctx parser.ICreateBenchmarkContext) interface{} {
	fmt.Println("VisitCreateBenchmark")
	return nil
}

func (e *Executor) VisitSearch(ctx parser.ISearchContext) interface{} {
	fmt.Println("VisitSearch")
	return nil
}

func (e *Executor) VisitModify(ctx parser.IModifyContext) interface{} {
	switch {
	case ctx.Insert() != nil:
		return e.VisitInsert(ctx.Insert())
	case ctx.Delete_() != nil:
		return e.VisitDelete(ctx.Delete_())
	}
	return nil
}

func (e *Executor) VisitInsert(ctx parser.IInsertContext) interface{} {
	fmt.Println("VisitInsert")
	return nil
}

func (e *Executor) VisitDelete(ctx parser.IDeleteContext) interface{} {
	fmt.Println("VisitDelete")
	return nil
}

func (e *Executor) VisitStart(ctx parser.IStartContext) interface{} {
	fmt.Println("VisitStart")
	return nil
}
