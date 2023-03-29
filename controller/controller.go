package controller

import "challenge07/service"

type Controller struct {
	service service.ServiceInterface
}

func NewController(service service.ServiceInterface) *Controller {
	return &Controller{service: service}
}
