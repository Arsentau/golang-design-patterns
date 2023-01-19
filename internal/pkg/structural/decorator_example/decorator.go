package decorator

import (
	"errors"
	"log"
)

func decorator(f func(s string) (string, error)) func(string) {
	return func(s string) {
		log.Println("Checking service status...")
		msg, err := f(s)
		if err != nil {
			log.Println(err.Error())
			return
		}
		log.Println("After:", msg)
	}
}

func CheckBadStatus(s string) (string, error) {
	log.Println(s)
	return "", errors.New("Could not connect to the server")
}

func CheckOkStatus(s string) (string, error) {
	log.Println(s)
	return "Online", nil
}

func Decorator() {

	// without decorator
	CheckBadStatus("NS1")

	// with decorator
	decorator(CheckBadStatus)("Route83")
	decorator(CheckOkStatus)("Cloudflare")

}

// This is a second way to do it, but I consider that this way introduces more complexity and is more difficult to maintain.

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
func Decorator2() {

	ns1Service := &Ns1Status{}
	logger := &Logger{
		service: ns1Service,
	}
	logger.LogStatus()
}
