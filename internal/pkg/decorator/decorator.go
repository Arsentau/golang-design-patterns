package decorator

import (
	"log"
)

type IService interface {
	CheckStatus() string
}

type Ns1Status struct{}

func (l *Ns1Status) CheckStatus() string {
	statusMessage := "Checking NS1 Server status"
	return statusMessage
}

// Decorator
type Logger struct {
	service IService
}

func (l *Logger) LogStatus() {
	statusMessage := l.service.CheckStatus()
	log.Println("From Decorator:", statusMessage)
}

// Decorator implementation
func Decorator() {
	ns1Service := &Ns1Status{}
	logger := &Logger{
		service: ns1Service,
	}
	logger.LogStatus()
}
