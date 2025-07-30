package index

type Iterator interface {
	Next() bool
	Key() []byte
	Value() []byte
	Error() error
	Close() error
}
