package computers

import (
	p "github.com/Arsentau/golang-design-patterns/internal/pkg/bridge/printers"
)

type Computer interface {
	Print()
	SetPrinter(p.Printer)
}
