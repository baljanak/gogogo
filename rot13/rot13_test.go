package rot13_test

import (
	"io"
	"os"
	"strings"

	"github.com/baljanak/gogogo/rot13"
)

func ExampleRot13Reader() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13.NewRot13Reader(s)
	io.Copy(os.Stdout, r)
	// Output:
	// You cracked the code!
}
