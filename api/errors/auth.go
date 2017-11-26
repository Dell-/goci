package errors

type Error struct {
	Message string `json:"message" description:"Error message"`
}

func LoginError() *Error {
	return &Error{Message: "Email or password is wrong."}
}

func LogoutError() *Error {
	return &Error{Message: "We could not process your request."}
}
