package http

import "github.com/wobusbz/wohuarou/http/controller"

func RegisterRoutes() {
	controller.NewBaseController.RegisterRoutes()
}
