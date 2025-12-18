package auth

type authService struct {
	userRepo UserRepo
	jwtToken string
}

func NewAuthService(userRepo UserRepo, jwtToken string) *authService {
	return &authService{
		userRepo: userRepo,
		jwtToken: jwtToken,
	}
}

func (a *authService) Register(email, username, password string) (string, error) {
	return "jwt-token", nil
}

func (a *authService) Login(email, password string) (string, error) {
	return "jwt-token", nil
}
