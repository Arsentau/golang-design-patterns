package main

import (
	servicesExample "github.com/Arsentau/golang-design-patterns/internal/pkg/bridge_services_example"
	deco "github.com/Arsentau/golang-design-patterns/internal/pkg/decorator_services_example"
	manual "github.com/Arsentau/golang-design-patterns/internal/pkg/dependency_injection/manual"
)

func main() {
	manual.DependencyInjection()
	servicesExample.Bridge()
	deco.Decorator()
}
