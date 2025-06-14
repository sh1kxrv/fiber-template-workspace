package rpc

type authService struct {
}

func NewAuthService() *authService {
	return &authService{}
}

func (s *authService) SetToken(token string, tgId string) (bool, error) {
	return false, nil
}

func (s *authService) GetTokenExists(token string) (bool, error) {
	return false, nil
}

func (s *authService) DeclineToken(token string) (bool, error) {
	return true, nil
}
