package computers

import (
	"fmt"

	p "github.com/Arsentau/golang-design-patterns/internal/pkg/bridge/printers"
)

type Windows struct {
	printer p.Printer
}

func (w *Windows) Print() {
	fmt.Println("Print request for windows")
	w.printer.PrintFile()
}

func (w *Windows) SetPrinter(p p.Printer) {
	w.printer = p
}
