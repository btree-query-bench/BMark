package btree

import (
	"github.com/btree-query-bench/bmark/index"
	"github.com/btree-query-bench/bmark/persist"
)

var _ index.Index = (*BPlusTree)(nil)

type BPlusTree struct {
	p    int
	root *BPlusTreeNode
}

type BPlusTreeNode struct {
}

func NewBPlusTree(p int) *BPlusTree {
	if p < 2 {
		p = 2
	}
	return &BPlusTree{
		p:    p,
		root: &BPlusTreeNode{},
	}
}

func (bt *BPlusTree) Insert(key, value int64) error {
	panic("implement me")
}

func (bt *BPlusTree) Get(key int64) (int64, error) {
	panic("implement me")
}

func (bt *BPlusTree) Delete(key int64) error {
	panic("implement me")
}

func (bt *BPlusTree) Range(start, end int64) (index.Iterator, error) {
	panic("implement me")
}

func (bt *BPlusTree) SaveTo(path string) error {
	return persist.Save(path, bt)
}

func (bt *BPlusTree) LoadFrom(path string) error {
	return persist.Load(path, bt)
}
