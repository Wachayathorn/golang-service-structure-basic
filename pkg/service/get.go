package service

func (s *Service) Get() string {
	s.Logger.Infof("service : Get")
	return "Hello World"
}
