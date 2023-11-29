package app

import "server/internal/app"

func main() {
	application := app.New()
	application.Run()
}
