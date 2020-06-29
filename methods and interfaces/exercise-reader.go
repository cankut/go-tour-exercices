package main

import "golang.org/x/tour/reader"

type MyReader struct {
	i int
}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func (r MyReader) Read(b []byte) (int, error) {
	b[r.i] = 'A'
	r.i += 1
	return 1, nil
}

func main() {
	reader.Validate(MyReader{})
}
