package api

func NewAuthService(authRepo AuthRepository) AuthService {
	return &authentication_service{
		authRepo: authRepo,
	}
}

func (a *authentication_service) LogIn(email, password string) (err error, access_token, refresh_token string) {
	return
}

func (a *authentication_service) ValidateCredentials(email, password string) (validity bool, err error) {
	return
}

func (a *authentication_service) GenerateAccessToken(email string, expiration int64) (signed_access_token string, err error) {
	return
}

func (a *authentication_service) GenerateRefreshToken(email string, customKey string) (signed_refresh_token string, err error) {
	return
}

func (a *authentication_service) ValidateAccessToken(access_token string) (status int) {
	return
}

func (a *authentication_service) ValidateRefreshToken(refresh_token, custom_key string) (validity bool) {
	return
}
