package modules

import (
	"github.com/facebookgo/inject"
)

type AppConfig struct {
	application *Moscato
}

func (app *AppConfig) config() *Moscato {
	var graph inject.Graph

	err := graph.Provide(
		&inject.Object{Value: NewMStable()},
		&inject.Object{Value: app.application},
		&inject.Object{Value: NewSecurity()})

	if err != nil {
		println(err)
		return nil
	}

	err = graph.Populate()
	if err != nil {
		println(err)
		return nil
	}
	return app.application
}
