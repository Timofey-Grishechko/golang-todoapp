package users_transport_http

type UsersHTTPHandler struct {
	userService UserService
}

type UserService interface {
}

func NewUserHTTPHandler(
	userService UserService,
) *UsersHTTPHandler {
	return &UsersHTTPHandler{
		userService: userService,
	}
}
