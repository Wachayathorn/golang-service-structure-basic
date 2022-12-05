package service

import "github.com/wachayathorn/golang-service-structure-basic/pkg/model"

func (s *Service) Post(req model.Post) (string, error) {
	if err := s.Validator.Struct(req); err != nil {
		return "", err
	}
	return req.Message, nil
}
