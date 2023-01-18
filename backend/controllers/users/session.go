package users

import (
	"errors"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/mqrc81/zeries/domain"
	"github.com/mqrc81/zeries/env"
	"net/http"
)

const (
	SessionKey                   = "session"
	SessionUserIdKey             = "userId"
	ContextUserKey               = "user"
	RememberLoginTokenCookieName = "remember_login"
)

func GetAuthenticatedUser(ctx echo.Context) (domain.User, error) {
	user, ok := ctx.Get(ContextUserKey).(domain.User)
	if !ok {
		return domain.User{}, errors.New("no user in session")
	}
	return user, nil
}

func IsAdmin(user domain.User) bool {
	for _, admin := range env.Config().Admins {
		if user.Username == admin {
			return true
		}
	}
	return false
}

func (c *usersController) authenticateUser(ctx echo.Context, user domain.User) error {
	sess, err := session.Get(SessionKey, ctx)
	if err == nil {
		sess.Values[SessionUserIdKey] = user.Id
		err = sess.Save(ctx.Request(), ctx.Response())
	}
	ctx.Set(ContextUserKey, user)

	rememberLoginToken := domain.CreateToken(domain.RememberLogin, user.Id)
	if err = c.tokensRepository.SaveOrReplace(rememberLoginToken); err != nil {
		return err
	}

	ctx.SetCookie(&http.Cookie{
		Name:     RememberLoginTokenCookieName,
		Value:    rememberLoginToken.Id,
		Expires:  rememberLoginToken.ExpiresAt,
		SameSite: http.SameSiteDefaultMode,
	})

	return err
}

func (c *usersController) unauthenticateUser(ctx echo.Context) error {
	if user, err := GetAuthenticatedUser(ctx); err == nil {
		return c.tokensRepository.DeleteByUserIdAndPurpose(user.Id, domain.RememberLogin)
	}
	sess, err := session.Get(SessionKey, ctx)
	if err == nil {
		delete(sess.Values, SessionUserIdKey)
		err = sess.Save(ctx.Request(), ctx.Response())
	}
	ctx.Set(ContextUserKey, nil)
	return err
}
