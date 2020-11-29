package consumer

import (
	"github.com/golang/protobuf/proto"
	"github.com/nsqio/go-nsq"
	"github.com/sirupsen/logrus"

	"github.com/tmvrus/service1/api"
)

var log = logrus.WithField("package", "consumer")

type Consumer struct {
	c                       *nsq.Consumer
	requeueUnmarshalError   bool
	suppressValidationError bool
	logUnknownEventError    bool
	userHandler             UserHandler
}

func NewConsumer(topic, channel string, nsqd []string, h UserHandler) (*Consumer, error) {
	c, err := nsq.NewConsumer(topic, channel, nil)
	if err != nil {
		return nil, err
	}

	if err := c.ConnectToNSQDs(nsqd); err != nil {
		return nil, err
	}

	userConsumer := &Consumer{
		userHandler:             h,
		suppressValidationError: true,
		logUnknownEventError:    true,
	}

	c.AddHandler(userConsumer)
	return userConsumer, nil

}

func (h Consumer) Stop() {
	h.c.Stop()
}

func (h Consumer) HandleMessage(message *nsq.Message) error {
	var userEvent api.UserEvent

	if err := proto.Unmarshal(message.Body, &userEvent); err != nil {
		log.WithError(err).Error("failed to unmarshal user event")
		if !h.requeueUnmarshalError {
			return err
		}

		return nil
	}

	if err := userEvent.Validate(); err != nil {
		log.WithError(err).Error("failed to validate user event")
		if !h.suppressValidationError {
			return err
		}

		return nil
	}

	if userEvent.UserCreate != nil {
		return h.userHandler.UserCreate(userEvent.UserCreate)
	}

	if userEvent.UserDelete != nil {
		return h.userHandler.UserDelete(userEvent.UserDelete)
	}

	if userEvent.UserSuspend != nil {
		return h.userHandler.UserSuspend(userEvent.UserSuspend)
	}

	l := log.Debug
	if !h.logUnknownEventError {
		l = log.Error
	}

	l("got unknown user event: %v", userEvent)
	return nil
}
