package example_service

import (
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)
	godotenv.Load()

	app := cli.App{
		Before: func(ctx *cli.Context) error {
			return nil
		},
		Commands: []*cli.Command{
			{
				Name: "grpc",
				Action: func(context *cli.Context) error {

					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		return
	}
}
