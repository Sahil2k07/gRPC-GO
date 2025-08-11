package error

type NotFoundError struct {
	Msg string
}

func (e *NotFoundError) Error() string {
	return e.Msg
}

type ValidationError struct {
	Msg string
}

func (e *ValidationError) Error() string {
	return e.Msg
}

type UnauthorizedError struct {
	Msg string
}

func (e *UnauthorizedError) Error() string {
	return e.Msg
}

type ForbiddenError struct {
	Msg string
}

func (e *ForbiddenError) Error() string {
	return e.Msg
}

type AlreadyExistsError struct {
	Msg string
}

func (e *AlreadyExistsError) Error() string {
	return e.Msg
}

type InternalError struct {
	Msg string
}

func (e *InternalError) Error() string {
	return e.Msg
}

// Optionally helper constructors

func NewNotFound(msg string) error {
	return &NotFoundError{Msg: msg}
}

func NewValidation(msg string) error {
	return &ValidationError{Msg: msg}
}

func NewUnauthorized(msg string) error {
	return &UnauthorizedError{Msg: msg}
}

func NewForbidden(msg string) error {
	return &ForbiddenError{Msg: msg}
}

func NewAlreadyExists(msg string) error {
	return &AlreadyExistsError{Msg: msg}
}

func NewInternalError(msg string) error {
	return &InternalError{Msg: msg}
}
