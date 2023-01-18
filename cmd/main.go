package main

import (
	"fmt"

	manual "github.com/Arsentau/golang-design-patterns/internal/pkg/dependency_injection/manual"
	servicesExample "github.com/Arsentau/golang-design-patterns/internal/pkg/structural/bridge_example"
	deco "github.com/Arsentau/golang-design-patterns/internal/pkg/structural/decorator_example"
)

func main() {
	separator := "\n====================\n"

	fmt.Println("DEPENDENCY INJECTION", separator)
	manual.DependencyInjection()
	fmt.Println(separator)

	fmt.Println("BRIDGE PATTERN", separator)
	servicesExample.Bridge()
	fmt.Println(separator)

	fmt.Println("DECORATOR PATTERN", separator)
	deco.Decorator()
	fmt.Println(separator)
}
