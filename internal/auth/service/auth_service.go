package service

// AuthService defines the contract for authentication business logic.
type AuthService interface {
	Signup(name, email, phone, password, role string) error
	Login(email, password string) (token string, err error)
}
