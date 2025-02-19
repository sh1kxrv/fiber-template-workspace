package rpc

// type authServer struct {
// 	authService *authService
// 	pb.UnimplementedAuthenticationServiceServer
// }

// func (s *authServer) SetToken(ctx context.Context, req *pb.SetTokenDataRequest) (*pb.SetTokenDataResponse, error) {
// 	result, err := s.authService.SetToken(req.Token, req.TelegramUserId)
// 	return &pb.SetTokenDataResponse{
// 		Status: result,
// 	}, err
// }

// func (s *authServer) GetTokenExists(ctx context.Context, req *pb.GetTokenExistsRequest) (*pb.GetTokenExistsResponse, error) {
// 	result, err := s.authService.GetTokenExists(req.Token)
// 	return &pb.GetTokenExistsResponse{
// 		Status: result,
// 	}, err
// }

// func (s *authServer) DeclineToken(ctx context.Context, req *pb.DeclineTokenRequest) (*pb.DeclineTokenResponse, error) {
// 	result, err := s.authService.DeclineToken(req.Token)
// 	return &pb.DeclineTokenResponse{
// 		Status: result,
// 	}, err
// }
