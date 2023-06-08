package error_msg

type Error struct {
	Error string `json:"error"`
}

const ErrorEmailAlreadyInUse = "User with this email already exists"
const ErrorUserWithIdNotFound = "User not found, by the given id"
