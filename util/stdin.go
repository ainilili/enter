package util

import (
	"bufio"
	"os"
)

var (
	in = bufio.NewReader(os.Stdin)
)

func Await() {
	_, _, _ = in.ReadRune()
}
