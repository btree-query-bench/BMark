package btree

import (
	"github.com/btree-query-bench/bmark/index"
	"github.com/btree-query-bench/bmark/persist"
)

var _ index.Index = (*BTree)(nil)

type BTree struct {
	t    int
	root *BTreeNode
}

type BTreeNode struct {
}

func NewBTree(t int) *BTree {
	if t < 2 {
		t = 2
	}
	return &BTree{
		t:    t,
		root: &BTreeNode{},
	}
}

func (bt *BTree) Insert(key, value int64) error {
	panic("implement me")
}

func (bt *BTree) Get(key int64) (int64, error) {
	panic("implement me")
}

func (bt *BTree) Delete(key int64) error {
	panic("implement me")
}

func (bt *BTree) Range(start, end int64) (index.Iterator, error) {
	panic("implement me")
}
func (bt *BTree) SaveTo(path string) error {
	return persist.Save(path, bt)
}

func (bt *BTree) LoadFrom(path string) error {
	return persist.Load(path, bt)
}
