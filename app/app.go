package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func GetApp() *cli.App {
	app := cli.NewApp()
	app.Name = "Cli web ip utils"
	app.Usage = "Use this to get info about Ip and servers in web"
	applyCommands(app)

	return app
}

func applyCommands(app *cli.App) {
	flag := cli.StringFlag{
		Name: "host",
	}
	commands := []cli.Command{
		{
			Name:        "ip",
			Description: "Get ip address",
			Flags:       []cli.Flag{flag},
			Action:      getIps,
		},
		{
			Name:        "server",
			Description: "Get server name",
			Flags:       []cli.Flag{flag},
			Action:      getServerName,
		},
	}
	app.Commands = commands
}

func getIps(ctx *cli.Context) {
	host := ctx.String("host")

	ips, err := net.LookupIP(host)

	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func getServerName(ctx *cli.Context) {
	host := ctx.String("host")

	servers, err := net.LookupNS(host)

	if err != nil {
		log.Fatal(err)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}

}
