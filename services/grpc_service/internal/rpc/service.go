package rpc

import "shared/cache/memcache"

type authService struct {
	cache *memcache.MemoryCache
}

func NewAuthService(cache *memcache.MemoryCache) *authService {
	return &authService{cache: cache}
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
