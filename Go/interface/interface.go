package file

type File interface {
	Read() string
	Write(w string)
}
