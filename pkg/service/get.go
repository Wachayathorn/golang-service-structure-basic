package service

import "github.com/wachayathorn/golang-service-structure-basic/pkg/model"

func (s *Service) Get() string {
	s.Logger.Infof("service : Get")

	if _, err := s.Post(model.Post{
		Message: "From Get",
	}); err != nil {
		return err.Error()
	}

	return "Hello World"
}
