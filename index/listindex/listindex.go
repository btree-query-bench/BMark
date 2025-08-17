package listindex

import (
	"errors"
	"fmt"
	"slices"

	"github.com/btree-query-bench/bmark/index"
	"github.com/btree-query-bench/bmark/persist"
)

var _ index.Index = (*ListIndex)(nil)

type Data struct {
	Key int64
	Val int64
}

type ListIndex struct {
	Data []Data
}

func NewListIndex() *ListIndex {
	return &ListIndex{
		Data: make([]Data, 0),
	}
}

func (l *ListIndex) Insert(key, value int64) error {
	val, i := l.search(key)
	if val != 0 {
		l.Data[i] = Data{
			Key: key,
			Val: value,
		}
	} else {
		l.Data = append(l.Data, Data{
			Key: key,
			Val: value,
		})
	}
	return nil
}

func (l *ListIndex) Get(key int64) (int64, error) {
	val, _ := l.search(key)
	if val == 0 {
		return 0, errors.New("could not find value based on key")
	}

	return val, nil
}

func (l *ListIndex) Delete(key int64) error {
	val, i := l.search(key)

	if val == 0 {
		return errors.New("could not find key and therefore not delete it")
	}

	// Deletes key/val from slice
	l.Data = slices.Delete(l.Data, i, i)
	return nil
}

func (l *ListIndex) Range(start, end int64) (index.Iterator, error) {
	panic("implement me")
}

func (l *ListIndex) SaveTo(path string) error {
	return persist.Save(path, l.Data)
}

func (l *ListIndex) LoadFrom(path string) error {
	return persist.Load(path, &l.Data)
}

func (l *ListIndex) search(key int64) (int64, int) {
	for i, data := range l.Data {
		if data.Key == key {
			return data.Val, i
		}
	}
	return 0, -1
}

func (l *ListIndex) Print() {
	for _, data := range l.Data {
		fmt.Printf("Key: %d, Value: %d\n", data.Key, data.Val)
	}
}
