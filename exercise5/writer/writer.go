package writer

type Writable interface {
	Write(Writer) error
	GetData() []string
}

type Writer interface {
	Write(Writable) error
}
