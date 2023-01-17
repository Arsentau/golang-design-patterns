package bridge

// Advantage
// Is important to notice that bridge design patter is useful to avoid complexity explosion

// Disadvantage
// The pattern itself incurs in a violation of the open close principle,
//     because if some new functionality is added to the parent struct a
//     new implementation of this logic needs to be done in the children structs.

type Computer interface {
	Print()
	SetPrinter(Printer)
}
