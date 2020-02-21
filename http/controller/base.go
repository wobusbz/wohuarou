package controller

import (
	"github.com/wobusbz/wohuarou/http/middleware"
	"net/http"
	"sync"
)

type BaseController struct {
	once *sync.Once
}

var NewBaseController = newBaseController()

func newBaseController() *BaseController {
	return &BaseController{
		once: new(sync.Once),
	}
}

func (c *BaseController) RegisterRoutes() {
	c.once.Do(func() {
		http.HandleFunc("/", middleware.Recover(c.Index))
	})
}
