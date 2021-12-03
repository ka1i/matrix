package app

import (
	"errors"
	"log"
	"os"

	"github.com/ka1i/matrix/internal/app/graphical"
	"github.com/ka1i/matrix/internal/pkg/usage"
	"github.com/ka1i/matrix/pkg/version"
)

type app struct {
	success int
	failure int
}

var App = GetApp()

func GetApp() *app {
	return &app{
		success: 0,
		failure: 137,
	}
}

func (app *app) processLaunch() error {
	graphical.UserInterface()
	return nil
}

func (app *app) argumentsParser(argv []string) error {
	var err error

	switch argv[0] {
	case "-f", "file":
		err = app.processLaunch()
	case "-h", "--help", "help":
		usage.Usage()
	case "-v", "--version", "version":
		version.Version.Print()
	default:
		err = errors.New("please check usage")
	}
	if err != nil {
		return err
	}
	return nil
}

func (app *app) Wispeeer() int {
	var err error
	if len(os.Args) > 1 {
		var argv = os.Args[1:]
		err = app.argumentsParser(argv)
	} else {
		err = app.processLaunch()
	}

	if err != nil {
		log.Printf("%v\n", err)
		return app.failure
	}
	return app.success
}
