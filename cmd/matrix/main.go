package main

import (
	"log"

	"github.com/ka1i/matrix/internal/app"
	"github.com/ka1i/matrix/pkg/version"
)

func init() {
	log.Println(version.Version.ToString())
}

func main() {
	app.App.Wispeeer()
}
