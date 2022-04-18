package main

import "salescrm/internal/app/cli"

func main() {
	cli.Migrate()
	cli.Seed()
	cli.ServeApplication()
}
