package Service

import "net/http"

type BookService struct {
	username string
	password string
	cookie   http.CookieJar
}

func NewBookService(username, password string) *BookService {
	return &BookService{
		username: username,
		password: password,
	}
}

func (i *BookService) Login() bool {
	return true
}
