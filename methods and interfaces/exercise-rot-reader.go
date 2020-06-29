package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(b []byte) (int, error) {
	
	chunk := make([]byte,1)
	
	n, err := rot.r.Read(chunk)
	
	if err == nil && n > 0 {
		c := chunk[0]
		if (c >= 'a' && c <= 'z') { 
			b[0] = (c-'a'+13) % ('z'-'a'+1)	+ 'a'
		} else if (c >= 'A' && c <= 'Z') {
			b[0] = (c-'A'+13) % ('Z'-'A'+1)	 + 'A'
		} else if (c >= '0' && c <= '9') {
			b[0] = (c-'0'+13) % ('9'-'0'+1)	+ '0'
		} else {
			b[0] = c	
		}
	}
	
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
