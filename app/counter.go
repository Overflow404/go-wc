package app

type Counter interface {
	Count(filename string, fileWrapper FileWrapper) (int64, error)
}
