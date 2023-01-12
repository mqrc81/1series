package users

import (
	"errors"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/mqrc81/zeries/domain"
)

const (
	SessionKey       = "session"
	SessionUserIdKey = "userId"
	SessionUserKey   = "user"
)

func GetUserFromSession(ctx echo.Context) (domain.User, error) {
	user, ok := ctx.Get(SessionUserKey).(domain.User)
	if !ok {
		return domain.User{}, errors.New("no user in session")
	}
	return user, nil
}

func AddUserToSession(ctx echo.Context, user domain.User) error {
	sess, err := session.Get(SessionKey, ctx)
	if err == nil {
		sess.Values[SessionUserIdKey] = user.Id
		err = sess.Save(ctx.Request(), ctx.Response())
	}
	return err
}

func RemoveUserFromSession(ctx echo.Context) error {
	sess, err := session.Get(SessionKey, ctx)
	if err == nil {
		delete(sess.Values, SessionUserIdKey)
		delete(sess.Values, SessionUserKey)
		err = sess.Save(ctx.Request(), ctx.Response())
	}
	return err
}
