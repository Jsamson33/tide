package main

import (
	"os"
	"github.com/Jsamson33/tide/internal/adapters/cli"
	"github.com/Jsamson33/tide/internal/adapters/file"
	"github.com/Jsamson33/tide/internal/domain"
)

func main() {
	repo := file.NewFileRepository()
	service := domain.NewTideService(repo)
	app := cli.NewApp(service)

	app.Run(os.Args)
}
