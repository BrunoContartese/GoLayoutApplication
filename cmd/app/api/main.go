package main

import "layoutapp/internal/app/cli"

func main() {
	cli.Migrate()
	cli.Seed()
	cli.ServeApplication()
}
