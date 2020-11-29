package service

import (
	"fmt"

	"github.com/tmvrus/service2/api"
	"github.com/tmvrus/service2/consumer"
)

type Handler struct {
}

func (h Handler) OfferCreate(event *api.OfferCreate) error {
	fmt.Printf("got OfferCreate: %v", event)
	return nil
}

func (h Handler) OfferDelete(event *api.OfferDelete) error {
	fmt.Printf("got OfferDelete: %v", event)
	return nil
}

func CreateOfferConsumer() {
	handler := Handler{}

	uc, err := consumer.NewConsumer("offer", "service1", []string{"host:port"}, handler)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	uc.Stop()
}
