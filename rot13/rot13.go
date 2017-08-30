package rot13

import (
	"io"
)

type rot13Reader struct {
	r io.Reader
}

func NewRot13Reader(r io.Reader) *rot13Reader {
	return &rot13Reader{r}
}

func rot13(b byte) byte {
	switch {
	case (b >= 'A' && b < 'N') || (b >= 'a' && b < 'n'):
		return b + 13
	case (b > 'M' && b <= 'Z') || (b > 'm' && b <= 'z'):
		return b - 13
	default:
		return b
	}
}

func (rot *rot13Reader) Read(bs []byte) (int, error) {
	n, err := rot.r.Read(bs)
	for i := 0; i < n; i++ {
		bs[i] = rot13(bs[i])
	}
	return n, err

}
