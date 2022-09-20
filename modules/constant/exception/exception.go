package exception

import "errors"

var (
	ErrRegisterFailed                 = errors.New("register failed")
	ErrInvalidEmailOrPassword         = errors.New("email or password is valid")
	ErrAccountAlreadyRegistered       = errors.New("account has already registered")
	ErrEmailAlreadyTaken              = errors.New("email has already taken")
	ErrJSONParseFailed                = errors.New("json parsing failed")
	ErrLoginFailed                    = errors.New("login failed")
	ErrJWTUnexpectedAlgorithm         = errors.New("unexpected signing method")
	ErrInvalidJWT                     = errors.New("invalid JWT")
	ErrDeclareRabbitMQExchangerFailed = errors.New("declare RabbitMQ exchanger failed")
	ErrDeclareRabbitMQQueueFailed     = errors.New("declare RabbitMQ queue failed")
	ErrPublishRabbitMQMessageFailed   = errors.New("publish RabbitMQ message failed")
	ErrInvalidData                    = errors.New("invalid data format")
)
