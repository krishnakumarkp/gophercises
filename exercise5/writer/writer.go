package writer

type Writer interface {
	Write(map[string]struct{}) error
}
