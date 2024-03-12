package decorator

import (
	"fmt"
	"testing"
)

type Printer interface {
	Print() string
}

type SimplePrinter struct{}

func (sp *SimplePrinter) Print() string {
	return "Hello, world!"
}

func BoldDecorator(p Printer) Printer {
	return PrinterFunc(func() string {
		return "<b>" + p.Print() + "</b>"
	})
}

func ColorDecorator(p Printer) Printer {
	return PrinterFunc(func() string {
		return "<span style=\"color:red\">" + p.Print() + "</span>"
	})
}

func EmojiDecorator(p Printer) Printer {
	return PrinterFunc(func() string {
		return p.Print() + "😄"
	})
}

type PrinterFunc func() string

func (pf PrinterFunc) Print() string {
	return pf()
}

func TestPrinter(t *testing.T) {
	simplePrinter := &SimplePrinter{}
	fmt.Println(simplePrinter.Print())

	BoldPrinter := BoldDecorator(simplePrinter)
	fmt.Println(BoldPrinter.Print())

	colorBoldPrinter := ColorDecorator(BoldDecorator(simplePrinter))
	fmt.Println(colorBoldPrinter.Print())

	emojiColorBoldPrinter := EmojiDecorator(ColorDecorator(BoldDecorator(simplePrinter)))
	fmt.Println(emojiColorBoldPrinter.Print())
}
