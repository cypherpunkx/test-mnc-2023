package exception

import "errors"

const (
	StatusBadRequest          = "Bad Request"
	StatusInternalServerError = "Internal Server Error"
	StatusUnauthorized        = "Unauthorized"
	StatusSuccess             = "Success"
	StatusFailed              = "Failed"
	StatusLoginSuccess        = "Login Successful!"
	StatusLogoutSuccess       = "Logged out Successful!"
	StatusRegistrationSuccess = "Registration Successful!"
	StatusTransactionSuccess  = "Transaction Successful!"
)

var (
	ErrUserNameExist             = errors.New("customer already exists")
	ErrEmailExist                = errors.New("email already exists")
	ErrCustomerDoesntExist       = errors.New("customer doesn't exists")
	ErrFailedCreate              = errors.New("failed to create")
	ErrFailedCreateToken         = errors.New("failed to create token")
	ErrInvalidParseToken         = errors.New("invalid to parse token")
	ErrInvalidTokenSigningMethod = errors.New("invalid token signing method")
	ErrInvalidToken              = errors.New("invalid token")
	ErrTokenRequired             = errors.New("token required")
	ErrTokenNotProvided          = errors.New("token not provided")
	ErrFailedGeneratePassword    = errors.New("failed to generate password")
	ErrInvalidUsernamePassword   = errors.New("invalid username password")
	ErrNotEnoughBalance          = errors.New("not enough balance")
	ErrInvalidFriend             = errors.New("invalid friend")
)
