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

func (bt *BTree) Insert(key, value []byte) error {
	panic("implement me")
}

func (bt *BTree) Get(key []byte) ([]byte, error) {
	panic("implement me")
}

func (bt *BTree) Delete(key []byte) error {
	panic("implement me")
}

func (bt *BTree) Range(start, end []byte) (index.Iterator, error) {
	panic("implement me")
}
func (bt *BTree) SaveTo(path string) error {
	return persist.Save(path, bt)
}

func (bt *BTree) LoadFrom(path string) error {
	return persist.Load(path, bt)
}
