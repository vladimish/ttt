package main

import (
	"github.com/vladimish/ttt/internal/app"
	"github.com/vladimish/ttt/internal/cfg"
	"github.com/vladimish/ttt/internal/sqlite"
)

func main() {
	db, err := sqlite.NewSqlite(cfg.Get().DataSource)
	if err != nil {
		panic(err)
	}

	a := app.NewApp(cfg.Get(), db)
	err = a.Run()
	if err != nil {
		panic(err)
	}
}
