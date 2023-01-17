package bridge

import (
	"fmt"

	c "github.com/Arsentau/golang-design-patterns/internal/pkg/bridge/computers"
	p "github.com/Arsentau/golang-design-patterns/internal/pkg/bridge/printers"
)

// Advantage
// Is important to notice that bridge design patter is useful to avoid complexity explosion

// Disadvantage
// The pattern itself incurs in a violation of the open close principle,
//     because if some new functionality is added to the parent struct a
//     new implementation of this logic needs to be done in the children structs.

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
