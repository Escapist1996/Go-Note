package main

import "reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (r Reader) Read(b []byte) (int, error) {
	b[0] = 'A'
	return 1,nil
}


func main() {
	reader.Validate(MyReader{})
	
}
