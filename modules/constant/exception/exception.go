package exception

import "errors"

var (
	ErrRegisterFailed           = errors.New("register failed")
	ErrInvalidEmailOrPassword   = errors.New("email or password is valid")
	ErrAccountAlreadyRegistered = errors.New("account has already registered")
	ErrEmailAlreadyTaken        = errors.New("email has already taken")
	ErrJSONParseFailed          = errors.New("json parsing failed")
	ErrLoginFailed              = errors.New("login failed")
	ErrJWTUnexpectedAlgorithm   = errors.New("unexpected signing method")
	ErrInvalidJWT               = errors.New("invalid JWT")
)
