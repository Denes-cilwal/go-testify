package main

import (
	"go-testify/app"

	"go.uber.org/fx"
)

func main() {
	// create a Fx application
	fx.New(app.Module).Run()
}
