package main

import (
	"fmt"

	servicesExample "github.com/Arsentau/golang-design-patterns/internal/pkg/bridge_services_example"
	deco "github.com/Arsentau/golang-design-patterns/internal/pkg/decorator_services_example"
	manual "github.com/Arsentau/golang-design-patterns/internal/pkg/dependency_injection/manual"
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
