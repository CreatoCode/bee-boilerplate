package user

import (
	"bee-boilerplate/public/net/http"
)

type controller struct{}

var ctr = controller{}

func RegisterRouter(router http.IRouterGroup) {
	router.Post("login", ctr.login, User{})
}

func (ctr controller) login(c *http.Context) (any, error) {
	g_logger.Debug("login")
	return nil, nil
}
