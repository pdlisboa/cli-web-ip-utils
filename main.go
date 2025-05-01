package main

import (
	"cli-web-ip-utils/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Start")

	app := app.GetApp()

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
