package error_msg

type Error struct {
	Error string `json:"error"`
}

const ErrorEmailAlreadyInUse = "User with this email already exists"
const ErrorUserWithIdNotFound = "User not found, by the given id"
const ErrorUserWithEmailNotFound = "User not found, by the given email"
const ErrorSpaceWithIdNotFound = "Space not found, by the given id"
const ErrorSpaceAccessWithIdsNotFound = "SpaceAccess not found, by the given id's"

const ErrorUnauthorizedNoSessionCookie = "Unauthorized: no session cookie provided"
const ErrorUnauthorizedSessionExpired = "Unauthorized: session expired"
const ErrorUnauthorizedNoSpaceAccess = "Unauthorized: no space access"
const ErrorUnauthorizedNoSessionNotFound = "Unauthorized: session not found"
