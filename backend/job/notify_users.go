package job

import (
	"github.com/mqrc81/zeries/repository"
	. "github.com/mqrc81/zeries/util"
)

func (e notifyUsersJob) execute() error {
	LogInfo("Yet to implement notify-users job")

	// TODO

	return nil
}

type notifyUsersJob struct {
	userRepository repository.UserRepository
	// mailgunClient mailgun.Client
}
