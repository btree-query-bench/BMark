package listindex

import (
	"bytes"
	"github.com/btree-query-bench/bmark/index"
)

var _ index.Index = (*ListIndex)(nil)

type Data struct {
	key []byte
	val []byte
}

type ListIndex struct {
	data []Data
}

func NewListIndex() *ListIndex {
	return &ListIndex{
		data: make([]Data, 0),
	}
}

func (l *ListIndex) Insert(key, value []byte) error {
	_, i, exists := l.search(key)
	if exists {
		l.data[i] = Data{
			key: append([]byte(nil), key...),
			val: append([]byte(nil), value...),
		}
	} else {
		l.data = append(l.data, Data{
			key: append([]byte(nil), key...),
			val: append([]byte(nil), value...),
		})
	}
	return nil
}

func (l *ListIndex) Get(key []byte) ([]byte, error) {
	panic("implement me")
}

func (l *ListIndex) Delete(key []byte) error {
	panic("implement me")
}

func (l *ListIndex) Range(start, end []byte) (index.Iterator, error) {
	panic("implement me")
}

func (l *ListIndex) SaveTo(path string) error {
	panic("implement me")
}

func (l *ListIndex) LoadFrom(path string) error {
	panic("implement me")
}

func (l *ListIndex) search(key []byte) ([]byte, int, bool) {
	for i, data := range l.data {
		if bytes.Equal(data.key, key) {
			return data.val, i, true
		}
	}
	return nil, -1, false
}
