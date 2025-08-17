package index

type Index interface {
	Insert(key, value int64) error
	Get(key int64) (int64, error)
	Delete(key int64) error
	Range(start, end int64) (Iterator, error)

	SaveTo(path string) error
	LoadFrom(path string) error
}
