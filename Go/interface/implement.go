package file

type File1 struct {
	Content string
}

func (file File1) Read() string {
	return file.Content
}

func (file *File1) Write(w string) {
	file.Content = w
}
