package users

import (
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/mqrc81/zeries/domain"
	"github.com/mqrc81/zeries/logger"
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
	user, ok := ctx.Get(SessionUserKey).(domain.User)
	if !ok {
		logger.Warning("Sessions not implemented properly yet. Defaulting to user with Id=1 & Username=marc")
		// TODO ms
		// err = errors.New("no user in session")
		user = domain.User{
			Id:       1,
			Username: "marc",
		}
	}
	return user, err
}
