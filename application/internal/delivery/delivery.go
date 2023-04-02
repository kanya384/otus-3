package delivery

import (
	"fmt"
	"otus/internal/service"

	"github.com/gin-gonic/gin"
)

type Delivery struct {
	service service.Service
	router  *gin.Engine

	options Options
}

type Options struct{}

func New(service service.Service, options Options) (*Delivery, error) {
	var d = &Delivery{
		service: service,
	}

	d.SetOptions(options)

	d.router = d.initRouter()

	return d, nil
}

func (d *Delivery) SetOptions(options Options) {
	if d.options != options {
		d.options = options
	}
}

func (d *Delivery) Run(port int) error {
	return d.router.Run(fmt.Sprintf(":%d", port))
}
