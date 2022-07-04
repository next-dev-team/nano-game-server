package main

import (
	"log"
	"nano-game-server/internal"
	"net/http"
	"os"
	"strings"

	"github.com/lonng/nano"
	"github.com/lonng/nano/component"
	"github.com/lonng/nano/serialize/json"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "Rolling Dice"
	app.Author = "nano authors"
	app.Version = "0.0.1"
	app.Copyright = "nano authors reserved"
	app.Usage = "rolling dice"

	// flags
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "addr",
			Value: ":3252",
			Usage: "game server address",
		},
	}

	app.Action = serve

	app.Run(os.Args)
}

func serve(ctx *cli.Context) error {
	components := &component.Components{}
	components.Register(internal.NewDiceManager(),
		component.WithName("dice"),
		component.WithNameFunc(strings.ToLower),
	)
	// components.Register(logic.NewWorld())

	// register all service
	options := []nano.Option{
		nano.WithIsWebsocket(true),
		nano.WithWSPath("game"),
		nano.WithComponents(components),
		nano.WithSerializer(json.NewSerializer()),
		nano.WithCheckOriginFunc(func(_ *http.Request) bool { return true }),
	}

	//nano.EnableDebug()
	log.SetFlags(log.LstdFlags | log.Llongfile)

	addr := ctx.String("addr")
	nano.Listen(addr, options...)
	return nil
}
