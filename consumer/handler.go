package consumer

import "github.com/tmvrus/service1/api"

var _ UserHandler = DropHandler{}

type UserHandler interface {
	UserCreate(event *api.UserCreate) error
	UserDelete(event *api.UserDelete) error
	UserSuspend(event *api.UserSuspend) error
}

type DropHandler struct{}

func (h DropHandler) UserCreate(event *api.UserCreate) error {
	log.Debugf("got DropHandler UserCreate call with %v", event)
	return nil
}

func (h DropHandler) UserDelete(event *api.UserDelete) error {
	log.Debugf("got DropHandler UserDelete call with %v", event)
	return nil
}

func (h DropHandler) UserSuspend(event *api.UserSuspend) error {
	log.Debugf("got DropHandler UserDelete call with %v", event)
	return nil
}
