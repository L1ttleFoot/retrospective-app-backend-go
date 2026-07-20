package core_service_jwt

type Jwt interface {
	AccessTocken() error
	RefreshToken() error
	ValidateToken() error
}
