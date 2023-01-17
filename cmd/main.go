package main

import (
	"github.com/Arsentau/golang-design-patterns/internal/pkg/bridge"
	manual "github.com/Arsentau/golang-design-patterns/internal/pkg/dependency_injection/manual"
)

func main() {
	manual.DependencyInjection()
	bridge.Bridge()
}
