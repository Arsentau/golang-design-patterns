package main

import (
	"github.com/Arsentau/golang-design-patterns/internal/pkg/bridge"
	deco "github.com/Arsentau/golang-design-patterns/internal/pkg/decorator"
	manual "github.com/Arsentau/golang-design-patterns/internal/pkg/dependency_injection/manual"
)

func main() {
	manual.DependencyInjection()
	bridge.Bridge()
	deco.Decorator()
}
