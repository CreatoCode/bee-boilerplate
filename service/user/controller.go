package user

import (
	http "bee-boilerplate/public/net/http/server"
	wechatLogin "bee-boilerplate/service/user/login/wechat"

	"github.com/google/uuid"
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

// 微信登录，并注册到本地
func (ctr controller) wecathLogin(c *http.Context) (any, error) {
	g_logger.Debug("login/wechat")
	if val, ok := c.Model.(map[string]interface{}); ok {
		if code, ok := val["code"].(string); ok {
			info, error := wechatLogin.SharedInstance().GetJsCode2Session(code)
			if nil != error {
				g_logger.Error("login/wechat:" + error.Error())
				return nil, error
			} else {
				u := User{UUID: uuid.NewString(), ThirdId: info.Openid, Name: info.Openid}
				u.CreateIfNotExist()
			}
		}
	}
	return nil, nil
}
