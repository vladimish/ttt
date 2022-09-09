package app

import (
	"github.com/vladimish/ttt/internal/cfg"
	"github.com/vladimish/ttt/internal/domain"
	"github.com/vladimish/ttt/internal/repositories"

	"github.com/c-bata/go-prompt"
)

type App struct {
	c *cfg.Config
	d *domain.Domain
}

func NewApp(config *cfg.Config, journal repositories.Journal) *App {
	d := domain.NewDomain(journal)
	return &App{
		c: config,
		d: d,
	}
}

func (a *App) Run() error {
	p := prompt.New(
		a.executor,
		a.completer,
		prompt.OptionPrefix(">>> "),
	)

	p.Run()

	return nil
}
