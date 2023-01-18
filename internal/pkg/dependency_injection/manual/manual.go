package manual

import "fmt"

type MyLogger struct {
	message string
}
type MyRepository struct {
	message string
}
type MyMessageBroker struct {
	message string
}

type Service struct {
	logger     MyLogger
	repository MyRepository
	broker     MyMessageBroker
}

func newService(logger MyLogger, repository MyRepository, broker MyMessageBroker) *Service {
	return &Service{
		logger:     logger,
		repository: repository,
		broker:     broker,
	}
}

func DependencyInjection() {
	logger := MyLogger{message: "Logging"}
	broker := MyMessageBroker{message: "Hi from Kafka"}
	repository := MyRepository{"Connected to repository"}

	service := newService(logger, repository, broker)

	fmt.Println(service.logger.message)
	fmt.Println(service.repository.message)
	fmt.Println(service.broker.message)

}
