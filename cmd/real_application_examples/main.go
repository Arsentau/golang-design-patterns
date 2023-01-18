package main

import (
	servicesExample "github.com/Arsentau/golang-design-patterns/internal/pkg/bridge_services_example"
	deco "github.com/Arsentau/golang-design-patterns/internal/pkg/decorator"
)

func main() {
	servicesExample.Bridge()
	deco.Decorator()
}
