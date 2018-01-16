package controller

import "github.com/StpDevelo/go-simple-project/pkg/app"

type ctrl struct {
	view app.View
}

// Config is a config controller
type Config struct {
	View app.View
}

// New is a create controller
func New(config Config) app.Controller {
	return &ctrl{
		view: config.View,
	}
}
