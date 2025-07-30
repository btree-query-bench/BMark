package index

type Index interface {
	Insert(key, value []byte) error
	Get(key []byte) ([]byte, error)
	Delete(key []byte) error
	Range(start, end []byte) (Iterator, error)

	SaveTo(path string) error
	LoadFrom(path string) error
}
