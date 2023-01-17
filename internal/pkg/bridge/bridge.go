package bridge

import (
	"fmt"

	c "github.com/Arsentau/golang-design-patterns/internal/pkg/bridge/computers"
	p "github.com/Arsentau/golang-design-patterns/internal/pkg/bridge/printers"
)

func Bridge() {

	hpPrinter := &p.Hp{}
	epsonPrinter := &p.Epson{}

	macComputer := &c.Mac{}

	macComputer.SetPrinter(hpPrinter)
	macComputer.Print()
	fmt.Println()

	macComputer.SetPrinter(epsonPrinter)
	macComputer.Print()
	fmt.Println()

	winComputer := &c.Windows{}

	winComputer.SetPrinter(hpPrinter)
	winComputer.Print()
	fmt.Println()

	winComputer.SetPrinter(epsonPrinter)
	winComputer.Print()
	fmt.Println()
}
