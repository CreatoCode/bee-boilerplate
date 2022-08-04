package user

import (
	"bee-boilerplate/public/net/http"
	wechatLogin "bee-boilerplate/service/user/login/wechat"
)

type controller struct{}

var ctr = controller{}

func RegisterRouter(router http.IRouterGroup) {
	router.Post("login", ctr.login, User{})
	router.Post("login/wechat", ctr.wecathLogin, make(map[string]string))
}

// 用户名密码登录
func (ctr controller) login(c *http.Context) (any, error) {
	g_logger.Debug("login")
	return nil, nil
}

// 微信登录
func (ctr controller) wecathLogin(c *http.Context) (any, error) {
	g_logger.Debug("login/wechat")
	if val, ok := c.Model.(map[string]interface{}); ok {
		if code, ok := val["code"].(string); ok {
			info, error := wechatLogin.SharedInstance().AppLogin(code)
			if nil != error {
				g_logger.Error("login/wechat:" + error.Error())
			}
			return info, error
		}
	}
	return nil, nil
}
