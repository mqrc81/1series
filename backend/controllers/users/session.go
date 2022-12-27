package users

import (
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/mqrc81/zeries/domain"
)

const (
	SessionKey       = "session"
	SessionUserIdKey = "userId"
	SessionUserKey   = "user"
)

func AddUserToSession(ctx echo.Context, user domain.User) (err error) {
	currentSession, err := session.Get(SessionKey, ctx)
	if err == nil {
		currentSession.Values[SessionUserIdKey] = user.Id
		err = currentSession.Save(ctx.Request(), ctx.Response())
	}
	return err
}

func GetUserFromSession(ctx echo.Context) (user domain.User, err error) {
	currentSession, err := session.Get(SessionKey, ctx)
	if err == nil {
		user = currentSession.Values[SessionUserKey].(domain.User)
	}
	return user, err
}
