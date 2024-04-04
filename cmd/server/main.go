package main

import (
	"context"
	"log"

	"github.com/BobrePatre/ProjectTemplate/internal/app"
)

func main() {
	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("failed to init app: %s", err.Error())
	}

	a.Run()
	if err != nil {
		log.Fatalf("failed to run app: %s", err.Error())
	}
}
