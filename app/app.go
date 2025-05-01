package app

import "github.com/urfave/cli"

func GetApp() *cli.App {
	app := cli.NewApp()
	app.Name = "Cli web ip utils"
	app.Usage = "Use this to get info about Ip and servers in web"

	return app
}
