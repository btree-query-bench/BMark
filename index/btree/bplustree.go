package btree

import (
	"github.com/btree-query-bench/bmark/index"
	"github.com/btree-query-bench/bmark/persist"
)

var _ index.Tree = (*BPlusTree)(nil)

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

func (bt *BPlusTree) Insert(key, value []byte) error {
	panic("implement me")
}

func (bt *BPlusTree) Get(key []byte) ([]byte, error) {
	panic("implement me")
}

func (bt *BPlusTree) Delete(key []byte) error {
	panic("implement me")
}

func (bt *BPlusTree) Range(start, end []byte) (index.Iterator, error) {
	panic("implement me")
}

func (bt *BPlusTree) SaveTo(path string) error {
	return persist.Save(path, bt)
}

func (bt *BPlusTree) LoadFrom(path string) error {
	return persist.Load(path, bt)
}
