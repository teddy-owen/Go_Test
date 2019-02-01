package stringutil

import(
	"github.com/common-nighthawk/go-figure"
)

func Ascii(s string) {
	myFigure := figure.NewFigure(s, "", true)
	myFigure.Print()
	return
}