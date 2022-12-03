package service

type ServiceInterface interface {
	Get() string
}

type Service struct{}

func Init() *Service {
	return &Service{}
}
