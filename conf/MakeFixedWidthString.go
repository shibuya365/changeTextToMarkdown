package conf

import (
	"bytes"

	"github.com/mattn/go-runewidth"
)

func MakeFixedWidthString(str string, length int) string {

	var buffer bytes.Buffer
	l := 0
	for _, c := range str {
		cl := runewidth.RuneWidth(c)
		if l+cl > length {
			break
		}
		buffer.WriteRune(c)
		l += cl
	}
	for i := 0; i < length-l; i++ {
		buffer.WriteRune(' ')
	}
	return buffer.String()
}
