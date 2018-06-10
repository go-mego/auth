package auth

import (
	"crypto/subtle"
	"net/http"

	"github.com/go-mego/mego"
)

// New 可以建立一組最基本的帳號密碼以供驗證外來請求。
func New(username string, password string) mego.HandlerFunc {
	return func(c *mego.Context) {
		user, pass, ok := c.Request.BasicAuth()
		if !ok || !secureCompare(user, pass, username, password) {
			c.Header("WWW-Authenticate", `Basic`)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Next()
	}
}

// NewFunc 能夠自訂帳號密碼的驗證函式來驗證外來請求。
func NewFunc(handler func(username string, password string) bool) mego.HandlerFunc {
	return func(c *mego.Context) {
		user, pass, ok := c.Request.BasicAuth()
		if !ok || !handler(user, pass) {
			c.Header("WWW-Authenticate", `Basic`)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Next()
	}
}

// Accounts 是多組帳號與密碼，供認證用。
type Accounts map[string]string

// NewAccounts 能使用多組帳號與密碼驗證外來請求，只要符合其中一組即為通過。
func NewAccounts(accounts *Accounts) mego.HandlerFunc {
	return func(c *mego.Context) {
		user, pass, ok := c.Request.BasicAuth()
		passed := func() bool {
			var ok bool
			for k, v := range *accounts {
				if secureCompare(user, pass, k, v) {
					ok = true
					break
				}
			}
			return ok
		}
		if !ok || !passed() {
			c.Header("WWW-Authenticate", `Basic`)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Next()
	}
}

// secureCompare 會以 `ConstantTimeCompare` 來比對資料以避免時間差攻擊。
func secureCompare(inputUsername, inputPassword, username, password string) bool {
	return subtle.ConstantTimeCompare([]byte(username), []byte(inputUsername)) == 1 && subtle.ConstantTimeCompare([]byte(password), []byte(inputPassword)) == 1
}
