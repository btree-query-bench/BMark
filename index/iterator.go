package index

type Iterator interface {
	Next() bool
	Key() int64
	Value() int64
	Error() error
	Close() error
}
